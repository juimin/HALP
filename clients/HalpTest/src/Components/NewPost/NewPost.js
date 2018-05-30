// This should be the root of all application components
// Everything runs under a stack navigation nexted from here

// Import required react components
import React, { Component } from 'react';
import { View, TouchableWithoutFeedback, Alert, Image, ScrollView } from 'react-native';
import { Picker, Button, Text } from 'native-base';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'
import ActionSheet from 'react-native-actionsheet'
import ImageResizer from 'react-native-image-resizer';
import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements';


// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
import { API_URL } from '../../Constants/Constants';

// import HALP compnents
import CanvasTest from '../Canvas/CanvasTest'
import HideableView from '../Helper/HideableView';
import HomeScreen from '../Home/HomeScreen';
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

const testboard = '5b077a0d0324ac00012a223a';
const testfavs = [testboard, '5b01b3017912ed0001434678']

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
          source: null,
          isHidden: true,
          imageURL: '',
          favorites: [],
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

    componentWillMount() {
      if (this.props.user) {
        this.loadFavorites();
        
      }
      //this.loadFavorites();
    }

    resetForm = () => {
      this.state = 
        {
          source: null,
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

    //load boards based on user subscriptions
    loadFavorites = () => {
      var x = fetch(API_URL + 'boards', {
        method: 'GET',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        }
      }).then(response => {
        if (response.status == 200) {
          //array containing objects with board name/id
          return response.json()
        // } else {
        //   Alert.alert("error getting user's favorites");
        //   this.setState({favorites: []});
        }
      }).then(data => {
          var favs = [];
          if (this.props.user.favorites.length == 0) {
            favs.push({
              key:   testboard,
              value: 'Your favorite boards will show up here!'
            })
          } else {
            this.props.user.favorites.forEach((element) => {
              //each element should be a boardID
              data.forEach((item) => {
                //each element should be a boardID
                if (item.id == element) {
                  favs.push({
                    key:   item.id,
                    value: item.title
                  });
                }
              });
            });
          }
          this.setState({favorites: favs});
      }).catch(err => {
        Alert.alert("Error getting user's favorites");
        this.setState({favorites: []});
      })
    }

    //submit
    submit = () => {
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
            this.props.navigation.state.params.returnData2(this.state.board)
            this.props.navigation.goBack(null);
          } else {
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
      if (!this.props.user) {
        return (
          <View style={Styles.home}>
            <Text>You must be logged in to make a post</Text>
          </View>
        )
      }

      let favItems = this.state.favorites.map( (f) => {
        return <Picker.Item key={f['key']} value={f['key']} label={f['value']} />
      });

      // react-native button example - replaced by native-base buttons
      // <Button color={Theme.colors.primaryColor}
      //           onPress={this.takePicture} 
      //       title = "Upload Image"/>

      return(
         <ScrollView>
            <Picker
              selectedValue={this.state.board}
              style={Styles.boardPicker}
              mode='dropdown'
              onValueChange={ (fav) => ( this.setState({board:fav}) ) } >
              <Picker.Item label="Choose board" value='' />
              {favItems}
            </Picker>
            <FormValidationMessage>{this.errorMessages.board}</FormValidationMessage>
            <FormLabel>Title *</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.state.title = text}}/>
            <FormValidationMessage>{this.errorMessages.title}</FormValidationMessage>
            <HideableView style={Styles.newPostView} hide={this.state.isHidden}><Image style={{width: 135, height: 192}} source = {this.state.source} /></HideableView>
            <View style={Styles.newPostView}>
            <Button rounded style={Styles.buttonTheme}
                onPress={this.takePicture} 
            ><Text>Upload Image</Text></Button></View>
            <FormLabel>Caption</FormLabel>
            <FormInput onChangeText={(text) => {this.state.caption = text}}/>
            <FormValidationMessage>{this.errorMessages.caption}</FormValidationMessage>
            <View style={Styles.newPostView}><Button rounded style={Styles.buttonTheme} onPress={this.submit}><Text>Post</Text></Button></View>
        </ScrollView>
      );
   }
}

export default connect(mapStateToProps)(NewPost)