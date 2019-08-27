package listener

import (
	"fgame/fgame/core/event"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	xiantilogic "fgame/fgame/game/xianti/logic"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	xiantitypes "fgame/fgame/game/xianti/types"
	"fgame/fgame/game/xianti/xianti"
	"fmt"
)

//玩家天魔体关联法宝皮肤激活
func playerTianMoTiToXianTi(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*tianmoeventtypes.PlayerTianMoTiUnitePiFuEventData)
	piFuType := eventData.GetPiFuType()
	piFuId := eventData.GetPiFuId()
	if piFuType != commontypes.AdvancedUnitePiFuTypeXianTi {
		return
	}
	xiantiTemplate := xianti.GetXianTiService().GetXianTi(int(piFuId))
	if xiantiTemplate == nil {
		return
	}
	if xiantiTemplate.GetTyp() != xiantitypes.XianTiTypeSkin {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	if manager.IsUnrealed(xiantiTemplate.TemplateId()) {
		return
	}
	manager.AddUnrealInfo(xiantiTemplate.TemplateId())
	//同步属性
	xiantilogic.XianTiPropertyChanged(pl)
	flag := manager.Unreal(xiantiTemplate.TemplateId())
	if !flag {
		panic(fmt.Errorf("xianti:幻化应该成功"))
	}
	scXianTiUnreal := pbutil.BuildSCXianTiUnreal(piFuId)
	pl.SendMsg(scXianTiUnreal)
	return
}

func init() {
	gameevent.AddEventListener(tianmoeventtypes.EventTypeTianMoUnitePiFu, event.EventListenerFunc(playerTianMoTiToXianTi))
}
