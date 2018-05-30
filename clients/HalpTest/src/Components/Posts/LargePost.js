import React, { Component } from 'react';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

import { API_URL } from '../../Constants/Constants' 

import { Image } from 'react-native';
import {
	View,
	Container,
	Right,
	Body,
	Left,
	Title,
	Subtitle,
	Header,
	Button,
	Icon,
	Thumbnail,
	FooterTab,
	Content,
	Text,
    ActionSheet,
    Card,
    CardItem,
} from 'native-base';

// Import react-redux connect 
import { connect } from "react-redux";
import { TouchableHighlight, TouchableWithoutFeedback } from 'react-native';
import { setActivePost, setUserAction } from '../../Redux/Actions';

const mapStateToProps = state => {
   return {
      user: state.AuthReducer.user,
      authToken: state.AuthReducer.authToken,
		password: state.AuthReducer.password,
   }
}

const mapDispatchToProps = dispatch => {
   return {
      setActivePost: post => dispatch(setActivePost(post)),
      setUser: usr => { dispatch(setUserAction(usr)) }
   }
}

class LargePost extends Component {
   constructor(props) {
      //props are primarily post objects
      super(props)
		this.state = {
         board: null
      }

      this.toggleBookmark = this.toggleBookmark.bind(this)

      // Get board related
      fetch( API_URL + "boards/single?id="+this.props.post.board_id,{
         method: "GET"
      }).then(response => {
         if (response.status == 200) {
            return response.json()
         } else {
            return null
         }
      }).then(board => {
         if (board != null) {
            this.setState({
               board: board
            })
         }
      })
    }

   hoursSince = (time) => {
      var original = new Date(time);
      var current = new Date();
      //console.log(current, original);
      //get difference in hours
      var hours = Math.round(Math.abs(current - original) / (60*60*1000));
      if (hours < 24) {
         return String(hours) + "h";
      }
      //over one day
      var days = Math.round(Math.abs(current - original) / (60*60*1000*24));
      if (days < 365) {
         return String(hours) + "d";
      } else {
         return String(days / 365) + "y " + String(days % 365) + "d"
      }
   }

   toggleBookmark() {
      // Check if we are logged in.
      if (this.props.user != null) {
         fetch(API_URL + "bookmarks", {
            method: "PATCH",
            headers: {
               'authorization': this.props.authToken,
               'Accept': 'application/json',
               'Content-Type': 'application/json',
            },
            body: JSON.stringify({
               "adding": !this.props.user.bookmarks.includes(this.props.post.id),
               "updateID": this.props.post.id
            })
         }).then(response => {
            if (response.status == 200) {
               return response.json()
            } else {
               return null
            }
         }).then(usr => {
            this.props.setUser(usr)
         }).catch(err => {
            console.log(err)
         })
      }
   }

    render() {
        let post = this.props.post;
        let photo = post.image_url.length != 0 ? {uri: post.image_url} : {uri: 'https://halp-staging.nyc3.digitaloceanspaces.com/Logo-09.png'};
        let boardName = this.state.board == null ? "Missing Board Name" : this.state.board.title
        let bookmarked = (this.props.user != null ) ? (this.props.user.bookmarks.includes(post.id)) ? true : false : false
        return(
            <Container key={post.id} style={Styles.largePost}>
                <Content style={{padding:"1%"}}>
                    <Card>
                        <CardItem  button onPress={() => {
                            this.props.setActivePost(post)
                            this.props.navigation.navigate("Post")
                        }}>
                            <Left>
                                <Body>
                                    <Text>{post.title}</Text>
                                    <Text note>{boardName} · {this.hoursSince(post.time_created)}</Text>
                                </Body>
                            </Left>
                        </CardItem>
                        <CardItem cardBody  button onPress={() => {
                            this.props.setActivePost(post);
                            this.props.navigation.navigate("Post");
                        }}>
                            <Image source={photo} style={{height: 200, width: null, flex: 1}}/>
                        </CardItem>
                        <CardItem button onPress={() => {
                            this.props.setActivePost(post);
                            this.props.navigation.navigate("Post")}
                        }>
                            <Left>
                                <View style={{flexDirection: "column"}}> 
                                    <Text note>{String(post.upvotes - post.downvotes) + " Points"}</Text>
                                    <Text note>{String(post.comments.length) + " Comments"}</Text>
                                </View>
                            </Left>
                            <Right>
                                <View style={{flexDirection: "row"}}>
                                    <Button transparent onPress={this.toggleBookmark}>
                                       <Icon active name="bookmark" style={
                                          {color: bookmarked ? "green": "gray"}
                                       } />
                                    </Button>
                                    <Button transparent> 
                                        <Icon name="more" />
                                    </Button>
                                </View>
                            </Right>
                        </CardItem>
                    </Card>
                </Content>
            </Container>
        )
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(LargePost)