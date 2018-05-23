import { LOGIN, LOGOUT } from '../ActionTypes';
import { loginAction, logoutAction } from '../Actions';

// Import the api URL
import { API_URL } from '../../Constants/Constants';

// Set the api endpoint handler
const endpoint = "";

// Set the initial state of this part
const initialState = {
	user: null,
	authToken: null,
	loggedIn: false
};

// This file exports the reducer used for login tasks
export default (state=initialState, action) => {
	switch (action.type) {
		case LOGIN:
			// In the case that we are logging in we need to perform the login and the do the setting of the state's auth token
			return Object.assign({}, state, {
				loggedIn: true
			})
		case LOGOUT:
			return state
		default:
			return state;
	}
};