package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Create(code Code, data interface{}, message ...string) Response {
	msg := code.Text()
	if len(message) > 0 {
		msg = message[0]
	}
	return Response{
		Code:    code.Code(),
		Message: msg,
		Data:    data,
	}
}
