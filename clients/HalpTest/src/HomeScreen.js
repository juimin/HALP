import React from 'react';
import { Button, StyleSheet, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
//import BottomNavigation, { Tab } from 'react-native-material-bottom-navigation'
import Icon from 'react-native-vector-icons/MaterialIcons'

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
      <View style={styles.container}>
        <Text></Text>
        <Button 
          color = "#F44336"
          title="Log in"
          onPress={() => this.props.navigation.navigate('Login')}
        />
        <Text></Text>
         <Button 
          color = "#F44336"
          title="Sign Up"
          onPress={() => this.props.navigation.navigate('Signup')}
        />
        <Text></Text>
        <Button 
              color = "#F44336"
              title="Try Me"
              onPress={() => this.setState({loggedin: true})}
            />
        <Text></Text>
        <Button
              color = "#F44336"
              title="Canvas Test"
              onPress={() => this.props.navigation.navigate('Canvas')}
            />
      
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
});

/*
<BottomNavigation
        labelColor="white"
        rippleColor="white"
        backgroundColor="#F44336"
        style={{
          height: 56,
          elevation: 8,
          position: 'absolute',
          left: 0,
          bottom: 0,
          right: 0
        }}
        onTabChange={newTabIndex => alert(`New Tab at position ${newTabIndex}`)}
      >
        <Tab
          label="Home"
          icon={<Icon size={24} color="white" name="home" />}
        />
        <Tab
          label="Search"
          icon={<Icon size={24} color="white" name="search" />}
        />
        <Tab
          label="Boards"
          icon={<Icon size={24} color="white" name="add-circle" />}
        />
        <Tab
          label="Account"
          icon={<Icon size={24} color="white" name="account-circle" />}
        />
        
      </BottomNavigation>
      */
