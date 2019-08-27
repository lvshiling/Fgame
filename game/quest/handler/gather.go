package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_GATHER_TYPE), dispatch.HandlerFunc(handleQuestGather))
}

//处理任务收集
func handleQuestGather(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理任务收集")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuestGather := msg.(*uipb.CSQuestGather)
	questId := csQuestGather.GetQuestId()
	itemId := csQuestGather.GetItemId()
	num := csQuestGather.GetNum()

	err = questGather(tpl, questId, itemId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"itemId":   itemId,
				"num":      num,
				"error":    err,
			}).Error("quest:处理任务收集,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"questId":  questId,
			"itemId":   itemId,
			"num":      num,
		}).Debug("quest:处理任务收集,完成")
	return nil
}

//收集
func questGather(pl player.Player, questId int32, itemId int32, num int32) (err error) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	//TODO 记录恶意刷的
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"itemId":   itemId,
				"num":      num,
			}).Warn("quest:处理任务收集,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest := manager.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	if quest == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:处理任务收集,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}

	flag := manager.IncreaseQuestData(questId, itemId, num)
	if !flag {
		panic("quest:处理任务收集,应该成功")
	}
	_, err = questlogic.CheckQuestIfFinish(pl, questId)
	if err != nil {
		return
	}
	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)
	scQuestGather := pbutil.BuildSCQuestGather(questId, itemId, num)
	pl.SendMsg(scQuestGather)
	return
}
