package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	playerproperty "fgame/fgame/game/property/player"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_ACCEPT_TYPE), dispatch.HandlerFunc(handleQuestAccept))
}

//处理任务接受
func handleQuestAccept(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理任务接受")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuestAccept := msg.(*uipb.CSQuestAccept)
	questId := csQuestAccept.GetQuestId()

	if questId <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务接受,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = questAccept(tpl, questId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"error":    err,
			}).Error("quest:处理任务接受,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"questId":  questId,
		}).Debug("quest:处理任务接受,完成")
	return nil
}

//接受
func questAccept(pl player.Player, questId int32) (err error) {
	//模板不存在
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	//TODO 记录恶意刷的
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务接受,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}

	//自动接取 不能手动
	if questTemplate.AutoAccept() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务接受,自动接取,不能手动")
		playerlogic.SendSystemMessage(pl, lang.QuestAcceptAuto)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest := manager.GetQuestByIdAndState(questtypes.QuestStateActive, questId)
	if quest == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务接受,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}

	//任务已经接受
	if !manager.ShouldAcceptQuest(questId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务接受,任务已经接受")
		playerlogic.SendSystemMessage(pl, lang.QuestAlreadyAccepted)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//判断各种条件
	//消耗物品
	consumeItemMap := questTemplate.GetConsumeItemMap(pl.GetRole())
	if len(consumeItemMap) != 0 {
		if !inventoryManager.HasEnoughItems(consumeItemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"questId":  questId,
				}).Warn("quest:处理任务接受,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}
	//等级判断
	level := pl.GetLevel()
	if level < questTemplate.MinLevel || level > questTemplate.MaxLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务接受,等级太低")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}
	if level > questTemplate.MaxLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务接受,等级太高")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooHigh)
		return
	}

	//银两判断
	if questTemplate.ConsumeSilver != 0 {
		if !propertyManager.HasEnoughSilver(int64(questTemplate.ConsumeSilver)) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"questId":  questId,
				}).Warn("quest:处理任务接受,银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//元宝
	if questTemplate.ConsumeGold != 0 {
		if !propertyManager.HasEnoughGold(int64(questTemplate.ConsumeGold), false) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"questId":  questId,
				}).Warn("quest:处理任务接受,元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}
	//TODO
	//转数判断
	//特殊条件判断

	if len(consumeItemMap) != 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonQuestAccept.String(), questTemplate.TemplateId())
		flag := inventoryManager.BatchRemove(consumeItemMap, commonlog.InventoryLogReasonQuestAccept, reasonText)
		if !flag {
			panic("quest:扣除消耗物品,应该成功")
		}
	}

	//银两判断
	if questTemplate.ConsumeSilver != 0 {
		reasonText := fmt.Sprintf(commonlog.SilverLogReasonQuestAccept.String(), questTemplate.TemplateId())
		flag := propertyManager.CostSilver(int64(questTemplate.ConsumeSilver), commonlog.SilverLogReasonQuestAccept, reasonText)
		if !flag {
			panic("quest:扣除银两,应该成功")
		}
	}
	//元宝
	if questTemplate.ConsumeGold != 0 {
		reasonText := fmt.Sprintf(commonlog.GoldLogReasonQuestAccept.String(), questTemplate.TemplateId())
		flag := propertyManager.CostGold(int64(questTemplate.ConsumeGold), false, commonlog.GoldLogReasonQuestAccept, reasonText)
		if !flag {
			panic("quest:扣除元宝,应该成功")
		}

	}
	flag := manager.AcceptQuest(questId)
	if !flag {
		panic(fmt.Errorf("quest:接受任务应该成功"))
	}

	questlogic.CheckAcceptQuestList(pl)
	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)

	scQuestAccept := pbutil.BuildSCQuestAccept(questId)
	pl.SendMsg(scQuestAccept)
	return
}
