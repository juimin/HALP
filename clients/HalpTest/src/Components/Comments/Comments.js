import React, { Component } from 'react'
import { Container, Header, Body, Title, Left, Right, Button, Text, Icon, Content } from 'native-base';
import { View, TouchableWithoutFeedback, Alert, Image, ScrollView } from 'react-native';

import ImageResizer from 'react-native-image-resizer';
import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements';

import { connect } from "react-redux";

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
import { API_URL } from '../../Constants/Constants';

import { setUserAction } from '../../Redux/Actions';

// import HALP compnents
import CanvasTest from '../Canvas/CanvasTest'
import HideableView from '../Helper/HideableView';

var ImagePicker = require('react-native-image-picker');
var options = {
    title: null,
    cameraType: 'back',
    mediaType: 'photo',
    rotation: 0,
    // storageOptions: {
    //   skipBackup: true,
    //   path: 'images'
    // }
    //maxWidth: 480,
    //maxHeight: 480,
}; 

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
      this.state = 
        {
          source: null,
          isHidden: true,
          imageURL: '',
          favorites: [],
          caption: '',
        };

        this.errors = {
          caption: false,
        };

        this.errorMessages = {
          caption: '',
        };
      this.postComment = this.postComment.bind(this)
      console.log(this.props.activePost)
   }

   resetForm = () => {
    this.state = 
      {
        source: null,
        isHidden: true,
        imageURL: '',
        caption: '',
      };

      this.errors = {
        caption: false,
        imageURL: false,
      };

      this.errorMessages = {
        caption: '',
        imageURL: '',
      };
  
  }

  validate = () => {
    var errored = false;

    if (this.state.caption.length == 0 && this.state.imageURL.length == 0) {
      this.errors.caption = true;
      this.errors.imageURL = true;
      this.errorMessages.caption = "Must have either image or caption";
      this.errorMessages.imageURL = "Must have either image or caption";
      errored = true;
    } else {
      this.errors.caption = false;
      this.errors.imageURL = false;
      this.errorMessages.caption = '';
      this.errorMessages.imageURL = '';
    }

    return !errored
  }

    postComment = () => {
        if (this.validate()){
            //console.log("post id:", this.props.navigation.state.params.post.id);
            var x = fetch(API_URL + 'comments/new', {
                method: 'POST',
                headers: {
                  'Accept': 'application/json',
                  'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                  "image_url": this.state.imageURL,
                  "content": this.state.caption,
                  "author_id": this.props.user.id,
                  "post_id": this.props.activePost.id,
                  })
              }).then(response => {
                if (response.status == 200) {
                  //Alert.alert('Post Success', 'Successfully submitted new post')
                  //this.props.navigation.state.params.returnData2(this.props.activePost.id)
                  this.props.navigation.goBack();
                } else {
                    console.log(response);
                    Alert.alert(
                        'Post Error',
                        'Please try again',
                        [
                         {text: 'OK', onPress: () => console.log('ok')},
                        ]
                     )
                 }
              }).catch(err => {
                 Alert.alert(
                    'Error getting response from server',
                    err,
                    [
                      {text: 'OK', onPress: () => console.log('ok')},
                    ]
                  )
              })
        } else {
            this.setState(this.state)
         }
    }

    returnData(url, externalurl) {
        Image.getSize(url, (width, height) => console.log(width, height));
        this.setState({source: {uri: url}, isHidden: false, imageURL: externalurl});
        console.log("success:", this.state.imageURL);
    }

    takePicture = () => {
        ImagePicker.launchCamera(options, (response) => {
          if (response.didCancel) {
            console.log('User cancelled image picker');
          }
          else if (response.error) {
            console.log('ImagePicker Error: ', response.error);
          }
          else {
            console.log('success! image HxW:', response.height, response.width)
            const { error, uri, originalRotation } = response
  
            if ( uri && !error ) {
              let rotation = 0
  
              if ( originalRotation === 90 ) {
                rotation = 90
              } else if ( originalRotation === 270 ) {
                rotation = -90
              }
  
              ImageResizer.createResizedImage( uri, 480, 640, "JPEG", 100, rotation )
                .then( ( { uri } ) => {
                  let source = {uri: response.uri };
                  this.setState({
                    source: source
                  });
                  this.props.navigation.navigate('Canvas', {source: source, returnData: this.returnData.bind(this)});
                } ).catch( err => {
                  console.log( err )
                  return Alert.alert( 'Unable to resize the photo', 'Please try again!' )
                } )
            }
           }
        });
      }


    usePicture = () => {
        console.log(this.props.navigation.state.params.source);
        if (this.props.navigation.state.params.source) {
          this.props.navigation.navigate('Canvas', {source: this.props.navigation.state.params.source, returnData: this.returnData.bind(this)});
        }
      }
    
   render() {

    if (!this.props.user) {
        return (
            <Container>
            <Header style={{
               backgroundColor: Theme.colors.primaryBackgroundColor
            }}>
               <Left>
                  <Button transparent onPress={() => this.props.navigation.goBack()}>
                     <Icon name='arrow-back' style={{color:"gray"}} />
                  </Button>
               </Left>
            </Header>
          <View style={Styles.home}>
            <Text>You must be logged in to make a post</Text>
          </View>
          </Container>
        )
    }
    
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
                     <Text style={!this.state.isHidden||this.state.caption.length != 0 ? {color: Theme.colors.primaryColor} : {color:"gray"}}>POST</Text>
                  </Button>
               </Right>
            </Header>
                <ScrollView>
                    <HideableView style={Styles.newPostView} hide={this.state.isHidden}><Image style={{width: 135, height: 192}} source = {this.state.source} /></HideableView>
                    <View style={Styles.newPostView}>
                        <Button rounded style={Styles.buttonTheme}
                            onPress={this.usePicture} 
                        ><Text>Use Existing Image</Text></Button>
                        <Button rounded style={Styles.buttonTheme}
                            onPress={this.takePicture} 
                        ><Text>Upload Image</Text></Button>
                    </View>
                    <FormLabel>Caption</FormLabel>
                    <FormInput onChangeText={(text) => {this.state.caption = text}}/>
                    <FormValidationMessage>{this.errorMessages.caption}</FormValidationMessage>
                    {/* <View style={Styles.newPostView}><Button rounded style={Styles.buttonTheme} onPress={this.submit}><Text>Post</Text></Button></View> */}
                </ScrollView>
         </Container>
      );
   }
}

export default connect(mapStateToProps, mapDispatchToProps)(Comments)