import 'package:auth_bloc/core/constant/app_color.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

ThemeData everblushTheme = ThemeData(
  brightness: Brightness.dark,
  scaffoldBackgroundColor: AppColor.background,
  primaryColor: AppColor.blue,
  hintColor: AppColor.cyan,
  appBarTheme: const AppBarTheme(
    backgroundColor: AppColor.background,
    foregroundColor: AppColor.foreground,
  ),
  textTheme: GoogleFonts.jetBrainsMonoTextTheme(
    const TextTheme(
      bodyMedium: TextStyle(color: AppColor.foreground),
      bodyLarge: TextStyle(color: AppColor.foreground),
    ),
  ),
  colorScheme: const ColorScheme.dark(
    primary: AppColor.blue,
    secondary: AppColor.cyan,
    surface: AppColor.background,
    error: AppColor.red,
    onPrimary: AppColor.foreground,
    onSecondary: AppColor.foreground,
    onSurface: AppColor.foreground,
    onError: AppColor.background,
  ),
  elevatedButtonTheme: ElevatedButtonThemeData(
    style: ElevatedButton.styleFrom(
      backgroundColor: AppColor.blue,
      foregroundColor: AppColor.background,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
    ),
  ),
);
