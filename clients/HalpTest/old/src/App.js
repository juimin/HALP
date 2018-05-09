import React, { Component } from 'react';
import { Button, StyleSheet, View, Text, TouchableWithoutFeedback} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'
import HomeScreen from './src/HomeScreen';
import SignupScreen from './src/SignupScreen';
import LoginScreen from './src/LoginScreen';
import CanvasTest from './src/CanvasTest';
import HomeNav from './src/HomeNav';
import SearchNav from './src/SearchNav';
import AccNav from './src/AccNav';
import BoardNav from './src/BoardNav';
import SettingsNav from './src/SettingsNav';

const activeTintColor = '#F44336';
const inactiveTintColor = 'gray';
const styles = StyleSheet.create({
  tabBar: {
    height: 49,
    flexDirection: 'row',
    borderTopWidth: StyleSheet.hairlineWidth,
    borderTopColor: 'rgba(0, 0, 0, .4)',
    backgroundColor: '#FFFFFF',
  },
  tab: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
});

class TabBar extends Component {
  renderItem = (route, index) => {
    const {
      navigation,
      jumpToIndex,
    } = this.props;

    const isNewPost = route.routeName === 'NewPost';

    const focused = index === navigation.state.index;
    const color = focused ? activeTintColor : inactiveTintColor;
    const size = 30;
    if (route.routeName == "HomeNav") {
      return (
        <TouchableWithoutFeedback
          key={route.key}
          style={styles.tab}
          onPress={() => jumpToIndex(index)}
        >
          <View style={styles.tab}>
            <Icon style={{ color }} size={size} name="home"/>
          </View>
        </TouchableWithoutFeedback>
      );
    }
    if (route.routeName == "SearchNav") {
      return (
        <TouchableWithoutFeedback
          key={route.key}
          style={styles.tab}
          onPress={() => jumpToIndex(index)}
        >
          <View style={styles.tab}>
            <Icon style={{ color }} size={size} name="search"/>
          </View>
        </TouchableWithoutFeedback>
      );
    }
    if (route.routeName == "AccNav") {
      return (
        <TouchableWithoutFeedback
          key={route.key}
          style={styles.tab}
          onPress={() => jumpToIndex(index)}
        >
          <View style={styles.tab}>
            <Icon style={{ color }} size={size} name="person"/>
          </View>
        </TouchableWithoutFeedback>
      );
    }
    if (route.routeName == "SettingsNav") {
      return (
        <TouchableWithoutFeedback
          key={route.key}
          style={styles.tab}
          onPress={() => jumpToIndex(index)}
        >
          <View style={styles.tab}>
            <Icon style={{ color }} size={size} name="settings"/>
          </View>
        </TouchableWithoutFeedback>
      );
    }
    return (
        <TouchableWithoutFeedback
          key={route.key}
          style={styles.tab}
          onPress={() => isNewPost ? navigation.navigate('NewPostModal') : jumpToIndex(index)}
        >
          <View style={styles.tab}>
            <Icon style={{ color }} style={styles.ico} size={size + 10} name="add-circle"/>
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
      <View style={styles.tabBar}>
        {routes && routes.map(this.renderItem)}
      </View>
    );
  }
}

//replace this with actual new post stuff
const Screen = (props) => (
  <View style={{ flex: 1, backgroundColor: '#fff', alignItems: 'center', justifyContent: 'center' }}>
    <Text>{props.title} Screen</Text>
  </View>
);

const Tabs = TabNavigator({
  HomeNav: {
    screen: (props) => <HomeNav />,
  },
  SearchNav: {
    screen: (props) => <SearchNav />,
  },
  NewPost: {
    screen: View,
  },
  AccNav: {
    screen: (props) => <AccNav />,
  },
  SettingsNav: {
    screen: (props) => <SettingsNav />,
  },
}, {
  tabBarPosition: 'bottom',
  tabBarComponent: TabBar,
});

const NewPostStack = StackNavigator({
  NewPost: {
    screen: (props) => <Screen title="New Post" {...props} />,
    navigationOptions: ({ navigation }) => ({
      headerTitle: 'New Post',
      headerLeft: (
        <Button
          title="Cancel"
          // Note that since we're going back to a different navigator (CaptureStack -> RootStack)
          // we need to pass `null` as an argument to goBack.
          onPress={() => navigation.goBack(null)}
        />
      ),
    }),
  },
})

/*
 * We need a root stack navigator with the mode set to modal so that we can open the capture screen
 * as a modal. Defaults to the Tabs navigator.
 */
const RootStack = StackNavigator({
  Tabs: {
    screen: Tabs,
  },
  NewPostModal: {
    screen: NewPostStack,
    navigationOptions: {
      gesturesEnabled: false,
    },
  },
}, {
  headerMode: 'none',
  mode: 'modal',
});

export default RootStack;
