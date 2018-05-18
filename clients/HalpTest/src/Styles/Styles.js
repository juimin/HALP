// Define a central location for styles
// We can keep thematic elements the same using this
import { StyleSheet } from 'react-native';

// Default Thematic Coloring so you can use it in multiple objects
import Theme from './Theme';
import GuestHome from '../Components/Home/GuestHome';

// Generate the stylesheet
export default StyleSheet.create({
   // Define Component Specific Styling
   home: {
      flex: 1,
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

   // Tile Scroll
   tileList: {
    flex: 1, 
    flexDirection: 'column',
    // What is the color for the background behind the tiles?
    backgroundColor: "#CACACA"
   },
   // Each Tile
   eachTile: {
    flex: 1,
    width: GuestHome.width - 10,
    margin: 7,
    height: 50,
    backgroundColor: "#FFFFFF"
   }
});