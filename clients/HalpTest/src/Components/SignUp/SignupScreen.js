// Import needed react dependancies
import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Define and export the component
export default class SignupScreen extends Component {
   render() {
      return (
         <View style={Styles.signup}>
            <Text>Sign Up Here! It worked Potato</Text>
         </View>
      );
   }
}