// Import react components
import React, { Component } from 'react';
import { Button, ScrollView, View, Text } from 'react-native';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import React Native Elements
import { Avatar, Header } from 'react-native-elements';

export default class AccountHeader extends Component {

   // Render the component
   render() {
      return(
         <View  style={Styles.accountHeader}>
            <Header />
            <Avatar
                  size="xlarge"
                  rounded
                  source={{uri: "https://s3.amazonaws.com/uifaces/faces/twitter/adhamdannaway/128.jpg"}}
                  onPress={() => console.log("Works!")}
                  activeOpacity={0.7}
            />
            <View style={Styles.accountStatBar}>

               <Text>Tomato</Text>
            </View>
         </View>
      );
   }
}