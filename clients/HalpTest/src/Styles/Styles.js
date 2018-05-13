// Define a central location for styles
// We can keep thematic elements the same using this
import { StyleSheet } from 'react-native';

// Default Thematic Coloring so you can use it in multiple objects
import Theme from './Theme';

// Generate the stylesheet
export default StyleSheet.create({
   // Define Component Specific Styling
   signup: {
      flex: 1,
      backgroundColor: Theme.primaryBackgroundColor,
      alignItems: 'center',
      justifyContent: 'center',
   },
});