package posts

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

//Add tests for various functions in posts.go

//Testing validation function
func TestValidate(t *testing.T) {
	cases := []struct {
		name           string
		input          *NewPost
		expectedOutput error
	}{
		{
			name:           "Empty",
			input:          &NewPost{},
			expectedOutput: fmt.Errorf("Please enter a title"),
		},
		{
			name: "valid - caption but no image",
			input: &NewPost{
				Title:    "Fun",
				ImageURL: "",
				Caption:  "How do I shot web?",
				AuthorID: bson.NewObjectId(),
				BoardID:  bson.NewObjectId(),
			},
			expectedOutput: nil,
		},
		{
			name: "valid - image but no caption",
			input: &NewPost{
				Title:    "Fun",
				ImageURL: "http://google.com",
				Caption:  "",
				AuthorID: bson.NewObjectId(),
				BoardID:  bson.NewObjectId(),
			},
			expectedOutput: nil,
		},
		{
			name: "!!!!!invalid - imageURL not valid",
			input: &NewPost{
				Title:    "Fun",
				ImageURL: "not a url",
				Caption:  "How do I shot web?",
				AuthorID: bson.NewObjectId(),
				BoardID:  bson.NewObjectId(),
			},
			expectedOutput: fmt.Errorf("Invalid Photo URL"),
		},
		{
			name: "invalid - neither image nor caption",
			input: &NewPost{
				Title:    "Fun",
				ImageURL: "",
				Caption:  "",
				AuthorID: bson.NewObjectId(),
				BoardID:  bson.NewObjectId(),
			},
			expectedOutput: fmt.Errorf("Please use either an image or a caption"),
		},
	}
	for _, c := range cases {
		result := c.input.Validate()
		if result != nil && c.expectedOutput != nil {
			if result.Error() != c.expectedOutput.Error() {
				t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
			}
		}
		if (result == nil || c.expectedOutput == nil) && !(result == nil && c.expectedOutput == nil) {
			t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
		}

	}
}

//TestToPost tests changing a NewPost to a Post
func TestToPost(t *testing.T) {
	cases := []struct {
		name           string
		input          *NewPost
		expectedOutput error
	}{
		{
			name: "validate Check - any error",
			input: &NewPost{
				Title:    "",
				ImageURL: "http://google.com",
				Caption:  "hello",
				AuthorID: bson.NewObjectId(),
				BoardID:  bson.NewObjectId(),
			},
			expectedOutput: fmt.Errorf("Please enter a title"),
		},
		{
			name: "valid input - no caption",
			input: &NewPost{
				Title:    "Hello",
				ImageURL: "http://google.com",
				Caption:  "",
				AuthorID: bson.NewObjectId(),
				BoardID:  bson.NewObjectId(),
			},
			expectedOutput: nil,
		},
		{
			name: "valid input - no image",
			input: &NewPost{
				Title:    "Hello",
				ImageURL: "",
				Caption:  "hi",
				AuthorID: bson.NewObjectId(),
				BoardID:  bson.NewObjectId(),
			},
			expectedOutput: nil,
		},
	}
	for _, c := range cases {
		_, result := c.input.ToPost()
		if result != nil && c.expectedOutput != nil {
			if result.Error() != c.expectedOutput.Error() {
				t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
			}
		}
		if (result == nil || c.expectedOutput == nil) && !(result == nil && c.expectedOutput == nil) {
			t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
		}
	}
}

