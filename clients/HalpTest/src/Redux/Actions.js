import { 
   LOGOUT, SETTOKEN, SETUSER, SAVEPASSWORD,
   GETBOARD, SETBOARD,
   SEARCH,
   ADD_SUBSCRIPTIONS,
   REMOVE_SUBSCRIPTIONS,
   GET_SUBSCRIPTIONS

} from './ActionTypes.js';

// These are action creators, you can use these to create actions
export const setTokenAction = (token) => ({type:SETTOKEN, payload: token})
export const logoutAction = () => ({type:LOGOUT, payload: null})
export const setUserAction = (usr) => ({type: SETUSER, payload: usr})
export const savePasswordAction = (pass) => ({type: SAVEPASSWORD, payload: pass})

// get board
export const getBoard = (board) => ({type: GETBOARD, payload: board})
export const setActiveBoard = (board) => ({type: SETBOARD, payload: board})

// SEARCH PAGE
export const searchPosts = term => ({type: SEARCH, searchTerm: term})

// Subscriptions
export const addSubscription = sub => ({type: ADD_SUBSCRIPTIONS, payload: sub})
export const removeSubscription = sub => ({type: REMOVE_SUBSCRIPTIONS, payload: sub})
export const getSubscriptions = () => ({type: GET_SUBSCRIPTIONS})

export default {
   // Home Screen actions
   logoutAction, setTokenAction, setUserAction, savePasswordAction,

   // BOARD
   setActiveBoard,
   // Search Actions
   searchPosts,
   // Subscriptions
   addSubscription, removeSubscription, getSubscriptions
}