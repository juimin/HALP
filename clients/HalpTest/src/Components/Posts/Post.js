import React, { Component } from 'react'
import { Container, Card, CardItem, Text, Subtitle, H2, View, Body, Button, Icon } from 'native-base';
import { TouchableWithoutFeedback } from 'react-native';

import { API_URL } from '../../Constants/Constants';
// Import react-redux connect 
import { connect } from "react-redux";
import { Image } from 'react-native';

const mapStateToProps = state => {
   return {
      activePost: state.PostReducer.activePost
   }
}

class Post extends Component {
   constructor(props) {
      super(props)
      this.state = {
         board: null
      }
      // Get board related
      fetch( API_URL + "boards/single?id="+this.props.activePost.board_id,{
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

   render() {

      let post = this.props.activePost;
      let photo = post.image_url.length != 0 ? {uri: post.image_url} : {uri: 'https://facebook.github.io/react-native/docs/assets/favicon.png'};
      let boardName = this.state.board == null ? "Missing Board Name" : this.state.board.title

      return(
         <Container>
            <Card>
               <CardItem header>
                  <View flexDirection={"column"}>
                     <H2>{this.props.activePost.title}</H2>
                     <Text note>{boardName} · {this.hoursSince(post.time_created)}</Text>
                  </View>
               </CardItem>
               <TouchableWithoutFeedback onPress={() => this.props.navigation.navigate('Image', {photosrc: photo})}>
               <CardItem cardBody>
                  <Image source={photo} style={{height: 200, width: null, flex: 1}}/>
               </CardItem>
               </TouchableWithoutFeedback>
               <CardItem>
                  <Body>
                     <Text>{this.props.activePost.caption}</Text>
                  </Body>
               </CardItem>
               <CardItem>
                  <View flexDirection={"row"}>
                     <View flexDirection={"column"}>
                        <Text note>{String(post.upvotes - post.downvotes) + " Points"}</Text>
                        <Text note>{String(post.comments.length) + " Comments"}</Text>
                     </View>
                     <Button transparent>
                        <Icon active name="arrow-round-up"/>
                     </Button>
                     <Button transparent>
                        <Icon active name="arrow-round-down"/>
                     </Button>
                     <Button transparent>
                        <Icon active name="undo"/>
                     </Button>
                     <Button transparent>
                        <Icon active name="bookmark"/>
                     </Button>
                  </View>
               </CardItem>
            </Card>

         </Container>
      );rr
   }
}

export default connect(mapStateToProps)(Post)