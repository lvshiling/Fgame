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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	playerproperty "fgame/fgame/game/property/player"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	pbutilquest "fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_FINISH_ALL_TYPE), dispatch.HandlerFunc(handleQuestTuMoFinishAll))
}

//处理屠魔任务一键完成信息
func handleQuestTuMoFinishAll(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理屠魔任务一键完成消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = questTuMoFinishAll(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("quest:处理屠魔任务一键完成消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理屠魔任务一键完成消息完成")
	return nil

}

//屠魔任务一键完成信息的逻辑
func questTuMoFinishAll(pl player.Player) (err error) {
	vipNum := pl.GetVip()
	vipLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoFinishVipLimit)
	if vipNum < vipLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"curVip":   vipNum,
				"vipLimit": vipLimit,
			}).Warn("quest:玩家VIP等级不足,不能一键完成")
		vipStr := fmt.Sprintf("%d", vipLimit)
		playerlogic.SendSystemMessage(pl, lang.QuestTuMoFinishAllVipNoEnough, vipStr)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	acceptQuestIds, leftNum, acceptNum := manager.GetTuMoAcceptQuestAndNeedNum()
	if leftNum < 0 {
		panic(fmt.Errorf("quest: secretFinish leftNum should be greater than -1"))
	}
	if leftNum == 0 && acceptNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("quest:已全部完成")
		playerlogic.SendSystemMessage(pl, lang.QuestTuMoFinishAll)
		return
	}

	costGold := questtemplate.GetQuestTemplateService().GetQuestTuMoFinishCostGold()
	costAcceptGold := questtemplate.GetQuestTemplateService().GetQuestTuMoImmediateFinishCostGold()
	needGold := int64(leftNum*costGold) + int64(acceptNum*costAcceptGold)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(int64(needGold), true)
	if !flag {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("quest:元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	questIdList := questtemplate.GetQuestTemplateService().GetQuestIdListForTuMoFinishAll(leftNum, pl.GetLevel())

	for _, questId := range acceptQuestIds {
		questIdList = append(questIdList, questId)
	}

	//消耗元宝
	reasonGoldText := commonlog.GoldLogReasonSecretCardFinishAllCost.String()
	flag = propertyManager.CostGold(needGold, true, commonlog.GoldLogReasonSecretCardFinishAllCost, reasonGoldText)
	if !flag {
		panic(fmt.Errorf("quest: secretFinish CostGold should be ok"))
	}

	totalNum := manager.UseTuMoNum(leftNum)
	//正在执行的状态置commit
	if len(acceptQuestIds) != 0 {
		//清空屠魔列表
		questList := manager.QuestToKeyComplete(acceptQuestIds)
		scQuestListUpdate := pbutilquest.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestListUpdate)
	}

	//一键完成对任务的特殊处理
	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishTuMo, 0, leftNum)

	itemMapList, err := questlogic.GiveQuestTuMoFinishAllReward(pl, questIdList, leftNum+acceptNum)
	if err != nil {
		return
	}

	scQuestTuMoFinishAll := pbutil.BuildSCQuestTuMoFinishAll(itemMapList, totalNum)
	pl.SendMsg(scQuestTuMoFinishAll)
	return
}
