// Import needed react dependancies
import React, { Component } from 'react';
import { Alert } from 'react-native';
import { StackNavigator } from 'react-navigation';

import {
   Container,
   Content,
   Form,
   Input,
   Item,
   Label,
	Text,
	H1,
   Button
} from 'native-base'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
import { API_URL } from '../../Constants/Constants';

import { FormLabel, FormInput, FormValidationMessage } from 'react-native-elements'
import Icon from 'react-native-vector-icons/FontAwesome';
// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { setTokenAction, setUserAction, savePasswordAction } from '../../Redux/Actions';

// Material UI Components

const endpoint = "users"

function validateEmail(email) {
   var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
   return re.test(String(email).toLowerCase());
}

const mapDispatchToProps = (dispatch) => {
   return {
      addAuthToken: token => { dispatch(setTokenAction({token})) },
      setUser: usr => { dispatch(setUserAction(usr)) },
      savePassword: pass => { dispatch(savePasswordAction({pass}))}
   }
}

// Define and export the component
class SignupScreen extends Component {
   constructor(props) {
      super(props)
      this.state = {
			values: {
				email: "",
				userName: "",
				password: "",
				passwordConf: "",
				firstName: "",
				lastName: "",
				occupation: ""
			},
			errored: {
				email: false,
				userName: false,
				password: false,
				passwordConf: false,
				firstName: false,
				lastName: false,
			},
			fieldLabels: {
				email: "Email",
				userName: "Username",
				password: "Password",
				passwordConf: "Confirm Password",
				firstName: "First Name",
				lastName: "Last Name",
				occupation: "Occupation (Optional)"
			},
			success: {
				email: false,
				userName: false,
				password: false,
				passwordConf: false,
				firstName: false,
				lastName: false,
			}
		}
      this.signup = this.signup.bind(this)
      this.validateForm = this.validateForm.bind(this)
   }

   // Validate the user input
   validateForm() {
		var tempErrored = {
			email: !validateEmail(this.state.values.email),
			userName: !(this.state.values.userName.length >= 6 && this.state.values.userName.length <= 26),
			password: !(this.state.values.password.length >= 6 && this.state.values.password.length <= 26),
			passwordConf: (this.state.values.password != this.state.values.passwordConf),
			firstName: (this.state.values.firstName.length == 0),
			lastName: (this.state.values.lastName.length == 0)
		}

		this.setState({
			values: this.state.values,
			errored: tempErrored,
			fieldLabels: {
				email: (tempErrored.email) ? "Email (Invalid Format)" : "Email",
				userName: (tempErrored.userName) ? "Username (must be 6 - 26 characters)" : "Username",
				password: (tempErrored.password) ? "Password (must be 6 - 26 characters)" : "Password",
				passwordConf: (tempErrored.passwordConf) ? "Confirmation Password (must match Password)" : "Confirm Password",
				firstName: (tempErrored.firstName) ? "First Name (Required)" : "First Name",
				lastName: (tempErrored.lastName) ? "Last Name (Required)" : "Last Name",
				occupation: "Occupation (Optional)"
			},
			success: {
				email: !tempErrored.email,
				userName: !tempErrored.userName,
				password: !tempErrored.password,
				passwordConf: (this.state.values.password == this.state.values.passwordConf),
				firstName: !tempErrored.firstName,
				lastName: !tempErrored.lastName
			}
		})

		var errored = tempErrored.email || tempErrored.userName || tempErrored.lastName || tempErrored.firstName || tempErrored.password || tempErrored.passwordConf
		return !errored
   }

