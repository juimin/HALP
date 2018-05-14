import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';


export default class Search extends Component {
  render() {
  	return (
      <View style={Styles.home}>
         <Text>This is the search home</Text>
      </View>
  	)
  }
}