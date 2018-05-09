import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

export default class SearchNav extends React.Component {
  static navigationOptions = {
    tabBarIcon: ({ tintColor }) => (<Icon size={28} name="search" style={{color:tintColor}}/>)
  }
 
  render() {
  	return (
  		<Text>this is a search page</Text>
  	)
  }
}