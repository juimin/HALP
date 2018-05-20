// This should be the root of all application components
// Everything runs under a stack navigation nexted from here

// Import required react components
import React, { Component } from 'react';
import { Button, View, Text, TouchableWithoutFeedback} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Define the new Post
export default class NewPost extends Component {
   render() {
      return(
         <View style={{ flex: 1, backgroundColor: '#fff', alignItems: 'center', justifyContent: 'center' }}>
            <Text>New Post Screen</Text>
         </View>
      );
   }
}