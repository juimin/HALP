import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

export default class LoginScreen extends Component {
  render() {
    const {goBack} = this.props.navigation;
    return (
      <View style={Styles.login}>
        <Text>Log In Here!</Text>
      </View>
    );
  }
}