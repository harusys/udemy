import 'package:flutter/material.dart';

class CategoryItem extends StatelessWidget {
  final String title;
  final Color color;

  const CategoryItem(
    this.title,
    this.color,
  );

  @override
  Widget build(BuildContext context) {
    return Container(
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
    );
  }
}
