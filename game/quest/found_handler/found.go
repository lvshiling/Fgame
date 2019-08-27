package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/secretcard/secretcard"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeTianJiPai, found.FoundObjDataHandlerFunc(getTianJiPaiFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeTuMo, found.FoundObjDataHandlerFunc(getTuMoFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeDailyQuest, found.FoundObjDataHandlerFunc(getDailyFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeAllianceDaily, found.FoundObjDataHandlerFunc(getAllianceDailyFoundParam))
}

func getTianJiPaiFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	resLevel, group = getParam(pl, questtypes.QuestTypeTianJiPai)
	maxTimes = secretcard.GetSecretCardService().GetConstSecretCardNum()
	return
}

func getDailyFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	resLevel, group = getParam(pl, questtypes.QuestTypeDaily)
	maxTimes = questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(questtypes.QuestDailyTagPerson)
	return
}

func getTuMoFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	resLevel, group = getParam(pl, questtypes.QuestTypeTuMo)
	maxTimes = questtemplate.GetQuestTemplateService().GetQuestTuMoInitialNum()
	return
}


func getAllianceDailyFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	resLevel, group = getParam(pl, questtypes.QuestTypeDailyAlliance)
	maxTimes = questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(questtypes.QuestDailyTagAlliance)
	return
}

func getParam(pl player.Player, typ questtypes.QuestType) (resLevel int32, group int32) {
	group = int32(1)
	resLevel = pl.GetLevel()
	return
}
