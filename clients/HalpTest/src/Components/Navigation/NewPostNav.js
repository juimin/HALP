// Import default react components
import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';
import { Button } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import HALP Components
import NewPost from '../NewPost/NewPost';
import CanvasTest from '../Canvas/CanvasTest';

// Import Stylesheet and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Export the Stack Navigator as a component
export default StackNavigator({
//    NewPost: {
//       screen: (props) => <NewPost title="New Post"/>,
//       navigationOptions: ({ navigation }) => ({
//          headerTitle: 'New Post',
//          headerLeft: (
//             // Add the Icon for canceling the Home
//             <Icon name="close" onPress={() => navigation.goBack(null)}/>
//          ),
//          headerRight: (
//             // Add the Icon for canceling the Home
//             <Button title="POST" onPress={() => navigation.goBack(null)}></Button>
//          ),
//       }),
//    },
//     Canvas: {
//         screen: CanvasTest,
//     },
    NewPost: {
        screen: NewPost,
        navigationOptions: ({ navigation }) => ({
            headerTitle: 'New Post',
            headerLeft: (
            // Add the Icon for canceling the Home
            <Icon size={25} name="close" onPress={() => navigation.goBack(null)}/>
            ),
            // headerRight: (
            // <Button title="POST" onPress={() => navigation.goBack(null)}></Button>
            // ),
        }),
    },
    Canvas: {
        screen: CanvasTest,
    }
},
{
    initialRouteName: 'NewPost',
    headerMode: 'screen',
 },
)