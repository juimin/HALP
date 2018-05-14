import { AUTH, DOWNLOAD, LOGIN } from './ActionTypes.js';

export const authAction = stuff => ({type: AUTH, payload: stuff});
export const downloadAction = stuff => ({type: DOWNLOAD, payload: stuff})
export const loginAction = toggle => ({type:LOGIN, setLogin: toggle})

export default {authAction, downloadAction, loginAction}