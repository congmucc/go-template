package utils

const (
	StatusOk  = 0
	StatusErr = 1
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *Result) SuccessWithData(message string, data interface{}) Result {
	return Result{
		Code:    StatusOk,
		Message: message,
		Data:    data,
	}
}

func (res *Result) SuccessWithMessage(message string) Result {
	return Result{
		Code:    StatusOk,
		Message: message,
		Data:    res.Data,
	}
}

func (res *Result) ErrorWithMessage(message string) Result {
	return Result{
		Code:    StatusErr,
		Message: message,
		Data:    res.Data,
	}
}
