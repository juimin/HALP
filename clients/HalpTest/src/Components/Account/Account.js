// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { ScrollView, View} from 'react-native';
import { TabNavigator } from 'react-navigation';
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

// Import React Native Elements
import { ButtonGroup } from 'react-native-elements';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

const mapStateToProps = (state) => {
	return {
		authToken: state.AuthReducer.authToken,
		password: state.AuthReducer.password,
		user: state.AuthReducer.user
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
		if (this.props.user == null) {
			return(
				<Container style={Styles.home}>
					<Button rounded style={Styles.button} 
						onPress={() => this.props.navigation.navigate('Login')}
					>
						<Text>Log In</Text>
					</Button>
				</Container>
			
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
									console.log(this.state)
									this.setState({
										menu: {
											selectedIndex: buttonIndex
										}
									});
									console.log(this.state)
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

export default connect(mapStateToProps)(Account)