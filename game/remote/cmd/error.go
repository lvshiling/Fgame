package cmd

import (
	"fmt"
)

type CmdError interface {
	error
	Code() ErrorCode
}

type ErrorCode int32

func (c ErrorCode) String() string {
	return errorCodeMap[c]
}

func (c ErrorCode) Error() string {
	return fmt.Sprintf("code:%d,desc:%s", c, c.String())
}

func (c ErrorCode) Code() ErrorCode {
	return c
}

var errorCodeMap = make(map[ErrorCode]string)

func MergeErrorCodeMap(codeMap map[ErrorCode]string) {
	for errorCode, errorDesc := range codeMap {
		errorCodeMap[errorCode] = errorDesc
	}
}

const (
	ErrorCodeCommon ErrorCode = 10000 * iota
	ErrorCodeMail
	ErrorCodePrivilegeCharge
)
