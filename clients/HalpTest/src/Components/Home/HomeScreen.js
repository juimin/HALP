import React, { Component } from 'react';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Import react-redux connect 
import { connect } from "react-redux";

// Import the different views based on user state
import GuestHome from './GuestHome';

import { 
   Container,
   Header,
   Body,
   Title,
   Right,
   Button,
   Content,
   Picker,
   Card,
   Icon
} from 'native-base';

const mapStateToProps = state => {
   return {
       boards: state.BoardReducer.boards,
       activeBoard: state.BoardReducer.activeBoard
   };
};

class HomeScreen extends Component {
   constructor(props) {
      super(props)
      this.state = {
         pickerIndex: 0
      }

      this.onValueChange = this.onValueChange.bind(this)
   }

   onValueChange(value) {
      this.setState({
        pickerIndex: value
      });
    }

   // Here we should run initialization scripts
   render() {
      console.log(this.props.boards)
      // This will be the same any user
      return (
         // <GuestHome {...this.props} />
         <Container>
            <Header style={{
               backgroundColor: Theme.colors.primaryBackgroundColor
            }}>
               <Body>
                  <Title style={{
                     color: Theme.colors.primaryColor,
                     alignSelf: "flex-end",
                     fontWeight: "bold"
                  }}>HALP</Title>
               </Body>
               <Right>
                  <Button transparent>
                     <Icon style={{color: 'gray'}} name='more'/>
                  </Button>
               </Right>
            </Header>
            <Content>
            <Picker
              mode="dropdown"
              selectedValue={this.state.pickerIndex}
              onValueChange={this.onValueChange}
              style={{width: "40%"}}
            >
              <Picker.Item label="New" value={0} />
              <Picker.Item label="Top" value={1} />
              <Picker.Item label="Comments" value={2} />
            </Picker>
            <Content>
               <Card></Card>
            </Content>
            </Content>

         </Container>
      );
   }
}

export default connect(mapStateToProps)(HomeScreen)