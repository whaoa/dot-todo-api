package logger

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/pkgerrors"
)

type ErrorStackMarshaler func(err error) interface{}

type state struct {
	b            []byte
	FullFilePath bool
}

func (s *state) Write(b []byte) (n int, err error) {
	s.b = b
	return len(b), nil
}
func (s *state) Width() (wid int, ok bool) {
	return 0, false
}
func (s *state) Precision() (prec int, ok bool) {
	return 0, false
}
func (s *state) Flag(c int) bool {
	return s.FullFilePath && c == '+'
}

func frameField(f errors.Frame, s *state, c rune) string {
	f.Format(s, c)
	return string(s.b)
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// 创建错误栈序列化处理函数
func createErrorStackMarshaler(fullFilePath bool) ErrorStackMarshaler {
	return func(err error) interface{} {
		e, ok := err.(stackTracer)
		if !ok {
			return nil
		}
		stack := e.StackTrace()
		out := make([]map[string]string, 0, len(stack))
		s := &state{FullFilePath: fullFilePath}

		for _, frame := range stack {
			out = append(out, map[string]string{
				pkgerrors.StackSourceFileName:     frameField(frame, s, 's'),
				pkgerrors.StackSourceLineName:     frameField(frame, s, 'd'),
				pkgerrors.StackSourceFunctionName: frameField(frame, s, 'n'),
			})
		}

		return out
	}
}
