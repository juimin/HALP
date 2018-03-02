import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { StackNavigator, DrawerNavigator } from 'react-navigation';
import HomeScreen from './src/HomeScreen';
import SignupScreen from './src/SignupScreen';
import LoginScreen from './src/LoginScreen';
//import PrimaryNav from './src/AppNavigation';



export default class App extends React.Component {
  render() {
    return <RootStack screenProps={{loggedin: false}}/>;
  }
}

//react-navigation

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
  },
  {
    initialRouteName: 'Home',
    headerMode: 'screen',

  },
);

