// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

export default class Account extends Component {
  static navigationOptions = {
    tabBarIcon: ({ tintColor }) => (<Icon size={28} name="person" style={{color:tintColor}}/>)
  }
 
  render() {
  	return (
  		<Text>this is an account page</Text>
  	)
  }
}