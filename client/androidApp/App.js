import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import { NavigationComponent } from 'react-native-material-bottom-navigation'
import HomeScreen from './src/HomeScreen';
import SignupScreen from './src/SignupScreen';
import LoginScreen from './src/LoginScreen';
//import PrimaryNav from './src/AppNavigation';
import CanvasTest from './src/CanvasTest';
import HomeNav from './src/HomeNav';
import SearchNav from './src/SearchNav';
import AccNav from './src/AccNav';
import BoardNav from './src/BoardNav';

export default class App extends React.Component {
  render() {
    return <RootTab screenProps={{loggedin: false}}/>;
  }
}

//react-navigation

// const RootStack = StackNavigator(
//   {
//     Home: {
//       screen: HomeScreen,
//       navigationOptions: { title: 'Home' },
//     },
//     Login: {
//       screen: LoginScreen,
//       navigationOptions: { title: 'Log In' },
//     },
//     Signup: {
//     	screen: SignupScreen,
//     	navigationOptions: { title: 'Sign Up' },
//     },
//     Canvas: {
//       screen: CanvasTest,
//       navigationOptions: {title: 'Canvas'},
//     },
//   },
//   {
//     initialRouteName: 'Home',
//     headerMode: 'screen',

//   },
// );

const RootTab = TabNavigator({
  HomeNav: { screen: HomeNav },
  SearchNav: { screen: SearchNav },
  BoardNav: { screen: BoardNav },
  AccNav: { screen: AccNav }
}, {
  tabBarComponent: NavigationComponent,
  tabBarPosition: 'bottom',
  tabBarOptions: {
    bottomNavigationOptions: {
      backgroundColor: "#F44336",
      labelColor: 'white',
      rippleColor: 'white',
      tabs: {
        HomeNav: {
          activeLabelColor: 'white', /*overwrites default label color*/
        },
        SearchNav: {
        },
        BoardNav: {
        },
        AccNav: {
        }
      }
    }
  }
})


