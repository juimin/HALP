// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

export default class Account extends Component {
  render() {
  	return (
      <View style={Styles.home}>
         <Text>Account Info for logged in User</Text>
      </View>
  	)
  }
}