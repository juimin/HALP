// This should be the root of all application components
// Everything runs under a stack navigation nexted from here

// Import required react components
import React, { Component } from 'react';
import { Button, View, Text, TouchableWithoutFeedback} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import Navigation Components
import NewPostStack from '../Navigation/NewPostNav';
import Tabs from '../Navigation/TabNav';

/*
 * We need a root stack navigator with the mode set to modal so that we can open the capture screen
 * as a modal. Defaults to the Tabs navigator.
 */
const AppRootStack = StackNavigator({
   // Add the tab navigator as a modal?
   Tabs: {
      screen: Tabs,
   },
   // Add this modal because it is separate from the main application navigation
   NewPostModal: {
      screen: NewPostStack,
      navigationOptions: {
         gesturesEnabled: false,
      },
   },
   }, {
   headerMode: 'none',
   mode: 'modal',
});

export default AppRootStack;