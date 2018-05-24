import { LOGIN, LOGOUT, SETUSER } from '../ActionTypes';
import { loginAction, logoutAction, setUserAction } from '../Actions';

// Import the api URL
import { API_URL } from '../../Constants/Constants';

// Set the api endpoint handler
const endpoint = "";

// Set the initial state of this part
const initialState = {
	user: null,
	authToken: null
};

// This file exports the reducer used for login tasks
export default (state=initialState, action) => {
	switch (action.type) {
		case LOGIN:
			// In the case that we are logging in we need to perform the login and the do the setting of the state's auth token
			return Object.assign({}, state, {
				authToken: action.payload,
				loggedIn: true
			})
		case LOGOUT:
			return Object.assign({}, state, {
				loggedIn: false,
				authToken: ""
			})
		case SETUSER:
			return Object.assign({}, state, {
				user: action.payload
			})
		default:
			return state;
	}
};