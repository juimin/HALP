// This file defines functions for HTTP Requests
import { API_URL } from '../Constants/Constants';
import { Alert } from 'react-native';

// A function that checks to see if the session has expired
const sessionExpired = (token) => {
   return fetch(API_URL + "/users/me", {
      method: "GET",
      headers: {
         'Accept': 'application/json',
         'Content-Type': 'application/json',
         'Authorization': token
      }
   }).then(response => {
      return response.status
   }).catch(err => {
      console.log(err)
   })
}

// Get a single board from the server based on object id
const getBoard = (item) => {
   return fetch(API_URL + "/boards/single?id=" + item, {
      method: "GET",
      headers: {
         'Accept': 'application/json',
         'Content-Type': 'application/json',
     }
   }).then(response => {
      if (response.status == 200) {
         return response.json()
      } else {
         return null
      }
   }).catch(err => {
      console.log(err)
   })
}

// Search for the boards through the API
const searchBoard = (type, term, auth) => {
   return fetch(API_URL + "/search?type=" + type + "&search=" + term, {
      method: "GET",
      headers: {
         'Accept': 'application/json',
         'Content-Type': 'application/json',
         'authorization': auth
     }
   }).then(response => {
      if (response.status == 200) {
         return response.json()
      } else {
         return null
      }
   }).catch(err => {
      console.log(err)
   })
}

// renew the session, gives back the auth and the user
const renewSession = (credentials) => {
   fetch(API_URL + "/sessions", {
      method: "POST",
      headers: {
         'Accept': 'application/json',
         'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials)
   }).then(response => {
      // Save token and password for later use
      return response.headers.get('authorization')
   }).catch(err => {
      console.log(err)
   })
}

export default {
   sessionExpired,
   getBoard,
   searchBoard,
   renewSession
}