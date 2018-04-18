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
	Upvotes     map[string]bool `json:"upvotes"`
	Downvotes   map[string]bool `json:"downvotes"`
	TotalVotes  int             `json:"total_votes"`
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
	Title      string          `json:"title"`
	Caption    string          `json:"caption"`
	ImageURL   string          `json:"image_url"`
	Upvotes    map[string]bool `json:"upvotes"`
	Downvotes  map[string]bool `json:"downvotes"`
	TotalVotes int             `json:"total_votes"`
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
		ID:          bson.NewObjectId(),
		Title:       np.Title,
		ImageURL:    "",
		Caption:     "",
		AuthorID:    np.AuthorID,
		Comments:    []bson.ObjectId{},
		BoardID:     bson.NewObjectId(),
		Upvotes:     map[string]bool{},
		Downvotes:   map[string]bool{},
		TotalVotes:  0,
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

//HasVote returns 1 if a user has upvoted, -1 if
//a user has downvoted, and 0 if a user has not
//voted
func (p *Post) HasVote(author bson.ObjectId) int {
	if val, ok := p.Upvotes[author.Hex()]; ok && val == true {
		return 1
	}
	if val, ok := p.Downvotes[author.Hex()]; ok && val == true {
		return -1
	}
	return 0
}

//Upvote modifies the current score (count of
//upvotes) for the post, returning a *PostUpdate
func (p *Post) Upvote(author bson.ObjectId) *PostUpdate {
	update := &PostUpdate{
		Title:      p.Title,
		ImageURL:   p.ImageURL,
		Caption:    p.Caption,
		Upvotes:    p.Upvotes,
		Downvotes:  p.Downvotes,
		TotalVotes: p.TotalVotes,
	}
	if p.HasVote(author) == -1 {
		update.Downvotes[author.Hex()] = false
	}
	if p.HasVote(author) != 1 {
		//checks if already upvoted
		update.Upvotes[author.Hex()] = true
		update.TotalVotes++

	}
	return update
}

//Downvote downvotes the post
func (p *Post) Downvote(author bson.ObjectId) *PostUpdate {
	update := &PostUpdate{
		Title:      p.Title,
		ImageURL:   p.ImageURL,
		Caption:    p.Caption,
		Upvotes:    p.Upvotes,
		Downvotes:  p.Downvotes,
		TotalVotes: p.TotalVotes,
	}
	if p.HasVote(author) == 1 {
		update.Downvotes[author.Hex()] = false
	}
	if p.HasVote(author) != -1 {
		update.Upvotes[author.Hex()] = true
		update.TotalVotes--
	}
	return update
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
	p.TotalVotes = updates.TotalVotes
	p.Upvotes = updates.Upvotes
	p.Downvotes = updates.Downvotes
	p.TimeEdited = time.Now()
	return nil
}

//AddComments adds comment IDs to the Post
//TODO: DOES NOT ACTUALLY GET STORED TO DB AT ALL SO... FIX
//change to "reply"? also can i just lump this into PostUpdate?
func (p *Post) AddComments(comment bson.ObjectId) {
	p.Comments = append(p.Comments, comment)
}
