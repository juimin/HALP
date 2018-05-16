// Import default react components
import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';

// Import Halp Components
import Search from '../Search/Search';

// Generate a stack for navigation
// Generally, this is the component that wraps the child components
// Specifically for this file, App.js will use this as a component because it allows for
// navigating between the Compoents listed
const RootStack = StackNavigator({
      Search: {
         screen: Search,
         navigationOptions: {
            title: "Search"
         }
      }
   },
   {
      initialRouteName: 'Search',
      headerMode: 'none',
   },
);

export default RootStack
