package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildISLogin(playerId int64, match bool) *crosspb.ISLogin {
	isLogin := &crosspb.ISLogin{}
	isLogin.Match = &match
	isLogin.PlayerId = &playerId
	return isLogin
}
