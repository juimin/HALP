// This should be the root of all application components
// Everything runs under a stack navigation nexted from here

// Import required react components
import React, { Component } from 'react';
import { Button, View, Text, TouchableWithoutFeedback, Alert, Image, ScrollView} from 'react-native';
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

// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

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
export default class NewPost extends Component {
    constructor(props) {
        super(props);
        this.state = 
        {
          source: require('../../Images/davint.png'),
          isHidden: true,
          imageURL: null,
        };
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
    // static navigationOptions = {
    //   headerTitle: 'New Post',
    //   headerLeft: (
    //     // Add the Icon for canceling the Home
    //     <Icon name="close" onPress={() => navigation.goBack(null)}/>
    //   ),
    //   headerRight: (
    //     // Add the Icon for canceling the Home
    //     <Button title="POST" onPress={() => navigation.goBack(null)}></Button>
    //   ),
    // }

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
    
    render() {
      //once auth works, only render page if user is logged in
      if (this.state.imageURL != null) {
        return(<Image style={{height: 320, width: 150}} source = {this.state.source} />)
      }
      return(
         <View style={Styles.newPostView}>
            <Button color={Theme.colors.primaryColor}
                onPress={this.takePiture} 
            title = "Upload Image"/>
            
            <HideableView hide={this.state.isHidden}><Image style={{height: 200, width: 100}} source = {this.state.source} /></HideableView>
        </View>
      );
   }
}


