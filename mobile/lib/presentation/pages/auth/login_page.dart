import 'package:auth_bloc/core/constant/app_color.dart';
import 'package:auth_bloc/core/utils/helper/asset_loader.dart';
import 'package:auth_bloc/logic/auth/bloc/auth_bloc.dart';
import 'package:auth_bloc/logic/auth/bloc/auth_event.dart';
import 'package:auth_bloc/logic/auth/bloc/auth_state.dart';
import 'package:auth_bloc/presentation/pages/auth/register_page.dart';
import 'package:auth_bloc/presentation/pages/homepage/home_page.dart';
import 'package:auth_bloc/presentation/widges/custom_button.dart';
import 'package:auth_bloc/presentation/widges/custom_form_field.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:gap/gap.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  void _onLogin() {
    final String username = usernameController.text.trim();
    final String password = passwordController.text.trim();
    context.read<AuthBloc>().add(AuthLoginRequested(username, password));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          "おかえり",
          style: Theme.of(
            context,
          ).textTheme.headlineSmall!.copyWith(fontWeight: FontWeight.w700),
        ),
      ),

      body: BlocListener<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is Authenticated) {
            Navigator.pushReplacement(
              context,
              MaterialPageRoute(
                builder: (context) {
                  return Homepage();
                },
              ),
            );
          } else if (state is AuthError) {
            ScaffoldMessenger.of(
              context,
            ).showSnackBar(SnackBar(content: Text(state.message)));
          }
        },
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                width: double.infinity,
                height: 200,
                decoration: BoxDecoration(
                  image: DecorationImage(
                    image: AssetImage(AssetLoader.pict1),
                    fit: BoxFit.cover,
                  ),
                  borderRadius: BorderRadius.all(Radius.circular(24)),
                ),
              ),
              Gap(12),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                spacing: 12,
                children: <Widget>[
                  Text(
                    "Welcome Back",
                    style: Theme.of(context).textTheme.headlineSmall!.copyWith(
                      fontWeight: FontWeight.w700,
                    ),
                  ),

                  ConstrainedBox(
                    constraints: BoxConstraints(maxWidth: 350),
                    child: Text(
                      "Dengan satu langkah mudah, kamu bisa kembali ke semua hal yang penting dan menyenangkan",
                      style: Theme.of(context).textTheme.bodyMedium!.copyWith(
                        color: AppColor.foreground.withAlpha(90),
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ),
                ],
              ),
              Gap(20),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                spacing: 12,
                children: [
                  Text(
                    "Username",
                    style: TextStyle(fontWeight: FontWeight.w700),
                  ),
                  CFormField(
                    controller: usernameController,
                    hintText: "Masukkan username",
                  ),
                ],
              ),
              Gap(12),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                spacing: 12,
                children: [
                  Text(
                    "Password",
                    style: TextStyle(fontWeight: FontWeight.w700),
                  ),
                  CFormField(
                    controller: passwordController,
                    secureText: true,
                    hintText: "Massukan password",
                  ),
                ],
              ),
              Align(
                alignment: Alignment.centerRight,
                child: TextButton(
                  onPressed: () {},
                  child: Text(
                    "Lupa Sandi?",
                    style: TextStyle(fontWeight: FontWeight.w700),
                  ),
                ),
              ),
              SizedBox(
                width: double.infinity,
                height: 50,
                child: CustomBtnVariant.normal(() => _onLogin(), "Login"),
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    "Belum punya akun?",
                    style: TextStyle(fontWeight: FontWeight.w700),
                  ),
                  TextButton(
                    onPressed: () {
                      Navigator.pushReplacement(
                        context,
                        MaterialPageRoute(
                          builder: (context) {
                            return RegisterPage();
                          },
                        ),
                      );
                    },
                    child: Text(
                      "Register",
                      style: TextStyle(fontWeight: FontWeight.w700),
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
