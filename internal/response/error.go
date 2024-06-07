package response

type ErrorResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BadRequestResponse(message string, code string) ErrorResponse {
	return ErrorResponse{
		Status:  400,
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func FormatBodyError() ErrorResponse {
	return ErrorResponse{
		Status:  400,
		Code:    "format_body_error",
		Message: "Format Body Expects To Be JSON Type",
		Data:    nil,
	}
}
