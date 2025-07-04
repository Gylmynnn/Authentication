package helper

type SuccessResponse struct {
	Success bool   `json:"success"`
	Status  int16  `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type FailedResponse struct {
	Success bool  `json:"success"`
	Status  int16 `json:"status"`
	Message any   `json:"message"`
}
