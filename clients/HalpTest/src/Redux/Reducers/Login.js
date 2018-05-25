import { LOGIN } from '../ActionTypes';
import { loginAction } from '../Actions';

// This file exports the reducer used for login tasks
export default (state = false, action) => {
   switch (action.type) {
      case LOGIN:
         return state.loggedin === action.payload ? state.loggedin : action.payload
      default:
      return state;
   }
 };