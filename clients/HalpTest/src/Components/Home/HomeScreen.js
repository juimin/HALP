import React, { Component } from 'react';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import react-redux connect 
import { connect } from "react-redux";

// Import the different views based on user state
import GuestHome from './GuestHome';
import LargePost from '../Posts/LargePost';


import { 
   Container,
   Header,
   Body,
   Title,
   Right,
   Button,
	Content,
	Text,
   Picker,
   Card,
   Icon
} from 'native-base';

const mapStateToProps = state => {
   return {
		posts: state.PostReducer.posts,
		activePost: state.PostReducer.activePost
   };
};

class HomeScreen extends Component {
   constructor(props) {
      super(props)
      this.state = {
         pickerIndex: 0,
         maxPosts: 20
      }
      this.onValueChange = this.onValueChange.bind(this)
      this.increaseMaxPosts = this.increaseMaxPosts.bind(this)
   }

   onValueChange(value) {
      this.setState({
        pickerIndex: value
      });
    }

    increaseMaxPosts() {

      this.setState({
        pickerIndex: this.state.pickerIndex,
        maxPosts: this.state.maxPosts + 20
      })
	 }
	 
	 // Initial Post getter defaulting to default sorting
	 componentWillMount() {
		console.log("Mounting")
		// Gettin posts
	 }

    componentWillUpdate() {
      console.log(this.state.maxPosts)
    }

   // Here we should run initialization scripts
   render() {
      // This will be the same any user
      return (
         // <GuestHome {...this.props} />
         <Container>
            <Header style={{
               backgroundColor: Theme.colors.primaryBackgroundColor
            }}>
               <Body>
                  <Title style={{
                     color: Theme.colors.primaryColor,
                     alignSelf: "flex-end",
                     fontWeight: "bold"
                  }}>HALP</Title>
               </Body>
               <Right>
                  <Button transparent>
                     <Icon style={{color: 'gray'}} name='more'/>
                  </Button>
               </Right>
            </Header>
            <Content>
              <Picker
                mode="dropdown"
                selectedValue={this.state.pickerIndex}
                onValueChange={this.onValueChange}
                style={{width: "40%"}}
              >
                <Picker.Item label="New" value={0} />
                <Picker.Item label="Top" value={1} />
                <Picker.Item label="Comments" value={2} />
              </Picker>
              <Content>
                	{
                    	this.props.posts.map((item, index) => {
								<LargePost post={item}/>
							})
               	}
              	</Content>
					<Button rounded style={Styles.button} onPress={this.increaseMaxPosts}>
						<Text>Get More Posts</Text>
					</Button>
            </Content>
         </Container>
      );
   }
}

export default connect(mapStateToProps)(HomeScreen)