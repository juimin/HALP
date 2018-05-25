// This file defines functions for HTTP Requests
import { API_URL } from '../Constants/Constants';


// A function that checks to see if the session has expired
export const sessionExpired = (token) => {
   return fetch(API_URL + "/users/me", {
      method: "GET",
      headers: {
         'Accept': 'application/json',
         'Content-Type': 'application/json',
         'Authorization': token
      }
   }).then(response => {
      return (response.status != 202)
   }).catch(err => {
      console.log(err)
      return true
   })
}