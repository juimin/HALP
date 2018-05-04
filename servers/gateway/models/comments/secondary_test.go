package comments

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestValidateSecondary(t *testing.T) {

	cases := []struct {
		name     string
		nc       *NewSecondaryComment
		expected error
	}{
		{
			name: "Proper New Comment",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - Invalid Author",
			nc: &NewSecondaryComment{
				AuthorID: "asdfasfdfd",
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Error: Invalid Author ID"),
		},
		{
			name: "New Comment - Invalid Post ID",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   "asdfasfdfd231--df?",
				Parent:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Error: Invalid Post ID"),
		},
		{
			name: "New Comment - No Content",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image or Content",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Cannot have no image or content in comment"),
		},
		{
			name: "New Comment - Bad Parent ID",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
				Parent:   "1232dd--d?Q12",
			},
			expected: fmt.Errorf("Error: Invalid Parent ID"),
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

func TestToCommentSecondary(t *testing.T) {

	cases := []struct {
		name     string
		nc       *NewSecondaryComment
		expected error
	}{
		{
			name: "Proper New Comment",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - Invalid Author",
			nc: &NewSecondaryComment{
				AuthorID: "asdfasfdfd",
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Error: Invalid Author ID"),
		},
		{
			name: "New Comment - Invalid Post ID",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "https://something.com",
				PostID:   "asdfasfdfd231--df?",
				Parent:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Error: Invalid Post ID"),
		},
		{
			name: "New Comment - No Content",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "https://something.com",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: nil,
		},
		{
			name: "New Comment - No Image or Content",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
				Parent:   bson.NewObjectId(),
			},
			expected: fmt.Errorf("Cannot have no image or content in comment"),
		},
		{
			name: "New Comment - Bad Parent ID",
			nc: &NewSecondaryComment{
				AuthorID: bson.NewObjectId(),
				Content:  "This is valid content",
				ImageURL: "",
				PostID:   bson.NewObjectId(),
				Parent:   "1232dd--d?Q12",
			},
			expected: fmt.Errorf("Error: Invalid Parent ID"),
		},
	}

	for _, c := range cases {
		comment, err := c.nc.ToSecondaryComment()
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
			if comment.Parent != c.nc.Parent {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.nc.Parent, comment.Parent)
			}
		}
	}
}

func TestVoteSecondary(t *testing.T) {

	cases := []struct {
		name          string
		comment       *SecondaryComment
		vote          *SecondaryCommentVote
		expectedUp    int
		expectedDown  int
		expectedError error
	}{
		{
			name: "Invalid Switching Vote Comment +1 +1",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   1,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: fmt.Errorf("Can't increment upvote and downvote"),
		},
		{
			name: "Invalid Switching Vote Comment -1 -1",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   -1,
				Downvote: -1,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: fmt.Errorf("Can't decrement both upvote and downvote"),
		},
		{
			name: "Valid Switching Vote Comment +1 -1",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   1,
				Downvote: -1,
			},
			expectedUp:    1,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment -1 +1",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   -1,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  1,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +0 +0",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   0,
				Downvote: 0,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +0 +1",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   0,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  1,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +1 +0",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   1,
				Downvote: 0,
			},
			expectedUp:    1,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "InValid Switching Vote Comment",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   123123,
				Downvote: 0,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "InValid Switching Vote Comment",
			comment: &SecondaryComment{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &SecondaryCommentVote{
				Upvote:   0,
				Downvote: 2312312,
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

func TestUpdateCommentSecondary(t *testing.T) {

	cases := []struct {
		name          string
		comment       *SecondaryComment
		updates       *SecondaryCommentUpdate
		expectedError error
	}{
		{
			name:    "Valid Update",
			comment: &SecondaryComment{},
			updates: &SecondaryCommentUpdate{
				ImageURL: "https://potato.com",
				Content:  "Something",
			},
			expectedError: nil,
		},
		{
			name:    "Bad Update",
			comment: &SecondaryComment{},
			updates: &SecondaryCommentUpdate{
				ImageURL: "",
				Content:  "",
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
		}
	}
}
