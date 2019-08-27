package collect_handler

import (
	"fgame/fgame/common/lang"
	collectnpc "fgame/fgame/game/collect/npc"
	longgonglogic "fgame/fgame/game/longgong/logic"
	longgongtemplate "fgame/fgame/game/longgong/template"
	playerlogic "fgame/fgame/game/player/logic"
	scene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	collectnpc.RegisterSceneCollectConditionHandler(scenetypes.SceneTypeLongGong, collectnpc.SceneCollectConditionHandlerFunc(longGongCollectCondition))
}

//龙宫探宝 采集条件
func longGongCollectCondition(pl scene.Player, cn scene.CollectNPC) bool {
	s := pl.GetScene()
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeLongGong {
		return false
	}

	sd := s.SceneDelegate()
	longgongSd, ok := sd.(longgonglogic.LongGongSceneData)
	if !ok {
		return false
	}

	bioType := cn.GetBiologyTemplate().GetBiologyScriptType()
	switch bioType {
	case scenetypes.BiologyScriptTypeLongGongTreasure:
		return collectLongGongTreasure(pl, longgongSd)
	default:
		break
	}

	return true
}

func collectLongGongTreasure(pl scene.Player, longgongSd longgonglogic.LongGongSceneData) bool {
	constTemp := longgongtemplate.GetLongGongTemplateService().GetLongGongConstTemplate()
	pCollectCount := longgongSd.GetPlayerTreasureCollectCount(pl.GetId())
	if pCollectCount >= constTemp.BossBeKillCaiJiPersonalCount {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"pCollectCount": pCollectCount,
				"CaiJiLimit":    constTemp.BossBeKillCaiJiPersonalCount,
			}).Warn("longgong:处理玩家采集黑龙财宝,采集次数不足")
		playerlogic.SendSystemMessage(pl, lang.LongGongCollectCountNoEnough)
		return false
	}

	return true
}
