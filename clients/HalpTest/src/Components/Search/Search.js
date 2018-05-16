import React, { Component } from 'react';
import { ScrollView, View } from 'react-native';

// Import themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import Elements
import { List, ListItem, SearchBar } from 'react-native-elements'

export default class Search extends Component {

   constructor(props) {
      super(props)
      // Set the initial state
      this.state = {
         results: [
            {
               title: "Potato 1"
            },
            {
               title: "Potato 2"
            }
         ],
         subscriptions: [],
         searchTerm: ""
      }

      // Bind the function search
      this.search = this.search.bind(this);
   }

   // Define the function that perform the search
   search(searchTerm) {
      // Should update the list when called
      // TODO: PERFORM THE SEARCH USING THE API

      // SET THE STATE WITH SEARCH RESULTS
      this.setState({
         results: this.state.results.concat([
            { title: searchTerm }
         ])
      });
   }

   render() {
      return (
         <View>
            <SearchBar 
               showLoading
               platform="android"
               placeholder="Search"
               lightTheme
               onChangeText={this.search}
               containerStyle={Styles.searchBar}
            />
            <ScrollView>
               <List containerStyle={Styles.searchList} >
                  {
                     this.state.results.map((item, i) => (
                        <ListItem key={i} title={item.title} />
                     ))
                  }
               </List>
            </ ScrollView>
         </View>
      )
   }
}