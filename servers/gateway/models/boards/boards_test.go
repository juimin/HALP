package boards

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestChangeSubscriberCount(t *testing.T) {

	cases := []struct {
		name           string
		board          *Board
		input          bool
		expectedOutput int
	}{
		{
			name: "Increment Subscriber",
			board: &Board{
				ID:          bson.NewObjectId(),
				Title:       "jfiwejiwjfw",
				Description: "Testing some cool stuff in go for boards and stuff",
				Image:       "http://imgur.reddit.government",
				Subscribers: 9000,
				Posts:       9000,
			},
			input:          true,
			expectedOutput: 9001,
		},
		{
			name: "Decrement Subscriber",
			board: &Board{
				ID:          bson.NewObjectId(),
				Title:       "jsdjfosjfosd",
				Description: "fjweofjwe",
				Image:       "http://hello.gov",
				Subscribers: 9000,
				Posts:       9000,
			},
			input:          false,
			expectedOutput: 8999,
		},
		{
			name: "Negative Subscribers",
			board: &Board{
				ID:          bson.NewObjectId(),
				Title:       "jsdjfosjfosd",
				Description: "fjweofjwe",
				Image:       "http://hello.gov",
				Subscribers: -2,
				Posts:       9000,
			},
			input:          false,
			expectedOutput: 0,
		},
	}

	for _, c := range cases {
		c.board.ChangeSubscriberCount(c.input)
		if c.board.Subscribers != c.expectedOutput {
			t.Errorf("%s: got %d but expected %d", c.name, c.board.Subscribers, c.expectedOutput)
		}
	}
}

func TestChangePostCount(t *testing.T) {
	cases := []struct {
		name           string
		board          *Board
		input          bool
		expectedOutput int
	}{
		{
			name: "Increment Post",
			board: &Board{
				ID:          bson.NewObjectId(),
				Title:       "jfiwejiwjfw",
				Description: "Testing some cool stuff in go for boards and stuff",
				Image:       "http://imgur.reddit.government",
				Subscribers: 9000,
				Posts:       9000,
			},
			input:          true,
			expectedOutput: 9001,
		},
		{
			name: "Increment Subscriber",
			board: &Board{
				ID:          bson.NewObjectId(),
				Title:       "jfiwejiwjfw",
				Description: "Testing some cool stuff in go for boards and stuff",
				Image:       "http://imgur.reddit.government",
				Subscribers: 9000,
				Posts:       9000,
			},
			input:          false,
			expectedOutput: 8999,
		},
		{
			name: "Increment Subscriber",
			board: &Board{
				ID:          bson.NewObjectId(),
				Title:       "jfiwejiwjfw",
				Description: "Testing some cool stuff in go for boards and stuff",
				Image:       "http://imgur.reddit.government",
				Subscribers: 9000,
				Posts:       -1,
			},
			input:          true,
			expectedOutput: 0,
		},
	}

	for _, c := range cases {
		c.board.ChangePostCount(c.input)
		if c.board.Posts != c.expectedOutput {
			t.Errorf("%s: got %d but expected %d", c.name, c.board.Posts, c.expectedOutput)
		}
	}
}

// Test the to board function
func TestToBoard(t *testing.T) {

	cases := []struct {
		name          string
		nb            *NewBoard
		expectedError error
	}{
		{
			name: "Valid Test",
			nb: &NewBoard{
				Title:       "ttest",
				Description: "This i s a test board",
				Image:       "https://google.com",
			},
			expectedError: nil,
		},
		{
			name: "No Title Test",
			nb: &NewBoard{
				Title:       "",
				Description: "This i s a test board",
				Image:       "https://google.com",
			},
			expectedError: fmt.Errorf("Please enter a title"),
		},
		{
			name: "No Desc Test",
			nb: &NewBoard{
				Title:       "Hold On",
				Description: "",
				Image:       "https://google.com",
			},
			expectedError: fmt.Errorf("Please enter a description"),
		},
		{
			name: "No Image Test",
			nb: &NewBoard{
				Title:       "Poop",
				Description: "lol",
				Image:       "",
			},
			expectedError: fmt.Errorf("Please use either an image or a caption"),
		},
		{
			name: "No Image URL Test",
			nb: &NewBoard{
				Title:       "Poop",
				Description: "lol",
				Image:       "??DSF?ASD",
			},
			expectedError: fmt.Errorf("Invalid Photo URL"),
		},
	}

	for _, c := range cases {
		b, err := c.nb.ToBoard()
		if err != nil {
			if c.expectedError != nil {
				if err.Error() != c.expectedError.Error() {
					t.Errorf("%s Failed: Expected %v but got %v", c.name, c.expectedError, err)
				}
			} else {
				t.Errorf("%s Failed: Expected success but got %v", c.name, err)
			}
		} else {
			if c.expectedError != nil {
				t.Errorf("%s Failed: Suceeded but should have failed with %v", c.name, c.expectedError)
			} else {
				if b.Title != c.nb.Title {
					t.Errorf("%s Failed: There should be matching data for %s and %s", c.name, b.Title, c.nb.Title)
				}
			}
		}
	}
}
