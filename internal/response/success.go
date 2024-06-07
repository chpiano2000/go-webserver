package response

type SuccessResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OK(data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  200,
		Code:    "Success",
		Message: "Success",
		Data:    data,
	}
}

func Created(code string, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  201,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Deleted(code string, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  200,
		Code:    code,
		Message: message,
		Data:    data,
	}
}
