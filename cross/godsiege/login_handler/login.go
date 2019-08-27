package login_handler

import (
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/godsiege/godsiege"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeGodSiegeQiLin, login.LogincHandlerFunc(godSiegeQiLinLogin))
	login.RegisterLoginHandler(crosstypes.CrossTypeGodSiegeHuoFeng, login.LogincHandlerFunc(godSiegeHuoFengLogin))
	login.RegisterLoginHandler(crosstypes.CrossTypeGodSiegeDuLong, login.LogincHandlerFunc(godSiegeDuLongLogin))
	login.RegisterLoginHandler(crosstypes.CrossTypeDenseWat, login.LogincHandlerFunc(denseWatLogin))
}

func godSiegeQiLinLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {
	return enterScene(pl, godsiegetypes.GodSiegeTypeQiLin)
}

func godSiegeHuoFengLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {
	return enterScene(pl, godsiegetypes.GodSiegeTypeHuoFeng)
}

func godSiegeDuLongLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {
	return enterScene(pl, godsiegetypes.GodSiegeTypeDuLong)
}

func denseWatLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {
	return enterScene(pl, godsiegetypes.GodSiegeTypeDenseWat)
}

func enterScene(pl scene.Player, godType godsiegetypes.GodSiegeType) bool {
	s := godsiege.GetGodSiegeService().GetGodSiegeScene(godType)
	if s == nil {
		log.WithFields(
			log.Fields{
				"godType": godType,
			}).Warn("godsiege:场景为空")
		return false
	}
	pos, flag := godsiege.GetGodSiegeService().GetRebornPos(godType, pl.GetId())
	if !flag {
		log.WithFields(
			log.Fields{
				"godType": godType,
			}).Warn("godsiege:场景为空")
		return false
	}

	return scenelogic.PlayerEnterScene(pl, s, pos)
}
