package collect_handler

import (
	"fgame/fgame/common/lang"
	collectnpc "fgame/fgame/game/collect/npc"
	moonlovetemplate "fgame/fgame/game/moonlove/template"
	playerlogic "fgame/fgame/game/player/logic"
	scene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	collectnpc.RegisterSceneCollectConditionHandler(scenetypes.SceneTypeYueXiaQingYuan, collectnpc.SceneCollectConditionHandlerFunc(collectCondition))
}

//月下情缘 采集条件
func collectCondition(pl scene.Player, cn scene.CollectNPC) bool {
	s := pl.GetScene()
	activityType, ok := s.MapTemplate().GetMapType().ToActivityType()
	if !ok {
		return false
	}

	moonloveTemp := moonlovetemplate.GetMoonloveTemplateService().GetMoonloveTemplate(pl.GetLevel())
	collectCount := pl.GetActivityTotalCollectCount(activityType)
	if collectCount >= moonloveTemp.CollectLimitCount {
		log.WithFields(
			log.Fields{
				"playerId":          pl.GetId(),
				"collectCount":      collectCount,
				"collectLimitCount": moonloveTemp.CollectLimitCount,
			}).Warn("moonlove:处理玩家采集,采集次数不足")
		playerlogic.SendSystemMessage(pl, lang.MoonloveCollectTimesNotEnough)
		return false
	}

	return true
}
