package listener

import (
	"fgame/fgame/core/event"
	buffcommon "fgame/fgame/game/buff/common"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	playerbuff "fgame/fgame/game/buff/player"
	bufftemplate "fgame/fgame/game/buff/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//buff更新
func buffUpdate(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)
	buffObject := data.(buffcommon.BuffObject)

	buffId := buffObject.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)

	//玩家保存数据
	if buffTemplate.IsSave() {
		switch pl := bo.(type) {
		case player.Player:
			buffManager := pl.GetPlayerDataManager(playertypes.PlayerBuffDataManagerType).(*playerbuff.PlayerBuffDataManager)
			buffManager.UpdateBuff(buffObject)
			break
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffUpdate, event.EventListenerFunc(buffUpdate))
}
