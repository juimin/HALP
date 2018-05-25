import { GETBOARD, SETBOARD } from '../ActionTypes';

const initialState = {
   boards: {},
   activeBoard: null,
}

export default (state=initialState, action) => {
   switch(action.type) {
      case GETBOARD:
         var b = state.boards
         b[action.payload.key] = action.payload.board
         return Object.assign({}, state, {
            boards: b
         })
      case SETBOARD:
         return Object.assign({}, state, {
            activeBoard: action.payload
         })
      default:
         return state
   }
}