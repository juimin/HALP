package comments

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

// Test the validate function in comments
func TestValidate(t *testing.T) {

	cases := []struct {
		name     string
		nc       *NewComment
		expected error
	}{
		{
			name: "Proper New Comment",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - Invalid Author",
			nc: &NewComment{
				AuthorID: "asdfasfdfd",
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Error: Invalid Author ID"),
		},
		{
			name: "New Comment - Invalid Post ID",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   "asdfasfdfd231--df?",
			},
			expected: fmt.Errorf("Error: Invalid Post ID"),
		},
		{
			name: "New Comment - No Content",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image or Content",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Cannot have no image or content in comment"),
		},
	}

	for _, c := range cases {
		err := c.nc.Validate()
		if err != c.expected {
			if err != nil {
				if c.expected != nil {
					if c.expected.Error() != err.Error() {
						t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expected, err)
					}
				}
			}
		}
	}
}

func TestToComment(t *testing.T) {

	cases := []struct {
		name     string
		nc       *NewComment
		expected error
	}{
		{
			name: "Proper New Comment",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - Invalid Author",
			nc: &NewComment{
				AuthorID: "asdfasfdfd",
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Error: Invalid Author ID"),
		},
		{
			name: "New Comment - Invalid Post ID",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   "asdfasfdfd231--df?",
			},
			expected: fmt.Errorf("Error: Invalid Post ID"),
		},
		{
			name: "New Comment - No Content",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image or Content",
			nc: &NewComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Cannot have no image or content in comment"),
		},
	}

	for _, c := range cases {
		comment, err := c.nc.ToComment()
		if err != c.expected {
			if err != nil {
				if c.expected != nil {
					if c.expected.Error() != err.Error() {
						t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expected, err)
					}
				}
			}
		}
		if comment != nil {
			if comment.AuthorID != c.nc.AuthorID {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.nc.AuthorID, comment.AuthorID)
			}
			if comment.Content != c.nc.Content {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.nc.Content, comment.Content)
			}
			if comment.ImageURL != c.nc.ImageURL {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.nc.ImageURL, comment.ImageURL)
			}
			if comment.PostID != c.nc.PostID {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.nc.PostID, comment.PostID)
			}
		}
	}
}

func TestVote(t *testing.T) {

	cases := []struct {
		name          string
		comment       *Comment
		vote          *CommentVote
		expectedUp    int
		expectedDown  int
		expectedError error
	}{
		{
			name: "Invalid Switching Vote Comment +1 +1",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   1,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: fmt.Errorf("Can't increment upvote and downvote"),
		},
		{
			name: "Invalid Switching Vote Comment -1 -1",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   -1,
				Downvote: -1,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: fmt.Errorf("Can't decrement both upvote and downvote"),
		},
		{
			name: "Valid Switching Vote Comment +1 -1",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   1,
				Downvote: -1,
			},
			expectedUp:    1,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment -1 +1",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   -1,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  1,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +0 +0",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   0,
				Downvote: 0,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +0 +1",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   0,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  1,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +1 +0",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   1,
				Downvote: 0,
			},
			expectedUp:    1,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Invalid Switching Vote Comment",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   1123123123,
				Downvote: 0,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Invalid Switching Vote Comment",
			comment: &Comment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &CommentVote{
				Upvote:   1,
				Downvote: 2320,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		err := c.comment.Vote(c.vote)
		if err != c.expectedError {
			if err != nil {
				if c.expectedError != nil {
					if c.expectedError.Error() != err.Error() {
						t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expectedError, err)
					}
				}
			}
		}
		if err == nil {
			if c.comment.Upvotes != c.expectedUp {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expectedUp, c.comment.Upvotes)
			}
			if c.comment.Downvotes != c.expectedDown {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expectedDown, c.comment.Downvotes)
			}
		}
	}
}

// Test Update Functions
func TestUpdateComment(t *testing.T) {

	cases := []struct {
		name          string
		comment       *Comment
		updates       *CommentUpdate
		expectedError error
	}{
		{
			name:    "Valid Update",
			comment: &Comment{},
			updates: &CommentUpdate{
				ImageURL: "https://potato.com",
				Comments: []bson.ObjectId{
					bson.NewObjectId(),
					bson.NewObjectId(),
				},
				Content: "Something",
			},
			expectedError: nil,
		},
		{
			name:    "Bad Update",
			comment: &Comment{},
			updates: &CommentUpdate{
				ImageURL: "",
				Comments: []bson.ObjectId{
					bson.NewObjectId(),
					bson.NewObjectId(),
				},
				Content: "",
			},
			expectedError: fmt.Errorf("We cannot set the comment to contain nothing"),
		},
	}

	for _, c := range cases {
		err := c.comment.Update(c.updates)
		if err != c.expectedError {
			if err != nil {
				if c.expectedError != nil {
					if c.expectedError.Error() != err.Error() {
						t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expectedError, err)
					}
				}
			}
		}
		if err == nil {
			if c.comment.ImageURL != c.updates.ImageURL {
				t.Errorf("Error updating image url")
			}
			if c.comment.Content != c.updates.Content {
				t.Errorf("Error updating Content")
			}
			for i, comm := range c.comment.Comments {
				if c.updates.Comments[i] != comm {
					t.Errorf("Error updating Comments list")
				}
			}
		}
	}
}
