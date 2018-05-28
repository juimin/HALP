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

import {setPictureSuccess} from '../../Redux/Actions';
 
import RNSketchCanvas from '@terrylinla/react-native-sketch-canvas';
import {captureRef, captureScreen} from "react-native-view-shot";
import ImageResizer from 'react-native-image-resizer';
import HideableView from '../Helper/HideableView';
import RNFetchBlob from 'react-native-fetch-blob'
global.Buffer = global.Buffer || require('buffer').Buffer

const mapStateToProps = (state) => {
  return {
    picture_success: state.PictureReducer.success
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    setSuccess: (setting) => { dispatch(setPictureSuccess(setting))}
  }
}


// aws-sdk to upload to digital ocean spaces
var AWS = require('aws-sdk/dist/aws-sdk-react-native'); 

// Configure client for use with Spaces
const spacesEndpoint = new AWS.Endpoint('nyc3.digitaloceanspaces.com');
const s3 = new AWS.S3({
    endpoint: spacesEndpoint,
    accessKeyId: 'E26M2XWYBF3XGE5PBPBE',
    secretAccessKey: 'Vw4bjoYa4kD8uWs2PKJPvITCzAWNioKsXGY1rcytOqw',

});
var myBucket = 'halp-staging';
//generates something that resembles bson id
var mongoObjectId = function () {
  var timestamp = (new Date().getTime() / 1000 | 0).toString(16);
  return timestamp + 'xxxxxxxxxxxxxxxx'.replace(/[x]/g, function() {
      return (Math.random() * 16 | 0).toString(16);
  }).toLowerCase();
};
var filename = String(mongoObjectId()) + '.jpg'; //change to post id? either way this will be the uploaded image filename
var uploadurl = 'https://' + myBucket + '.nyc3.digitaloceanspaces.com/'+ filename;

export default class CanvasTest extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      source: null,
      isHidden: false,
    };
  }

  upFile = (data) => {
    var buff = new Buffer(data, 'base64')
    console.log("uploading2...")
    var params = {ACL: "public-read", Body: buff, Bucket: myBucket, Key: filename, 
      Metadata: {
        'Content-Type': 'image/jpeg'
      }
    };
    s3.upload(params, function(err, data) {
      if (err) {
          console.log(err)
      } else {
          console.log("Successfully uploaded data to halp-staging");
          //console.log(data.Location);
        }
    });
  }

  uploadObject = (uri) => {
    //this.props.navigation.state.params.returnData(uri);
    console.log("uploading...");
    //console.log(this.props);
    RNFetchBlob.fs.readFile(uri, 'base64').then(data => this.upFile(data));
    //console.log(uploadurl);
    this.props.navigation.state.params.returnData(uri, uploadurl);
  }
  
  captureScreenFunction=()=>{
    captureRef(this.refs["image"],
    {
      format: "jpg",
      quality: 1,
      // height: 1136,
      // width: 640,
      // result: "base64",
    })
    .then(
      //uri => this.props.navigation.state.params.returnData(uri),
      uri => this.uploadObject(uri),
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
            saveComponent={<HideableView hide={this.state.isHidden}><TouchableHighlight onPress={() => {
              this.setState({isHidden: true}, this.captureScreenFunction);
              }}><View style={styles.functionButton}><Text style={{color: 'white'}}>Done</Text></View></TouchableHighlight></HideableView>}
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
 