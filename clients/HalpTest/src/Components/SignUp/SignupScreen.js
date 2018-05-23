// Import needed react dependancies
import React, { Component } from 'react';
import { ScrollView, Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
import API_URL from '../../Constants/Constants';

import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements'
import Icon from 'react-native-vector-icons/FontAwesome';

const endpoint = "users"

// Define and export the component
export default class SignupScreen extends Component {
   constructor(props) {
      super(props)
      this.state = {
         email: "",
         userName: "",
         password: "",
         passwordConf: "",
         firstName: "",
         lastName: "",
         occupation: ""
      }

      this.signup = this.signup.bind(this)
   }

   // sign up function
   signup() {
      fetch(API_URL + endpoint, {
         method: "POST",
         headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(this.state)
      }).then(response => {
         console.log(response)
      }).catch(err => {
         console.log(err)
      })
   }

   render() {
      return (
         <ScrollView>
            <FormLabel>Email</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.state.email = text}}/>
            <FormLabel>Username</FormLabel>
            <FormInput onChangeText={(text) => {this.state.userName = text}}/>
            <FormLabel>Password</FormLabel>
            <FormInput onChangeText={(text) => {this.state.password = text}}/>
            <FormLabel>Password Confirm</FormLabel>
            <FormInput onChangeText={(text) => {this.state.passwordConf = text}}/>
            <FormLabel>First Name</FormLabel>
            <FormInput onChangeText={(text) => {this.state.firstName = text}}/>
            <FormLabel>Last Name</FormLabel>
            <FormInput onChangeText={(text) => {this.state.lastName = text}}/>
            <FormLabel>Occupation</FormLabel>
            <FormInput onChangeText={(text) => {this.state.occupation = text}}/>
            <Button title="Submit" onPress={this.signup}></Button>
         </ScrollView>
      );
   }
}