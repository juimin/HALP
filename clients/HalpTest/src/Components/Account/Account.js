// This page should be the account page for the user's account information

// Import react components
import React, { Component } from 'react';
import { Button, ScrollView, View, Text } from 'react-native';
import { TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import React Native Elements
import { ButtonGroup } from 'react-native-elements';

// Import the header
import AccountHeader from './AccountHeader';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

export default class Account extends Component {

   constructor()  {
      super()
      this.state = {
         selectedIndex: 0
      }
      this.switchSelection = this.switchSelection.bind(this)
   }

   switchSelection(selectedIndex) {
      // Set the index so that the component knows where we are
      this.setState({selectedIndex})
   }

   render() {
      const { selectedIndex } = this.state  
      return (
         <View>
            <AccountHeader />
            <ButtonGroup
               onPress={this.switchSelection}
               selectedIndex={selectedIndex}
               buttons={['Saved', 'Posts', 'Comments', 'History']}
               containerStyle={Styles.accountNavButtons}
            />
            <ScrollView>
               {
                  // Apend cards here
                }
            </ScrollView>
         </View>
      )
  }
}