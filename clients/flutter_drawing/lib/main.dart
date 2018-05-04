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
          child: new CustomPaint(
            painter: new Sky(),
            child: new Center(
              child: new Text(
                'Once upon a time...',
                style: const TextStyle(
                  fontSize: 40.0,
                  fontWeight: FontWeight.w900,
                  color: const Color(0xFFFFFFFF),
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}

class NavBar extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new BottomNavigationBar(
      
    );
  }
}

