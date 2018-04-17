package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/JuiMin/HALP/servers/gateway/models/posts"
)

//NewPostHandler handles requests related to Posts
//POST /posts/new
func (cr *ContextReceiver) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		errorMessage := ""
		canProceed := true
		status := http.StatusAccepted
		if r.Body == nil {
			errorMessage = "Error: Could not decode request body"
			status = http.StatusBadRequest
			canProceed = false
		}
		if canProceed {
			newPost := &posts.NewPost{}
			err := json.NewDecoder(r.Body).Decode(newPost)
			if err != nil {
				errorMessage = "Error: could not decode request body"
				status = http.StatusBadRequest
				canProceed = false
			}
			if canProceed {
				//Validate the Post
				err = newPost.Validate()
				if err != nil {
					errorMessage = "Error: Could not validate new post"
					status = http.StatusConflict
					canProceed = false
				}
			}
			//don't need to check for duplicate posts as with users
			if canProceed {
				thisPost, err := cr.PostStore.Insert(newPost)
				if err != nil {
					fmt.Printf("Could not insert the post: %v", err)
					status = http.StatusInternalServerError
				}
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(&thisPost)
			}
		}
		if !canProceed {
			fmt.Printf(errorMessage + "\n")
			w.WriteHeader(status)
			w.Write([]byte(errorMessage))
		}
	}
}

//UpdatePostHandler should update a post with a PostUpdate
//PATCH /posts/update?id=<id>
func (cr *ContextReceiver) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	//post id delivered as part of post/put/patch request?
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		//status := http.StatusAccepted
		canProceed := true
		//get post by id?
		updates := &posts.PostUpdate{}
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			canProceed = false
		}

		//not possible to do posts/update/<id>
		//must we do posts/update?id=<id>?
		//or can we have client send postid?
		id := r.URL.Query().Get("id")
		//string to bson.ObjectId?
		if len(id) == 0 || !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
			canProceed = false
		}
		if canProceed {
			postid := bson.ObjectId(id)

			//get post from db
			post, err := cr.PostStore.GetByID(postid)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				canProceed = false
			}

			//postnew := &posts.Post{post}
			if canProceed {
				err := json.NewDecoder(r.Body).Decode(updates)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					canProceed = false
				}
				//apply updates
				if canProceed {
					//to the struct
					err = post.ApplyUpdates(updates)
					if err != nil {
						fmt.Printf("Could not update post state\n")
						canProceed = false
					}
					//to the user store
					err = cr.PostStore.PostUpdate(postid, updates)
					if err != nil {
						fmt.Printf("Could not update post in database\n")
						canProceed = false
					}
				}
				if canProceed {
					w.WriteHeader(http.StatusAccepted)
					// Output the updated user back
					json.NewEncoder(w).Encode(post)
				}

			}
		}

		if !canProceed {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

//GetPostHandler returns the content of a single post
//GET /posts/get?id=<id>
func (cr *ContextReceiver) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		id := r.URL.Query().Get("id")
		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
		}
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
		}
		postid := bson.ObjectId(id)
		post, err := cr.PostStore.GetByID(postid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(post)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
