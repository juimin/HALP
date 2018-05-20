import { SEARCH, GET_SUBSCRIPTIONS, ADD_SUBSCRIPTIONS, REMOVE_SUBSCRIPTIONS } from '../ActionTypes';
import { loginAction } from '../Actions';

import API_URL from '../../Constants/Constants';

// This file exports the reducer used for login tasks
export default (state={}, action) => {
   switch (action.type) {
      case SEARCH:
         // Here we should perform the action that updates the state (return a new state)
         // HTTP REQUEST
         var request = new Request(API_URL); 
         // Perform the fetch
         var headerOptions = { method: 'GET',
               headers: myHeaders,
               mode: 'cors',
               cache: 'default' };
         return fetch(request, headerOptions).then(
            (response) => {
               return response.blob()
            }
         ).then(
            (response) => {
               return Object.assign({}, state, {
                  searchPostResults: response
               })
            }
         )
      case GET_SUBSCRIPTIONS:
         return Object.assign({}, state, {
            boardSubscriptions: state.subscriptions.add(action.payload)
         })
      case ADD_SUBSCRIPTIONS:
         return state
      case REMOVE_SUBSCRIPTIONS:
         return state
      default:
         return state;
   }
 };