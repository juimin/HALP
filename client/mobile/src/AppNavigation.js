import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { StackNavigator, DrawerNavigator } from 'react-navigation';
import HomeScreen from './src/HomeScreen';
import SignupScreen from './src/SignupScreen';
import LoginScreen from './src/LoginScreen';
import SignupTest from './src/SignupTest';
import LoginTest from './src/LoginTest';


//https://shift.infinite.red/react-navigation-drawer-tutorial-a802fc3ee6dc

const PrimaryNav = StackNavigator({
  loginStack: { screen: LoginStack },
  drawerStack: { screen: DrawerNavigation }
}, {
  // Default config for all screens
  headerMode: 'none',
  title: 'Main',
  initialRouteName: 'loginStack'
});

//login stack
const LoginStack = StackNavigator({
  loginScreen: { screen: LoginTest },
  signupScreen: { screen: SignupTest },
  {
  headerMode: 'screen',
  navigationOptions: {
    headerStyle: {backgroundColor: '#E73536'},
    title: 'You are not logged in',
    headerTintColor: 'white'
  }
});

const DrawerNavigation = StackNavigator({
  DrawerStack: { screen: DrawerStack }
}, {
  headerMode: 'screen',
  navigationOptions: ({navigation}) => ({
    headerStyle: {backgroundColor: '#4C3E54'},
    title: 'Welcome!',
    headerTintColor: 'white',
  })
});

const DrawerStack = DrawerNavigator({
  screen1: { screen: LoginTest },
  screen2: { screen: SignupTest },
});

export default PrimaryNav