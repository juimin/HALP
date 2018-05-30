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
         var p = state.posts
         b.push(action.payload.post)
         return Object.assign({}, state, {
            boards: b
         })
      default:
         return state
   }
}