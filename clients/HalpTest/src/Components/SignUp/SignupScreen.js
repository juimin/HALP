// Import needed react dependancies
import React, { Component } from 'react';
import { Alert, ScrollView, Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
import { API_URL } from '../../Constants/Constants';

import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements'
import Icon from 'react-native-vector-icons/FontAwesome';
// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { loginAction, setUserAction } from '../../Redux/Actions';

// Material UI Components

const endpoint = "users"

function validateEmail(email) {
   var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
   return re.test(String(email).toLowerCase());
}

const mapDispatchToProps = (dispatch) => {
   return {
      addAuthToken: token => { dispatch(loginAction(token)) },
      setUser: usr => { dispatch(setUserAction(usr)) }
   }
}

// Define and export the component
class SignupScreen extends Component {
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

      this.errors = {
         email: false,
         userName: false,
         password: false,
         passwordConf: false,
         firstName: false,
         lastName: false,
         occupation: false
      }

      this.errorMessages = {
         email: "",
         userName: "",
         password: "",
         passwordConf: "",
         firstName: "",
         lastName: "",
         occupation: ""
      }

      this.signup = this.signup.bind(this)
      this.validate = this.validate.bind(this)
      this.resetForm = this.resetForm.bind(this)
   }

   // Reset the form state
   resetForm() {
      this.state = {
         email: "",
         userName: "",
         password: "",
         passwordConf: "",
         firstName: "",
         lastName: "",
         occupation: ""
      }

      this.errors = {
         email: false,
         userName: false,
         password: false,
         passwordConf: false,
         firstName: false,
         lastName: false,
         occupation: false
      }

      this.errorMessages = {
         email: "",
         userName: "",
         password: "",
         passwordConf: "",
         firstName: "",
         lastName: "",
         occupation: ""
      }
   }

   // Validate the user input
   validate() {
      var errored = false
      // Validate email
      if (!validateEmail(this.state.email)) {
         this.errors.email = true
         this.errorMessages.email = "Email is invalid"
         errored = true
      } else {
         this.errors.email = false
         this.errorMessages.email = ""
      }

      if (this.state.userName.length < 6) {
         this.errors.userName = true
         this.errorMessages.userName = "Username must be at least 6 characters"
         errored = true
      } else {
         this.errors.userName = false
         this.errorMessages.userName = ""
      }

      if (this.state.password.length < 6) {
         this.errors.password = true
         this.errorMessages.password = "Password must be at least 6 characters"
         errored = true
      } else {
         this.errors.password = false
         this.errorMessages.password = ""
      }

      if (this.state.password != this.state.passwordConf) {
         this.errors.passwordConf = true
         this.errorMessages.password = "PasswordConf does not match password"
         errored = true
      } else {
         this.errorMessages.passwordConf = ""
         this.errors.passwordConf = false
      }

      if (this.state.firstName.length <= 0) {
         this.errorMessages.firstName = "Must Enter a first Name"
         this.errors.firstName = true
         errored = true
      } else {
         this.errorMessages.firstName = ""
         this.errors.firstName = false
      }

      if (this.state.lastName.length <= 0) {
         this.errorMessages.lastName = "Must Enter a first Name"
         this.errors.lastName = true
         errored = true
      } else {
         this.errorMessages.lastName = ""
         this.errors.lastName = false
      }

      return !errored
   }

   // sign up function
   signup() {
      // Validate the input
      if (this.validate()) {
         var x = fetch(API_URL + endpoint, {
            method: "POST",
            headers: {
               'Accept': 'application/json',
               'Content-Type': 'application/json',
           },
           body: JSON.stringify(this.state)
         }).then(response => {
            if (response.status == 201) {
               this.props.addAuthToken(response.headers.get("authorization"))
               this.props.navigation.goBack()
            } else {
               // Something went wrong with the server
               Alert.alert(
                  'Sign Up Error',
                  'A problem arose when signing up. Please try again',
                  [
                    {text: 'OK', onPress: () => console.log('OK Pressed')},
                  ]
                )
            }
         }).catch(err => {
            Alert.alert(
               'Error getting response from server',
               err,
               [
                 {text: 'OK', onPress: () => console.log('OK Pressed')},
               ]
             )
         })
      } else {
         // RErender the component
         this.setState(this.state)
      }
   }

   render() {
      return (
         <ScrollView>
            <FormLabel>Email *</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.state.email = text}}/>
            <FormValidationMessage>{this.errorMessages.email}</FormValidationMessage>
            <FormLabel>Username *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.userName = text}}/>
            <FormValidationMessage>{this.errorMessages.userName}</FormValidationMessage>
            <FormLabel>Password *</FormLabel>
            <FormInput secureTextEntry={true} onChangeText={(text) => {this.state.password = text}}/>
            <FormValidationMessage>{this.errorMessages.password}</FormValidationMessage>
            <FormLabel>Password Confirm *</FormLabel>
            <FormInput secureTextEntry={true} onChangeText={(text) => {this.state.passwordConf = text}}/>
            <FormValidationMessage>{this.errorMessages.passwordConf}</FormValidationMessage>
            <FormLabel>First Name *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.firstName = text}}/>
            <FormValidationMessage>{this.errorMessages.firstName}</FormValidationMessage>
            <FormLabel>Last Name *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.lastName = text}}/>
            <FormValidationMessage>{this.errorMessages.lastName}</FormValidationMessage>
            <FormLabel>Occupation</FormLabel>
            <FormInput onChangeText={(text) => {this.state.occupation = text}}/>
            <Button title="Submit" onPress={this.signup}></Button>
         </ScrollView>
      );
   }
}

export default connect(null, mapDispatchToProps)(SignupScreen)