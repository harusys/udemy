import 'package:flutter/material.dart';
import 'package:meals_app/screens/main_drawer.dart';

class FilterScreen extends StatelessWidget {
  static const routeName = '/filters';

  const FilterScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Your Filters!'),
      ),
      drawer: MainDrawer(),
      body: Center(
        child: Text('Filters!'),
      ),
    );
  }
}
