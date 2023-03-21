import 'package:flutter/material.dart';
import 'package:meals_app/models/meal.dart';

class FavoritesScreen extends StatelessWidget {
  final List<Meal> favoriteMeals;

  const FavoritesScreen({
    super.key,
    required this.favoriteMeals,
  });

  @override
  Widget build(BuildContext context) {
    return const Center(
      child: Text('The Favorites'),
    );
  }
}
