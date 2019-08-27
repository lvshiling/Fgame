package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeQuest, command.CommandHandlerFunc(handleQuest))
}

func handleQuest(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理任务")
	if len(c.Args[0]) <= 0 {
		log.Warn("gm:处理任务,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	questIdStr := c.Args[0]
	questId, err := strconv.ParseInt(questIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"questId": questIdStr,
			}).Warn("gm:处理任务,questId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//TODO 修改物品数量
	err = modifyQuest(pl, int32(questId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"questId": questId,
			}).Warn("gm:处理任务,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"questId": questId,
		}).Debug("gm:处理任务,完成")
	return
}

func modifyQuest(p scene.Player, questId int32) (err error) {
	pl := p.(player.Player)
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"questId": questId,
			}).Warn("gm:处理任务,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questManager.GMModifyQuestId(questId)

	questlogic.CheckInitQuestList(pl)
	questlogic.CheckAcceptQuestList(pl)

	err = questlogic.CheckCommitQuestList(pl)
	questList := make([]*playerquest.PlayerQuestObject, 0, 16)
	activeQuests := questManager.GetQuestMap(questtypes.QuestStateActive)
	for _, quest := range activeQuests {
		questList = append(questList, quest)
	}
	acceptQuests := questManager.GetQuestMap(questtypes.QuestStateAccept)
	for _, quest := range acceptQuests {
		questList = append(questList, quest)
	}
	finishQuests := questManager.GetQuestMap(questtypes.QuestStateFinish)
	for _, quest := range finishQuests {
		questList = append(questList, quest)
	}

	finish := false
	mainQuestId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMainQuestId)
	mainQuest := questManager.GetQuestById(mainQuestId)
	if mainQuest != nil && mainQuest.QuestState == questtypes.QuestStateCommit {
		finish = true
	}

	scQuestList := pbutil.BuildSCQuestList(questList, finish)
	pl.SendMsg(scQuestList)
	return
}
