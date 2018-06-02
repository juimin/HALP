import React, { Component } from 'react';
import { ScrollView, Switch, Slider, TouchableHighlight, View } from 'react-native';
import { Container, Header, Content, List, ListItem, CheckBox, Text, Body, Button } from 'native-base';
import { TabNavigator, StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import Themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme'

// Export the Component
export default class Settings extends Component {
   constructor(props) {
      super(props)
      this.state = {
         setting1: false,
         setting2: false,
         setting3: false,
      }
   }
   
   
   render() {
      return (
         <ScrollView >
            <Text style={Styles.settingTitle}>Settings</Text>
               <Content>
               <ListItem>
                  <CheckBox checked={this.state.setting1} onPress={() => {this.setState({
                  setting1: !this.state.setting1,
                  setting2: this.state.setting2,
                  setting3: this.state.setting3
               })}}/>
                  <Body>
                     <Text>Setting 1</Text>
                  </Body>
               </ListItem>
               </Content>
            <Text style={Styles.settingTitle}>Other Settings</Text>
            <Container>
                <Content>
                <List>
                    <ListItem onPress={() => this.props.navigation.navigate('Web', {uri: 'https://github.com/JuiMin/HALP'})}>
                    <Text>GitHub</Text>
                    </ListItem>
                    <ListItem onPress={() => this.props.navigation.navigate('Web', {uri: 'https://halpapp.github.io'})}>
                    <Text>Project Website</Text>
                    </ListItem>
                    <ListItem onPress={() => this.props.navigation.navigate('About')}>
                    <Text>About</Text>
                    </ListItem>
                </List>
                </Content>
            </Container>
         </ScrollView>
      )
   }
}