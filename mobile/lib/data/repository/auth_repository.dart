import 'package:auth_bloc/core/utils/helper/res_formatter.dart';
import 'package:auth_bloc/logic/auth/models/login_model.dart';
import 'package:auth_bloc/logic/auth/models/register_model.dart';
import 'package:auth_bloc/logic/auth/models/user_model.dart';
import 'package:dio/dio.dart';

class AuthRepository {
  final Dio _dio = Dio(BaseOptions(baseUrl: 'http://10.0.2.2:3000/api'));

  Future<UserModel> login(LoginModel req) async {
    final response = await _dio.post<ResponseModel>(
      '/login',
      data: {'username': req.username, 'password': req.password},
    );

    if (response.statusCode == 200 && response.data!.success) {
      final UserModel data = UserModel.fromJson(response.data!.data);
      return data;
    } else {
      throw Exception(response.data!.message);
    }
  }

  Future<UserModel> register(RegisterModel req) async {
    final response = await _dio.post(
      '/register',
      data: {
        'name': req.name,
        'username': req.username,
        'email': req.email,
        'password': req.password,
      },
    );
    print(response);
    if (response.statusCode == 201) {
      final UserModel data = UserModel.fromJson(response.data['data']);
      print(data);
      return data;
    } else {
      throw Exception(response.data!.message);
    }
  }

  // Future<Map<String, dynamic>> getProfile(String token) async {
  //   final response = await _dio.get(
  //     '/me',
  //     options: Options(headers: {'Authorization': 'Bearer $token'}),
  //   );
  //
  //   if (response.statusCode == 200 && response.data['success']) {
  //     return response.data['data'];
  //   } else {
  //     throw Exception('Gagal mengambil data profil');
  //   }
  // }
}
