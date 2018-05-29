// Import required react components
import React, { Component } from 'react';
import { Button, View, Text, TouchableWithoutFeedback} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import Stylesheet and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

export default class TabBar extends Component {
    returnData2 = (board) => {
        console.log("success returning board to TabBar.js navigator:", board);
    }

    renderItem = (route, index) => {
      const {
         navigation,
         jumpToIndex,
      } = this.props;

      const isNewPost = route.routeName === 'NewPost';

      const focused = index === navigation.state.index;
      const color = focused ? Theme.colors.activeTintColor : Theme.colors.inactiveTintColor;
      const size = 30;
      if (route.routeName == "HomeNav") {
         return (
         <TouchableWithoutFeedback
            key={route.key}
            style={Styles.navigationTab}
            onPress={() => jumpToIndex(index)}
         >
            <View style={Styles.navigationTab}>
               <Icon style={{ color }} size={size} name="home"/>
            </View>
         </TouchableWithoutFeedback>
         );
      }
      if (route.routeName == "SearchNav") {
         return (
         <TouchableWithoutFeedback
            key={route.key}
            style={Styles.navigationTab}
            onPress={() => jumpToIndex(index)}
         >
            <View style={Styles.navigationTab}>
               <Icon style={{ color }} size={size} name="search"/>
            </View>
         </TouchableWithoutFeedback>
         );
      }
      if (route.routeName == "AccNav") {
         return (
         <TouchableWithoutFeedback
            key={route.key}
            style={Styles.navigationTab}
            onPress={() => jumpToIndex(index)}
         >
            <View style={Styles.navigationTab}>
               <Icon style={{ color }} size={size} name="person"/>
            </View>
         </TouchableWithoutFeedback>
         );
      }
      if (route.routeName == "SettingsNav") {
         return (
         <TouchableWithoutFeedback
            key={route.key}
            style={Styles.navigationTab}
            onPress={() => jumpToIndex(index)}
         >
            <View style={Styles.navigationTab}>
               <Icon style={{ color }} size={size} name="settings"/>
            </View>
         </TouchableWithoutFeedback>
         );
      }
      return (
         <TouchableWithoutFeedback
            key={route.key}
            style={Styles.navigationTab}
            onPress={() => isNewPost ? navigation.navigate('NewPostModal', {returnData2: this.returnData2.bind(this)}) : jumpToIndex(index)}
         >
            <View style={Styles.navigationTab}>
               <Icon style={{ color }} size={size + 7} name="add-circle"/>
            </View>
         </TouchableWithoutFeedback>
      );
  };

   render() {
      const {
         navigation,
      } = this.props;

      const {
         routes,
      } = navigation.state;

      return (
         <View style={Styles.navigationBar}>
         {routes && routes.map(this.renderItem)}
         </View>
      );
   }
}