package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertitle "fgame/fgame/game/title/player"
)

//结婚开启
func funcOpen(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	funcType := data.(funcopentypes.FuncOpenType)
	if funcType != funcopentypes.FuncOpenTypeMarry {
		return
	}
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(pl.GetRole(), pl.GetSex())
	manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	manager.TempTitleRemove(playerCreateTemplate.TitleId)

	return
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(funcOpen))
}
