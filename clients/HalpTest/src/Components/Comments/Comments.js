import React, { Component } from 'react'
import { Container, Header, Body, Title, Left, Right, Button, Text, Icon } from 'native-base';

import { connect } from "react-redux";

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

import { setUserAction } from '../../Redux/Actions';

const mapStateToProps = state => {
   return {
      activePost: state.PostReducer.activePost,
      user: state.AuthReducer.user,
      authToken: state.AuthReducer.authToken,
		password: state.AuthReducer.password,
   }
}

const mapDispatchToProps = dispatch => {
   return {
      setUser: usr => { dispatch(setUserAction(usr)) }
   }
}

class Comments extends Component {
   constructor(props) {
      super(props)
      this.state = {}
      this.postComment = this.postComment.bind(this)
      console.log(this.props.activePost)
   }

   postComment() {

   }

   render() {
      return(
         <Container>
            <Header style={{
               backgroundColor: Theme.colors.primaryBackgroundColor
            }}>
               <Left>
                  <Button transparent onPress={() => this.props.navigation.goBack()}>
                     <Icon name='arrow-back' style={{color:"gray"}} />
                  </Button>
               </Left>
               <Body>
                  <Title style={{
                     color: Theme.colors.primaryColor,
                     alignSelf: "flex-start",
                     fontWeight: "bold"
                  }}>Reply</Title>
               </Body>
               <Right>
                  <Button transparent onPress={this.postComment}>
                     <Text style={{color:"gray"}}>POST</Text>
                  </Button>
               </Right>
            </Header>
         </Container>
      );
   }
}

export default connect(mapStateToProps, mapDispatchToProps)(Comments)