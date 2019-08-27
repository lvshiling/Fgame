package listener

import (
	"fgame/fgame/core/event"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
	fabaotypes "fgame/fgame/game/fabao/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	"fmt"
)

//玩家噬魂幡关联法宝皮肤激活
func playerShiHunFanFaBao(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*shihunfaneventtypes.PlayerShiHunFanUnitePiFuEventData)
	piFuType := eventData.GetPiFuType()
	piFuId := eventData.GetPiFuId()
	if piFuType != commontypes.AdvancedUnitePiFuTypeFaBao {
		return
	}
	fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(piFuId))
	if fabaoTemplate == nil {
		return
	}
	if fabaoTemplate.GetTyp() != fabaotypes.FaBaoTypeSkin {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)

	if manager.IsUnrealed(fabaoTemplate.TemplateId()) {
		return
	}
	manager.AddUnrealInfo(fabaoTemplate.TemplateId())
	//同步属性
	fabaologic.FaBaoPropertyChanged(pl)
	flag := manager.Unreal(fabaoTemplate.TemplateId())
	if !flag {
		panic(fmt.Errorf("fabao:幻化应该成功"))
	}
	scFaBaoUnreal := pbutil.BuildSCFaBaoUnreal(piFuId)
	pl.SendMsg(scFaBaoUnreal)
	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanUnitePiFu, event.EventListenerFunc(playerShiHunFanFaBao))
}
