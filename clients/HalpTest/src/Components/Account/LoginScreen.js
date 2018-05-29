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

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

class NewUser extends Component {
   constructor(props) {
      super(props)
   }

   render() {
      return(
         <Container style={Styles.home}>
            <Thumbnail square source={{uri: uri}} />
            <Container style={{
               marginTop: "25%",
               width: "90%"
            }}>
               <Form>
                  <Item>
                     <Label>Email</Label>
                     <Input/>
                  </Item>
               </Form>
            </Container>
            <Button rounded style={Styles.button} 
               onPress={() => this.props.navigation.navigate('Login')}
            >
               <Text>Log In</Text>
            </Button>
            <Button rounded style={Styles.button} 
               onPress={() => this.props.navigation.navigate('Signup')}
            >
               <Text>Sign Up</Text>
            </Button>
         </Container>
      )
   }
}

export default NewUser