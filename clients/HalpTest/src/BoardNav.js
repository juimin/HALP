import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

export default class BoardNav extends React.Component {
  static navigationOptions = {
    tabBarIcon: ({ tintColor }) => (<Icon size={28} name="add-circle" style={{color:tintColor}}/>)
  }
 
  render() {
  	return (
  		<Text>this is a navigation page</Text>
  	)
  }
}