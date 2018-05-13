import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import Themes
import Theme from '../../Styles/Theme'

export default class Settings extends React.Component {
   // Set the icon color when selected ??
   // What is this actually for?
   static navigationOptions = {
      tabBarIcon: (tintColor) => (
         <Icon size={Theme.tabBar.iconSize} name="settings" style={{color:tintColor}}/>
      )
   }
 
   render() {
      return (
         <Text>this is the settings page</Text>
      )
   }
}

