import React, { Component } from 'react';
import { ScrollView, Text, View } from 'react-native';

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
         subscriptions: [
            {
               title: "Subscription 1"
            },
            {
               title: "Subscription 2"
            },
            {
               title: "Subscription 3"
            }
         ],
         searchTerm: ""
      }

      // Bind the function search
      this.search = this.search.bind(this);
   }

   // Define the function that perform the search
   search(input) {
      // Should update the list when called
      // TODO: PERFORM THE SEARCH USING THE API

      // SET THE STATE WITH SEARCH RESULTS
      this.setState({
         results: this.state.results.concat([
            { 
               title: input
            }
         ]),
         searchTerm: input
      });
   }

   render() {
      if (this.state.searchTerm.length == 0) {
         return(
            <View style={Styles.searchScreen}>
               <SearchBar 
                  showLoading
                  placeholder="Search"
                  lightTheme
                  onChangeText={this.search}
                  containerStyle={Styles.searchBar}
               />
               <Text>Subscriptions</Text>
               <ScrollView>
                  <List containerStyle={Styles.searchList} >
                     {
                        this.state.subscriptions.map((item, i) => (
                           <ListItem key={i} title={item.title} containerStyle={Styles.searchListItem}/>
                        ))
                     }
                  </List>
               </ ScrollView>
            </View>
         )
      } else {
         return (
            <View style={Styles.searchScreen}>
               <SearchBar 
                  showLoading
                  placeholder="Search"
                  lightTheme
                  onChangeText={this.search}
                  containerStyle={Styles.searchBar}
               />
               <ScrollView>
                  <Text style={Styles.searchTitle}>Search Results</Text>
                  <List containerStyle={Styles.searchList} >
                     {
                        this.state.results.map((item, i) => (
                           <ListItem key={i} title={item.title} containerStyle={Styles.searchListItem}/>
                        ))
                     }
                  </List>
               </ ScrollView>
            </View>
         )
      }
   }
}