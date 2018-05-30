// GuestHome describes the home screen seen by a guest user.
// A guest user should be defined as a user that has yet to create an account
// or is not yet loged in.

// Import React Components
import React, { Component } from 'react';
import { Button, View, Text, ScrollView, Image } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'
import { Card } from 'react-native-elements';

// Import react-redux connect 
import { connect } from "react-redux";

//testing change back to fetchPosts
import { API } from "../../Redux/Api";
// import { bindActionCreators } from 'redux';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

const mapStateToProps = state => {
    return {
        boards: state.BoardReducer.boards,
        activeBoard: state.BoardReducer.activeBoard
    };
};

class HardPosts extends Component {
    render() {
        return (
           <View style={Styles.tileList}>
           <Card>
               <View style={Styles.eachTile}>
                    <Card title={this.props.title} />
                    <Image style={Styles.tileImage} resizeMode="cover" source={{ uri: this.props.image_url }} />
               </View>
           </Card> 
           </View>
        )
    }
}

class BoardsList extends React.Component {
    componentDidMount() {
       console.log("numero uno", API.items)
       //this.props.dispatch(API.state.apiCall());
    }
    
    render() {
      const { error, loading, boards } = this.props;
      return (
        <ScrollView>
            <HardPosts id={1} upvotes={200} downvotes={3} comments={54} title="How do I eat?" author_id="alexis" board_id="Food" time_created={22} image_url='https://food.fnr.sndimg.com/content/dam/images/food/fullset/2018/6/0/FN_snapchat_coachella_wingman%20.jpeg.rend.hgtvcom.616.462.suffix/1523633513292.jpeg' />
            <HardPosts id={2} upvotes={45} downvotes={200} comments={23} title="Why is my car on fire?" author_id="teeler" board_id="Auto" time_created={35} image_url="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSBxOhnyuaLSuWIHCG4fBqKrfwJEOZrMxAh-fTwCp1W_m-Eq5P7Uw" />
            <HardPosts id={3} upvotes={56} downvotes={12} comments={75} title="How hot is the hot potato?" author_id="jumbotron" board_id="Potatoes" time_created={40} image_url="https://scontent.fsea1-1.fna.fbcdn.net/v/t1.0-1/p200x200/14291794_1171930419548193_8185068610699574818_n.jpg?_nc_cat=0&oh=392bea7d0e4a1dc83b5bd4b45f230ec2&oe=5BC3F8E9&efg=eyJhZG1pc3Npb25fY29udHJvbCI6MCwidXBsb2FkZXJfaWQiOiI3OTAxMjc0Mjc3Mjg0OTYifQ%3D%3D" />
            <HardPosts id={4} upvotes={809} downvotes={7} comments={1} title="Where is Carmen San Diego?" author_id="rickDsanchez" board_id="Where Ya At?" time_created={67} image_url="https://images.rapgenius.com/78844aada7bacd1807df3c54a34462da.372x450x1.jpg" />
            <HardPosts id={5} upvotes={34} downvotes={5} comments={58} title="Where Waldo at?" author_id="mortymcfly" board_id="Where Ya At?" time_created={80} image_url="https://static1.squarespace.com/static/56438e3fe4b0c2d5ac1d4d26/5643945ae4b0eadf5c6537e9/5664d039e4b058c26c239306/1525966385398/maps_troy.jpg?format=1500w" />
        </ScrollView>
      )
    }
}
export default connect(mapStateToProps)(BoardsList);