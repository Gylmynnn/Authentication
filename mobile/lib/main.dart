import 'package:auth_bloc/data/repository/auth_repository.dart';
import 'package:auth_bloc/logic/auth/bloc/auth_bloc.dart';
import 'package:auth_bloc/logic/auth/bloc/auth_event.dart';
import 'package:auth_bloc/presentation/pages/auth/login_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

void main() {
  runApp(const App());
}

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return RepositoryProvider(
      create: (_) => AuthRepository(),
      child: BlocProvider(
        create: (ctx) {
          return AuthBloc(authRepository: ctx.read<AuthRepository>())
            ..add(AuthCheckStatus());
        },
        child: MaterialApp(
          title: 'Auth with bloc',
          theme: ThemeData(
            colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
          ),
          home: LoginPage(),
        ),
      ),
    );
  }
}
