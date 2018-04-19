package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"

	"gopkg.in/mgo.v2/bson"
)

// BoardsAllHandler returns all the boards that exist
func (cr *ContextReceiver) BoardsAllHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	canProceed := true
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		canProceed = false
	}
	if canProceed {
		allBoards, err := cr.BoardStore.GetAllBoards()
		if err != nil {
			status = http.StatusInternalServerError
			canProceed = false
		}
		if canProceed {
			// Encodes all boards into the response
			json.NewEncoder(w).Encode(allBoards)
			w.WriteHeader(status)
		}
	}
}

// SingleBoardHandler gets a board ID and returns the corresponding board
func (cr *ContextReceiver) SingleBoardHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	canProceed := true
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		canProceed = false
	}
	if canProceed {
		boardID := r.URL.Query().Get("id")
		board, err := cr.BoardStore.GetByID(bson.ObjectIdHex(boardID))
		if err != nil {
			status = http.StatusInternalServerError
			canProceed = false
		}
		if canProceed {
			// Encodes the board into the response
			json.NewEncoder(w).Encode(board)
			w.WriteHeader(status)
		}
	}
}

// UpdatePostHandler gets a board id and updates the given board post count
func (cr *ContextReceiver) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
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
				changeToStore := &boards.UpdatePost{}
				changeToStore.Post = board.Subscribers
				cr.BoardStore.UpdatePostCount(boardID, changeToStore)
				status = http.StatusAccepted
			}
		}
	}
	w.WriteHeader(status)
}

// UpdateSubscriberHandler gets a board id and updates the given board post count
func (cr *ContextReceiver) UpdateSubscriberHandler(w http.ResponseWriter, r *http.Request) {
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
				changeToStore.Sub = board.Subscribers
				cr.BoardStore.UpdateSubscriberCount(boardID, changeToStore)
				status = http.StatusAccepted
			}
		}
	}
	w.WriteHeader(status)
}
