package collect_handler

import (
	"fgame/fgame/common/lang"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	scene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	xiantaologic "fgame/fgame/game/xiantao/logic"
	playerxiantao "fgame/fgame/game/xiantao/player"
	xiantaotemplate "fgame/fgame/game/xiantao/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	collectnpc.RegisterSceneCollectConditionHandler(scenetypes.SceneTypeXianTaoDaHui, collectnpc.SceneCollectConditionHandlerFunc(xianTaoCollectCondition))
}

//仙桃大会采集条件
func xianTaoCollectCondition(p scene.Player, cn scene.CollectNPC) bool {
	pl, ok := p.(player.Player)
	if !ok {
		return false
	}
	s := pl.GetScene()
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		return false
	}

	sd := s.SceneDelegate()
	xiantaoSd, ok := sd.(xiantaologic.XianTaoSceneData)
	if !ok {
		return false
	}

	xianTaoService := xiantaotemplate.GetXianTaoTemplateService()
	constTemp := xianTaoService.GetXianTaoConstTemplate()
	pCollectCount := xiantaoSd.GetPlayerCollectCount(pl.GetId())
	if pCollectCount >= constTemp.CaiJiLimit {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"pCollectCount": pCollectCount,
				"CaiJiLimit":    constTemp.CaiJiLimit,
			}).Warn("xiantao:处理玩家采集仙桃,采集次数不足")
		playerlogic.SendSystemMessage(pl, lang.XianTaoCollectCountNoEnough)
		return false
	}

	xianTaoManager := pl.GetPlayerDataManager(types.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoObject := xianTaoManager.GetXianTaoObject()
	bioType := cn.GetBiologyTemplate().GetBiologyScriptType()
	switch bioType {
	case scenetypes.BiologyScriptTypeXianTaoQianNianCollect:
		if xianTaoObject.HighPeachCount >= constTemp.XianTaoMax {
			log.WithFields(
				log.Fields{
					"playerId":       pl.GetId(),
					"XianTaoMax":     constTemp.XianTaoMax,
					"HighPeachCount": xianTaoObject.HighPeachCount,
				}).Warn("xiantao:处理玩家采集仙桃,千年仙桃数量上限")
			playerlogic.SendSystemMessage(pl, lang.XianTaoCollectCountNoEnough)
			return false
		}
		break
	case scenetypes.BiologyScriptTypeXianTaoBaiNianCollect:
		if xianTaoObject.JuniorPeachCount >= constTemp.XianTaoMax {
			log.WithFields(
				log.Fields{
					"playerId":         pl.GetId(),
					"XianTaoMax":       constTemp.XianTaoMax,
					"JuniorPeachCount": xianTaoObject.JuniorPeachCount,
				}).Warn("xiantao:处理玩家采集仙桃,百年仙桃数量上限")
			playerlogic.SendSystemMessage(pl, lang.XianTaoCollectCountNoEnough)
			return false
		}
		break
	default:
		return false
	}

	return true
}
