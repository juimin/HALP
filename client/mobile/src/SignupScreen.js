import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

export default class SignupScreen extends React.Component {
  render() {
    const {goBack} = this.props.navigation;
    return (
      <View style={styles.container}>
        <Text>Sign Up Here!</Text>
        <Button
	        		title="Go Back"
	        		onPress={() => goBack()}
	        	/>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});