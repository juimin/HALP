import { 
   LOGOUT, SETTOKEN, SETUSER, SAVEPASSWORD,
   GETBOARD, SETBOARD,
   ADD_SUBSCRIPTIONS,
   REMOVE_SUBSCRIPTIONS,
   GET_SUBSCRIPTIONS,
   MAKE_POST,
   PICTURE_SUCCESS

} from './ActionTypes.js';

// These are action creators, you can use these to create actions
export const setTokenAction = (token) => ({type:SETTOKEN, payload: token})
export const logoutAction = () => ({type:LOGOUT, payload: null})
export const setUserAction = (usr) => ({type: SETUSER, payload: usr})
export const savePasswordAction = (pass) => ({type: SAVEPASSWORD, payload: pass})

// get board
export const getBoard = (board) => ({type: GETBOARD, payload: board})
export const setActiveBoard = (board) => ({type: SETBOARD, payload: board})

// Subscriptions
export const addSubscription = sub => ({type: ADD_SUBSCRIPTIONS, payload: sub})
export const removeSubscription = sub => ({type: REMOVE_SUBSCRIPTIONS, payload: sub})
export const getSubscriptions = () => ({type: GET_SUBSCRIPTIONS})

//POSTS
export const makePost = (post) => ({type: MAKE_POST, payload: post})
export const setPictureSuccess = (pizza) => ({type: PICTURE_SUCCESS, payload: pizza})

export default {
   // Home Screen actions
   logoutAction, setTokenAction, setUserAction, savePasswordAction,

   // BOARD
   setActiveBoard,
   // Subscriptions
   addSubscription, removeSubscription, getSubscriptions,
   // Posts
   makePost,
   setPictureSuccess
}