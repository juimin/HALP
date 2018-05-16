// Import default react components
import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';

// Import Halp Components
import HomeScreen from '../Home/HomeScreen';
import SignupScreen from '../SignUp/SignupScreen';
import LoginScreen from '../Login/LoginScreen';
import CanvasTest from '../Canvas/CanvasTest';
import GuestHome from '../Home/GuestHome';
import UserHome from '../Home/UserHome';

// Generate a stack for navigation
// Generally, this is the component that wraps the child components
// Specifically for this file, App.js will use this as a component because it allows for
// navigating between the Compoents listed
const loggedin = false;

const RootStack = StackNavigator({
      Home: {
         screen: loggedin ? UserHome : GuestHome,
         navigationOptions: {
            title: "HALP",
            titleStyle: {
                  textALign: 'center'
            }
         }
      },
      UserHome: {
         screen: UserHome
      },
      GuestHome: {
         screen: GuestHome
      },
      Login: {
         screen: LoginScreen,
      },
      Signup: {
         screen: SignupScreen,
      },
      Canvas: {
         screen: CanvasTest,
      },
   },
   {
      initialRouteName: 'Home',
      headerMode: 'screen',
   },
);

export default RootStack
