import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import Themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme'

// Export the Component
export default class Settings extends Component {
   render() {
      return (
         <View style={Styles.home}>
            <Text>Settings for logged in User</Text>
         </View>
      )
   }
}