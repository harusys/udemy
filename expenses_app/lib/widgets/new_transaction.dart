import 'package:flutter/material.dart';

class NewTransaction extends StatelessWidget {
  // String? titleInput;
  // String? amountInput;
  final titleController = TextEditingController();
  final amountController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Card(
      // Card はマージンしか定義できないため、パディングを使うために Contianer を挟む
      elevation: 5,
      child: Container(
        padding: EdgeInsets.all(10),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.end,
          // 入力フォームを作る
          children: <Widget>[
            TextField(
              decoration: InputDecoration(labelText: 'Title'),
              controller: titleController,
              // onChanged: (val) => titleInput = val,
            ),
            TextField(
              decoration: InputDecoration(labelText: 'Amount'),
              controller: amountController,
              // onChanged: (val) => amountInput = val,
            ),
            TextButton(
              child: Text('Add Transaction'),
              onPressed: () {
                // print(titleInput);
                // print(amountInput);
                print(titleController.text);
                print(amountController.text);
              },
            ),
          ],
        ),
      ),
    );
  }
}
