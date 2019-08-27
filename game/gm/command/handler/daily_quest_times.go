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
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeDailyQuest, command.CommandHandlerFunc(handleDailyQuest))

}

func handleDailyQuest(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	timesStr := c.Args[0]
	times, err := strconv.ParseInt(timesStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"timesStr": timesStr,
				"error":    err,
			}).Warn("gm:处理设置日环任务次数,times不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	dailyTagStr := c.Args[1]
	dailyTagInt, err := strconv.ParseInt(dailyTagStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"timesStr": timesStr,
				"dailyTag": dailyTagStr,
				"error":    err,
			}).Warn("gm:处理设置日环任务,dailyTagInt不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	dailyTag := questtypes.QuestDailyTag(dailyTagInt)
	if !dailyTag.Valid() {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"timesStr": timesStr,
				"dailyTag": dailyTagStr,
				"error":    err,
			}).Warn("gm:处理设置日环任务,dailyTagInt无效")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	dailyTimes := questtypes.QuestDailyType(times)
	maxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(dailyTag)
	if dailyTimes < questtypes.QuestDailyTypeMin || dailyTimes > questtypes.QuestDailyType(maxTimes) {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"timesStr": timesStr,
				"error":    err,
			}).Warn("gm:处理设置日环任务次数,times小于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questId, quest := manager.GMSetDailyQuestTimes(dailyTimes)
	if questId != 0 {
		_, err := questlogic.FinishDailyQuestReward(pl, questId, false)
		if err != nil {
			return err
		}
		qu := manager.GetQuestById(questId)
		if qu != nil {
			scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
			pl.SendMsg(scQuestUpdate)
		}
	}

	dailyObj := manager.GetDailyObj(dailyTag)
	if quest != nil {
		scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), dailyObj.GetSeqId(), dailyObj.GetTimes())
		pl.SendMsg(scQuestDailSeq)
		scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
		pl.SendMsg(scQuestUpdate)
	}

	return
}