   // sign up function
   signup() {
      // Validate the input
      if (this.validateForm()) {
         fetch(API_URL + "users", {
            method: "POST",
            headers: {
               'Accept': 'application/json',
               'Content-Type': 'application/json',
           },
           body: JSON.stringify(this.state.values)
         }).then(response => {
            if (response.status == 201) {
					this.props.addAuthToken(response.headers.get("authorization"))
					this.props.savePassword(this.state.password)
					return response.json()
				} else if (response.status == 409) {
					return response.text()
            } else {
					return null
            }
         }).then(resp => {
				if (resp != null) {
					if (resp == "Error: Email already Taken" || resp == "Error: Username already taken") {
						// Something went wrong with the server
						Alert.alert(
							'Sign Up Error',
							resp,
							[
								{text: 'OK', onPress: () => console.log('OK Pressed')},
							]
						)
					} else {
						// Save the user to the thing
						this.props.setUser(resp)
						// Save the password
						this.props.navigation.goBack()
					}
				} else {
               // Something went wrong with the server
               Alert.alert(
                  'Sign Up Error',
                  'Server is experiencing problems. Try again later.',
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
			// Rerender the component
         this.validateForm()
      }
   }

   render() {
		console.log(this.state)
      return (
			<Container>
				<Content style={{paddingRight:"5%"}}>
					<Form>
						<Item floatingLabel error={this.state.errored.email} success={this.state.success.email}>
                     <Label>{this.state.fieldLabels.email}</Label>
                     <Input onChangeText={text => {
                        this.setState({
									errored: this.state.errored,
									fieldLabels: this.state.fieldLabels,
									success: this.state.success,
									values: {
										email: text,
										userName: this.state.values.userName,
										password: this.state.values.password,
										passwordConf: this.state.values.passwordConf,
										firstName: this.state.values.firstName,
										lastName: this.state.values.lastName,
										occupation: this.state.occupation
									}
                        })
                     }}/>
                  </Item>
						<Item floatingLabel error={this.state.errored.userName} success={this.state.success.userName}>
                     <Label>{this.state.fieldLabels.userName}</Label>
                     <Input onChangeText={text => {
                        this.setState({
									errored: this.state.errored,
									success: this.state.success,
									fieldLabels: this.state.fieldLabels,
									values: {
										email: this.state.values.email,
										userName: text,
										password: this.state.values.password,
										passwordConf: this.state.values.passwordConf,
										firstName: this.state.values.firstName,
										lastName: this.state.values.lastName,
										occupation: this.state.values.occupation
									}
								})
                     }}/>
                  </Item>
						<Item floatingLabel error={this.state.errored.password} success={this.state.success.password}>
                     <Label>{this.state.fieldLabels.password}</Label>
                     <Input secureTextEntry onChangeText={text => {
                        this.setState({
									errored: this.state.errored,
									fieldLabels: this.state.fieldLabels,
									success: this.state.success,
									values: {
										email: this.state.values.email,
										userName: this.state.values.userName,
										password: text,
										passwordConf: this.state.values.passwordConf,
										firstName: this.state.values.firstName,
										lastName: this.state.values.lastName,
										occupation: this.state.values.occupation
									}
                        })
                     }}/>
                  </Item>
						<Item floatingLabel error={this.state.errored.passwordConf} success={this.state.success.passwordConf}>
                     <Label>{this.state.fieldLabels.passwordConf}</Label>
                     <Input secureTextEntry onChangeText={text => {
                        this.setState({
									errored: this.state.errored,
									fieldLabels: this.state.fieldLabels,
									success: this.state.success,
									values: {
										email: this.state.values.email,
										userName: this.state.values.userName,
										password: this.state.values.password,
										passwordConf: text,
										firstName: this.state.values.firstName,
										lastName: this.state.values.lastName,
										occupation: this.state.values.occupation
									}
                        })
                     }}/>
                  </Item>
						<Item floatingLabel error={this.state.errored.firstName} success={this.state.success.firstName}>
                     <Label>{this.state.fieldLabels.firstName}</Label>
                     <Input onChangeText={text => {
                        this.setState({
									errored: this.state.errored,
									success: this.state.success,
									fieldLabels: this.state.fieldLabels,
									values: {
										email: this.state.values.email,
										userName: this.state.values.userName,
										password: this.state.values.password,
										passwordConf: this.state.values.passwordConf,
										firstName: text,
										lastName: this.state.values.lastName,
										occupation: this.state.values.occupation
									}

                        })
                     }}/>
                  </Item>
						<Item floatingLabel error={this.state.errored.lastName} success={this.state.success.lastName}>
                     <Label>{this.state.fieldLabels.lastName}</Label>
                     <Input onChangeText={text => {
                        this.setState({
									errored: this.state.errored,
									fieldLabels: this.state.fieldLabels,
									success: this.state.success,
									values: {
										email: this.state.values.email,
										userName: this.state.values.userName,
										password: this.state.values.password,
										passwordConf: this.state.values.passwordConf,
										firstName: this.state.values.firstName,
										lastName: text,
										occupation: this.state.values.occupation
									}
                           
                        })
                     }}/>
                  </Item>
						<Item floatingLabel>
                     <Label>{this.state.fieldLabels.occupation}</Label>
                     <Input onChangeText={text => {
                        this.setState({
									errored: this.state.errored,
									fieldLabels: this.state.fieldLabels,
									success: this.state.success,
									values: {
										email: this.state.values.email,
										userName: this.state.values.userName,
										password: this.state.values.password,
										passwordConf: this.state.values.passwordConf,
										firstName: this.state.values.firstName,
										lastName: this.state.values.lastName,
										occupation: text,
									}
                        })
                     }}/>
                  </Item>
					</Form>
					<Content style={{marginTop: "7%"}}>
					<Button rounded style={Styles.button} onPress={this.signup}>
						<Text>Submit</Text>
					</Button>
					</Content>
				</Content>
			</Container>
      );
   }
}

export default connect(null, mapDispatchToProps)(SignupScreen)