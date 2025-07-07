import 'package:auth_bloc/core/constant/app_color.dart';
import 'package:flutter/material.dart';

class CBtn extends StatelessWidget {
  const CBtn({
    super.key,
    required this.onTap,
    required this.btnName,
    this.btnStyle,
  });

  final VoidCallback onTap;
  final String btnName;
  final ButtonStyle? btnStyle;

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: onTap,
      style: btnStyle,
      child: Text(btnName, style: TextStyle(fontWeight: FontWeight.w700)),
    );
  }
}

class CustomBtnVariant {
  static Widget normal(VoidCallback onTap, String btnName) {
    return CBtn(onTap: onTap, btnName: btnName);
  }

  static Widget success(VoidCallback onTap, String btnName) {
    return CBtn(
      onTap: onTap,
      btnName: btnName,
      btnStyle: ButtonStyle(
        backgroundColor: WidgetStatePropertyAll(AppColor.green),
      ),
    );
  }

  static Widget danger(VoidCallback onTap, String btnName) {
    return CBtn(
      onTap: onTap,
      btnName: btnName,
      btnStyle: ButtonStyle(
        backgroundColor: WidgetStatePropertyAll(AppColor.red),
      ),
    );
  }
}
