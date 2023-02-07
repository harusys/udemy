import 'package:flutter/material.dart';
import 'package:meals_app/category_meals_screen.dart';
// import 'package:meals_app/category_meals_screen.dart';

class CategoryItem extends StatelessWidget {
  final String id;
  final String title;
  final Color color;

  const CategoryItem(
    this.id,
    this.title,
    this.color,
  );

  void selectCategory(BuildContext ctx) {
    Navigator.of(ctx).pushNamed(
      CategoryMealsScreen.routeName,
      arguments: {
        'id': id,
        'title': title,
      },
    );
    // Navigator.of(ctx).push(
    //   MaterialPageRoute(
    //     builder: (_) => CategoryMealsScreen(id, title),
    //   ),
    // );
  }

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () => selectCategory(context),
      splashColor: Theme.of(context).primaryColor,
      borderRadius: BorderRadius.circular(15),
      child: Container(
        padding: const EdgeInsets.all(15),
        child: Text(
          title,
          style: Theme.of(context).textTheme.titleMedium,
        ),
        decoration: BoxDecoration(
            // グラデーションをつける
            gradient: LinearGradient(
              colors: [
                // 最終値は透明度を上げる
                color.withOpacity(0.7),
                color,
              ],
              // 左上から右下にかけてグラデーション
              begin: Alignment.topLeft,
              end: Alignment.bottomRight,
            ),
            // 角を丸くする
            borderRadius: BorderRadius.circular(15)),
      ),
    );
  }
}
