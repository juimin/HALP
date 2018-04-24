package indexes

import "testing"
import "gopkg.in/mgo.v2/bson"
import "fmt"

//TODO: implement automated tests for your trie data structure
func TestTrieInsert(t *testing.T) {
	cases := []struct {
		testName       string
		inputKey       string
		inputVal       bson.ObjectId
		expectedOutput error
	}{
		{
			testName:       "Test Input",
			inputKey:       "Derek",
			inputVal:       bson.NewObjectId(),
			expectedOutput: nil,
		},
		{
			testName:       "Test Empty",
			inputKey:       "",
			inputVal:       bson.NewObjectId(),
			expectedOutput: fmt.Errorf("Error, empty insert"),
		},
		{
			testName:       "Test Empty",
			inputKey:       "d",
			inputVal:       bson.NewObjectId(),
			expectedOutput: nil,
		},
	}

	// Create a test Trie
	myTrie := &TrieNode{
		values: make(map[string]bson.ObjectId),
		next:   make(map[rune]*TrieNode),
	}

	for _, c := range cases {
		err := myTrie.Insert(c.inputKey, c.inputVal, 0)
		if err != c.expectedOutput {
			if err != nil && c.expectedOutput != nil {
				if err.Error() != c.expectedOutput.Error() {
					t.Errorf("%s: got %v but expected %v", c.testName, err, c.expectedOutput)
				}
			} else {
				t.Errorf("%s: got %v but expected %v", c.testName, err, c.expectedOutput)
			}
		}
	}
}

func TestManyTrieInsert(t *testing.T) {
	cases := []struct {
		testName       string
		inputKeys      []string
		expectedOutput error
	}{
		{
			testName: "Test a lot of Input",
			inputKeys: []string{
				"Derek",
				"Gold",
				"Deer",
				"Golf",
				"Potato",
				"Dead",
				"Go",
				"Git",
				"Get",
				"Police",
				"Pray",
				"Prince",
			},
			expectedOutput: nil,
		},
	}

	// Create a test Trie
	myTrie := &TrieNode{
		values: make(map[string]bson.ObjectId),
		next:   make(map[rune]*TrieNode),
	}

	for _, c := range cases {
		for _, v := range c.inputKeys {
			err := myTrie.Insert(v, bson.NewObjectId(), 0)
			if err != c.expectedOutput {
				t.Errorf("%s: got %v but expected %v", c.testName, err, c.expectedOutput)
			}
		}
	}
}

func TestTrieRemove(t *testing.T) {
	cases := []struct {
		testName       string
		inputKey       string
		inputVal       bson.ObjectId
		expectedOutput error
	}{
		{
			testName:       "Test Input",
			inputKey:       "Derek",
			inputVal:       bson.NewObjectId(),
			expectedOutput: nil,
		},
		{
			testName:       "Test Empty",
			inputKey:       "",
			inputVal:       bson.NewObjectId(),
			expectedOutput: fmt.Errorf("This key doesn't exist"),
		},

		{
			testName:       "Test Invalid Key",
			inputKey:       "Derekf",
			inputVal:       bson.NewObjectId(),
			expectedOutput: fmt.Errorf("This key doesn't exist"),
		},
	}

	// Create a test Trie
	myTrie := &TrieNode{
		values: make(map[string]bson.ObjectId),
		next:   make(map[rune]*TrieNode),
	}

	inputKeys :=
		[]string{
			"Derek",
			"Gold",
			"Deer",
			"Golf",
			"Potato",
			"Dead",
			"Go",
			"Git",
			"Get",
			"Police",
			"Pray",
			"Prince",
		}

	for _, s := range inputKeys {
		err := myTrie.Insert(s, bson.NewObjectId(), 0)
		if err != nil {
			t.Errorf("Error inserting")
		}
	}

	// Seed values
	for _, c := range cases {
		err := myTrie.Remove(c.inputKey)
		if err != c.expectedOutput {
			if err != nil && c.expectedOutput != nil {
				if err.Error() != c.expectedOutput.Error() {
					t.Errorf("%s: got %v but expected %v", c.testName, err, c.expectedOutput)
				}
			} else {
				t.Errorf("%s: got %v but expected %v", c.testName, err, c.expectedOutput)
			}
		}
	}
}

func TestTrieNValues(t *testing.T) {
	cases := []struct {
		testName       string
		inputKeys      []string
		search         string
		n              int
		expectedOutput error
	}{
		{
			testName: "Test Input",
			inputKeys: []string{
				"Derek",
				"Gold",
				"Deer",
				"Golf",
				"Potato",
				"Dead",
				"Go",
				"Git",
				"Get",
				"Police",
				"Pray",
				"Prince",
			},
			search:         "d",
			n:              4,
			expectedOutput: nil,
		},
		{
			testName: "Test Two",
			inputKeys: []string{
				"Derek",
				"Gold",
				"Deer",
				"Golf",
				"Potato",
				"Dead",
				"Go",
				"Git",
				"Get",
				"Police",
				"Pray",
				"Prince",
			},
			search:         "p",
			n:              2,
			expectedOutput: nil,
		},
		{
			testName: "Test Two",
			inputKeys: []string{
				"Derek",
				"Gold",
				"Deer",
				"Golf",
				"Potato",
				"Dead",
				"Go",
				"Git",
				"Get",
				"Police",
				"Pray",
				"Prince",
			},
			search:         "",
			n:              5,
			expectedOutput: fmt.Errorf("Empty search"),
		},
		{
			testName: "Test Two",
			inputKeys: []string{
				"Derek",
				"Gold",
				"Deer",
				"Golf",
				"Potato",
				"Dead",
				"Go",
				"Git",
				"Get",
				"Police",
				"Pray",
				"Prince",
			},
			search:         "sandwich",
			n:              5,
			expectedOutput: fmt.Errorf("This search string is not contained"),
		},
		{
			testName: "Test Two",
			inputKeys: []string{
				"Derek",
				"Gold",
				"God",
				"Golf",
				"Goal",
				"Good",
				"Goo",
				"Go",
				"Git",
				"Get",
				"Gore",
				"Pray",
				"Prince",
			},
			search:         "go",
			n:              8,
			expectedOutput: nil,
		},
	}

	for _, c := range cases {
		// Create a test Trie
		myTrie := &TrieNode{
			values: make(map[string]bson.ObjectId),
			next:   make(map[rune]*TrieNode),
		}

		// Insert all the values
		for _, v := range c.inputKeys {
			err := myTrie.Insert(v, bson.NewObjectId(), 0)
			if err != nil {
				fmt.Println(v)
			}
		}

		// Test
		result, err := myTrie.NValues(c.search, c.n, 0)
		if err != c.expectedOutput {
			if err != nil && c.expectedOutput != nil {
				if err.Error() != c.expectedOutput.Error() {
					t.Errorf("%s: got %v but expected %v", c.testName, err, c.expectedOutput)
				}
			} else {
				t.Errorf("%s: got %v but expected %v", c.testName, err, c.expectedOutput)
			}
		}
		fmt.Println()
		fmt.Printf("%v", result)
		fmt.Println()
	}
}
