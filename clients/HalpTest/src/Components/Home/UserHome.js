// UserHome describes the home screen seen by a known user.
// This just means a user that has logged in.

// Import React Components
import React from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Export the default class
export default class UserHome extends React.Component {
   constructor(props) {
      super(props);
   }

   render() {
      return(
         <View style={Styles.home}>
            <Text>Dashboard for logged in User</Text>
         </View>
      )
   }
}