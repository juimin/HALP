// Import the Application bundle in the App component
// as well as react native app registry to run
import { AppRegistry } from 'react-native';
import App from './src/Components/App/App';

// Import redux
import { createStore } from 'redux'
import Reducer from './src/Redux/Reducer';

// Create the redux store
const store = createStore(Reducer)

// Register the HALP test component
AppRegistry.registerComponent(
   // 'HalpTest', () => <Provider store={store}>App</Provider>
   'HalpTest', () => App
);
