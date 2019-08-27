package listener

import (
	"fgame/fgame/core/event"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	gameevent "fgame/fgame/game/event"
)

//buff更新
func buffUpdate(target event.EventTarget, data event.EventData) (err error) {
	//发送保存事件

	// bo := target.(scene.BattleObject)
	// buffObject := data.(buffcommon.BuffObject)

	// buffId := buffObject.GetBuffId()
	// buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)

	// //玩家保存数据
	// if buffTemplate.IsSave() {
	// 	switch pl := bo.(type) {
	// 	case player.Player:
	// 		buffManager := pl.GetPlayerDataManager(playertypes.PlayerBuffDataManagerType).(*playerbuff.PlayerBuffDataManager)
	// 		buffManager.UpdateBuff(buffObject)
	// 		break
	// 	}
	// }

	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffUpdate, event.EventListenerFunc(buffUpdate))
}
