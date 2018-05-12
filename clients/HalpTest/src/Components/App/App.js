// Import react
import React, { Component } from 'react';
// Import Navigation and React Native Elements
import { Button, StyleSheet, View, Text, TouchableWithoutFeedback} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';

// Import Components from our Source Files
import Icon from 'react-native-vector-icons/MaterialIcons'
import HomeScreen from '../Home/HomeScreen';
import SignupScreen from '../Signup/SignupScreen';
import LoginScreen from '../Login/LoginScreen';
import CanvasTest from '../Canvas/CanvasTest';
import HomeNav from '../Home/HomeNav';
import SearchNav from '../Search/SearchNav';
import AccNav from '../Account/AccNav';
import BoardNav from '../Boards/BoardNav';
import SettingsNav from '../Settings/SettingsNav';
import NewPost from '../NewPost/NewPost';
import Tabs from '../Tabs/Tabs';

// We need a root stack navigator with the mode set to modal so that we can open the capture screen
// a modal. Defaults to the Tabs navigator.
const App = StackNavigator(
  // Route Config Map
  {
    // Add the tabs to the navigator
    Tabs: {
      screen: Tabs,
    },
    // Add the Modal for new posts as a custom element in the navigator
    NewPostModal: {
      screen: NewPost,
      navigationOptions: {
        gesturesEnabled: false,
      },
    },
  }, 
  // Stack Config
  {
    headerMode: 'none',
    mode: 'modal',
  }
);

export default App;
