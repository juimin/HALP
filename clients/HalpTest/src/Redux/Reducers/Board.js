import { GETBOARD } from '../ActionTypes';

const initialState = {
   boards: {}
}

export default (state=initialState, action) => {
   switch(action.type) {
      case GETBOARD:
         var b = state.boards
         b[action.payload.key] = action.payload.board
         return Object.assign({}, state, {
            boards: b
         })
      default:
         return state
   }
}