//TestApplyUpdates tests applying updates to an existing post
func TestApplyUpdates(t *testing.T) {
	cases := []struct {
		name           string
		input          *PostUpdate
		expectedOutput error
	}{
		{
			name: "empty title",
			input: &PostUpdate{
				Title:    "",
				ImageURL: "",
				Caption:  "yes",
			},
			expectedOutput: fmt.Errorf("Title cannot be empty"),
		},
		{
			name: "no caption or image",
			input: &PostUpdate{
				Title:    "yello",
				ImageURL: "",
				Caption:  "",
			},
			expectedOutput: fmt.Errorf("Must have either a caption or an image"),
		},
		{
			name: "invalid image url",
			input: &PostUpdate{
				Title:    "yello",
				ImageURL: "googleurl",
				Caption:  "",
			},
			expectedOutput: fmt.Errorf("Invalid Photo URL"),
		},
		{
			name: "valid - update caption",
			input: &PostUpdate{
				Title:    "yello",
				ImageURL: "",
				Caption:  "testing",
			},
			expectedOutput: nil,
		},
		{
			name: "valid - update image url",
			input: &PostUpdate{
				Title:    "yello",
				ImageURL: "http://google.com",
				Caption:  "",
			},
			expectedOutput: nil,
		},
	}
	testPost := &Post{}
	for _, c := range cases {
		result := testPost.ApplyUpdates(c.input)
		if result != nil && c.expectedOutput != nil {
			if result.Error() != c.expectedOutput.Error() {
				t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
			}
		}
		if (result == nil || c.expectedOutput == nil) && !(result == nil && c.expectedOutput == nil) {
			t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
		}
	}
}

//Test AddComments tests adding comment ids to a post
func TestAddComments(t *testing.T) {
	testslice := []bson.ObjectId{}
	testcomment := bson.NewObjectId()
	testslice = append(testslice, testcomment)
	cases := []struct {
		name           string
		input          bson.ObjectId
		expectedOutput []bson.ObjectId
	}{
		{
			name:           "add comment",
			input:          testcomment,
			expectedOutput: testslice,
		},
	}
	testPost := &Post{}
	for _, c := range cases {
		testPost.AddComment(c.input)
		output := testPost.Comments
		for i, num := range output {
			if num != c.expectedOutput[i] {
				t.Errorf("%s: got %s but expected %s", c.name, num, c.expectedOutput[i])
			}
		}
	}
}

// TestApplyVotes tests the system used for applying votes to a post
func TestApplyVotes(t *testing.T) {

	cases := []struct {
		name          string
		post          *Post
		vote          *PostVote
		expectedUp    int
		expectedDown  int
		expectedError error
	}{
		{
			name: "Invalid Switching Vote Comment +1 +1",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   1,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: fmt.Errorf("The updates to both upvotes and downvotes cannot be the same"),
		},
		{
			name: "Invalid Switching Vote Comment -1 -1",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   -1,
				Downvote: -1,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: fmt.Errorf("The updates to both upvotes and downvotes cannot be the same"),
		},
		{
			name: "Valid Switching Vote Comment +1 -1",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   1,
				Downvote: -1,
			},
			expectedUp:    1,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment -1 +1",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   -1,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  1,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +0 +0",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   0,
				Downvote: 0,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: fmt.Errorf("The updates to both upvotes and downvotes cannot be the same"),
		},
		{
			name: "Valid Switching Vote Comment +0 +1",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   0,
				Downvote: 1,
			},
			expectedUp:    0,
			expectedDown:  1,
			expectedError: nil,
		},
		{
			name: "Valid Switching Vote Comment +1 +0",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   1,
				Downvote: 0,
			},
			expectedUp:    1,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Invalid Switching Vote Comment",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   1123123123,
				Downvote: 0,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: nil,
		},
		{
			name: "Invalid Switching Vote Comment",
			post: &Post{
				Upvotes:   0,
				Downvotes: 0,
			},
			vote: &PostVote{
				Upvote:   1,
				Downvote: 2320,
			},
			expectedUp:    0,
			expectedDown:  0,
			expectedError: nil,
		},
	}

	for _, c := range cases {
		err := c.post.ApplyVotes(c.vote)
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
			if c.post.Upvotes != c.expectedUp {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expectedUp, c.post.Upvotes)
			}
			if c.post.Downvotes != c.expectedDown {
				t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expectedDown, c.post.Downvotes)
			}
		}
	}
}
