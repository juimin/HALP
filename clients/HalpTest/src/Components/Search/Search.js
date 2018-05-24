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
      // Bind the function search
      this.search = this.search.bind(this);
   }

   // Define the function that perform the search
   search(input) {
      // Should update the list when called
      // TODO: PERFORM THE SEARCH USING THE API
      // SET THE STATE WITH SEARCH RESULTS
      if (input.length == 0) {
         this.setState({
            searching: "Subscriptions",
            items: this.state.items
         });
      } else {
         this.setState({
            searching: "Results",
            items: {
               Subscriptions: this.state.items.Subscriptions,
               Results: this.state.items.Results.concat([{
                  title: input,
                  image_url: "data:image/png;base64,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"
               }])
            }
         });
      }
   }

   render() {
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
               <Text style={Styles.searchTitle}>{this.state.searching}</Text>
               <List containerStyle={Styles.searchList} >
                  {
                     this.state.items[this.state.searching].map((item, i) => (
                        <ListItem
                           roundAvatar
                           avatar={{uri:item.image_url}}
                           key={i}
                           title={item.title} 
                           containerStyle={Styles.searchListItem}
                           onPress={() => this.props.navigation.navigate('Board')}
                        />
                     ))
                  }
               </List>
            </ ScrollView>
         </View>
      )
   }
}