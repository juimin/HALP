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

// //TestHasUpvote tests getting how the user has voted on a post
// func TestHasUpvote(t *testing.T) {
// 	testuser := bson.NewObjectId()
// 	mapwithuser := map[string]bool{}
// 	mapwithuser[testuser.Hex()] = true
// 	cases := []struct {
// 		name           string
// 		input          *Post
// 		expectedOutput int
// 	}{
// 		{
// 			name: "no votes",
// 			input: &Post{
// 				ID:          bson.NewObjectId(),
// 				Title:       "",
// 				ImageURL:    "",
// 				Caption:     "",
// 				Comments:    []bson.ObjectId{},
// 				BoardID:     bson.NewObjectId(),
// 				Upvotes:     map[string]bool{},
// 				Downvotes:   map[string]bool{},
// 				TotalVotes:  0,
// 				TimeCreated: time.Now(),
// 				TimeEdited:  time.Now()},
// 			expectedOutput: 0,
// 		},
// 		{
// 			name: "user has upvote",
// 			input: &Post{
// 				ID:          bson.NewObjectId(),
// 				Title:       "",
// 				ImageURL:    "",
// 				Caption:     "",
// 				Comments:    []bson.ObjectId{},
// 				BoardID:     bson.NewObjectId(),
// 				Upvotes:     mapwithuser,
// 				Downvotes:   map[string]bool{},
// 				TotalVotes:  0,
// 				TimeCreated: time.Now(),
// 				TimeEdited:  time.Now(),
// 			},
// 			expectedOutput: 1,
// 		},
// 		{
// 			name: "user has downvote",
// 			input: &Post{
// 				ID:          bson.NewObjectId(),
// 				Title:       "",
// 				ImageURL:    "",
// 				Caption:     "",
// 				Comments:    []bson.ObjectId{},
// 				BoardID:     bson.NewObjectId(),
// 				Upvotes:     map[string]bool{},
// 				Downvotes:   mapwithuser,
// 				TotalVotes:  0,
// 				TimeCreated: time.Now(),
// 				TimeEdited:  time.Now(),
// 			},
// 			expectedOutput: -1,
// 		},
// 	}
// 	for _, c := range cases {
// 		if output := c.input.HasVote(testuser); c.expectedOutput != output {
// 			t.Errorf("%s: got %d but expected %d", c.name, output, c.expectedOutput)
// 		}
// 	}
// }

// //TestUpvote tests upvoting a post
// func TestUpvote(t *testing.T) {
// 	testuser1 := bson.NewObjectId()
// 	testuser2 := bson.NewObjectId()
// 	testmap1 := map[string]bool{}
// 	testmap1[testuser1.Hex()] = true
// 	testmap2 := map[string]bool{}
// 	testmap2[testuser1.Hex()] = true
// 	testmap2[testuser2.Hex()] = true
// 	cases := []struct {
// 		name           string
// 		input          bson.ObjectId
// 		expectedOutput int //represents total votes
// 		expectedMap    map[string]bool
// 	}{
// 		{
// 			name:           "first upvote",
// 			input:          testuser1,
// 			expectedOutput: 1,
// 			expectedMap:    testmap1,
// 		},
// 		{
// 			name:           "second upvote",
// 			input:          testuser2,
// 			expectedOutput: 2,
// 			expectedMap:    testmap2,
// 		},
// 	}
// 	testPost := &Post{
// 		Title:     "testing",
// 		Caption:   "testing",
// 		Upvotes:   map[string]bool{},
// 		Downvotes: map[string]bool{},
// 	}
// 	for _, c := range cases {
// 		//testpost.upvote should return a postupdate
// 		//with the updated upvotes map and total votes
// 		//then appy that to testpost
// 		testPost.ApplyUpdates(testPost.Upvote(c.input))
// 		output := testPost.TotalVotes
// 		if output != c.expectedOutput {
// 			t.Errorf("%s: got %d but expected %d", c.name, output, c.expectedOutput)
// 		}
// 		outmap := testPost.Upvotes
// 		for k, v := range outmap {
// 			if v != c.expectedMap[k] {
// 				t.Errorf("%s: got (%s, %s) but expected (%s, %s)", c.name, k, strconv.FormatBool(v), k, strconv.FormatBool(c.expectedMap[k]))
// 			}
// 		}
// 	}
// }

