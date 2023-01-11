import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter App',
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Flutter App'),
      ),
      body: Column(
        children: <Widget>[
          // Card 配下の Text を Container でラップする（Card は子ウィジェットの大きさに依存）
          Card(
            color: Colors.blue,
            child: Container(
              width: double.infinity,
              child: Text('CHART!'),
            ),
            elevation: 5,
          ),
          // Card を Container でラップする
          Container(
            width: double.infinity,
            child: Card(
              child: Text('LIST OF TX'),
            ),
          ),
        ],
      ),
    );
  }
}
