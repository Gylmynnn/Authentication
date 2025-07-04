import 'package:auth_bloc/data/repository/auth_repository.dart';
import 'package:auth_bloc/logic/auth/models/login_model.dart';
import 'package:auth_bloc/logic/auth/models/register_model.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'auth_event.dart';
import 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final AuthRepository authRepository;

  AuthBloc({required this.authRepository}) : super(AuthInitial()) {
    on<AuthLoginRequested>(_onLogin);
    on<AuthRegisterRequested>(_onRegister);
    on<AuthLogoutRequested>(_onLogout);
    on<AuthCheckStatus>(_onCheckStatus);
  }

  void _onLogin(AuthLoginRequested event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    final LoginModel req = LoginModel(
      username: event.username,
      password: event.password,
    );
    try {
      final response = await authRepository.login(req);
      final String token = response.token ?? "";
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('token', token);
      emit(Authenticated(token, response));
    } catch (e) {
      emit(AuthError(e.toString()));
    }
  }

  void _onRegister(AuthRegisterRequested event, Emitter<AuthState> emit) async {
    emit(AuthLoading());
    final RegisterModel req = RegisterModel(
      name: event.name,
      username: event.username,
      email: event.email,
      password: event.password,
    );
    try {
      await authRepository.register(req);
      emit(Unauthenticated());
    } catch (e) {
      emit(AuthError(e.toString()));
    }
  }

  void _onLogout(AuthLogoutRequested event, Emitter<AuthState> emit) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('token');
    emit(Unauthenticated());
  }

  void _onCheckStatus(AuthCheckStatus event, Emitter<AuthState> emit) async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('token');
    if (token != null) {
      try {
        emit(Authenticated(token, null));
      } catch (e) {
        emit(Unauthenticated());
      }
    } else {
      emit(Unauthenticated());
    }
  }
}
