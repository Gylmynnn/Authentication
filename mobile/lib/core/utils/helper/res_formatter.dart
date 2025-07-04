class ResponseModel {
  final bool success;
  final int status;
  final String message;
  final dynamic data;

  ResponseModel({
    required this.success,
    required this.status,
    required this.message,
    required this.data,
  });

  factory ResponseModel.fromJson(Map<String, dynamic> json) => ResponseModel(
    success: json['success'],
    status: json['status'],
    message: json['message'],
    data: json['data'],
  );
}
