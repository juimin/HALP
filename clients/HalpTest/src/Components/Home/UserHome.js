// UserHome describes the home screen seen by a known user.
// This just means a user that has logged in.

// Import React Components
import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import redux
import { connect } from 'redux'
import { loginAction } from '../../Redux/Actions'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Export the default class
class UserHome extends Component {
   render() {
      return(
         <View style={Styles.home}>
            <Text>Dashboard for logged in User</Text>
         </View>
      )
   }
}

export default UserHome