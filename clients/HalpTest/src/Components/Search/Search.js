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
import Requests from '../../Requests/Requests';

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
		setActiveBoard: (board) => { dispatch(ReduxActions.setActiveBoard(board))},
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
		// Bind the functions to this
		this.search = this.search.bind(this)
		this.load = this.load.bind(this)
   }

   componentWillMount() {
		// Check for user to see if we need to restore the session token
      if (this.props.authToken != "") {
			// Check for session expiration
			Requests.sessionExpired(this.props.authToken).then(response => {
				if (response != 202) {
					Requests.renewSession({
						email: this.props.user.email,
						password: this.props.password
					}).then(response => {
						if (response != null) {
							this.props.restoreToken(response)
						}
					}).catch(err => {
						console.log(err)
					})
				}
			}).catch(err => {
				console.log(err)
			})
			// Load Subscriptions
			this.load();
      }
	}

	load() {
		var items = [];
		if (this.props.user != null) {
			this.props.user.favorites.forEach((item, index) => {
				Requests.getBoard(item).then(board => {
					if (board != null) {
						items.push(board)
						this.setState({
							searchTerm: "",
							items: items
						})
					}
				});
			});
		}
	}
	
	search(text) {
		// Search
		if (text == "") {
			this.load()
		} else {
			// Perform the search
			var items = []
			Requests.searchBoard("BOARD", text, this.props.authToken).then(results => {
				if (results != null) {
					results.forEach((item, index) => {
						Requests.getBoard(item).then(board => {
							if (board != null) {
								items.push(board)
								this.setState({
									searchTerm: text,
									items: items
								});
							}
						});
					});
				} else {
					this.setState({
						searchTerm: text,
						items: []
					});
				}
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
                     this.state.items.map((item, i) => (
                        <ListItem
                           roundAvatar
                           avatar={{uri:item.image}}
                           key={i}
                           title={item.title} 
                           containerStyle={Styles.searchListItem}
                           onPress={() => {
										this.props.setActiveBoard(item)
										this.props.navigation.navigate('Board')
									}
									}
                        />
                     ))
                  }
               </List>
					<Text style={Styles.searchTitle}>{(this.state.items.length == 0) ? "No Boards Found": ""}</Text>
            </ ScrollView>
         </View>
      )
	}
}

export default connect(mapStateToProps, mapDispatchToProps)(Search)