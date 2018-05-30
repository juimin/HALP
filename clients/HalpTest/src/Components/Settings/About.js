import React, { Component } from 'react';
import { View, Container, Header, Content, Text, Thumbnail } from 'native-base';

// Import Themes
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme'

export default class About extends Component {
  render() {
    return (
        <Container style={Styles.home}>
        <Container style={{
            marginTop: "25%",
            marginRight: "4%",
            width: "90%",
         }}>
            <Thumbnail large source={require("../../Images/Logo-09.png")} style={{ alignSelf: "center"}} />
            <Content>
                <Text style={{textAlign: 'center'}}>We are Team HALP.</Text>
            </Content>
        </Container>
        </Container>
    );
  }
}