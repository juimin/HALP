import React, { Component } from 'react';
import { ScrollView, Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements'
import Icon from 'react-native-vector-icons/FontAwesome';

export default class LoginScreen extends Component {
   // See if we can submit
   // If we can then we should route away
   constructor(props) {
      super(props)
      this.state = {
         email: "",
         password: ""
      }
   }

   login() {
      console.log(this.state)
   }

      
   render() {
      return (
         <View style={Styles.signup}>
            <FormLabel>Email</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.state.email = text}}/>
            <FormLabel>Password</FormLabel>
            <FormInput onChangeText={(text) => {this.state.password = text}}/>
            <Button title="Log In" onPress={()=> this.login()}></Button>
         </View>
      );
   }
}