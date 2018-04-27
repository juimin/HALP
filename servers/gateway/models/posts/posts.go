package posts

import (
	"fmt"
	"net/url"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Post is a post
type Post struct {
	ID          bson.ObjectId   `json:"id" bson:"_id"`
	Title       string          `json:"title"`
	ImageURL    string          `json:"image_url"`
	Caption     string          `json:"caption"`
	AuthorID    bson.ObjectId   `json:"author_id"`
	Comments    []bson.ObjectId `json:"comments"`
	BoardID     bson.ObjectId   `json:"board_id"`
	Upvotes     int             `json:"upvotes"`
	Downvotes   int             `json:"downvotes"`
	TimeCreated time.Time       `json:"time_created"`
	TimeEdited  time.Time       `json:"time_edited"`
}

//NewPost is a new post
type NewPost struct {
	Title    string        `json:"title"`
	ImageURL string        `json:"image_url"`
	Caption  string        `json:"caption"`
	AuthorID bson.ObjectId `json:"author_id"`
	BoardID  bson.ObjectId `json:"board_id"`
}

//PostUpdate represents allowed updates to a post
type PostUpdate struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	ImageURL string `json:"image_url"`
}

// PostVote updater
type PostVote struct {
	Upvote   int `json:"upvotes"`
	Downvote int `json:"downvotes"`
}

//Validate confirms that a new post contains a title
//and at least one of either an image or a caption
func (np *NewPost) Validate() error {
	//Check for a title
	if len(np.Title) == 0 {
		return fmt.Errorf("Please enter a title")
	}
	//Check for either caption or image
	if len(np.Caption) == 0 && len(np.ImageURL) == 0 {
		return fmt.Errorf("Please use either an image or a caption")
	}
	//if image, verify url
	if len(np.ImageURL) > 0 {
		if _, err := url.ParseRequestURI(np.ImageURL); err != nil {
			return fmt.Errorf("Invalid Photo URL")
		}
	}
	return nil
}

//ToPost converts the NewPost to a real Post
func (np *NewPost) ToPost() (*Post, error) {
	err := np.Validate()
	if err != nil {
		return nil, err
	}
	post := &Post{
		ID:        bson.NewObjectId(),
		Title:     np.Title,
		ImageURL:  "",
		Caption:   "",
		AuthorID:  np.AuthorID,
		Comments:  []bson.ObjectId{},
		BoardID:   bson.NewObjectId(),
		Upvotes:   0,
		Downvotes: 0,
		//TotalVotes:  0,
		TimeCreated: time.Now(),
		TimeEdited:  time.Now(),
	}

	if len(np.ImageURL) > 0 {
		post.ImageURL = np.ImageURL
	}
	if len(np.Caption) > 0 {
		post.Caption = np.Caption
	}
	return post, nil
}

//ApplyUpdates applies the post updates to the post
func (p *Post) ApplyUpdates(updates *PostUpdate) error {
	if len(updates.Title) == 0 {
		return fmt.Errorf("Title cannot be empty")
	}
	if len(updates.Caption) == 0 && len(updates.ImageURL) == 0 {
		return fmt.Errorf("Must have either a caption or an image")
	}
	if len(updates.Title) > 0 {
		p.Title = updates.Title
	}
	if len(updates.Caption) > 0 {
		p.Caption = updates.Caption
	}
	if len(updates.ImageURL) > 0 {
		if _, err := url.ParseRequestURI(updates.ImageURL); err != nil {
			return fmt.Errorf("Invalid Photo URL")
		}
		p.ImageURL = updates.ImageURL
	}
	// p.TotalVotes = updates.TotalVotes
	// p.Upvotes = updates.Upvotes
	// p.Downvotes = updates.Downvotes
	p.TimeEdited = time.Now()
	return nil
}

//AddComment adds comment IDs to the Post
//change to "reply"? also can i just lump this into PostUpdate?
func (p *Post) AddComment(comment bson.ObjectId) {
	p.Comments = append(p.Comments, comment)
}

// ApplyVotes takes care of applying the votes later
func (p *Post) ApplyVotes(updates *PostVote) error {
	// Check if the updates are within the defined bounds
	if updates.Downvote >= 1 || updates.Downvote <= -1 {
		return fmt.Errorf("Downvotes are out of bounds: %d", updates.Downvote)
	}
	if updates.Upvote >= 1 || updates.Upvote <= -1 {
		return fmt.Errorf("Upvotes are out of bounds: %d", updates.Upvote)
	}
	if updates.Downvote == updates.Upvote {
		return fmt.Errorf("The updates to both upvotes and downvotes cannot be the same")
	}
	// Make the update
	p.Upvotes += updates.Upvote
	p.Downvotes += updates.Downvote

	return nil
}
