// Import default react components
import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';

// Import Halp Components
import Settings from '../Settings/Settings';
import Web from '../Settings/Web';
import About from '../Settings/About';

// Generate a stack for navigation
// Generally, this is the component that wraps the child components
// Specifically for this file, App.js will use this as a component because it allows for
// navigating between the Compoents listed
const RootStack = StackNavigator({
      Settings: {
         screen: Settings,
         navigationOptions: {
            title: "Settings"
         }
      },
      Web: {
            screen: Web,
            navigationOptions: {
                  title: "Browser"
            }
      },
      About: {
            screen: About,
            navigationOptions: {
                  title: "About Us"
            }
      }
   },
   {
      initialRouteName: 'Settings',
      headerMode: 'screen',
   },
);

export default RootStack
