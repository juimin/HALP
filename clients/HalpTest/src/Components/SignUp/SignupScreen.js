// Import needed react dependancies
import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

// Import stylesheet
import Styles from '../../Styles/Styles';

export default class SignupScreen extends React.Component {
  render() {
    const {goBack} = this.props.navigation;
    return (
      <View style={Styles.signup}>
        <Text>Sign Up Here! It worked</Text>
        <Button color = "#F44336"
	        		title="Go Back"
	        		onPress={() => goBack()}
	        	/>
      </View>
    );
  }
}