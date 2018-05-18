//BoardTiles describes the individual boards of the home page.

// Import React Components
import React, { Component } from 'react';
import { Button, View, Text, ScrollView } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

class BoardTiles extends Component {
    render() {
       return(
          //<ScrollView style={Styles.home}>
            <ScrollView style={Styles.eachTile}>
                <View style={Styles.home}>
                    <Text>Dashboard testing SDFSDFSDFSD</Text>
                    <Text>Good Stuff</Text>
                </View>
            </ScrollView>
       )
    }
 }
 
 export default BoardTiles