// GuestHome describes the home screen seen by a guest user.
// A guest user should be defined as a user that has yet to create an account
// or is not yet loged in.

// Import React Components
import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// TODO FIX THIS

// Export the default class
export default class GuestHome extends Component {
   constructor(props) {
      super(props);
      this.state = {loggedin: false};
   }

   render() {
      return(
         <View style={Styles.home}>
            <Button color={Theme.colors.primaryColor} title="Log in"
               onPress={() => this.props.navigation.navigate('Login')}
            />
            <Button color={Theme.colors.primaryColor} title="Sign Up"
               onPress={() => this.props.navigation.navigate('Signup')}
            />
            <Button color={Theme.colors.primaryColor} title="Try Me"
               onPress={() => this.setState({loggedin: true})}
            />
            <Button color={Theme.colors.primaryColor} title="Canvas Test"
               onPress={() => this.props.navigation.navigate('Canvas')}
            />
         </View>
      )
   }
}