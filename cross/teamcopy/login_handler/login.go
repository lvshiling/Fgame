package login_handler

import (
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/teamcopy/teamcopy"
	crosstypes "fgame/fgame/game/cross/types"
	scenelogic "fgame/fgame/game/scene/logic"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeTeamCopy, login.LogincHandlerFunc(teamCopyLogin))
}

func teamCopyLogin(pl *player.Player, ct crosstypes.CrossType, args ...string) bool {
	sd := teamcopy.GetTeamCopyService().GetTeamCopySceneData(pl)
	if sd == nil {
		log.WithFields(
			log.Fields{}).Warn("teamcopy:组队副本场景不存在")
		return false
	}
	// 获取pos
	teamObj := sd.GetTeamObj()
	teamPurpose := teamObj.GetTeamPurpose()

	pos, flag := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyBorn(teamPurpose)
	if !flag {
		return false
	}
	scenelogic.PlayerEnterScene(pl, sd.GetScene(), pos)
	log.WithFields(
		log.Fields{
			"playerName":  pl.GetName(),
			"teamPurpose": pl.GetTeamPurpose(),
		}).Warn("teamcopy:玩家进场组队副本")
	return true
}
