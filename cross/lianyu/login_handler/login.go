package login_handler

import (
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/lianyu/lianyu"
	scenelogic "fgame/fgame/game/scene/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeLianYu, login.LogincHandlerFunc(lainYuLogin))
}

func lainYuLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {
	s := lianyu.GetLianYuService().GetLianYuScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"crossType": ct,
			}).Warn("login:无间炼狱登陆失败，场景不存在")
		return false
	}
	pos, flag := lianyu.GetLianYuService().GetRebornPos(pl.GetId())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"crossType": ct,
			}).Warn("login:无间炼狱登陆失败，获取出生位置失败")
		return false
	}

	scenelogic.PlayerEnterScene(pl, s, pos)
	return true
}
