import { 
ADDPOSTS,
SETACTIVEPOST
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
      case SETACTIVEPOST:
         return Object.assign({}, state, {
            activePost: action.payload
         })
      default:
         return state
   }
}