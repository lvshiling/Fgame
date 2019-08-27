package collect_handler

import (
	"fgame/fgame/common/lang"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	collectnpc "fgame/fgame/game/collect/npc"
	playerlogic "fgame/fgame/game/player/logic"
	scene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	collectnpc.RegisterSceneCollectConditionHandler(scenetypes.SceneTypeChengZhan, collectnpc.SceneCollectConditionHandlerFunc(collectCondition))
}

//神域之战 采集条件
func collectCondition(pl scene.Player, cn scene.CollectNPC) bool {
	s := pl.GetScene()
	sd := s.SceneDelegate().(alliancescene.AllianceSceneData)

	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	// 玉玺
	if cn.GetBiologyTemplate().Id == warTemp.GetYuXiBiologyTemp().Id {
		if sd.GetCurrentDefendAllianceId() == pl.GetAllianceId() {
			log.WithFields(log.Fields{
				"currentAllianceId": sd.GetCurrentDefendAllianceId(),
				"playerId":          pl.GetId(),
				"allianceId":        pl.GetAllianceId(),
			}).Warn("alliance:当前玩家仙盟是守方仙盟")
			playerlogic.SendSystemMessage(pl, lang.AllianceScenePlayerIsDefence)
			return false
		}
	}

	return true
}
