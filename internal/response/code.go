package response

import (
	"fmt"
	"net/http"
)

type Code struct {
	code int
	http int
	text string
}

func (c Code) String() string {
	return fmt.Sprintf("%d(%d%s)", c.http, c.code, c.text)
}
func (c Code) Status() int {
	return c.http
}
func (c Code) Code() int {
	return c.code
}
func (c Code) Text() string {
	return c.text
}

var (
	OK          = Code{code: 0, http: http.StatusOK, text: "正常"}
	ServerError = Code{code: 10101, http: http.StatusInternalServerError, text: "内部服务器错误"}
)
