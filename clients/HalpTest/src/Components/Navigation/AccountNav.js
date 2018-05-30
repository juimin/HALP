// Import default react components
import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';

// Import Halp Components
import Account from '../Account/Account';
import SearchNav from './SearchNav';
import SignupScreen from '../SignUp/SignupScreen';

// Generate a stack for navigation
// Generally, this is the component that wraps the child components
// Specifically for this file, App.js will use this as a component because it allows for
// navigating between the Compoents listed
const RootStack = StackNavigator({
      Account: {
			  screen: Account,
			  navigationOptions: {
          header: null
        }
      },
      Signup: {
			  screen: SignupScreen,
      }
   },
   {
      initialRouteName: 'Account',
      headerMode: 'screen',
   },
);

export default RootStack
