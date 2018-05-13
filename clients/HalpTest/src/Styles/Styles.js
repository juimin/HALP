// Define a central location for styles
// We can keep thematic elements the same using this
import { StyleSheet } from 'react-native';

// Default Thematic Coloring so you can use it in multiple objects
import Colors from './Colors';

// Generate the stylesheet
export default StyleSheet.create({
   // Define Component Specific Styling
   signup: {
      flex: 1,
      backgroundColor: Colors.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },

});