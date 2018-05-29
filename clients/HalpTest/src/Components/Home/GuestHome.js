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

// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

// Export the default class
class GuestHome extends Component {
   render() {
      return(
         <View style={Styles.home}>
            <Button color={Theme.colors.primaryColor} title="Sign Up"
               onPress={() => this.props.navigation.navigate('Signup')}
            />
            {/* <Button color={Theme.colors.primaryColor} title="Canvas Test"
               onPress={() => this.props.navigation.navigate('Canvas')}
            /> */}
         </View>
      )
   }
}

export default GuestHome