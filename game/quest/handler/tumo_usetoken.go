package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_USETOKEN_TYPE), dispatch.HandlerFunc(handleQuestTuMoUseToken))
}

//处理使用屠魔令
func handleQuestTuMoUseToken(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理使用屠魔令")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuestTuMoUseToken := msg.(*uipb.CSQuestTuMoUseToken)
	token := csQuestTuMoUseToken.GetToken()

	err = questTuMoUseToken(tpl, questtypes.QuestLevelType(token))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"token":    token,
				"error":    err,
			}).Error("quest:处理使用屠魔令,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"token":    token,
		}).Debug("quest:处理使用屠魔令,完成")
	return nil
}

//使用屠魔令
func questTuMoUseToken(pl player.Player, token questtypes.QuestLevelType) (err error) {
	//校验参数
	if !token.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"token":    token,
			}).Warn("quest:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	//任务列表是否已满
	flag := manager.IfTuMoListReachLimit()
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"token":    token,
			}).Warn("quest:屠魔任务列表已满,请先完成并交付")
		playerlogic.SendSystemMessage(pl, lang.QuestTuMoListReachLimit)
		return
	}

	//屠魔次数是否达上限
	flag = manager.IfTuMoNumReachLimit()
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"token":    token,
			}).Warn("quest:今日屠魔次数已用完")
		playerlogic.SendSystemMessage(pl, lang.QuestTuMoNumReachLimit)
		return
	}

	//获取屠魔令所需道具
	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeTuMoLing, token.ItemTumoSubType())
	tuMoItem := int32(itemTemplate.TemplateId())
	//判断物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag = inventoryManager.HasEnoughItem(tuMoItem, 1)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"token":    token,
			}).Warn("quest:屠魔令数量不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	questId, flag := manager.GetTuMoQuestIdByToken(token)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"token":    token,
			}).Warn("quest:当前服务器繁忙，建议先完成接取到的屠魔任务")
		playerlogic.SendSystemMessage(pl, lang.QuestTuMoQuestIdNotGet)
		return
	}

	//消耗屠魔令
	reasonText := commonlog.InventoryLogReasonUseToken.String()
	flag = inventoryManager.UseItem(tuMoItem, 1, commonlog.InventoryLogReasonUseToken, reasonText)
	if !flag {
		panic(fmt.Errorf("quest: questTuMoUseToken UseItem should be ok"))
	}
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)

	//接取任务
	flag = manager.AcceptTumoQuest(questId)
	if !flag {
		panic(fmt.Errorf("quest: questTuMoUseToken AcceptTumoQuest should be ok"))
	}

	quest := manager.GetTuMoQuestByLevelAndId(token, questId)
	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)

	scQuestTuMoUseToken := pbutil.BuildSCQuestTuMoUseToken(int32(token), questId)
	pl.SendMsg(scQuestTuMoUseToken)
	return

}
