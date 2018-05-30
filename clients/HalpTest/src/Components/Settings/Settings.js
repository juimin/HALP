import React, { Component } from 'react';
import { ScrollView, Switch, Slider, TouchableHighlight, View } from 'react-native';
import { Container, Header, Content, List, ListItem, Text, Button } from 'native-base';
import { TabNavigator, StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import Themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme'

// Export the Component
export default class Settings extends Component {
   render() {
      return (
         <ScrollView >
            <Text style={Styles.settingTitle}>Toggle Settings</Text>
            <Switch />
            <Switch />
            <Switch />
            <Switch />
            <Text style={Styles.settingTitle}>Other Settings</Text>
            <Slider />
            <Slider />
            <Slider />
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