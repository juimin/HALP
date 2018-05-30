// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { ButtonGroup } from 'react-native-elements';
import { ScrollView } from 'react-native';
import {
	View,
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

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import ReduxActions from '../../Redux/Actions';

//Import HALP components
import SubscribeButton from './SubscribeButton';
import CompactPost from '../Posts/CompactPost';
import LargePost from '../Posts/LargePost';


const mapStateToProps = (state) => {
	return {
      authToken: state.AuthReducer.authToken,
      user: state.AuthReducer.user,
      password: state.AuthReducer.password,
      activeBoard: state.BoardReducer.activeBoard
	}
}

class Board extends Component {
	constructor(props) {
		super(props)
		this.state = {
			  menu: {
				  selectedIndex: -1
			  },
			  subscribed: this.isSubscribed(),
			  subscribers: this.props.activeBoard.subscribers,
			  posts: []
		}
	}

	componentWillMount = () => {
		this.fetchPosts();
	}

	isSubscribed = () => {
		if (!this.props.user) {
			return null
		}
		return this.props.user.favorites.includes(this.props.activeBoard.id);
    }

	returnData = (sub) => {
		this.setState({subscribed: sub});
		sub ? this.setState({subscribers: this.state.subscribers++}) : this.setState({subscribers: this.state.subscribers--});
		console.log(this.state.subscribers);
	}

	fetchPosts = () => {
		var x = fetch('https://staging.halp.derekwang.net/posts/get/board?id=' + this.props.activeBoard.id, {
			method: 'GET',
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json',
			}
		}).then(response => {
			return response.json()
		}).then(data => {
			this.setState({posts: data})
		})
	}


	render() {

		let postItems = this.state.posts.reverse().map( (post, i) => {
			console.log(post);
			return <LargePost key={i} post={post} />
		  });

		return (
			<ScrollView>
			  	<Header style={Styles.boardHeader}>	
					<Right>
						<SubscribeButton user={this.props.user} board={this.props.activeBoard} subbed={this.isSubscribed()} authToken={this.props.authToken} returnData={this.returnData.bind(this)}/>
						<Button transparent>
							<Icon name='create' />
						</Button>
						<Button transparent
							onPress={() =>
							ActionSheet.show(
								{
									options: ['Cancel'],
									cancelButtonIndex: 0,
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
				<Header span style={Styles.boardHeader}>
					<Left>
					  <Thumbnail style={Styles.accountThumbnail} large source={{uri: "https://facebook.github.io/react-native/docs/assets/favicon.png"}} />
					</Left>
					<Body style={Styles.accountTitle}>
						<Title>{this.props.activeBoard.title}</Title>
						{/* <Right style={Styles.boardSubButton}>
							<SubscribeButton user={this.props.user} board={this.props.activeBoard} subbed={this.isSubscribed()} authToken={this.props.authToken} returnData={this.returnData.bind(this)}/>
						</Right>  */}
						<Subtitle style={Styles.boardSubs}>{this.state.subscribers} subscribers</Subtitle>
						<Subtitle style={Styles.boardDesc}>{this.props.activeBoard.description}</Subtitle>
					</Body>
				</Header>
				<View>
					{postItems}
				</View>
		</ScrollView>
			
		)
	}
}

export default connect(mapStateToProps)(Board)