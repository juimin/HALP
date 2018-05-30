import React, { Component } from 'react';
import { WebView } from 'react-native';

export default class Web extends Component {
  render() {
    return (
      <WebView
        source={{uri: this.props.navigation.state.params.uri}}
      />
    );
  }
}