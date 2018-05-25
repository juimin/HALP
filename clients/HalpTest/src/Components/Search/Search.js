import React, { Component } from 'react';
import { ScrollView, Text, View } from 'react-native';

// Import themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';
import { API_URL } from '../../Constants/Constants';

// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import ReduxActions from '../../Redux/Actions';

// Import Elements
import { List, ListItem, SearchBar } from 'react-native-elements'

// Import Requests
import { sessionExpired } from '../../Requests/Requests';

const mapStateToProps = (state) => {
	return {
      authToken: state.AuthReducer.authToken,
      user: state.AuthReducer.user,
      password: state.AuthReducer.password
	}
}

const mapDispatchToProps = (dispatch) => {
   return {
		restoreToken: (token) => { dispatch(ReduxActions.setTokenAction(token)) },
		logout: () => { dispatch(ReduxActions.logoutAction())}
   }
}

class Search extends Component {
   constructor(props) {
      super(props)
      // Generate a class variable for input, we don't need this in global state
		this.state = { 
			searchTerm: "",
			items: []
		}

		this.search = this.search.bind(this)
   }

   onComponentWillMount() {
      // Check for user to see if we need to restore the session token
      if (this.props.authToken != "") {
         // Check for session expiration
         if (sessionExpired(this.props.authToken)) {
				// Re-obtain a session
				fetch(API_URL + "/sessions", {
					method: "POST",
					headers: {
						'Accept': 'application/json',
						'Content-Type': 'application/json',
				  },
				  body: JSON.stringify({
					  email: this.props.user.email,
					  password: this.props.password
				  })
				}).then(response => {
					if (response.status == 202) {
						// Save token and password for later use
						this.props.addAuthToken(response.headers.get("authorization"))
						return response.json()
					} else {
						// Something went wrong with the server
						Alert.alert(
							'Session Error',
							'Error Acquiring Session',
							[ {text: 'OK', onPress: () => console.log('OK Pressed')} ]
						)
						return null
					}
				}).then(user => {
					if (user != null) {
						// Save the user to the thing
						this.props.setUser(user)
					} else {
						// Log out if we can't make a new session
						this.props.logout()
					}
				}).catch(err => {
					Alert.alert(
						'Error getting response from server',
						err,
						[ {text: 'OK', onPress: () => console.log('OK Pressed')} ]
					)
				})
         }
      }
	}
	
	search(text) {
		// Search
		var items = []
		if (text == "") {
			// Get Subscriptions
			if (this.props.user != null) {
				// Append each board
				this.props.user.favorites.forEach((item, index) => {
					fetch(API_URL + "/boards/single?id=" + item, {
						method: "GET",
						headers: {
							'Accept': 'application/json',
							'Content-Type': 'application/json',
					  }
					}).then(response => {
						if (response.status == 200) {
							return response.json()
						} else {
							return null
						}
					}).then(body => {
						if (body != null) {
							items.push(body)
						}
					}).catch(err => {
						console.log(err)
					})
				});
			}
			this.setState({
				searchTerm: text,
				items: items
			})
		} else {
			// Perform the search
			
			// Set the search term
			this.setState({
				searchTerm: text,
				items: items
			})
		}

	}

   render() {
      return (
         <View style={Styles.searchScreen}>
            <SearchBar 
               showLoading
               placeholder="Search"
               lightTheme
               onChangeText={(text) => this.search(text)}
               containerStyle={Styles.searchBar}
            />
            <ScrollView>
               <Text style={Styles.searchTitle}>{(this.state.searchTerm == "") ? "Subscriptions": "Results"}</Text>
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

export default connect(mapStateToProps)(Search)