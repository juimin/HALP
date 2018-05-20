import { 
   AUTH, DOWNLOAD, LOGIN, SEARCH, UPDATE,
   ADD_SUBSCRIPTIONS,
   REMOVE_SUBSCRIPTIONS,
   GET_SUBSCRIPTIONS

} from './ActionTypes.js';

// These are action creators, you can use these to create actions
export const authAction = stuff => ({type: AUTH, payload: stuff});
export const downloadAction = stuff => ({type: DOWNLOAD, payload: stuff})
export const loginAction = toggle => ({type:LOGIN, payload: toggle})

// SEARCH PAGE
export const searchPosts = searchTerm => ({type: SEARCH, payload: searchTerm})
export const addSubscription = sub => ({type: ADD_SUBSCRIPTIONS, payload: sub})
export const removeSubscription = sub => ({type: REMOVE_SUBSCRIPTIONS, payload: sub})
export const getSubscriptions = something => ({type: GET_SUBSCRIPTIONS, payload: null})

export default {
   // Home Screen actions
   authAction, downloadAction, loginAction, searchPosts,
   // Search Actions
   searchPosts, addSubscription, removeSubscription, getSubscriptions
}