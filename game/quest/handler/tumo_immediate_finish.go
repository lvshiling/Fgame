package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_IMMEDIATE_FINISH_TYPE), dispatch.HandlerFunc(handleQuestTuMoImmediateFinish))
}

//处理屠魔任务直接完成信息
func handleQuestTuMoImmediateFinish(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理屠魔任务直接完成消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csQuestTuMoImmediate := msg.(*uipb.CSQuestTuMoImmediate)
	questId := csQuestTuMoImmediate.GetQuestId()

	err = questTuMoImmediateFinish(tpl, questId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"error":    err,
			}).Error("quest:处理屠魔任务直接完成消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理屠魔任务直接完成消息完成")
	return nil

}

//屠魔任务直接完成信息的逻辑
func questTuMoImmediateFinish(pl player.Player, questId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	questType := questTemplate.GetQuestType()
	if questType != questtypes.QuestTypeTuMo {
		return
	}
	qu := manager.GetQuestById(questId)
	if qu == nil {
		return
	}
	if qu.QuestState != questtypes.QuestStateAccept {
		return
	}

	costGold := questtemplate.GetQuestTemplateService().GetQuestTuMoImmediateFinishCostGold()
	needGold := int64(1 * costGold)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(int64(needGold), true)
	if !flag {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("quest:元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	reasonGoldText := commonlog.GoldLogReasonTuMoImmediateFinishCost.String()
	flag = propertyManager.CostGold(needGold, true, commonlog.GoldLogReasonTuMoImmediateFinishCost, reasonGoldText)
	if !flag {
		panic(fmt.Errorf("quest: questTuMoImmediateFinish CostGold should be ok"))
	}
	//同步属性
	propertylogic.SnapChangedProperty(pl)

	//正在执行的状态置commit
	questList, flag := manager.QuestImmediateFinish(questId)
	if !flag {
		return
	}
	scQuestUpdate := pbutilquest.BuildSCQuestUpdate(questList)
	pl.SendMsg(scQuestUpdate)

	//直接完成对任务的特殊处理
	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishTuMo, 0, 1)

	itemMap, err := questlogic.GiveQuestTuMoImmediateFinishReward(pl, questId)
	if err != nil {
		return
	}

	scQuestTuMoImmediateFinish := pbutil.BuildSCQuestTuMoImmediateFinish(questId, itemMap)
	pl.SendMsg(scQuestTuMoImmediateFinish)
	return
}
