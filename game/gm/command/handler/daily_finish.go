package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeDailyFinish, command.CommandHandlerFunc(handleDailyQuestFinish))

}

func handleDailyQuestFinish(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
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
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	dailyObj := manager.GetDailyObj(dailyTag)
	dailyTemplate := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, dailyObj.GetSeqId())
	if dailyTemplate == nil {
		return
	}

	qu := manager.GetQuestById(dailyTemplate.GetQuestId())
	if qu == nil {
		return
	}
	if qu.QuestState != questtypes.QuestStateAccept {
		return
	}

	questTempalte := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
	if questTempalte == nil {
		return
	}

	questSubType := questTempalte.GetQuestSubType()
	demandId := int32(0)
	for id, _ := range questTempalte.GetQuestDemandMap(pl.GetRole()) {
		demandId = id
		break
	}

	err = questlogic.FillQuestData(pl, questSubType, demandId)
	return
}
