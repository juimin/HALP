package boards

import (
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
