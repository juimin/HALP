import React, { Component } from 'react';
import { ScrollView, Button, Switch, Slider, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import Themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme'

// Export the Component
export default class Settings extends Component {
   render() {
      return (
         <ScrollView >
            <Text style={Styles.settingTitle}>Toggle Settings</Text>
            <Switch />
            <Switch />
            <Switch />
            <Switch />
            <Text style={Styles.settingTitle}>Other Settings Settings</Text>
            <Slider />
            <Slider />
            <Slider />
         </ScrollView>
      )
   }
}