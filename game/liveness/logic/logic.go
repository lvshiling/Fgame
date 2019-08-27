package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	playerliveness "fgame/fgame/game/liveness/player"
	livenesstemplate "fgame/fgame/game/liveness/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//活跃度宝箱奖励
func GiveLivenessBoxReward(pl player.Player, starTemplate *gametemplate.HuoYueBoxTemplate) (isReturn bool) {
	rewData := starTemplate.GetRewData()
	rewItemMap := starTemplate.GetRewItemMap()
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//奖励物品
	if len(rewItemMap) != 0 {
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			isReturn = true
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		inventoryLogReason := commonlog.InventoryLogReasonLivenessStarRew
		reasonText := inventoryLogReason.String()
		flag = inventoryManager.BatchAdd(rewItemMap, inventoryLogReason, reasonText)
		if !flag {
			panic(fmt.Errorf("liveness: GiveLivenessBoxReward BatchAdd should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//添加奖励属性
	if rewData != nil {
		goldLog := commonlog.GoldLogReasonLivenessStarRew
		silverLog := commonlog.SilverLogReasonLivenessStarRew
		levelReason := commonlog.LevelLogReasonLivenessStarRew
		goldReasonText := goldLog.String()
		silverReasonText := silverLog.String()
		reasonText := levelReason.String()
		flag := propertyManager.AddRewData(rewData, goldLog, goldReasonText, silverLog, silverReasonText, levelReason, reasonText)
		if !flag {
			panic(fmt.Errorf("liveness: GiveLivenessBoxReward AddRewData should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}
	return
}

//活跃度任务提交
func LivenessQuestCommit(pl player.Player, questId int32) (err error) {
	log.WithFields(
		log.Fields{
			"liveness": questId,
		}).Info("任务id")

	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	questType := questTemplate.GetQuestType()
	if questType != questtypes.QuestTypeLiveness {
		return
	}

	livenessTempalte := livenesstemplate.GetHuoYueTempalteService().GetHuoYueTemplate(questId)
	if livenessTempalte == nil {
		return
	}
	_, flag := livenesstemplate.GetHuoYueTempalteService().GetHuoYueLevelTemplate(questId, pl.GetLevel())
	if !flag {
		return
	}

	liveNessmanager := pl.GetPlayerDataManager(types.PlayerLivenessDataManagerType).(*playerliveness.PlayerLivenessDataManager)
	curNum := liveNessmanager.AddQuestNum(questId)
	if curNum >= livenessTempalte.RewardCountLimit {
		return
	}
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest := questManager.CommitLivenessResetInit(questId)
	if quest != nil {
		questlogic.CheckInitQuest(pl, quest)
		scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
		pl.SendMsg(scQuestUpdate)
	}
	return
}
