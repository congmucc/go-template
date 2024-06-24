package result

var (
	OK = Response{
		Code:    0,
		Message: "操作成功",
		Data:    nil,
	}
	ERR = Response{
		Code:    1,
		Message: "操作失败",
		Data:    nil,
	}
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (res *Response) SuccessWithData(data interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    data,
	}
}

func (res *Response) SuccessWithMessage(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    res.Data,
	}
}

func (res *Response) SuccessWithDataAndMessage(message string, data interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    data,
	}
}

func (res *Response) ErrorWithMessage(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    res.Data,
	}
}
