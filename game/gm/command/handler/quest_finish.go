package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	"fgame/fgame/game/scene/scene"
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeQuestFinish, command.CommandHandlerFunc(handleQuestFinish))
}

func handleQuestFinish(pl scene.Player, c *command.Command) (err error) {
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
	err = finishQuest(pl, int32(questId))
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

func finishQuest(p scene.Player, questId int32) (err error) {
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
	questManager.FinishQuest(questId)
	if questTemplate.AutoCommit() {
		flag := questManager.CommitQuest(questId, false)
		if !flag {
			panic(fmt.Errorf("quest:提交任务应该ok"))
		}
		//TODO发送奖励
		questlogic.GiveQuestCommitReward(pl, questId)
	}

	quest := questManager.GetQuestById(questId)
	scMsg := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scMsg)
	return
}
