package battle

import (
	"fgame/fgame/common/lang"
	chuangshiscene "fgame/fgame/cross/chuangshi/scene"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	collectnpc "fgame/fgame/game/collect/npc"
	playerlogic "fgame/fgame/game/player/logic"
	scene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	collectnpc.RegisterSceneCollectConditionHandler(scenetypes.SceneTypeChuangShiZhiZhanFuShu, collectnpc.SceneCollectConditionHandlerFunc(collectCondition))
}

//采集条件
func collectCondition(pl scene.Player, cn scene.CollectNPC) bool {
	s := pl.GetScene()
	sd := s.SceneDelegate().(chuangshiscene.FuShuSceneData)

	initDefenCamp := sd.GetInitDefendCampType()
	warTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(initDefenCamp)
	// 玉玺
	if cn.GetBiologyTemplate().Id == warTemp.GetYuXiBiologyTemp().Id {
		if sd.GetCurrentDefendCampType() == pl.GetCamp() {
			log.WithFields(
				log.Fields{
					"playerId":         pl.GetId(),
					"currentDefenCamp": sd.GetCurrentDefendCampType(),
					"plCamp":           pl.GetCamp(),
				}).Warn("chuangshi:当前玩家仙盟是守方仙盟")
			playerlogic.SendSystemMessage(pl, lang.AllianceScenePlayerIsDefence)
			return false
		}
	}

	return true
}
