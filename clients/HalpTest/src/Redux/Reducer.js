// Import all our reducers
import AuthReducer from './Reducers/Auth';
import SearchReducer from './Reducers/Search';
import BoardsReducer from './Reducers/Boards'
import BoardReducer from './Reducers/Board';

// Import the reducer combination
import { combineReducers } from 'redux'

// We can use the combine reducers to allow separation in our reducers while coding
// and then bring them all together as the single app reducer here
export default combineReducers({
   AuthReducer,
   BoardsReducer,
   BoardReducer
})