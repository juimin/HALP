// Import needed react dependancies
import React, { Component } from 'react';
import { ScrollView, Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';


import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements'
import Icon from 'react-native-vector-icons/FontAwesome';

// Define and export the component
export default class SignupScreen extends Component {
   render() {
      return (
         <ScrollView>
            <FormLabel>Email</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.alterState(text)}}/>
            <FormLabel>Username</FormLabel>
            <FormInput onChangeText={(text) => {this.alterState(text)}}/>
            <FormLabel>Password</FormLabel>
            <FormInput onChangeText={(text) => {this.alterState(text)}}/>
            <FormLabel>Password Confirm</FormLabel>
            <FormInput onChangeText={(text) => {this.alterState(text)}}/>
            <FormLabel>First Name</FormLabel>
            <FormInput onChangeText={(text) => {this.alterState(text)}}/>
            <FormLabel>Last Name</FormLabel>
            <FormInput onChangeText={(text) => {this.alterState(text)}}/>
            <FormLabel>Occupation</FormLabel>
            <FormInput onChangeText={(text) => {this.alterState(text)}}/>
            <Button title="Submit" onPress={()=>{console.log("?")}}></Button>
         </ScrollView>
      );
   }
}