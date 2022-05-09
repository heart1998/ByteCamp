package errcode

import "fmt"

var codes = make(map[int]bool)

type Err interface {
	error
	Code() int
	Msg() string
}

type Code struct {
	code int
	msg  string
}

func NewCode(code int, msg string) Err {
	if codes[code] {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = true
	return &Code{
		code: code,
		msg:  msg,
	}
}

func (code *Code) Error() string {
	return fmt.Sprintf("错误码:%d,错误信息:%s", code.code, code.msg)
}

func (code *Code) Code() int {
	return code.code
}

func (code *Code) Msg() string {
	return code.msg
}
