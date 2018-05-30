import { 
   LOGOUT, SETTOKEN, SETUSER, SAVEPASSWORD,
   ADDBOARD, SETACTIVEBOARD,
   ADDPOSTS,
   ADD_SUBSCRIPTIONS,
   REMOVE_SUBSCRIPTIONS,
   GET_SUBSCRIPTIONS,
   FETCH_BOARDS_BEGIN, 
   FETCH_BOARDS_SUCCESS,
   FETCH_BOARDS_FAILIURE,
   MAKE_POST,
   PICTURE_SUCCESS

} from './ActionTypes.js';

// These are action creators, you can use these to create actions
export const setTokenAction = (token) => ({type:SETTOKEN, payload: token})
export const logoutAction = () => ({type:LOGOUT, payload: null})
export const setUserAction = (usr) => ({type: SETUSER, payload: usr})
export const savePasswordAction = (pass) => ({type: SAVEPASSWORD, payload: pass})

// get board
export const addBoard = (board) => ({type: ADDBOARD, payload: board})
export const setActiveBoard = (board) => ({type: SETACTIVEBOARD, payload: board})

export const addPosts = (post) => ({type: ADDPOSTS, payload: post})

// Subscriptions
export const addSubscription = sub => ({type: ADD_SUBSCRIPTIONS, payload: sub})
export const removeSubscription = sub => ({type: REMOVE_SUBSCRIPTIONS, payload: sub})
export const getSubscriptions = () => ({type: GET_SUBSCRIPTIONS})

// HomeScreen Boards 
export const fetchBoardsBeginAction = () => ({type: FETCH_BOARDS_BEGIN})
export const fetchBoardsSuccessAction = (boards) => ({type: FETCH_BOARDS_SUCCESS, payload: boards })
export const fetchBoardsFailiureAction = (error) => ({type: FETCH_BOARDS_FAILIURE, payload: error })

//POSTS
export const makePost = (post) => ({type: MAKE_POST, payload: post})
export const setPictureSuccess = (pizza) => ({type: PICTURE_SUCCESS, payload: pizza})

export default {
   // Home Screen actions
   logoutAction, setTokenAction, setUserAction, savePasswordAction,
   // BOARD
   addBoard, setActiveBoard,
   // pOSTS
   addPosts,
   // Subscriptions
   addSubscription, removeSubscription, getSubscriptions,
   // HomeScreen Boards
   fetchBoardsBeginAction, fetchBoardsSuccessAction, fetchBoardsFailiureAction,
   // Posts
   makePost,
   setPictureSuccess
}
