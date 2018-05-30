// This should be the root of all application components
// Everything runs under a stack navigation nexted from here

// Import required react components
import React, { Component } from 'react';
import { Button, View, Text, TouchableWithoutFeedback} from 'react-native';
import { StackNavigator, DrawerNavigator, TabNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

import { Root } from 'native-base';

// Import Navigation Components
import NewPostStack from '../Navigation/NewPostNav';
import Tabs from '../Navigation/TabNav';

// Import redux
import { createStore } from 'redux'
import { Provider } from 'react-redux'
import Reducer from '../../Redux/Reducer';

// Create the redux store
const store = createStore(Reducer)

/*
 * We need a root stack navigator with the mode set to modal so that we can open the capture screen
 * as a modal. Defaults to the Tabs navigator.
 */

const AppRootStack = StackNavigator({
   // Add the tab navigator as a modal?
   Tabs: {
      screen: Tabs,
      navigationOptions: {
         gesturesEnabled: false
      }
   },
   // Add this modal because it is separate from the main application navigation
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

class App extends Component {
   render() {
      return(
         <Provider store={store}>
            <Root>
               <AppRootStack />
            </Root>
         </Provider>
      );
   }

}

// Wrap the redux aroiund the app
export default App;