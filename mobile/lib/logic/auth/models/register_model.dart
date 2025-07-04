class RegisterModel {
  final String name;
  final String username;
  final String email;
  final String password;

  RegisterModel({
    required this.name,
    required this.username,
    required this.email,
    required this.password,
  });

  factory RegisterModel.fromJson(Map<String, dynamic> json) => RegisterModel(
    name: json['name'],
    username: json['username'],
    email: json['email'],
    password: json['password'],
  );
}
