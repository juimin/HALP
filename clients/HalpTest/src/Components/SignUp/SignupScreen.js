// Import needed react dependancies
import React, { Component } from 'react';
import { ScrollView, Button, View, Text } from 'react-native';
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
import { loginAction } from '../../Redux/Actions';

const endpoint = "users"

function validateEmail(email) {
   var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
   return re.test(String(email).toLowerCase());
}

const mapDispatchToProps = (dispatch) => {
   return {
      login: token => { dispatch(loginAction(token)) },
      setUser: usr => { dispatch(setUser(usr)) }
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
   }

   // Validate the user input
   validate() {
      var errored = false
      // Validate email
      if (!validateEmail(this.state.email)) {
         this.errors.email = true
         errored = true
      }

      if (this.state.userName.length < 6) {
         this.errors.userName = true
         errored = true
      }

      if (this.state.password.length < 6) {
         this.errors.password = true
         errored = true
      }

      if (this.state.passwordConf.length < 6) {
         this.errors.passwordConf = true
         errored = true
      }

      if (this.state.password != this.state.passwordConf) {
         this.errors.passwordConf = true
         errored = true
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
               console.log(response.headers.get('authorization'))
               this.props.login(response.headers.get("authorization"))
               console.log(JSON.parse(response.body))
               // this.props.setUser(response.json())
               this.props.navigation.goBack()
            } else {
               console.log(response)
            }
         }).catch(err => {
            console.log(err)
         })
      } else {
         console.log(this.errors)
      }
   }

   render() {
      return (
         <ScrollView>
            <FormLabel>Email *</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.state.email = text}}/>
            <FormLabel>Username *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.userName = text}}/>
            <FormLabel>Password *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.password = text}}/>
            <FormLabel>Password Confirm *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.passwordConf = text}}/>
            <FormLabel>First Name *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.firstName = text}}/>
            <FormLabel>Last Name *</FormLabel>
            <FormInput onChangeText={(text) => {this.state.lastName = text}}/>
            <FormLabel>Occupation</FormLabel>
            <FormInput onChangeText={(text) => {this.state.occupation = text}}/>
            <Button title="Submit" onPress={this.signup}></Button>
         </ScrollView>
      );
   }
}

export default connect(null, mapDispatchToProps)(SignupScreen)