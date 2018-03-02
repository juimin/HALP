import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';

export default class HomeScreen extends React.Component {
  constructor(props) {
    super(props);
    this.state = {loggedin: false};
  }

  render() {
  	const {goBack} = this.props.navigation;
    if (this.state.loggedin) {
    	return(
    		<View style={styles.container}>
    			<Text>Dashboard</Text>
    			<Button
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
      <View style={styles.container}>
        <Text></Text>
        <Button
          title="Log in"
          onPress={() => this.props.navigation.navigate('Login')}
        />
        <Text></Text>
         <Button
          title="Sign Up"
          onPress={() => this.props.navigation.navigate('Signup')}
        />
        <Text></Text>
        <Button
	        		title="Try Me"
	        		onPress={() => this.setState({loggedin: true})}
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