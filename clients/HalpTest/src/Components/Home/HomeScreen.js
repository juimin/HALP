import React from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

export default class HomeScreen extends React.Component {
   constructor(props) {
      super(props);
      this.state = {loggedin: false};
   }

   render() {
      const {goBack} = this.props.navigation;
      if (this.state.loggedin) {
         return(
            <View style={Styles.home}>
               <Text>Dashboard</Text>
               <Button color = "#F44336"
                  title="Go Back"
                  onPress={() => {
                     this.setState({loggedin: false});
                  }}
               />
            </View>
         )
      }
      //if not logged in
      return (
         <View style={Styles.home}>
         <Text></Text>
         <Button 
            color={Theme.colors.primaryColor}
            title="Log in"
            onPress={() => this.props.navigation.navigate('Login')}
         />
         <Text></Text>
            <Button 
               color={Theme.colors.primaryColor}
            title="Sign Up"
            onPress={() => this.props.navigation.navigate('Signup')}
         />
         <Text></Text>
         <Button 
               color={Theme.colors.primaryColor}
               title="Try Me"
               onPress={() => this.setState({loggedin: true})}
               />
         <Text></Text>
         <Button
               color={Theme.colors.primaryColor}
               title="Canvas Test"
               onPress={() => this.props.navigation.navigate('Canvas')}
               />
         
         </View>
      );
   }
}