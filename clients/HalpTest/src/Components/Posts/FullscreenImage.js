import React, { Component } from 'react';
import { Image, ScrollView } from 'react-native';

// Import the styles and themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

export default class Web extends Component {
  render() {
    console.log(this.props.navigation.state.params.photosrc);
    return (
      <ScrollView>
        <Image source={this.props.navigation.state.params.photosrc} style={Styles.fullImage}/>
      </ScrollView>
    );
  }
}
