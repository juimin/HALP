import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator, StackNavigator } from 'react-navigation';
//import BottomNavigation, { Tab } from 'react-native-material-bottom-navigation'
import Icon from 'react-native-vector-icons/MaterialIcons'
import HomeScreen from './HomeScreen';
import SignupScreen from './SignupScreen';
import LoginScreen from './LoginScreen';
import CanvasTest from './CanvasTest';

export default class HomeNav extends React.Component {
  static navigationOptions = {
    tabBarIcon: ({ tintColor }) => (<Icon size={28} name="home" style={{color:tintColor}}/>)
  }
 
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