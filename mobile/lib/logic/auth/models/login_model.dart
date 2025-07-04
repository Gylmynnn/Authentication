class LoginModel {
  final String username;
  final String password;

  LoginModel({required this.username, required this.password});

  factory LoginModel.fromJson(Map<String, dynamic> json) =>
      LoginModel(username: json['username'], password: json['password']);
}
