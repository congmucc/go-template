package utils

const (
	StatusOk  = 0
	StatusErr = 1
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *Response) SuccessWithData(message string, data interface{}) Response {
	return Response{
		Code:    StatusOk,
		Message: message,
		Data:    data,
	}
}

func (res *Response) SuccessWithMessage(message string) Response {
	return Response{
		Code:    StatusOk,
		Message: message,
		Data:    res.Data,
	}
}

func (res *Response) ErrorWithMessage(message string) Response {
	return Response{
		Code:    StatusErr,
		Message: message,
		Data:    res.Data,
	}
}
