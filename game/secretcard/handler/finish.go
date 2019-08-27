package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_FINISH_TYPE), dispatch.HandlerFunc(handleSecretFinish))
}

//处理天机牌一键完成信息
func handleSecretFinish(s session.Session, msg interface{}) (err error) {
	log.Debug("secretcard:处理天机牌一键完成消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = secretFinish(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("secretcard:处理天机牌一键完成消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("secretcard:处理天机牌一键完成消息完成")
	return nil

}

//天机牌一键完成信息的逻辑
func secretFinish(pl player.Player) (err error) {
	vipNum := pl.GetVip()
	vipLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTianJiPaiFinishVipLimit)
	if vipNum < vipLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"curVip":   vipNum,
				"vipLimit": vipLimit,
			}).Warn("secretcard:玩家VIP等级不足,不能一键完成")
		vipStr := fmt.Sprintf("%d", vipLimit)
		playerlogic.SendSystemMessage(pl, lang.SecretCardFinishAllVipNoEnough, vipStr)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	leftNum := manager.GetLeftCardNum()
	if leftNum < 0 {
		panic(fmt.Errorf("secretcard: secretFinish leftNum  should be greater than -1"))
	}
	cardObj := manager.GetSecretCard()
	curCardId := cardObj.CardId
	cardMap := cardObj.CardMap
	needNum := leftNum
	noCommitNum := leftNum
	questIdList := make([]int32, 0, 8)
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	//获取任务接取进行中
	if curCardId != 0 {
		questId, _, _ := secretcard.GetSecretCardService().GetQuestIdByCardId(curCardId)
		questState := questManager.GetQuestById(questId).QuestState
		if questState == questtypes.QuestStateAccept {
			needNum++
			questIdList = append(questIdList, questId)
		}
		if questState != questtypes.QuestStateCommit {
			noCommitNum++
		}
	}
	//翻牌了未接取任务
	if len(cardMap) != 0 {
		needNum++
		noCommitNum++
	}
	if needNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("secretcard:已全部完成")
		playerlogic.SendSystemMessage(pl, lang.SecretCardFinishAll)
		return
	}

	costGold := secretcard.GetSecretCardService().GetConstSecretCardCostGold()
	needGold := int64(needNum * costGold)
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
	if len(questIdList) != 0 {
		questList := questManager.QuestToKeyComplete(questIdList)
		scQuestListUpdate := pbutilquest.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestListUpdate)
	}

	//天机牌一键完成奖励
	dropItemMap, addStar, leftBoxList, err := secretcardlogic.GiveSecretCardFinishAllReward(pl, cardObj, leftNum)

	//增加星数
	manager.SecretCardFinishAll(leftNum, addStar, leftBoxList)

	//一键完成对任务的特殊处理
	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishSecretCard, 0, leftNum)

	// 领取事件
	data := questeventtypes.CreateQuestFinishAllEventData(questtypes.QuestTypeTianJiPai, noCommitNum)
	gameevent.Emit(questeventtypes.EventTypeQuestFinishAll, pl, data)

	//同步属性
	propertylogic.SnapChangedProperty(pl)

	scQuestSecretFinish := pbutil.BuildSCQuestSecretFinish(dropItemMap, true)
	pl.SendMsg(scQuestSecretFinish)
	return
}
