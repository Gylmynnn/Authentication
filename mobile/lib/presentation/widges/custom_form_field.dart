import 'package:auth_bloc/core/constant/app_color.dart';
import 'package:flutter/material.dart';

class CFormField extends StatelessWidget {
  const CFormField({
    super.key,
    this.initvalue,
    this.suffixIcon,
    this.hintText,
    required this.controller,
    this.secureText = false,
  });
  final TextEditingController controller;
  final String? initvalue;
  final bool secureText;
  final Widget? suffixIcon;
  final String? hintText;

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      controller: controller,
      onTapOutside: (PointerDownEvent e) {
        return FocusManager.instance.primaryFocus?.unfocus();
      },
      initialValue: initvalue,
      obscureText: secureText,
      obscuringCharacter: "*",
      style: TextStyle(color: AppColor.foreground, fontWeight: FontWeight.w700),
      decoration: InputDecoration(
        hintStyle: Theme.of(context).textTheme.bodyMedium!.copyWith(
          fontWeight: FontWeight.w700,
          color: AppColor.foreground.withAlpha(90),
        ),
        hintText: hintText,
        suffixIcon: suffixIcon,
        contentPadding: const EdgeInsets.symmetric(
          vertical: 14,
          horizontal: 14,
        ),
        filled: true,
        fillColor: AppColor.black,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide.none,
        ),
      ),
    );
  }
}
