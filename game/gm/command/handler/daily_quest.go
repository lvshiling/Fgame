package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeDailyQuestSet, command.CommandHandlerFunc(handleDailyQuestSet))

}

func handleDailyQuestSet(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	dailyTagStr := c.Args[0]
	dailyTagInt, err := strconv.ParseInt(dailyTagStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
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
				"dailyTag": dailyTagStr,
				"error":    err,
			}).Warn("gm:处理设置日环任务,dailyTagInt无效")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	seqIdStr := c.Args[1]
	seqId, err := strconv.ParseInt(seqIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"dailyTag": dailyTagStr,
				"error":    err,
			}).Warn("gm:处理设置日环任务,dailyTagInt不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	updateQuestList := manager.GMSetDailyQuest(dailyTag, int32(seqId))
	if len(updateQuestList) == 0 {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"dailyTag": dailyTagStr,
				"error":    err,
			}).Warn("gm:处理设置日环任务,dailyTagInt不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	dailyObj := manager.GetDailyObj(dailyTag)
	scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), dailyObj.GetSeqId(), dailyObj.GetTimes())
	pl.SendMsg(scQuestDailSeq)

	scQuestUpdate := pbutil.BuildSCQuestListUpdate(updateQuestList)
	pl.SendMsg(scQuestUpdate)
	return
}
