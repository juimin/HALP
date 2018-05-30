// Define a central location for styles
// We can keep thematic elements the same using this
import { StyleSheet } from 'react-native';

// Default Thematic Coloring so you can use it in multiple objects
import Theme from './Theme';
import Styles from './Styles';

// Generate the stylesheet
export default StyleSheet.create({
   // Define Component Specific Styling
   home: {
      flex: 1,
      backgroundColor: Theme.colors.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },

   login: {
      flex: 1,
      backgroundColor: Theme.colors.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },

   signup: {
      flex: 1,
      backgroundColor: Theme.colors.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },

   // Navigation Bar from the default view
   navigationBar: {
      height: 49,
      flexDirection: 'row',
      borderTopWidth: StyleSheet.hairlineWidth,
      borderTopColor: 'rgba(0, 0, 0, .4)',
      backgroundColor: '#FFFFFF',
   },

   // Navigation Tabs
   navigationTab: {
      flex: 1,
      alignItems: 'center',
      justifyContent: 'center',
   },
   
   // New Post View
   newPostView: { 
        flex: 1, 
        // backgroundColor: '#fff', 
        alignItems: 'center', 
        justifyContent: 'center',
        padding: 5
    },

   searchScreen: {
      backgroundColor: Theme.colors.primaryBackgroundColor
   },

   // Search bar
   searchBar: {
      backgroundColor: Theme.colors.primaryBackgroundColor,
      height: 49,
      width: "100%",
      borderBottomColor: Theme.colors.primaryBackgroundColor
   },

   searchList: {
      backgroundColor: Theme.colors.primaryBackgroundColor,
      marginTop: 0,
      borderColor: Theme.colors.primaryBackgroundColor
   },

   searchListItem: {
      backgroundColor: Theme.colors.primaryBackgroundColor,
      borderBottomWidth: 0
   },

   searchTitle: {
      margin: 10,
      color: 'black'
   },

   accountNavButtons: {
      width: "100%",
      marginLeft: 0,
      marginTop: 0,
      padding: 0
   },

   accountHeader: {
      marginLeft: 0,
      width: "100%",
      padding: 0
   },

   accountStatBar: {
      height: "40%"
   },

   signinFormInput: {
      borderBottomColor: Theme.colors.inactiveTintColor
   },

   boardPicker: {
    height: 50,
    width: 200,
   },

   buttonTheme: {
    backgroundColor: Theme.colors.primaryColor,
    alignSelf: 'center',
    },

    closeIcon: {
        padding: 5
    },

   settingTitle: {
      margin: 10,
      fontSize: 20,
      color: 'black'
   },

   accountThumbnail: {
      marginLeft: "20%",
      marginTop: "25%"
   },

   accountTitle: {
      marginLeft: "5%",
   },

   accountHeader: {
      backgroundColor: Theme.colors.secondaryColor,
      borderBottomColor: Theme.colors.secondaryColor,
   },

   accountHeaderTwo: {
      backgroundColor: Theme.colors.secondaryColor,
      borderBottomColor: Theme.colors.secondaryColor,
      marginLeft: 0,
      width: "100%",
      padding: 0,
      height: 120,
   },

   statsBar: {
      //marginTop: "25%",
      flex: 1,
      flexDirection: 'row',
      width: "100%"
   },

   eachStat: {
    width: 50, 
    height: 50,
    backgroundColor: 'white',
    alignItems: 'center'
   },

   allTabs: {
    backgroundColor: 'white'
   },

   eachTab: {
    backgroundColor: 'white'
   },

    boardHeader: {
        marginLeft: 0,
        width: "100%",
        padding: 0,
        backgroundColor: Theme.colors.secondaryColor,
        borderBottomColor: Theme.colors.secondaryColor
    },

    boardSubs: {
        color: Theme.colors.secondaryTextColor
    },

    boardDesc: {
        textAlign: 'left'
    },

    boardSubButton: {
        alignSelf: 'flex-end',
    },

    subscribeButtonColor: {
        marginTop: "3%",
        backgroundColor: Theme.colors.primaryColor,
        marginLeft: "3%",
        height: "80%",
        justifyContent: 'center'
    },

    subscribeButton: {
        marginTop: "3%",
        marginLeft: "3%",
        backgroundColor: "gray",
        height: "80%",
        justifyContent: 'center'
    },

    postThumb: {
        width: 50,
        height: 50,
    },

    compactPostText: {
        textAlign: 'left',
    },

    cardButton: {
        width: '75%',
    },

    largePost: {
      height: 350,
    },

    fullImage: {
        height: 435,
        width: null,
        flex: 1
    },

    button: {
      backgroundColor: Theme.colors.primaryColor,
      alignSelf: "center",
   },

   savedNothing: {
     alignItems: 'center',
     width: "100%"
   },

   savedSomething: {
     alignItems: 'center',
     width: "100%"
   },

   tabImages: {
       width: null, 
       flex: 1
   },

   tabHeader: {

   },
   
   tabCard: {
      flex: 1,
      height: 170,
      width: "98%"
   }
   
});