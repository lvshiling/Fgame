package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	crosslogic "fgame/fgame/game/cross/logic"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shareboss/shareboss"
	"fgame/fgame/game/team/team"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	xianzuncardlogic "fgame/fgame/game/xianzuncard/logic"
	playerzhenxi "fgame/fgame/game/zhenxi/player"
	zhenxitemplate "fgame/fgame/game/zhenxi/template"

	"fmt"

	log "github.com/Sirupsen/logrus"
)

func CheckPlayerIfCanZhenXiBossChallenge(pl player.Player, biologyId int32) (flag bool) {
	// if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeCangJingGe) {
	// 	return
	// }

	// huiYuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	// isHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	// tempHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	// if !(isHuiYuan || tempHuiYuan) {
	// 	return
	// }

	// if !playerlogic.CheckCanEnterScene(pl) {
	// 	return
	// }

	// bossTemp := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(biologyId)
	// if bossTemp == nil {
	// 	return
	// }

	// s := scene.GetSceneService().GetCangJingGeSceneByMapId(bossTemp.MapId)
	// if s == nil {
	// 	return
	// }
	// flag = true
	return
}

func HandleKillZhenXiBoss(pl player.Player, bossBiologyId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeZhenXiBoss) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": bossBiologyId,
			}).Warn("zhenxi:珍稀boss挑战请求，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	//战斗状态
	if pl.IsPvpBattle() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bossId":   bossBiologyId,
			}).Warn("scene:处理进入世界场景,正在pvp")
		playerlogic.SendSystemMessage(pl, lang.PlayerInPVP)
		return
	}

	// // 跨服中
	// if pl.IsCross() {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"bossId":   bossBiologyId,
	// 		}).Warn("scene:处理进入世界场景,正在跨服场景")
	// 	playerlogic.SendSystemMessage(pl, lang.PlayerInCross)
	// 	return
	// }

	bossTemp := zhenxitemplate.GetZhenXiTemplateService().GetZhenXiBossTemplateByBiologyId(bossBiologyId)
	if bossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": bossBiologyId,
			}).Warn("zhenxi:boss挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	typ := worldbosstypes.BossTypeZhenXi
	boss := shareboss.GetShareBossService().GetShareBoss(typ, bossBiologyId)
	if boss == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"bossBiologyId": bossBiologyId,
			}).Warn("zhenxi:boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	myTeam := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	if myTeam != nil && myTeam.IsMatch() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shareBoss:珍稀boss匹配中")
		playerlogic.SendSystemMessage(pl, lang.TeamInMatch)
		return
	}

	//判断消耗
	bossUseTemplate := zhenxitemplate.GetZhenXiTemplateService().GetBossUseTemplate(bossTemp.MapId)
	if bossUseTemplate != nil {
		useItemMap := bossUseTemplate.GetUseItemMap()
		if IfFreeEnter(pl) {
			goto Enter
		}

		if len(useItemMap) != 0 {
			inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
			flag := inventoryManager.HasEnoughItems(useItemMap)
			if !flag {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
					}).Warn("zhenxi:物品不足")
				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
				return
			}
			reason := commonlog.InventoryLogReasonZhenXiCost
			reasonText := fmt.Sprintf(reason.String(), bossTemp.GetMapTemplate().Name)
			flag = inventoryManager.BatchRemove(useItemMap, reason, reasonText)
			if !flag {
				panic("zhenxi:移除物品应该成功")
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

Enter:
	pl.PlayerBossReset(typ)
	zhenXinManager := pl.GetPlayerDataManager(playertypes.PlayerZhenXiDataManagerType).(*playerzhenxi.PlayerZhenXiDataManager)
	zhenXinManager.EnterZhenXiBoss()

	crossType := typ.CrossType()
	arg := fmt.Sprintf("%d", bossBiologyId)
	crosslogic.PlayerEnterCross(pl, crossType, arg)

	scChallengeWorldBoss := pbutil.BuildSCChallengeWorldBoss(boss.GetPosition(), int32(typ))
	pl.SendMsg(scChallengeWorldBoss)
	return
}

//珍稀boss是否免费进入
func IfFreeEnter(pl player.Player) bool {
	zhenXinManager := pl.GetPlayerDataManager(playertypes.PlayerZhenXiDataManagerType).(*playerzhenxi.PlayerZhenXiDataManager)
	zhenXiInfo := zhenXinManager.GetPlayerZhenXiObject()
	freeTimes := xianzuncardlogic.ZhenXiBossFreeTimes(pl)
	return zhenXiInfo.GetEnterTimes() < freeTimes
}
