package indexes

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//TODO: implement a trie data structure that stores
//keys of type string and values of type bson.ObjectId

// MAX_VALUES contains the max number of values that can be associated at 	this
// node before it needs to be converted into a trie node
const maxValues = 5

// TrieNode is a struct detailing nodes of a hybrid trie
type TrieNode struct {
	values map[string]bson.ObjectId
	next   map[rune]*TrieNode
}

// NewSearchTrie returns the root node of as earch trie
func NewSearchTrie() *TrieNode {
	return &TrieNode{
		values: make(map[string]bson.ObjectId),
		next:   make(map[rune]*TrieNode),
	}
}

// Insert takes in a key as a string and a pointer to the bson objectId
// value and inserts the value into the trie
// If an error is observed, the error is returned
func (tn *TrieNode) Insert(key string, value bson.ObjectId, index int) error {
	// If there are no next nodes, then we are at a leaf node
	// If the key is empty, then you can't really do anything
	if len(key) == 0 {
		return fmt.Errorf("Error, empty insert")
	}
	// Runes for the given key
	key = strings.ToLower(key)
	keyRunes := []rune(key)
	// Check if we are at a leaf node
	// If we are, then we can append
	if len(tn.next) == 0 || index == len(key)-1 {
		// Since this is a hybrid trie, we need to check if the values
		// At the leaf node exceed the limit
		tn.values[key[index:len(keyRunes)]] = value
		return tn.PutIn(key, value, index)
	}
	// This isn't a leaf node but your token does exist
	if _, exists := tn.next[keyRunes[index]]; exists == false {
		// NOt a leaf node and the token doesn't exist
		tn.next[keyRunes[index]] = &TrieNode{
			values: make(map[string]bson.ObjectId),
			next:   make(map[rune]*TrieNode),
		}
	}
	return tn.next[keyRunes[index]].Insert(key, value, index+1)
}

// PutIn Appends to values
func (tn *TrieNode) PutIn(key string, value bson.ObjectId, index int) error {
	// Check if the current node's values are full
	// Append the value since we have space
	// Append the value and the key to the node
	if maxValues < len(tn.values) {
		// If the size is equivalent, we must separate the nodes
		// For eachKeyStore key value pair, check if there is a node for it
		// and if there is, append the key value, if not make the node
		// and then append
		for k, v := range tn.values {
			// If the key is bigger than 1 character, move it
			if len(k) > 1 {
				first := []rune(k)
				if tn.next[first[0]] == nil {
					// The node doesn't exist so append a new one keyed to the new rune
					tn.next[first[0]] = &TrieNode{
						values: make(map[string]bson.ObjectId),
						next:   make(map[rune]*TrieNode),
					}
				}
				tn.next[first[0]].values[string(first[1:len(first)])] = v
				// Remove the value from the current node
				delete(tn.values, k)
			}
		}
	}
	return nil
}

// Remove the key and value from the trie
func (tn *TrieNode) Remove(key string) error {
	key = strings.ToLower(key)
	keyRunes := []rune(key)
	// This is a leaf node
	// If the value exists
	_, exists := tn.values[key]
	if exists {
		delete(tn.values, key)
		return nil
		// Value doesn't exist, return error
	}
	// This is not a leaf node
	if len(keyRunes) > 0 && len(tn.next) > 0 {
		return tn.next[keyRunes[0]].Remove(string(keyRunes[1:len(keyRunes)]))
	}
	return fmt.Errorf("This key doesn't exist")
}

// NValues Returns the first N values of the search term
func (tn *TrieNode) NValues(key string, n int, index int) ([]*bson.ObjectId, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Empty search")
	}
	key = strings.ToLower(key)
	if len(key) == index || len(tn.next) == 0 {
		// We have searched as far as the search term goes
		// From  here we can do the depth first search
		values := tn.SearchN(key, n, index)
		if len(values) > 0 {
			return values, nil
		}
	}
	// If this is not a leaf node
	if len(tn.next) != 0 {
		// Traverse to the next node
		runes := []rune(key)
		// Check if the next node exists, if it does then we can search more
		if tn.next[runes[index]] != nil {
			return tn.next[runes[index]].NValues(key, n, index+1)
		}
	}
	// If it is a leaf node, our search term exceeds the trie's knowledge
	// So we can't return anything yet

	// Node doesn't exist so there  is nothing that matches the search terms
	return nil, fmt.Errorf("This search string is not contained")
}

// SearchN returns a depth first search of n potential values
func (tn *TrieNode) SearchN(key string, n int, index int) []*bson.ObjectId {
	// If we get to here, then we don't have enough values yet
	output := []*bson.ObjectId{}
	// Get all the values at this level
	for s := range tn.values {
		ptr := tn.values[s]
		match := true

		if len(key)-1 > len(s) {
			match = false
		} else {
			for i := index; i < len(key); i++ {
				if key[i] != s[i-index] {
					match = false
				}
			}
		}

		fmt.Printf("%s -- %s index: %d match: %v\n", key, s, index, match)
		if match {
			seen := false
			for _, outputid := range output {
				if (*outputid) == ptr {
					seen = true
				}
			}
			if seen == false {
				output = append(output, &ptr)
			}
		}
		if len(output) == n {
			return output
		}
	}
	// Depth first search
	if len(tn.next) > 0 {
		for r := range tn.next {
			values := tn.next[r].SearchN(key, n-len(output), index)
			for _, id := range values {
				output = append(output, id)
			}
			if len(output) == n {
				return output
			}
		}
	}
	// If we get to here, then output has nothing to look for
	// so return it
	return output
}
