package common

import (
	"fgame/fgame/common/lang"
	"fmt"
)

//通用错误代码
type Error interface {
	error
	Code() lang.LangCode
}

type commonError struct {
	code lang.LangCode
}

func (ce *commonError) Error() string {
	return fmt.Sprintf("error: code:%d", ce.code)
}

func (ce *commonError) Code() lang.LangCode {
	return ce.code
}

func CodeError(code lang.LangCode) error {
	cErr := &commonError{
		code: code,
	}
	return cErr
}
