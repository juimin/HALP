import { SETTOKEN, LOGOUT, SETUSER, SAVEPASSWORD } from '../ActionTypes';

// Set the initial state of this part
const initialState = {
	user: null,
	authToken: "",
	password: ""
};

// This file exports the reducer used for login tasks
export default (state=initialState, action) => {
	switch (action.type) {
		case SETTOKEN:
			// In the case that we are logging in we need to perform the login and the do the setting of the state's auth token
			return Object.assign({}, state, {
				authToken: action.payload
			})
		case LOGOUT:
			return Object.assign({}, state, {
				authToken: "",
				user: null,
				password: ""
			})
		case SAVEPASSWORD:
			return Object.assign({}, state, {
				password: action.payload	
			})
		case SETUSER:
			return Object.assign({}, state, {
				user: action.payload
			})
		default:
			return state;
	}
};