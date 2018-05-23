import { 
   LOGOUT, LOGIN, SETUSER,
   SEARCH,
   ADD_SUBSCRIPTIONS,
   REMOVE_SUBSCRIPTIONS,
   GET_SUBSCRIPTIONS

} from './ActionTypes.js';

// These are action creators, you can use these to create actions
export const loginAction = (token) => ({type:LOGIN, payload: token})
export const logoutAction = () => ({type:LOGOUT, payload: null})
export const setUserAction = (usr) => ({type: SETUSER, payload: usr})

// SEARCH PAGE
export const searchPosts = term => ({type: SEARCH, searchTerm: term})

// Subscriptions
export const addSubscription = sub => ({type: ADD_SUBSCRIPTIONS, payload: sub})
export const removeSubscription = sub => ({type: REMOVE_SUBSCRIPTIONS, payload: sub})
export const getSubscriptions = () => ({type: GET_SUBSCRIPTIONS})

export default {
   // Home Screen actions
   logoutAction, loginAction,
   // Search Actions
   searchPosts,
   // Subscriptions
   addSubscription, removeSubscription, getSubscriptions
}