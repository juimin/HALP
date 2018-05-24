// GuestHome describes the home screen seen by a guest user.
// A guest user should be defined as a user that has yet to create an account
// or is not yet loged in.

// Import React Components
import React, { Component } from 'react';
import { Button, View, Text, ScrollView } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'
import BoardTiles from './BoardTiles';

// Import react-redux connect 
import { connect } from "react-redux";
import { fetchBoards } from "../../Redux/ListBoardsOperation.js";
import { bindActionCreators } from 'redux';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

// Export the default class
// export default class GuestHome extends Component {

//      render() {
//       return(
//         /* <View style={Styles.home}>
//             <Button color={Theme.colors.primaryColor} title="Log in"
//             onPress={() => this.props.navigation.navigate('Login')}
//             />
//             <Button color={Theme.colors.primaryColor} title="Sign Up"
//             onPress={() => this.props.navigation.navigate('Signup')}
//             />
//             <Button color={Theme.colors.primaryColor} title="Try Me"
//             onPress={() => this.setState({loggedin: true})}
//             />
//             <Button color={Theme.colors.primaryColor} title="Canvas Test"
//             onPress={() => this.props.navigation.navigate('Canvas')}
//             />
//         </View> */
//         <ScrollView style={Styles.tileList}>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//             <BoardTiles/>
//         </ScrollView>
//       )
//    }
// }

class BoardsList extends Component {
    componentDidMount() {
      this.props.dispatch(fetchBoards());
    }
  
    render() {
      const { error, loading, boards } = this.props;
      
      if (error) {
        return <div>Error! {error.message}</div>;
      }
  
      if (loading) {
        return <div>Loading...</div>;
      }
  
      return (
        <ul>
          {boards.map(boards =>
            <li key={boards.id}>{boards.title}</li>
          )}
        </ul>
      );
    }
  }
  
  const mapStateToProps = state => ({
    products: state.boards.items,
    loading: state.boards.loading,
    error: state.boards.error
  });
  
  export default connect(mapStateToProps)(BoardsList);