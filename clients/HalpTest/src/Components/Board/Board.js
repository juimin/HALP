// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import ReduxActions from '../../Redux/Actions';


const mapStateToProps = (state) => {
	return {
      authToken: state.AuthReducer.authToken,
      user: state.AuthReducer.user,
      password: state.AuthReducer.password,
      activeBoard: state.BoardReducer.activeBoard
	}
}

class Board extends Component {
	render() {
		return (
			<View style={Styles.home}>
				<Text>{JSON.stringify(this.props.activeBoard)}</Text>
				<Text>Something</Text>
			</View>
		)
	}
}

export default connect(mapStateToProps)(Board)