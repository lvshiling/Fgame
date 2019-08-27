package relive_handler

import (
	"fgame/fgame/common/lang"
	teamcopylogic "fgame/fgame/cross/teamcopy/logic"
	teamscene "fgame/fgame/cross/teamcopy/scene"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterReliveEntryPointHandler(scenetypes.SceneTypeCrossTeamCopy, scene.ReliveEntryPointHandlerFunc(ReliveEntryPoint))
}

//回进入点复活
func ReliveEntryPoint(pl scene.Player) (flag bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	sd, ok := s.SceneDelegate().(teamscene.TeamCopySceneData)
	if !ok {
		return
	}
	teamObj := sd.GetTeamObj()
	purpose := teamObj.GetTeamPurpose()
	reliveTime := sd.GetReliveTime(pl)

	teamCopyTemplate := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyTempalte(purpose)
	if teamCopyTemplate == nil {
		return
	}
	maxReliveTime := teamCopyTemplate.ResurrectionNumber
	if reliveTime >= maxReliveTime {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"reliveTime":    reliveTime,
				"maxReliveTime": maxReliveTime,
			}).Warn("teamcopy:复活超过最大次数")
		playerlogic.SendSystemMessage(pl, lang.ReliveBaseExceedMaxTimes)
		return false
	}

	teamcopylogic.Reborn(sd, pl)
	return true
}
