package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	pbutilquest "fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	secretcardlogic "fgame/fgame/game/secretcard/logic"
	"fgame/fgame/game/secretcard/pbutil"
	playersecretcard "fgame/fgame/game/secretcard/player"
	"fgame/fgame/game/secretcard/secretcard"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_IMMEDIATE_FINISH_TYPE), dispatch.HandlerFunc(handleSecretImmediateFinish))
}

//处理天机牌直接完成信息
func handleSecretImmediateFinish(s session.Session, msg interface{}) (err error) {
	log.Debug("secretcard:处理天机牌直接完成消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = secretImmediateFinish(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("secretcard:处理天机牌直接完成消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("secretcard:处理天机牌直接完成消息完成")
	return nil

}

//天机牌直接完成信息的逻辑
func secretImmediateFinish(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	cardObj := manager.GetSecretCard()
	curCardId := cardObj.CardId
	cardMap := cardObj.CardMap
	questId := int32(0)
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	//避免发多次
	if curCardId == 0 && len(cardMap) == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("secretcard:请先接取任务")
		return
	}
	//翻牌了未接取任务
	if len(cardMap) != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("secretcard:请先接取任务")
		playerlogic.SendSystemMessage(pl, lang.SecretImmediateFinish)
		return
	}
	//获取任务接取进行中
	if curCardId != 0 {
		questId, _, _ = secretcard.GetSecretCardService().GetQuestIdByCardId(curCardId)
		quest := questManager.GetQuestById(questId)
		if quest != nil {
			questState := quest.QuestState
			if questState != questtypes.QuestStateAccept {
				return
			}
		}
	}

	costGold := secretcard.GetSecretCardService().GetConstSecretCardCostGold()
	needGold := int64(1 * costGold)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(int64(needGold), true)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("secretcard:元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	reasonGoldText := commonlog.GoldLogReasonSecretCardFinishAllCost.String()
	flag = propertyManager.CostGold(needGold, true, commonlog.GoldLogReasonSecretCardFinishAllCost, reasonGoldText)
	if !flag {
		panic(fmt.Errorf("secretcard: secretFinish CostGold should be ok"))
	}

	//正在执行的状态置commit
	if questId != 0 {
		qu, flag := questManager.QuestImmediateFinish(questId)
		if !flag {
			return
		}
		scQuestUpdate := pbutilquest.BuildSCQuestUpdate(qu)
		pl.SendMsg(scQuestUpdate)
	}

	//天机牌直接完成奖励
	dropItemDataList, err := secretcardlogic.GiveSecretCardImmediateFinishReward(pl, cardObj)

	//直接完成对任务的特殊处理
	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishSecretCard, 0, 1)

	manager.ImmediateFinish()

	// 资源找回事件
	data := questeventtypes.CreateQuestFinishAllEventData(questtypes.QuestTypeTianJiPai, 1)
	gameevent.Emit(questeventtypes.EventTypeQuestFinishAll, pl, data)
	//同步属性
	propertylogic.SnapChangedProperty(pl)
	scQuestSecretImmediateFinish := pbutil.BuildSCQuestSecretImmediateFinish(dropItemDataList)
	pl.SendMsg(scQuestSecretImmediateFinish)
	return
}
