import { 
	ADDPOSTS 
} from '../ActionTypes';

const initialState = {
   posts: [],
   activePost: null,
}

export default (state=initialState, action) => {
   switch(action.type) {
      case ADDPOSTS:
         return Object.assign({}, state, {
            posts: action.payload
         })
      default:
         return state
   }
}