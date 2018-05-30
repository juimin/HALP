// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { ScrollView } from 'react-native';
import {
	Container,
	Right,
	Body,
	Left,
	Title,
	Subtitle,
	Header,
	Button,
	Icon,
	Thumbnail,
	FooterTab,
	Content,
	Text,
	ActionSheet,
} from 'native-base';

// Import Component pieces
import LoginScreen from './LoginScreen';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { setTokenAction, setUserAction, savePasswordAction } from '../../Redux/Actions';

const mapStateToProps = (state) => {
	return {
		authToken: state.AuthReducer.authToken,
		password: state.AuthReducer.password,
		user: state.AuthReducer.user
	}
}

const mapDispatchToProps = (dispatch) => {
   return {
      addAuthToken: token => { dispatch(setTokenAction(token)) },
      setUser: usr => { dispatch(setUserAction(usr)) },
		savePassword: pass => { dispatch(savePasswordAction(pass))}
   }
}

class Account extends Component {

   constructor(props) {
      super(props)
      this.state = {
			menu: {
				selectedIndex: -1
			}
      }
	}

   render() {
		// iF THE USER IS NOT SIGNED IN...
		if (this.props.user == null) {
			return(
				<LoginScreen {...this.props} />
			);
		}

      return (
        <Container>
			  	<Header style={Styles.accountHeader}>	
					<Right>
						<Button transparent>
							<Icon name='create' />
						</Button>
						<Button transparent
							onPress={() =>
							ActionSheet.show(
								{
									options: ['Log Out' , 'Cancel'],
									cancelButtonIndex: 1,
									title: "Options"
								},
								buttonIndex => {
									this.props.addAuthToken("")
									this.props.setUser(null)
									this.props.savePassword("")
									this.setState({
										menu: {
											selectedIndex: buttonIndex
										}
									});
								}
							)}
						>
							<Icon name='more' />
						</Button>
					</Right>
				</Header>
				<Header span style={Styles.accountHeader}>
					<Left>
					  <Thumbnail style={Styles.accountThumbnail} large source={{uri: "https://facebook.github.io/react-native/docs/assets/favicon.png"}} />
					</Left>
					<Body style={Styles.accountTitle}>
						<Title>Name</Title>
						<Subtitle>Filler</Subtitle>
					</Body>
				</Header>
        </Container>
      )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Account)