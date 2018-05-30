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

    hoursSince = (time) => {
        var original = new Date(time);
        var current = new Date();
        //console.log(current, original);
        //get difference in hours
        var hours = Math.round(Math.abs(current - original) / (60*60*1000));
        if (hours < 24) {
            return String(hours) + "h ago";
        }
        //over one day
        hours = Math.round(Math.abs(current - original) / (60*60*1000*24));
        return String(hours) + "d ago";
    }

    render() {
        
        let post = this.props.post;
        let photo = post.image_url.length != 0 ? {uri: post.image_url} : {uri: 'https://facebook.github.io/react-native/docs/assets/favicon.png'};

        return(
            <Container key={post.id} style={Styles.largePost}>
                <Content style={{padding:"3%"}}>
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
                        <Button transparent style={Styles.cardButton}>
                        <Icon active name="thumbs-up" />
                        <Text>{String(post.upvotes - post.downvotes)}</Text>
                        </Button>
                    </Left>
                    <Body>
                        <Button transparent style={Styles.cardButton}>
                        <Icon active name="chatbubbles" />
                        <Text>{String(post.comments.length)}</Text>
                        </Button>
                    </Body>
                    <Right>
                        <Text>{this.hoursSince(post.time_created)}</Text>
                    </Right>
                    </CardItem>
                </Card>
                </Content>
            </Container>
        )
    }
}