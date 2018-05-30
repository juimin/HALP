//This is the tab for a given user's saved AKA bookmarked posts

// Import the styles and themes
import Styles from '../../../Styles/Styles';
import Theme from '../../../Styles/Theme';

// Import react components
import React, { Component } from 'react';
import { ScrollView, View, Image } from 'react-native';

// Connect 
import { connect } from 'react-redux';

// Use Native Base
import { Text, Container, Content, Left, ActionSheet, Header, Card, Right, Body, CardItem, Thumbnail, Button, Icon } from 'native-base';

import { setActivePost } from '../../../Redux/Actions'; 
import { API_URL } from '../../../Constants/Constants' 

const mapStateToProps = (state) => {
	return {
		authToken: state.AuthReducer.authToken,
		password: state.AuthReducer.password,
		user: state.AuthReducer.user
	}
}

const mapDispatchToProps = dispatch => {
    return {
        setActivePost: post => dispatch(setActivePost(post))
    }
}

class Saved extends Component {

    constructor(props) {
        super(props)
        // Later you can just replace everything with user.bookmarks
        this.state = {
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

    render() {
        let post = this.props.post;
        // TODO: if (this.props.user.bookmarks == null) {
        if (this.state == null) {
            return (
                <Container>
                    <Content style={Styles.savedNothing}>
                        <View>
                            <Text>Bookmark some posts before you can see them here!</Text>
                        </View>
                    </Content>
                </Container>
            )
        } else {
            return (
                <Container style={{height: 330}}>
                    {/* <Header style={Styles.tabHeader} /> */}
                    <Content>
                    <ScrollView>
                        {this.state.testPosts.map(testPost =>
                            <Card style={Styles.tabCard}>
                                <CardItem style={{}}>
                                    <Left>
                                        <Body>
                                        <Text>{testPost.title}</Text>
                                        <Text note>{testPost.board_id} · {testPost.author_id} · {testPost.time_created}</Text>
                                        </Body>
                                    </Left>
                                    <Right>
                                        <Thumbnail square large source={{uri: testPost.image_url}} />
                                    </Right>
                                    </CardItem>
                                    <CardItem cardBody>
                                        <Image source={{uri: testPost.image_url}} style={Styles.tabImages}/>
                                    </CardItem>
                                    <CardItem>
                                </CardItem>
                                {/* <CardItem button onPress={() => {
                                    this.props.setActivePost(post);
                                    this.props.navigation.navigate("Post")}
                                }> */}
                                <CardItem style={{height:30}}>
                                    <Left>
                                        <View style={{flexDirection: "column"}}> 
                                            {/* <Text note>{String(post.upvotes - post.downvotes) + " Points"}</Text>
                                            <Text note>{String(post.comments.length) + " Comments"}</Text> */}
                                            <Text note>{"100" + " Points"}</Text>
                                            <Text note>{"150" + " Comments"}</Text>
                                        </View>
                                    </Left>
                                    <Right>
                                        <View style={{flexDirection: "row", width:'20%'}}>
                                            <Button transparent>
                                                <Icon active name="bookmark" />
                                            </Button>
                                            <Button transparent> 
                                                <Icon name="more" />
                                            </Button>
                                        </View>
                                    </Right>
                                </CardItem>
                            </Card>
                        )}
                    </ScrollView>
                    </Content>
                </Container>
            )
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Saved)
// export default Saved