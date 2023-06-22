package admin

type Error struct {
	msg  string
	code int
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() int {
	return e.code
}

func NewError(msg string, code int) *Error {
	return &Error{
		msg:  msg,
		code: code,
	}
}
