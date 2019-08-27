package boss_handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	arenabosstemplate "fgame/fgame/game/arenaboss/template"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	crosslogic "fgame/fgame/game/cross/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shareboss/shareboss"
	"fgame/fgame/game/team/team"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeArena, worldboss.KillBossHandlerFunc(killArenaBoss))
}

func killArenaBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
	//战斗状态
	if pl.IsPvpBattle() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bossId":   biologyId,
			}).Warn("arenaboss:处理进入世界场景,正在pvp")
		playerlogic.SendSystemMessage(pl, lang.PlayerInPVP)
		return
	}

	bossTemp := arenabosstemplate.GetArenaBossTemplateService().GetArenaBossTemplateByBiologyId(biologyId)
	if bossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arenaboss:boss不存在")

		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	myTeam := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	if myTeam != nil && myTeam.IsMatch() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arenaboss:3v3匹配中")
		playerlogic.SendSystemMessage(pl, lang.TeamInMatch)
		return
	}

	boss := shareboss.GetShareBossService().GetShareBoss(typ, biologyId)
	if boss == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"bossBiologyId": biologyId,
			}).Warn("arenaboss:boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断消耗
	itemId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeShengShouItemId)
	itemNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeShengShouItemNum)

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag := inventoryManager.HasEnoughItem(itemId, itemNum)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arenaboss:物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}
	reason := commonlog.InventoryLogReasonArenaBossCost
	reasonText := fmt.Sprintf(reason.String())
	flag = inventoryManager.UseItem(itemId, itemNum, reason, reasonText)
	if !flag {
		panic("zhenxi:移除物品应该成功")
	}
	inventorylogic.SnapInventoryChanged(pl)
	//重置复活次数
	pl.PlayerBossReset(typ)

	crossType := typ.CrossType()
	//TODO 判断是否在跨服中
	arg := fmt.Sprintf("%d", biologyId)
	crosslogic.PlayerEnterCross(pl, crossType, arg)

	scChallengeWorldBoss := pbutil.BuildSCChallengeWorldBoss(boss.GetPosition(), int32(typ))
	pl.SendMsg(scChallengeWorldBoss)
}
