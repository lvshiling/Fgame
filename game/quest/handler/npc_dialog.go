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
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_NPC_DIALOG_TYPE), dispatch.HandlerFunc(handleQuestNPCDialog))
}

//处理任务对话
func handleQuestNPCDialog(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理任务对话")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuestNPCDialog := msg.(*uipb.CSQuestNPCDialog)
	questId := csQuestNPCDialog.GetQuestId()
	npcId := csQuestNPCDialog.GetNpcId()

	err = questNPCDialog(tpl, questId, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"npcId":    npcId,
				"error":    err,
			}).Error("quest:处理任务对话,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"questId":  questId,
			"npcId":    npcId,
		}).Debug("quest:处理任务对话,完成")
	return nil
}

//任务对话
func questNPCDialog(pl player.Player, questId int32, npcId int32) (err error) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	//TODO 记录恶意刷的
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"npcId":    npcId,
			}).Warn("quest:处理任务对话,任务不存在")
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
			}).Warn("quest:处理任务对话,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}

	flag := manager.IncreaseQuestData(questId, npcId, 1)
	if !flag {
		panic("quest:处理任务对话,应该成功")
	}

	//TODO 检测任务是否完成
	_, err = questlogic.CheckQuestIfFinish(pl, questId)
	if err != nil {
		return
	}
	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)
	scQuestNPCDialog := pbutil.BuildSCQuestNPCDialog(questId, npcId)
	pl.SendMsg(scQuestNPCDialog)
	return
}
