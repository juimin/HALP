import React, { Component } from 'react'

import {
   Container,
   Content,
   Form,
   Input,
   Item,
   Label,
   Text,
   Thumbnail,
   Button
} from 'native-base'

import { Alert } from 'react-native'
import { API_URL } from '../../Constants/Constants';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

class NewUser extends Component {
   constructor(props) {
      super(props)
      // Set local login state. We don't need redux for this
      this.state = {
         email: "",
         password: ""
      }

      this.login = this.login.bind(this)
   }

   login() {
      fetch(API_URL + "sessions", {
         method: 'POST',
         headers: {
             'Accept': 'application/json',
             'Content-Type': 'application/json',
         },
         body: JSON.stringify(this.state)
      }).then(response => {
         if (response.status == 202) {
				// Save token and password for later use
            this.props.addAuthToken(response.headers.get("authorization"))
            this.props.savePassword(this.state.password)
            return response.json()
         } else {
            // Something went wrong with the server
            Alert.alert(
               'Error',
               'Invalid username and password combination',
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
      return(
         <Container style={Styles.home}>
            <Container style={{
               marginTop: "25%",
               marginRight: "4%",
               width: "90%",
            }}>
               <Thumbnail large source={require("../../Images/Logo-09.png")} style={{ alignSelf: "center"}} />
               <Content>
               <Form>
                  <Item floatingLabel>
                     <Label>Email</Label>
                     <Input onChangeText={text => {
                        this.setState({
                           email: text,
                           password: this.state.password
                        })
                     }}/>
                  </Item>
                  <Item floatingLabel>
                     <Label>Password</Label>
                     <Input secureTextEntry={true} onChangeText={text => {
                        this.setState({
                           email: this.state.email,
                           password: text
                        })
                     }}/>
                  </Item>
                  <Content style={{marginTop: "10%"}}>
                     <Content>
                        <Button rounded style={Styles.button} 
                           onPress={this.login}
                        >
                           <Text>Log In</Text>
                        </Button>
                     </Content>
                     <Content style={{marginTop:"10%"}}>
                        <Text style={{alignSelf: "center", marginBottom: "2%"}}>New to HALP? Make an account here</Text>
                        <Button rounded style={Styles.button} 
                           onPress={() => this.props.navigation.navigate('Signup')}
                        >
                           <Text>Sign Up</Text>
                        </Button>
                     </Content>
                  </Content>
                  </Form>
               </Content>
            </Container>
         </Container>
      )
   }
}

export default NewUser