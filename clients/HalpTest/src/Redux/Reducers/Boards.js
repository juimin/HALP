// Boards reducer
import {
    FETCH_BOARDS_BEGIN,
    FETCH_BOARDS_SUCCESS,
    FETCH_BOARDS_FAILIURE
} from '../ActionTypes.js';

const initialState = {
    items: [],
    loading: false,
    error: null
};

export default function boardsReducer(state = initialState, action) {
    switch(action.type) {
      case FETCH_BOARDS_BEGIN:
        // Mark the state as "loading" so we can show a spinner or something
        return {
          ...state,
          loading: true,
          error: null
        };
  
      case FETCH_BOARDS_SUCCESS:
        // All done: set loading "false".
        return {
          ...state,
          loading: false,
          items: action.payload.boards
        };
  
      case FETCH_BOARDS_FAILIURE:
        // The request failed, but it did stop, so set loading to "false".
        // Save the error, and we can display it somewhere
        // Since it failed, we don't have items to display anymore, so set it empty.
        return {
          ...state,
          loading: false,
          error: action.payload.error,
          items: []
        };
  
      default:
        // ALWAYS have a default case in a reducer
        return state;
    }
  }