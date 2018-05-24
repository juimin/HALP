import React from 'react';
import {
  StyleSheet,
  Text,
  View,
  Alert,
  Image,
  ImageBackground,
  TouchableHighlight,
} from 'react-native';
 
import RNSketchCanvas from '@terrylinla/react-native-sketch-canvas';
import {captureRef, captureScreen} from "react-native-view-shot";
import ImageResizer from 'react-native-image-resizer';
import HideableView from '../App/HideableView';
 

export default class CanvasTest extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      error: null,
      res: null,
      value: {
        format: "jpg",
        quality: 0.8,
        snapshotContentContainer: false
      },
      source: null,
      isHidden: false,

    };
  }
  
  captureScreenFunction=()=>{
    captureRef(this.refs["image"],
    {
      format: "jpg",
      quality: 1,
      // height: 1136,
      // width: 640,
    })
    .then(
      uri => this.props.navigation.state.params.returnData(uri), 
      console.log("saving"), 
      this.props.navigation.goBack(),
      error => console.error("Oops, Something Went Wrong", error)
    );
  }

render() {
    return (
      <View style={styles.container}>
        <View
        ref="image" collapsable={false}>
          <ImageBackground
          source={this.props.navigation.state.params.source}
          style={[styles.container,]}
          imageStyle={{height: '100%',}}>
          <RNSketchCanvas
            containerStyle={{ backgroundColor: 'transparent', flex: 1 }}
            canvasStyle={{ backgroundColor: 'transparent', flex: 1 }}
            defaultStrokeIndex={0}
            defaultStrokeWidth={5}
            undoComponent={<HideableView hide={this.state.isHidden} style={styles.functionButton}><Text style={{color: 'white'}}>Undo</Text></HideableView>}
            clearComponent={<HideableView hide={this.state.isHidden} style={styles.functionButton}><Text style={{color: 'white'}}>Clear</Text></HideableView>}
            strokeComponent={color => (
              <View style={[{ backgroundColor: color }, styles.strokeColorButton, this.state.isHidden && styles.hidden]} />
            )}
            strokeSelectedComponent={(color, index, changed) => {
              return (
                <View style={[{ backgroundColor: color, borderWidth: 2 }, styles.strokeColorButton]} />
              )
            }}
            strokeWidthComponent={(w) => {
              return (<HideableView hide={this.state.isHidden} style={styles.strokeWidthButton}>
                <View style={{
                  backgroundColor: 'white', marginHorizontal: 2.5,
                  width: Math.sqrt(w / 3) * 10, height: Math.sqrt(w / 3) * 10, borderRadius: Math.sqrt(w / 3) * 10 / 2
                }} />
              </HideableView>
            )}}
            saveComponent={<TouchableHighlight onPress={() => {
              this.setState({isHidden: true}, this.captureScreenFunction);
              }}><HideableView hide={this.state.isHidden} style={styles.functionButton}><Text style={{color: 'white'}}>Done</Text></HideableView></TouchableHighlight>}
            savePreference={() => {
              return {
                folder: 'RNSketchCanvas',
                filename: String(Math.ceil(Math.random() * 100000000)),
                transparent: false,
                imageType: 'png'
              }
            }}
          />
          </ImageBackground>
        </View>
      </View>
    );
  }
}
 
const styles = StyleSheet.create({
  container: {
    alignSelf: "stretch", flex: 1, justifyContent: 'center', alignItems: 'center', backgroundColor: '#F5FCFF',
  },
  strokeColorButton: {
    marginHorizontal: .05, marginVertical: 8, width: 30, height: 30, borderRadius: 15,
  },
  strokeWidthButton: {
    marginHorizontal: 2.5, marginVertical: 8, width: 30, height: 30, borderRadius: 15,
    justifyContent: 'center', alignItems: 'center', backgroundColor: '#F44336'
  },
  functionButton: {
    marginHorizontal: 2.5, marginVertical: 8, height: 30, width: 60,
    backgroundColor: '#F44336', justifyContent: 'center', alignItems: 'center', borderRadius: 5,
  },
  backgroundImage: {
    flex: 1,
    alignSelf: 'stretch',
    width: null,
  }
});
 