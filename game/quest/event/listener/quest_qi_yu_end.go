package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	questeventtypes "fgame/fgame/game/quest/event/types"
	playerquest "fgame/fgame/game/quest/player"
	questplayer "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
)

//奇遇任务结束
func questQiYuEnd(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	qiyu, ok := data.(*questplayer.PlayerQiYuObject)
	if !ok {
		return
	}
	qiyuId := qiyu.GetQiYuId()
	qiYuTemplate := questtemplate.GetQuestTemplateService().GetQiYuTemplate(qiyuId)
	if qiYuTemplate == nil {
		return
	}

	if qiyu.GetIsFinish() != 1 {
		return
	}

	if qiyu.GetIsReceive() != 0 {
		return
	}

	rewItemMap := qiYuTemplate.GetRewEmailItemMap()
	if qiYuTemplate.RewExpPoint > 0 {
		exp := propertylogic.ExpPointConvertExp(qiYuTemplate.RewExpPoint, pl.GetLevel())
		rewItemMap[constanttypes.ExpItem] += int32(exp)
	}
	title := lang.GetLangService().ReadLang(lang.QuestQiYuEndMailTitle)
	content := lang.GetLangService().ReadLang(lang.QuestQiYuEndMailContent)
	emaillogic.AddEmail(pl, title, content, rewItemMap)

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	manager.ReceiveQiYu(qiyuId)

	return
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestQiYuEnd, event.EventListenerFunc(questQiYuEnd))
}
