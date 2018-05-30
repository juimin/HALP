import React, { Component } from 'react';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

import { Image } from 'react-native';
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
    Card,
    CardItem,
} from 'native-base';

import { TouchableHighlight, TouchableWithoutFeedback } from 'react-native';

export default class LargePost extends Component {
    constructor(props) {
        //props are primarily post objects
        super(props)
		this.state = {
			
        }
    }

    render() {
        
        let post = this.props.post;
        let photo = post.image_url.length != 0 ? {uri: post.image_url} : {uri: 'https://facebook.github.io/react-native/docs/assets/favicon.png'};

        return(
            <Container key={post.id} style={Styles.largePost}>
                <Content>
                <Card>
                    <CardItem>
                    <Left>
                        <Body>
                        <Text>{post.title}</Text>
                        <Text note>{post.caption}</Text>
                        </Body>
                    </Left>
                    </CardItem>
                    <CardItem cardBody>
                    <Image source={photo} style={{height: 200, width: null, flex: 1}}/>
                    </CardItem>
                    <CardItem>
                    <Left>
                        <Button transparent>
                        <Icon active name="thumbs-up" />
                        <Text>{String(post.upvotes - post.downvotes)} Likes</Text>
                        </Button>
                    </Left>
                    <Body>
                        <Button transparent>
                        <Icon active name="chatbubbles" />
                        <Text>{String(post.comments.length)} Comments</Text>
                        </Button>
                    </Body>
                    <Right>
                        <Text>11h ago</Text>
                    </Right>
                    </CardItem>
                </Card>
                </Content>
            </Container>
        )
    }
}