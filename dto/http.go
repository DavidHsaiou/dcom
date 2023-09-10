package dto

type Request struct {
	Data any
}

type Response struct {
	ResultCode int
	Message    string
	Data       any
}

var (
	ResultCodeSuccess = 0
)

var resultMsgMap = map[int]string{
	ResultCodeSuccess: "success",
}

func NewResponse(resultCode int, data any) *Response {
	return &Response{
		ResultCode: resultCode,
		Message:    resultMsgMap[resultCode],
		Data:       data,
	}
}
