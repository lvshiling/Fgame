package battle

import (
	"fgame/fgame/common/lang"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	relivelogic "fgame/fgame/game/relive/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterReliveHandler(scenetypes.SceneTypeChengZhan, scene.ReliveHandlerFunc(Relive))
	scene.RegisterReliveHandler(scenetypes.SceneTypeHuangGong, scene.ReliveHandlerFunc(Relive))
}

//城战原地复活
func Relive(pl scene.Player, autoBuy bool) {
	s := pl.GetScene()
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}

	tpl, ok := pl.(player.Player)
	if !ok {
		return
	}
	manager := tpl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceObj := manager.GetPlayerAllianceObject()
	reliveTime := allianceObj.GetReliveTime()

	maxReliveTime := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().RebornSitu
	if reliveTime >= maxReliveTime {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"reliveTime":    reliveTime,
				"maxReliveTime": maxReliveTime,
			}).Warn("alliance:复活超过最大次数")
		playerlogic.SendSystemMessage(pl, lang.ReliveBaseYuanDiMaxTimes)
		return
	}

	sucess := relivelogic.Relive(pl, autoBuy)
	if !sucess {
		return
	}
	manager.AddReliveTime()
}
