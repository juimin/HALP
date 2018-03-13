import 'package:flutter/material.dart';

void main() => runApp(new DrawingWidget());

class DrawingWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      title: 'Drawing Widget',
      home: new Scaffold(
        appBar: new AppBar(
          title: new Text('Drawing'),
        ),
        body: new Center(
          child: new Text('Hello World'),
        ),
      ),
    );
  }
}