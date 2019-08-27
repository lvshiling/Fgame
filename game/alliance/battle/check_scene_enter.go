package battle

import (
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeChengZhan, scene.CheckEnterSceneHandlerFunc(chenZhanCheckEnterScene))
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeHuangGong, scene.CheckEnterSceneHandlerFunc(huangGongCheckEnterScene))
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeAllianceBoss, scene.CheckEnterSceneHandlerFunc(bossCheckEnterScene))
}

// 城外
func chenZhanCheckEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	enterPos = pos
	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeJiuXiaoChengZhan) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	if pl.GetAllianceId() == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，没有加入仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}

	// switch enterType {
	// case scenetypes.SceneEnterTypePortal:
	// 	{
	// 		now := global.GetGame().GetTimeService().Now()
	// 		if (s.GetEndTime() - now) <= int64(alliancetemplate.GetAllianceTemplateService().GetWarTemplate().HuanggongTime) {
	// 			log.WithFields(
	// 				log.Fields{
	// 					"playerId": pl.GetId(),
	// 					"mapId":    s.MapId(),
	// 				}).Warn("scene:处理进入场景,皇宫已经关闭")
	// 			playerlogic.SendSystemMessage(pl, lang.AllianceHuangGongCloseCannotExit)
	// 			return
	// 		}
	// 		flag = true
	// 		return
	// 	}
	// }
	flag = true
	return
}

//皇宫
func huangGongCheckEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	enterPos = pos
	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if pl.GetAllianceId() == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，没有加入仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}

	switch enterType {
	case scenetypes.SceneEnterTypeCommon:
		{
			flag = true
			return
		}
	case scenetypes.SceneEnterTypePortal:
		{
			// now := global.GetGame().GetTimeService().Now()
			// if (s.GetEndTime() - now) <= int64(alliancetemplate.GetAllianceTemplateService().GetWarTemplate().HuanggongTime) {
			// 	log.WithFields(
			// 		log.Fields{
			// 			"playerId": pl.GetId(),
			// 			"mapId":    s.MapId(),
			// 		}).Warn("scene:处理进入场景,皇宫已经关闭")
			// 	playerlogic.SendSystemMessage(pl, lang.AllianceHuangGongCloseCannotEnter)
			// 	return
			// }
			flag = true
			return
		}
	default:
		{
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    s.MapId(),
				}).Warn("scene:处理进入场景,皇宫已经关闭")
			playerlogic.SendSystemMessage(pl, lang.AllianceMemberNotRescue)
			return
		}
	}
}

//仙盟boss
func bossCheckEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	enterPos = pos
	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeAllianceBoss) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	if pl.GetAllianceId() == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，没有加入仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}

	flag = true
	return
}
