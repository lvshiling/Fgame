package login_handler

import (
	"fgame/fgame/cross/arenapvp/arenapvp"
	"fgame/fgame/cross/arenapvp/pbutil"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	crosstypes "fgame/fgame/game/cross/types"
	scenelogic "fgame/fgame/game/scene/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeArenapvp, login.LogincHandlerFunc(arenapvpLogin))
}

func arenapvpLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {

	battleS := arenapvp.GetArenapvpService().GetArenapvpBattleScene(pl.GetId())
	if battleS != nil {
		sd, ok := battleS.SceneDelegate().(arenapvpscene.ArenapvpBattleSceneData)
		if !ok {
			return false
		}
		bornPos := sd.GetEnterPos(pl.GetId())
		if !scenelogic.PlayerEnterScene(pl, battleS, bornPos) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("login:pvp重连场景失败")
			return false
		}
	} else {
		s := arenapvp.GetArenapvpService().GetArenapvpSceneElection()
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"crossType": ct,
				}).Warn("login:pvp海选登陆失败，场景不存在")
			return false
		}

		pos := s.MapTemplate().RandomPosition()
		if !scenelogic.PlayerEnterScene(pl, s, pos) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("login:pvp进入海选场景失败")
			return false
		}
	}

	isMsg := pbutil.BuildISArenapvpAttendSuccess()
	pl.SendMsg(isMsg)

	return true
}
