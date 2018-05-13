import React from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import the different views based on user state
import UserHome from './UserHome';
import GuestHome from './GuestHome';

export default class HomeScreen extends React.Component {
   // Here we should run initialization scripts
   constructor(props) {
      super(props);
      this.state = {loggedin: true};
   }

   render() {
      const {goBack} = this.props.navigation;
      if (this.state.loggedin) {
         return (
            <UserHome />
         );
      }
      //if not logged in
      return (
         <GuestHome />
      );
   }
}