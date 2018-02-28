import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import BottomNavigation, { Tab } from 'react-native-material-bottom-navigation'
import Icon from 'react-native-vector-icons/MaterialIcons'

export default class AccNav extends React.Component {
  static navigationOptions = {
    tabBarLabel: 'Account',
    tabBarIcon: () => (<Icon size={24} color="white" name="account-circle" />)
  }
 
  render() {
  	return (
  		<Text>this is an account page</Text>
  	)
  }
}