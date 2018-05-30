package users

import (
	"fmt"
	"net/mail"

	"gopkg.in/mgo.v2/bson"
)

// UserUpdate represents allowed updates to a user profile
type UserUpdate struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

// FavoritesUpdate holds the new list of favorites from the client
// Adding boolean determines if the update is a removal or an addition
type FavoritesUpdate struct {
	Adding   bool          `json:"adding"`
	UpdateID bson.ObjectId `json:"updateID"`
}

// BookmarksUpdate holds the new list of favorites from the clinet
type BookmarksUpdate struct {
	Adding   bool          `json:"adding"`
	UpdateID bson.ObjectId `json:"updateID"`
}

// PostVoting handles voting on posts
type PostVoting struct {
	Vote   bool          `json:"vote"`
	PostID bson.ObjectId `json:"postId"`
	Remove bool          `json:"remove"`
}

// CommentVoting handles voting on cvomments
type CommentVoting struct {
	Vote      bool          `json:"vote"`
	CommentID bson.ObjectId `json:"commentID"`
	Remove    bool          `json:"remove"`
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *UserUpdate) error {
	if len(updates.FirstName) == 0 || len(updates.LastName) == 0 {
		return fmt.Errorf("Invalid input. First and last name must both have a non-zero length")
	}

	// We can't deal with empty emails either because this is not optional
	if len(updates.Email) == 0 {
		return fmt.Errorf("Invalid Input. Email cannot be empty")
	}

	// Check Email valid
	if _, err := mail.ParseAddress(updates.Email); err != nil {
		return fmt.Errorf("Invalid input. Email not a valid email")
	}

	// We aren't dealing with occupation because it is optional
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName
	u.Email = updates.Email
	u.Occupation = updates.Occupation

	return nil
}

// UpdateFavorites updates the user's favorites with the given update
func (u *User) UpdateFavorites(updates *FavoritesUpdate) error {
	if updates.Adding {
		u.Favorites = append(u.Favorites, updates.UpdateID)
		return nil
	}
	success := false
	for idx, item := range u.Favorites {
		if item == updates.UpdateID {
			// Eliminate the item from the slice
			u.Favorites = append(u.Favorites[:idx], u.Favorites[idx+1:]...)
			success = true
		}
	}
	if !success {
		return fmt.Errorf("Could not remove the item from the favorites")
	}
	return nil
}

// UpdateBookmarks updates the user's favorites with the given update
func (u *User) UpdateBookmarks(updates *BookmarksUpdate) error {
	if updates.Adding {
		u.Bookmarks = append(u.Bookmarks, updates.UpdateID)
		return nil
	}
	success := false
	for idx, item := range u.Bookmarks {
		if item == updates.UpdateID {
			// Eliminate the item from the slice
			u.Bookmarks = append(u.Bookmarks[:idx], u.Bookmarks[idx+1:]...)
			success = true
		}
	}
	if !success {
		return fmt.Errorf("Could not remove the item from the favorites")
	}
	return nil
}

// PostVote handles updating the user with a vote
func (u *User) PostVote(updates *PostVoting) error {
	// If there is a thing in here, then we need to figure out what to do
	if _, ok := u.PostVotes[updates.PostID.Hex()]; ok {
		if updates.Remove {
			// Remove the entry
			delete(u.PostVotes, updates.PostID.Hex())
		} else {
			u.PostVotes[updates.PostID.Hex()] = updates.Vote
		}
	} else {
		if !updates.Remove {
			u.PostVotes[updates.PostID.Hex()] = updates.Vote
		}
	}
	return nil
}

// CommentVote handles updating the user with a vote
func (u *User) CommentVote(updates *CommentVoting) error {
	// If there is a thing in here, then we need to figure out what to do
	if _, ok := u.CommentVotes[updates.CommentID.Hex()]; ok {
		if updates.Remove {
			// Remove the entry
			delete(u.PostVotes, updates.CommentID.Hex())
		} else {
			u.PostVotes[updates.CommentID.Hex()] = updates.Vote
		}
	} else {
		if !updates.Remove {
			u.PostVotes[updates.CommentID.Hex()] = updates.Vote
		}
	}
	return nil
}
