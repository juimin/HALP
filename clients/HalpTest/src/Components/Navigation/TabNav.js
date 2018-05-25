// TabNav defines the Tab Navigator that we need when using the navigation bar at the bottom
// of the app. The actual content of this is in the TabBar Component

import { View } from 'react-native';
import { TabNavigator } from 'react-navigation';

// Import HALP Components
import HomeNav from './HomeNav';
import SearchNav from './SearchNav';
import AccountNav from './AccountNav';
import SettingsNav from './SettingsNav';
import TabBar from '../TabBar/TabBar';
import NewPost from '../NewPost/NewPost';

export default TabNavigator(
   {
      HomeNav: {
         screen: HomeNav,
      },
      SearchNav: {
         screen: SearchNav,
      },
      NewPost: {
         screen: View,
      },
      AccNav: {
         screen: AccountNav,
      },
      SettingsNav: {
         screen: SettingsNav,
      },
   }, 
   {
      tabBarPosition: 'bottom',
      tabBarComponent: TabBar,
   }
);