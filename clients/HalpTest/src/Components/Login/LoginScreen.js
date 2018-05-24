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

const endpoint = "sessions"

const mapDispatchToProps = (dispatch) => {
   return {
      addAuthToken: token => { dispatch(loginAction(token)) },
      setUser: usr => { dispatch(setUserAction(usr)) }
   }
}

class LoginScreen extends Component {
   // See if we can submit
   // If we can then we should route away
   constructor(props) {
      super(props)
      this.state = {
         email: "",
         password: ""
      }

      this.login = this.login.bind(this)
   }

   login() {
      fetch(API_URL + endpoint, {
         method: 'POST',
         headers: {
             'Accept': 'application/json',
             'Content-Type': 'application/json',
         },
         body: JSON.stringify(this.state)
      }).then(response => {
         if (response.status == 202) {
            this.props.addAuthToken(response.headers.get("authorization"))
            return response.json()
         } else {
            // Something went wrong with the server
            Alert.alert(
               'Sign Up Error',
               response.status.toString(),
               [
                  {text: 'OK', onPress: () => console.log('OK Pressed')},
               ]
            )
            return null
         }
      }).then(user => {
         if (user != null) {
            // Save the user to the thing
            this.props.setUser(user)
            this.props.navigation.goBack()
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
   }

   render() {
      return (
         <View style={Styles.signup}>
            <FormLabel>Email</FormLabel>
            <FormInput style={Styles.signinFormInput} onChangeText={(text) => {this.state.email = text}}/>
            <FormLabel>Password</FormLabel>
            <FormInput  secureTextEntry={true}  onChangeText={(text) => {this.state.password = text}}/>
            <Button title="Log In" onPress={this.login}></Button>
         </View>
      );
   }
}

export default connect(null, mapDispatchToProps)(LoginScreen)