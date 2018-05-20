// TabNav defines the Tab Navigator that we need when using the navigation bar at the bottom
// of the app. The actual content of this is in the TabBar Component

import { View } from 'react-native';
import { TabNavigator } from 'react-navigation';

// Import HALP Components
import HomeNav from '../Navigation/HomeNav';
import Search from '../Search/Search';
import Account from '../Account/Account';
import BoardNav from '../Board/BoardNav';
import Settings from '../Settings/Settings';
import TabBar from '../TabBar/TabBar';
import NewPost from '../NewPost/NewPost';

export default TabNavigator(
   {
      HomeNav: {
         screen: HomeNav,
      },
      SearchNav: {
         screen: Search,
      },
      NewPost: {
         screen: View,
      },
      AccNav: {
         screen: Account,
      },
      SettingsNav: {
         screen: Settings,
      },
   }, 
   {
      tabBarPosition: 'bottom',
      tabBarComponent: TabBar,
   }
);