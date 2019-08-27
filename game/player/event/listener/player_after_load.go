package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	propertylogic "fgame/fgame/game/property/logic"
)

func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	if !p.IsGetNewReward() {
		newTitle := lang.GetLangService().ReadLang(lang.NewTitle)
		newContent := lang.GetLangService().ReadLang(lang.NewContent)
		attachment := make(map[int32]int32)
		itemId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeNewGift)
		attachment[itemId] = 1
		emaillogic.AddEmail(p, newTitle, newContent, attachment)
		//领取新手礼包
		p.GetNewReward()
	}
	//通知客户端属性变化
	propertylogic.SnapChangedProperty(p)
	return nil
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
