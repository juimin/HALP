import React, { Component } from 'react';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

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

import { TouchableHighlight, TouchableWithoutHighlight } from 'react-native';

export default class CompactPost extends Component {
    constructor(props) {
        //props are primarily post objects
        super(props)
		this.state = {
			
        }
    }

    render() {
        
        let post = this.props.post;

        return(
            <Container key={post.id} onPress={() => console.log(post.id + "pressed")}>
			<Content>
                <Left><Thumbnail style={Styles.postThumb} large source={ post.image_url.length != 0 ? {uri: post.image_url} : {uri: 'https://facebook.github.io/react-native/docs/assets/favicon.png'}}/></Left>
                <View><Text style={Styles.compactPostText}>{post.title}</Text></View>
                <View><Text style={Styles.compactPostText}>{post.caption}</Text></View>
            </Content>
			</Container>
        )
    }
}