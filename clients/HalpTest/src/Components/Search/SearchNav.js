import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'
import { COLOR, Toolbar } from 'react-native-material-ui';


export default class SearchNav extends React.Component {

  constructor(props) {
    super(props);
    this.state = { text: 'Useless Placeholder' };
  }

  static navigationOptions = {
    tabBarIcon: ({ tintColor }) => (<Icon size={28} name="search" style={{color:tintColor}}/>)
  }
 
  render() {
    return (
      <Toolbar
        uiTheme={uiTheme}
        leftElement="menu"
        centerElement="Searchable"
        searchable={{
          autoFocus: true,
          placeholder: 'Search',
        }}
      />
    );  
  }
}