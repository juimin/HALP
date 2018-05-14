// Import default react components
import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';

// Import Halp Components
import HomeScreen from '../Home/HomeScreen';
import SignupScreen from '../SignUp/SignupScreen';
import LoginScreen from '../Login/LoginScreen';
import CanvasTest from '../Canvas/CanvasTest';

// Generate a stack for navigation
// Generally, this is the component that wraps the child components
// Specifically for this file, App.js will use this as a component because it allows for
// navigating between the Compoents listed
const RootStack = StackNavigator(
   {
      Home: {
         screen: HomeScreen,
         navigationOptions: { title: 'Home' },
      },
      Login: {
         screen: LoginScreen,
         navigationOptions: { title: 'Log In' },
      },
      Signup: {
         screen: SignupScreen,
         navigationOptions: { title: 'Sign Up' },
      },
      Canvas: {
         screen: CanvasTest,
         navigationOptions: {title: 'Edit Image'},
      },
   },
   {
      initialRouteName: 'Home',
      headerMode: 'screen',
   },
);

// Export the navigator
export default class HomeNav extends Component {
   render() {
      return <RootStack screenProps={{loggedin: false}}/>;
   }
}