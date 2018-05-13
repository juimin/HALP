// This should be the root of all application components
// Everything runs under a stack navigation nexted from here

// Import required react components
import React, { Component } from 'react';
import { Button, View, Text, TouchableWithoutFeedback} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import HALP Components
import HomeScreen from '../Home/HomeScreen';
import HomeNav from '../Navigation/HomeNav';
import Search from '../Search/Search';
import Account from '../Account/Account';
import BoardNav from '../Board/BoardNav';
import Settings from '../Settings/Settings';

// Import Stylesheet and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

class TabBar extends Component {
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
          onPress={() => isNewPost ? navigation.navigate('NewPostModal') : jumpToIndex(index)}
        >
          <View style={Styles.navigationTab}>
            <Icon style={{ color }} size={size + 10} name="add-circle"/>
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
    screen: (props) => <Search />,
  },
  NewPost: {
    screen: View,
  },
  AccNav: {
    screen: (props) => <Account />,
  },
  SettingsNav: {
    screen: (props) => <Settings />,
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
