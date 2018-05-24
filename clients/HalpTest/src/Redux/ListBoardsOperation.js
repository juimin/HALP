// ListBoardsOperation.js
import { fetchBoardsBeginAction, fetchBoardsSuccessAction, fetchBoardsFailiureAction } from './Actions.js';

// Board operation for getting all boards for the home screen
// TODO: Not necessary but we can place all operations here or just divide all operations seperately?

export function fetchBoards() {
    return dispatch => {
        dispatch(fetchBoardsBeginAction());
        return fetch("https://staging.halp.derekwang.net/boards")
            .then(handleErrors)
            .then(res => res.json())
            .then(json => {
                dispatch(fetchBoardsSuccessAction(json.boards))
                return json.boards;
            })
            .catch(error => dispatch(fetchBoardsFailiureAction(error)))
    };
}

// Hnadling errors because fetch does not handle errors properly 
export function handleErrors() {
    if (!response.ok) {
        throw Error(response.statusText);
    }
    return response;
}

