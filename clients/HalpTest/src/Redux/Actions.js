import { AUTH, DOWNLOAD, LOGIN, FETCH_BOARDS_BEGIN, FETCH_BOARDS_SUCCESS, FETCH_BOARDS_FAILIURE } from './ActionTypes.js';

export const authAction = stuff => ({type: AUTH, payload: stuff});
export const downloadAction = stuff => ({type: DOWNLOAD, payload: stuff})
export const loginAction = toggle => ({type:LOGIN, setLogin: toggle})

// Boards actions 
export const fetchBoardsBeginAction = () => ({type: FETCH_BOARDS_BEGIN})
export const fetchBoardsSuccessAction = boards => ({type: FETCH_BOARDS_SUCCESS, payload: { boards }});
export const fetchBoardsFailiureAction = error => ({type: FETCH_BOARDS_FAILIURE, payload: { error }});

export default { authAction, downloadAction, loginAction, fetchBoardsBeginAction, fetchBoardsSuccessAction, fetchBoardsFailiureAction }