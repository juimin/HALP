import React, { Component } from 'react';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

import { Button, Text, Icon } from 'native-base'

export default class SubscribeButton extends Component {
    constructor(props) {
		super(props)
		this.state = {
              user: this.props.user,
              board: this.props.board,
              subscribed: this.props.subbed
		}
    }

    // isSubscribed = () => {
    //     return this.props.user.favorites.includes(this.props.board.id);
    // }

    

    updateUser = (add) => {
        var x = fetch('https://staging.halp.derekwang.net/favorites', {
            method: 'PATCH',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': this.props.authToken, 
            },
            body: JSON.stringify({
                "adding": add,
                "updateID": this.props.board.id
            })
        }).then(response => {
            if (response.status == 200) {
                console.log("user success")
                return response.json()
            } else {
				console.log(response.status, response.statusText);
			}
        }).then(data => {
            //console.log(data);
        }).catch(err => {
            console.log("error updating user", err)
        });
    }

    updateBoard = (add) => {
        console.log(this.props.board.id)
        var y = fetch('https://staging.halp.derekwang.net/boards/updatesubscriber?id=' + this.props.board.id, {
            method: 'PATCH',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "temp": add,
            })
        }).then(response => {
            if (response.status == 200) {
                console.log("board success")
                return response.json()
            } else {
				console.log(response.status, response.statusText);
			}
        }).then(data => {
            //console.log(data);
        }).catch(err => {
            console.log("error updating board", err)
        });
    }

    subscribe = () => {
        if (!this.state.subscribed) {
            console.log('subscribing');
            console.log('before', this.props.user.favorites);
            if (!this.props.user.favorites.includes(this.props.board.id)) {
                this.props.user.favorites.push(this.props.board.id);
            }
            console.log('after', this.props.user.favorites);
            this.state.subscribed = true;
            this.props.returnData(true);
            this.updateUser(true);
            //this.updateBoard(true);
        }
    }

    unsubscribe = () => {
        if (this.state.subscribed) {
            console.log('unsubscribing');
            console.log('before', this.props.user.favorites);
            if (this.props.user.favorites.includes(this.props.board.id)) {
                this.props.user.favorites.splice(this.props.user.favorites.indexOf(this.props.board.id));
            }
            console.log('after', this.props.user.favorites);
            this.state.subscribed = false;
            this.props.returnData(false);
            this.updateUser(false);
            //this.updateBoard(false);
        }
    }
    
    render() {
        if (!this.props.user) {
            return(
                <Button disabled iconLeft style={Styles.subscribeButton}><Icon type="MaterialIcons" name="add" /><Text>Subscribe</Text></Button>
            )
        }
        if (this.state.subscribed) {
            return (
                <Button iconLeft style={Styles.subscribeButtonColor} onPress={() => this.unsubscribe()}><Icon type="MaterialIcons" name="remove" /><Text>Unsubscribe</Text></Button>
            )
        }
        
        return (
            <Button iconLeft style={Styles.subscribeButtonColor} onPress={() => this.subscribe()}><Icon type="MaterialIcons" name="add" /><Text>Subscribe</Text></Button>
        )
    }

}