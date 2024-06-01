package responses

const (
	SUCCESS = "1"
	FAIL    = "0"
	ERROR   = "-1"
)

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func (a *ApiResponse) SuccessWithDataAndMessage(data interface{}, message string) ApiResponse {
	return ApiResponse{
		data, SUCCESS, message, true,
	}
}

func (a *ApiResponse) FailWithMessage(message string) ApiResponse {
	return ApiResponse{
		nil, FAIL, message, false,
	}
}

func (a *ApiResponse) SuccessDefault() ApiResponse {
	return ApiResponse{
		nil, SUCCESS, "", true,
	}
}

func (a *ApiResponse) Error() ApiResponse {
	return ApiResponse{
		nil, ERROR, "服务器忙，请稍后重试", false,
	}
}
