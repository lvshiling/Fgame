package exception

import (
	"fgame/fgame/common/lang"
)

type ExceptionCode int32

const (
	ExceptionCodeKickout                ExceptionCode = iota + 1
	ExceptionCodeServerBusy                           //服务器繁忙
	ExceptionCodePlayerLoginTimeout                   //登陆超时
	ExceptionCodePlayerStateException                 //状态异常
	ExceptionCodePlayerTokenInvalid                   //玩家登陆失效
	ExceptionCodePlayerLoginSameTime                  //玩家同时登陆
	ExceptionCodeRegisterClose                        //已经关闭注册
	ExceptionCodeAccountArgumentInvalid               //账户参数无效
	ExceptionCodeServerNoOpen                         //服务器还没开始
	ExceptionCodeAccountForbid                        //账户或ip被封禁

)

var (
	exceptionCodeLangMap = map[ExceptionCode]lang.LangCode{
		ExceptionCodeServerBusy:             lang.ServerBusy,
		ExceptionCodePlayerLoginTimeout:     lang.ExceptionPlayerAuthTimeout,
		ExceptionCodePlayerStateException:   lang.ExceptionPlayerStateError,
		ExceptionCodePlayerTokenInvalid:     lang.AccountNoExist,
		ExceptionCodePlayerLoginSameTime:    lang.AccountLoginAtSameTime,
		ExceptionCodeRegisterClose:          lang.AccountRegisterClose,
		ExceptionCodeAccountArgumentInvalid: lang.AccountArgumentInvalid,
		ExceptionCodeServerNoOpen:           lang.ServerNoOpen,
		ExceptionCodeAccountForbid:          lang.AccountOrIpForbid,
	}
)

func (c ExceptionCode) LangCode() lang.LangCode {
	return exceptionCodeLangMap[c]
}