// //TestDownvote tests downvoting a post
// func TestDownvote(t *testing.T) {
// 	testuser1 := bson.NewObjectId()
// 	testuser2 := bson.NewObjectId()
// 	testmap1 := map[string]bool{}
// 	testmap1[testuser1.Hex()] = true
// 	testmap2 := map[string]bool{}
// 	testmap2[testuser1.Hex()] = true
// 	testmap2[testuser2.Hex()] = true
// 	cases := []struct {
// 		name           string
// 		input          bson.ObjectId
// 		expectedOutput int //represents total votes
// 		expectedMap    map[string]bool
// 	}{
// 		{
// 			name:           "first downvote",
// 			input:          testuser1,
// 			expectedOutput: -1,
// 			expectedMap:    testmap1,
// 		},
// 		{
// 			name:           "second downvote",
// 			input:          testuser2,
// 			expectedOutput: -2,
// 			expectedMap:    testmap2,
// 		},
// 	}
// 	testPost := &Post{
// 		Title:     "testing",
// 		Caption:   "testing",
// 		Upvotes:   map[string]bool{},
// 		Downvotes: map[string]bool{},
// 	}
// 	for _, c := range cases {
// 		testPost.ApplyUpdates(testPost.Downvote(c.input))
// 		output := testPost.TotalVotes
// 		if output != c.expectedOutput {
// 			t.Errorf("%s: got %d but expected %d", c.name, output, c.expectedOutput)
// 		}
// 		outmap := testPost.Downvotes
// 		for k, v := range outmap {
// 			if v != c.expectedMap[k] {
// 				t.Errorf("%s: got (%s, %s) but expected (%s, %s)", c.name, k, strconv.FormatBool(v), k, strconv.FormatBool(c.expectedMap[k]))
// 			}
// 		}
// 	}
// }

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
		testPost.AddComments(c.input)
		output := testPost.Comments
		for i, num := range output {
			if num != c.expectedOutput[i] {
				t.Errorf("%s: got %s but expected %s", c.name, num, c.expectedOutput[i])
			}
		}
	}
}

//Test Upvoting
func TestUpvote(t *testing.T) {
	cases := []struct {
		name           string
		expectedOutput int
	}{
		{
			name:           "add upvote",
			expectedOutput: 1,
		},
	}
	testPost := &Post{}
	for _, c := range cases {
		testPost.Upvote()
		if testPost.Upvotes != c.expectedOutput {
			t.Errorf("%s: got %d upvotes but expected %d upvotes", c.name, testPost.Upvotes, c.expectedOutput)
		}

	}
}

//Test removing updates from post
func TestRemoveUpvote(t *testing.T) {
	cases := []struct {
		name           string
		expectedOutput int
	}{
		{
			name:           "removing upvote",
			expectedOutput: 0,
		},
	}
	testPost := &Post{
		Upvotes: 1,
	}
	for _, c := range cases {
		testPost.RemoveUpvote()
		if testPost.Upvotes != c.expectedOutput {
			t.Errorf("%s: got %d upvotes but expected %d upvotes", c.name, testPost.Upvotes, c.expectedOutput)
		}

	}
}

//Test Downvoting
func TestDownvote(t *testing.T) {
	cases := []struct {
		name           string
		expectedOutput int
	}{
		{
			name:           "add downvote",
			expectedOutput: 1,
		},
	}
	testPost := &Post{}
	for _, c := range cases {
		testPost.Downvote()
		if testPost.Downvotes != c.expectedOutput {
			t.Errorf("%s: got %d downvotes but expected %d upvotes", c.name, testPost.Upvotes, c.expectedOutput)
		}

	}
}

//Test removing downvotes from post
func TestRemoveDownvote(t *testing.T) {
	cases := []struct {
		name           string
		expectedOutput int
	}{
		{
			name:           "removing downvote",
			expectedOutput: 0,
		},
	}
	testPost := &Post{
		Downvotes: 1,
	}
	for _, c := range cases {
		testPost.RemoveDownvote()
		if testPost.Downvotes != c.expectedOutput {
			t.Errorf("%s: got %d downvotes but expected %d downvotes", c.name, testPost.Downvotes, c.expectedOutput)
		}

	}
}
