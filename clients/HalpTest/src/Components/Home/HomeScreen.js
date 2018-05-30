import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import the different views based on user state
import UserHome from './UserHome';
import GuestHome from './GuestHome';
import HomeNav from '../Navigation/HomeNav';

// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

const mapStateToProps = (state) => {
	return {
		user: state.AuthReducer.user
	}
}

class HomeScreen extends Component {
   // Here we should run initialization scripts
   render() {
      if (this.props.user != null) {
         return (
            <UserHome {...this.props} />
         );
      }
      //if not logged in
      return (
         <GuestHome {...this.props} />
      );
   }
}

export default connect(mapStateToProps)(HomeScreen)
