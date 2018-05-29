import { MAKE_POST, PICTURE_SUCCESS } from '../ActionTypes';

// Set the initial state of this part
const initialState = {
	success: false
};

// This file exports the reducer used for login tasks
export default (state=initialState, action) => {
	switch (action.type) {
		case PICTURE_SUCCESS:
            return Object.assign({}, state, {
                success: action.payload
            })
		default:
			return state;
	}
};