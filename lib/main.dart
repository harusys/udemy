import 'package:flutter/material.dart';

import './question.dart';
import './answer.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    // TODO: implement createState
    return _MyAppState();
  }
}

class _MyAppState extends State<MyApp> {
  static const questions = [
    {
      'questionText': 'What\'s your favorite color?',
      'answer': ['くろ', 'あか', 'みどり', 'しろ'],
    },
    {
      'questionText': 'What\'s your favorite animal?',
      'answer': ['ウサギ', 'ヘビ', 'ゾウ', 'ライオン'],
    },
  ];

  var _questionIndex = 0;

  void _answerQuestion() {
    setState(() {
      _questionIndex += 1;
    });
    print(_questionIndex);
    if (_questionIndex < questions.length) {
      print('We have more questions!');
    } else {
      print('No more questions!');
    }
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text('My First App'),
        ),
        body: _questionIndex < questions.length
            ? Column(
                children: <Widget>[
                  Question(
                    questions[_questionIndex]['questionText'] as String,
                  ),
                  ...(questions[_questionIndex]['answer'] as List<String>)
                      .map((answer) {
                    return Answer(answer, _answerQuestion);
                  }).toList()
                ],
              )
            : Center(
                child: Text('You did it!'),
              ),
      ),
    );
  }
}
