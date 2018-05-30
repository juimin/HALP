import React, { Component } from 'react';
import { Image, ScrollView } from 'react-native';

export default class Web extends Component {
  render() {
    console.log(this.props.navigation.state.params.photosrc);
    return (
      <ScrollView>
        <Image source={this.props.navigation.state.params.photosrc} style={{height: 455, width: null, flex: 1}}/>
      </ScrollView>
    );
  }
}
