// Import all our reducers
import loginReducer from './Reducers/Login';
import boardReducer from './Reducers/Boards';
import AuthReducer from './Reducers/Auth';
import SearchReducer from './Reducers/Search';

// Import the reducer combination
import { combineReducers } from 'redux'

// We can use the combine reducers to allow separation in our reducers while coding
// and then bring them all together as the single app reducer here
export default combineReducers({
   loginReducer,
   boardReducer,
   AuthReducer
})