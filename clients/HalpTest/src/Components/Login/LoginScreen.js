import React, { Component } from 'react';
import { ScrollView, Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// import t from 'tcomb-form-native';

// const Form = t.form.Form;

import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements'
import Icon from 'react-native-vector-icons/FontAwesome';

export default class LoginScreen extends Component {
   // This doesn't need global state because it's just a form
   constructor(props) {
      super(props)
      this.state = {
         newUser: {

         },
         errors: {

         }
      }

      // Bind the on click Function
      this.alterState = this.alterState.bind(this)
   }

   // Alter the state
   alterState(text) {
      console.log(text)
   }

   // See if we can submit
   // If we can then we should route away
   submit() {

   }
      
   render() {
      return (
         <View style={Styles.signup}>
            <FormLabel>Email</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.alterState(text)}}/>
            <FormLabel>Password</FormLabel>
            <FormInput onChangeText={(text) => {this.alterState(text)}}/>
            <Button title="Log In" onPress={()=>{console.log("?")}}></Button>
         </View>
      );
   }
}