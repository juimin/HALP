// This should be the root of all application components
// Everything runs under a stack navigation nexted from here

// Import required react components
import React, { Component } from 'react';
import { Button, View, Text, TouchableWithoutFeedback, Alert, Image, ScrollView, Picker} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'
import ActionSheet from 'react-native-actionsheet'
import ImageResizer from 'react-native-image-resizer';
import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements'


// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
import { API_URL } from '../../Constants/Constants';

// import HALP compnents
import CanvasTest from '../Canvas/CanvasTest'
import HideableView from '../Helper/HideableView';
import HomeScreen from '../Home/GuestHome';
import HomeNav from '../Navigation/HomeNav';


// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

const mapStateToProps = (state) => {
	return {
    user: state.AuthReducer.user
	}
}

//user.favorites should contain list of boards
//right now it's hardcoded but replace this later


var mongoObjectId = () => {
  var timestamp = (new Date().getTime() / 1000 | 0).toString(16);
  return timestamp + 'xxxxxxxxxxxxxxxx'.replace(/[x]/g, () => {
      return (Math.random() * 16 | 0).toString(16);
  }).toLowerCase();
};

const testboard = '5b077a0d0324ac00012a223a';

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
// Define the new Post
class NewPost extends Component {
    constructor(props) {
        super(props);
        this.state = 
        {
          source: require('../../Images/davint.png'),
          isHidden: true,
          imageURL: '',
          board: '',
          title: '',
          caption: '',
        };

        this.errors = {
          board: false,
          title: false,
          caption: false,
        };

        this.errorMessages = {
          board: '',
          title: '',
          caption: '',
        };
    }

    resetForm = () => {
      this.state = 
        {
          source: require('../../Images/davint.png'),
          isHidden: true,
          imageURL: '',
          board: '',
          title: '',
          caption: '',
        };

        this.errors = {
          board: false,
          title: false,
          caption: false,
          imageURL: false,
        };

        this.errorMessages = {
          board: '',
          title: '',
          caption: '',
          imageURL: '',
        };
    
    }

    validate = () => {
      var errored = false;
      if (this.state.title.length == 0) {
        this.errors.title = true;
        this.errorMessages.title = "Title cannot be left blank";
        errored = true;
      } else {
        this.errors.title = false;
        this.errorMessages.title = '';
      }

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

      if (this.state.board.length == 0) {
        this.errors.board = true;
        this.errorMessages.board = "Choose a board to post on";
        errored = true;
      } else {
        this.errors.board = false;
        this.errorMessages.board = false;
      }

      return !errored
    }

    //submit
    submit = () => {
      console.log('current user', this.props.user)
      console.log('title', this.state.title);
      console.log('imageURL', this.state.imageURL);
      console.log('caption', this.state.caption);
      console.log('board', this.state.board);
      if (this.validate()) {
        var x = fetch(API_URL + 'posts/new', {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            "title": this.state.title,
            "image_url": this.state.imageURL,
            "caption": this.state.caption,
            "author_id": this.props.user.id,
            "board_id": this.state.board
            })
        }).then(response => {
          if (response.status == 201) {
            //Alert.alert('Post Success', 'Successfully submitted new post')
            this.props.navigation.goBack(null);
          } else {
              console.log('response', response)
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

    takePiture = () => {
      ImagePicker.launchCamera(options, (response) => {
        console.log('Response = ', response);
      
        if (response.didCancel) {
          console.log('User cancelled image picker');
        }
        else if (response.error) {
          console.log('ImagePicker Error: ', response.error);
        }
        else {
          console.log('success');
          console.log(response.height, response.width)
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
    
    showActionSheet = () => {
        this.ActionSheet.show()
    }

    // ActionSheet demo
    // <Button color={Theme.colors.primaryColor} 
    //             onPress={this.showActionSheet} title="shiet"/>
    //             <ActionSheet
    //             ref={o => this.ActionSheet = o}
    //             title={'make your choice:'}
    //             options={['davin', 'derek', 'Cancel']}
    //             cancelButtonIndex={2}
    //             destructiveButtonIndex={1}
    //             onPress={(index) => { /* do something */ }}
    //             />

    //stupid way to send data back from child component without redux
    //pass returnData while navigating
    returnData(url, externalurl) {
      Image.getSize(url, (width, height) => console.log(width, height));
      this.setState({source: {uri: url}, isHidden: false, imageURL: externalurl});
      console.log("success:", this.state.imageURL);
      }
    
    //image size is 1080 * 1536 so /8 to fit photo to display in the form
    render() {
      if (this.props.user == null) {
        return (
          <View style={Styles.home}>
            <Text>You must be logged in to make a post</Text>
          </View>
        )
      }
      //for now just using other forms' style
      //also need to generate list of picker.items for user's boards
      return(
         <ScrollView>
            <Picker
              selectedValue={this.state.board}
              style={{ height: 50, width: 200 }}
              mode='dropdown'
              onValueChange={(itemValue, itemIndex) => this.setState({board: itemValue})}>
              <Picker.Item label="Choose board" value='' />
              <Picker.Item label="Testing" value={testboard} />
              <Picker.Item label="Cooking" value='5b01b3017912ed0001434678' />
            </Picker>
            <FormValidationMessage>{this.errorMessages.board}</FormValidationMessage>
            <FormLabel>Title *</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.state.title = text}}/>
            <FormValidationMessage>{this.errorMessages.title}</FormValidationMessage>
            <HideableView hide={this.state.isHidden}><Image style={{width: 135, height: 192}} source = {this.state.source} /></HideableView>
            <Button color={Theme.colors.primaryColor}
                onPress={this.takePiture} 
            title = "Upload Image"/>
            <FormLabel>Caption</FormLabel>
            <FormInput onChangeText={(text) => {this.state.caption = text}}/>
            <FormValidationMessage>{this.errorMessages.caption}</FormValidationMessage>
            <Button color={Theme.colors.primaryColor} title="Post" onPress={this.submit}></Button>
        </ScrollView>
      );
   }
}

export default connect(mapStateToProps)(NewPost)