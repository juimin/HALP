//https://github.com/keshavkaul/react-native-sketch-view

import React, { Component } from 'react';
import {
    View,
    Text,
    TouchableHighlight
} from 'react-native';
import SketchView from 'react-native-sketch-view';
import { StackNavigator } from 'react-navigation';

const sketchViewConstants = SketchView.constants;

const tools = {};

tools[sketchViewConstants.toolType.pen.id] = {
    id: sketchViewConstants.toolType.pen.id,
    name: sketchViewConstants.toolType.pen.name,
    nextId: sketchViewConstants.toolType.eraser.id
};
tools[sketchViewConstants.toolType.eraser.id] = {
    id: sketchViewConstants.toolType.eraser.id,
    name: sketchViewConstants.toolType.eraser.name,
    nextId: sketchViewConstants.toolType.pen.id
};

class HandNote extends Component {

    constructor(props) {
        super(props);
        this.state = {
            toolSelected: sketchViewConstants.toolType.pen.id
        };
    }

    isEraserToolSelected() {
        return this.state.toolSelected === sketchViewConstants.toolType.eraser.id;
    }

    toolChangeClick() {
        this.setState({toolSelected: tools[this.state.toolSelected].nextId});
    }

    getToolName() {
        return tools[this.state.toolSelected].name;
    }

    onSketchSave(saveEvent) {
        this.props.onSave && this.props.onSave(saveEvent);
    }

    render() {
        return (
            <View style={{flex: 1, flexDirection: 'column'}}>
                <SketchView style={{flex: 1, backgroundColor: 'white'}} ref="sketchRef" 
                selectedTool={this.state.toolSelected} 
                onSaveSketch={this.onSketchSave.bind(this)}
                localSourceImagePath={this.props.localSourceImagePath}/>
				
                <View style={{ flexDirection: 'row', backgroundColor: '#EEE'}}>
                    <TouchableHighlight underlayColor={"#CCC"} style={{ flex: 1, alignItems: 'center', paddingVertical:20 }} onPress={() => { this.refs.sketchRef.clearSketch() }}>
                        <Text style={{color:'#888',fontWeight:'600'}}>CLEAR</Text>
                    </TouchableHighlight>
                    <TouchableHighlight underlayColor={"#CCC"} style={{ flex: 1, alignItems: 'center', paddingVertical:20, borderLeftWidth:1, borderRightWidth:1, borderColor:'#DDD' }} onPress={() => { this.refs.sketchRef.saveSketch() }}>
                        <Text style={{color:'#888',fontWeight:'600'}}>SAVE</Text>
                    </TouchableHighlight>
                    <TouchableHighlight underlayColor={"#CCC"} style={{ flex: 1, justifyContent:'center', alignItems: 'center', backgroundColor:this.isEraserToolSelected() ? "#CCC" : "rgba(0,0,0,0)" }} onPress={this.toolChangeClick.bind(this)}>
						<Text style={{color:'#888',fontWeight:'600'}}>ERASER</Text>
                    </TouchableHighlight>
                </View>
            </View>
        );
    }
}

export default HandNote;