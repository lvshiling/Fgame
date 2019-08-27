package login_handler

import (
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	crosstypes "fgame/fgame/game/cross/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/shenmo/shenmo"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeShenMoWar, login.LogincHandlerFunc(shenMoLogin))
}

func shenMoLogin(pl *player.Player, ct crosstypes.CrossType, crossArgs ...string) (flag bool) {
	s := shenmo.GetShenMoService().GetShenMoScene()
	if s == nil {
		return false
	}

	pos := s.MapTemplate().GetBornPos()
	if !scenelogic.PlayerEnterScene(pl, s, pos) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tulong:进入场景失败")
		return false
	}

	
	return true
}
