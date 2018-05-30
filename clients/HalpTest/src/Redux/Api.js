import { fetchBoardsSuccessAction, fetchBoardsBeginAction, fetchBoardsFailiureAction } from "./Actions";
import React, { Component } from 'react';

// const mapStateToProps = state => state;

export default class API extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            items: [],
            loading: false,
            apiCall: this.fetchBoards()
        };
    }
    
    fetchBoards() {
        dispatch => {
            dispatch(fetchBoardsBeginAction());
            fetch("https://staging.halp.derekwang.net/boards")
            .then(handleErrors)
            .then(res => res.json())
            .then(json => {
                dispatch(fetchBoardsSuccessAction(json));
                this.setState({ items: json,
                                loading: false
                })
                return json;
            })
            .catch(error => dispatch(fetchBoardsFailiureAction(error)));
        }
        return (fetchBoardsSuccessAction(this.state.items))
    };

    // Handle HTTP errors since fetch won't.
    handleErrors(response) {
    if (!response.ok) {
        throw Error(response.statusText);
    }
    return response;
    }
}

//for testing diddly squat
export function hardFetchPosts() {
    const hardPosts = [
        {
            id: 1,
            upvotes: 200,
            downvotes: 3,
            comments: 54,
            title: "How do I eat?",
            author_id: "alexis",
            board_id: "Food",
            time_created: 22, 
            image_url: 'https://food.fnr.sndimg.com/content/dam/images/food/fullset/2018/6/0/FN_snapchat_coachella_wingman%20.jpeg.rend.hgtvcom.616.462.suffix/1523633513292.jpeg'
        },
        {
            id: 2,
            upvotes: 45,
            downvotes: 200,
            comments: 23,
            title: "Why is my car on fire?",
            author_id: "teeler",
            board_id: "Auto",
            time_created: 35,
            image_url: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSBxOhnyuaLSuWIHCG4fBqKrfwJEOZrMxAh-fTwCp1W_m-Eq5P7Uw"
        },
        {
            id: 3,
            upvotes: 56,
            downvotes: 12,
            comments: 75,
            title: "How hot is the hot potato?", 
            author_id: "jumbotron", 
            board_id: "Potatoes", 
            time_created: 40,
            image_url: "https://scontent.fsea1-1.fna.fbcdn.net/v/t1.0-1/p200x200/14291794_1171930419548193_8185068610699574818_n.jpg?_nc_cat=0&oh=392bea7d0e4a1dc83b5bd4b45f230ec2&oe=5BC3F8E9&efg=eyJhZG1pc3Npb25fY29udHJvbCI6MCwidXBsb2FkZXJfaWQiOiI3OTAxMjc0Mjc3Mjg0OTYifQ%3D%3D"
        },
        {
            id: 4,
            upvotes: 809,
            downvotes: 7,
            comments: 1,
            title: "Where is Carmen San Diego?", 
            author_id: "rickDsanchez",
            board_id: "Where Ya At?",
            time_created: 67,
            image_url: "https://images.rapgenius.com/78844aada7bacd1807df3c54a34462da.372x450x1.jpg"
        },
        {
            id: 5,
            upvotes: 34,
            downvotes: 5,
            comments: 58,
            title: "Where Waldo at?",
            author_id: "mortymcfly",
            board_id: "Where Ya At?",
            time_created: 80,
            image_url: "https://static1.squarespace.com/static/56438e3fe4b0c2d5ac1d4d26/5643945ae4b0eadf5c6537e9/5664d039e4b058c26c239306/1525966385398/maps_troy.jpg?format=1500w"
        }
    ];

    // return dispatch => {
    //     dispatch(fetchPostsSuccessAction(hardPosts));
    //     return hardPosts;
    // };
    return fetchPostsSuccessAction(hardPosts);
}