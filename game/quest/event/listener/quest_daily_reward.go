package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	questeventtypes "fgame/fgame/game/quest/event/types"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
)

//任务奖励
func questDailyReward(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	dailyObj := data.(*playerquest.PlayerDailyObject)
	seqId := dailyObj.GetSeqId()
	dailyTag := dailyObj.GetDailyTag()
	dailyTempalte := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(dailyTag, seqId)
	if dailyTempalte == nil {
		return
	}
	emailMap := dailyTempalte.GetEmailItemMap()
	rewData := dailyTempalte.GetRewData()

	itemMap := propertylogic.GetItemMapWithExpPoint(rewData.RewExpPoint, pl.GetLevel(), emailMap)
	if len(itemMap) == 0 {
		return
	}

	//个人日环额外处理
	if dailyTag == questtypes.QuestDailyTagPerson {
		itemId, num := droptemplate.GetDropTemplateService().GetDropItem(dailyTempalte.GetDropId())
		if itemId != 0 && num != 0 {
			itemMap[itemId] += num
		}
	}

	emailTitle := lang.GetLangService().ReadLang(lang.QuestDailyEmailTitle)
	emailContent := lang.GetLangService().ReadLang(lang.QuestDailyEmailContent)
	emaillogic.AddEmail(pl, emailTitle, emailContent, itemMap)
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestDailyReward, event.EventListenerFunc(questDailyReward))
}
