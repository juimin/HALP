package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"

	"gopkg.in/mgo.v2/bson"
)

// BoardsAllHandler returns all the boards that exist
func (cr *ContextReceiver) BoardsAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		allBoards, err := cr.BoardStore.GetAllBoards()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			// Encodes all boards into the response
			json.NewEncoder(w).Encode(allBoards)
			w.WriteHeader(http.StatusOK)
		}
	}
}

// SingleBoardHandler gets a board ID and returns the corresponding board
func (cr *ContextReceiver) SingleBoardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		boardID := r.URL.Query().Get("id")
		if bson.IsObjectIdHex(boardID) {
			board, err := cr.BoardStore.GetByID(bson.ObjectIdHex(boardID))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				json.NewEncoder(w).Encode(board)
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

// UpdatePostCountHandler gets a board id and updates the given board post count
func (cr *ContextReceiver) UpdatePostCountHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	canProceed := true
	if r.Method != "PATCH" {
		status = http.StatusMethodNotAllowed
		canProceed = false
	}
	if r.Body == nil {
		status = http.StatusBadRequest
		canProceed = false
	}
	if canProceed {
		boardID := bson.ObjectIdHex(r.URL.Query().Get("id"))
		board, err := cr.BoardStore.GetByID(boardID)
		if err != nil {
			status = http.StatusInternalServerError
			canProceed = false
		}
		if canProceed {
			// object to store new post information
			update := &boards.TempBoolStore{}
			err = json.NewDecoder(r.Body).Decode(update)
			if err != nil {
				status = http.StatusNotAcceptable
				canProceed = false
			}
			if canProceed {
				board.ChangePostCount(update.TempSubPost)
				//new value of subs
				changeToStore := &boards.UpdatePost{}
				changeToStore.Posts = board.Posts
				cr.BoardStore.UpdatePostCount(boardID, changeToStore)
				status = http.StatusOK
			}
		}
	}
	w.WriteHeader(status)
}

// UpdateSubscriberCountHandler gets a board id and updates the given board post count
func (cr *ContextReceiver) UpdateSubscriberCountHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	canProceed := true
	if r.Method != "PATCH" {
		status = http.StatusMethodNotAllowed
		canProceed = false
	}
	if r.Body == nil {
		status = http.StatusBadRequest
		canProceed = false
	}
	if canProceed {
		boardID := bson.ObjectIdHex(r.URL.Query().Get("id"))
		board, err := cr.BoardStore.GetByID(boardID)
		if err != nil {
			status = http.StatusInternalServerError
			canProceed = false
		}
		if canProceed {
			// object to store new post information
			update := &boards.TempBoolStore{}
			err = json.NewDecoder(r.Body).Decode(update)
			if err != nil {
				status = http.StatusBadRequest
				canProceed = false
			}
			if canProceed {
				board.ChangeSubscriberCount(update.TempSubPost)
				//new value of subs
				changeToStore := &boards.UpdateSubscriber{}
				changeToStore.Subscribers = board.Subscribers
				cr.BoardStore.UpdateSubscriberCount(boardID, changeToStore)
				status = http.StatusOK
			}
		}
	}
	w.WriteHeader(status)
}
