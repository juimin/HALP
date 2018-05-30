// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { ScrollView, View, StyleSheet } from 'react-native';

// For the tabs
import TabPost from '../Posts/TabPost';

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
	ListItem,
	Grid,
	Col,
	Tabs,
	Tab,
	Card
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
			},
			test: {
				testPosts: [
					{
						id: '5b0e2dc93f33260001ab06ed',
						title: "Real Life DAMn",
						image_url: "http://s0.hulkshare.com/song_images/original/2/d/d/2dd00ab2a1e7d193ab7e6dc3bfae813f.jpg?dd=1388552400",
						caption: "I am the realest caption",
						author_id: '5b0e01ee00031000019fc400',
						comments: {good: "23423",bod: "242342",greatest: "242342"},
						board_id: '5b01b3017912ed0001434678',
						upvotes: 22,
						downvotes: 302,
						time_created: '2018-05-30T04:51:21.809Z',
						time_edited: '2018-05-30T04:51:21.809Z'
					},
					{
						id: '5b0e2dc93f33260001ab06ed',
						title: "Real Life DAMn",
						image_url: "http://s0.hulkshare.com/song_images/original/2/d/d/2dd00ab2a1e7d193ab7e6dc3bfae813f.jpg?dd=1388552400",
						caption: "I am the realest caption",
						author_id: '5b0e01ee00031000019fc400',
						comments: {good: "23423",bod: "242342",greatest: "242342"},
						board_id: '5b01b3017912ed0001434678',
						upvotes: 22,
						downvotes: 302,
						time_created: '2018-05-30T04:51:21.809Z',
						time_edited: '2018-05-30T04:51:21.809Z'
					},
					{
						id: '5b0e2dc93f33260001ab06ed',
						title: "Real Life DAMn",
						image_url: "http://s0.hulkshare.com/song_images/original/2/d/d/2dd00ab2a1e7d193ab7e6dc3bfae813f.jpg?dd=1388552400",
						caption: "I am the realest caption",
						author_id: '5b0e01ee00031000019fc400',
						comments: {good: "23423",bod: "242342",greatest: "242342"},
						board_id: '5b01b3017912ed0001434678',
						upvotes: 22,
						downvotes: 302,
						time_created: '2018-05-30T04:51:21.809Z',
						time_edited: '2018-05-30T04:51:21.809Z'
					},
					{
						id: '5b0e2dc93f33260001ab06ed',
						title: "Real Life DAMn",
						image_url: "http://s0.hulkshare.com/song_images/original/2/d/d/2dd00ab2a1e7d193ab7e6dc3bfae813f.jpg?dd=1388552400",
						caption: "I am the realest caption",
						author_id: '5b0e01ee00031000019fc400',
						comments: {good: "23423",bod: "242342",greatest: "242342"},
						board_id: '5b01b3017912ed0001434678',
						upvotes: 22,
						downvotes: 302,
						time_created: '2018-05-30T04:51:21.809Z',
						time_edited: '2018-05-30T04:51:21.809Z'
					},
					{
						id: '5b0e2dc93f33260001ab06ed',
						title: "Real Life DAMn",
						image_url: "http://s0.hulkshare.com/song_images/original/2/d/d/2dd00ab2a1e7d193ab7e6dc3bfae813f.jpg?dd=1388552400",
						caption: "I am the realest caption",
						author_id: '5b0e01ee00031000019fc400',
						comments: {good: "23423",bod: "242342",greatest: "242342"},
						board_id: '5b01b3017912ed0001434678',
						upvotes: 22,
						downvotes: 302,
						time_created: '2018-05-30T04:51:21.809Z',
						time_edited: '2018-05-30T04:51:21.809Z'
					}
				]
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
				<Header style={Styles.accountHeaderTwo}>
					<Left>
					  <Thumbnail style={Styles.accountThumbnail} large source={{uri: "https://facebook.github.io/react-native/docs/assets/favicon.png"}} />
					</Left>
					<Body style={Styles.accountTitle}>
						<Title>{this.props.user.firstName + " " + this.props.user.lastName}</Title>
						<Subtitle></Subtitle>
					</Body>
				</Header>
				<Content>
					<Grid style={Styles.statsBar} style={{borderTopWidth:1, borderTopColor:'gray'}}>
						<Col style={Styles.eachStat}>
							<Text>{this.props.user.favorites.length + this.props.user.bookmarks.length}</Text>
							<Text>Points {console.log("pp tape", this.props.user)}</Text>
						</Col>
						<Col style={Styles.eachStat}>
							<Text>{Object.keys(this.props.user.postvotes).length}</Text>
							<Text>Posts</Text>
						</Col>
						<Col style={Styles.eachStat}>
							<Text>{Object.keys(this.props.user.commentvotes).length}</Text>
							<Text>Comments</Text>
						</Col>
					</Grid>
					<Tabs initialPage={0} tabStyle={{color: '#f44336'}} style={{borderTopWidth:1, borderTopColor:'#D3D3D3'}}>
						<Tab heading="Saved" tabStyle={{backgroundColor: 'white'}} textStyle={{color: 'gray'}} activeTabStyle={{backgroundColor: 'white'}} activeTextStyle={{color: '#f44336', fontWeight: 'bold'}}>
							<View style={{overflow: 'scroll'}}>
								{
									this.state.test.testPosts.map((item, i) => {
										return <TabPost key={i} post={item} {...this.props}/>
									})
								}
							</View>
						</Tab>
						<Tab heading="Posts" tabStyle={{backgroundColor: 'white'}} textStyle={{color: 'gray'}} activeTabStyle={{backgroundColor: 'white'}} activeTextStyle={{color: '#f44336', fontWeight: 'bold'}}>
							<View style={{overflow: 'scroll'}}>
								{
									this.state.test.testPosts.map((item, i) => {
										return <TabPost key={i} post={item} {...this.props}/>
									})
								}
							</View>
						</Tab>
						<Tab heading="Comments" tabStyle={{backgroundColor: 'white'}} textStyle={{color: 'gray'}} activeTabStyle={{backgroundColor: 'white'}} activeTextStyle={{color: '#f44336', fontWeight: 'bold'}}>
							<View>
								{
									this.state.test.testPosts.map((item, i) => {
										return <TabPost key={i} post={item} {...this.props}/>
									})
								}
							</View>
						</Tab>
						<Tab heading="History" tabStyle={{backgroundColor: 'white'}} textStyle={{color: 'gray'}} activeTabStyle={{backgroundColor: 'white'}} activeTextStyle={{color: '#f44336', fontWeight: 'bold'}}>
							<View>
								{
									this.state.test.testPosts.map((item, i) => {
										return <TabPost key={i} post={item} {...this.props}/>
									})
								}
							</View>
						</Tab>
					</Tabs>
				</Content>
		</Container>
      )
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Account)