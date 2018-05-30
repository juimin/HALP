import { 
	ADDBOARD, SETACTIVEBOARD
} from '../ActionTypes';

const initialState = {
   boards: {},
   activeBoard: null,
}

export default (state=initialState, action) => {
   switch(action.type) {
      case ADDBOARD:
         var b = state.boards
         b[action.payload.key] = action.payload.board
         return Object.assign({}, state, {
            boards: b
         })
      case SETACTIVEBOARD:
         return Object.assign({}, state, {
            activeBoard: action.payload
         })
      default:
         return state
   }
}