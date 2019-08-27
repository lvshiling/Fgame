package login

import (
	"fgame/fgame/cross/player/player"
	crosstypes "fgame/fgame/game/cross/types"
	"fmt"
)

type LoginHandler interface {
	Login(pl *player.Player, crossType crosstypes.CrossType, args ...string) bool
}

type LogincHandlerFunc func(pl *player.Player, crossType crosstypes.CrossType, args ...string) bool

func (f LogincHandlerFunc) Login(pl *player.Player, crossType crosstypes.CrossType, args ...string) bool {
	return f(pl, crossType, args...)
}

var (
	loginHandleMap = map[crosstypes.CrossType]LoginHandler{}
)

func RegisterLoginHandler(ct crosstypes.CrossType, h LoginHandler) {
	_, exist := loginHandleMap[ct]
	if exist {
		panic(fmt.Errorf("重复注册跨服登陆%s", ct.String()))
	}
	loginHandleMap[ct] = h
}

func GetLoginHandler(ct crosstypes.CrossType) LoginHandler {
	h, ok := loginHandleMap[ct]
	if !ok {
		return nil
	}
	return h
}
