package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiuxianbooktypes "fgame/fgame/game/welfare/xiuxianbook/types"
	"fgame/fgame/game/welfare/xiuxianbook/xiuxianbook"
)

// 元神金装脱下来的信息
func playerGoldequipTakeOffStrength(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldequipData, ok := data.(*goldequipeventtypes.PlayerGoldEquipStatusEventData)
	if !ok {
		return
	}

	subType := welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength
	level := goldequipData.GetStrengthenLevel()
	err = handlePlayerGoldequipTakeOff(pl, subType, level)
	return
}

func playerGoldequipTakeOffOpenLight(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldequipData, ok := data.(*goldequipeventtypes.PlayerGoldEquipStatusEventData)
	if !ok {
		return
	}

	subType := welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight
	level := goldequipData.GetOpenlightLevel()
	err = handlePlayerGoldequipTakeOff(pl, subType, level)
	return
}

func playerGoldequipTakeOffUpStar(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldequipData, ok := data.(*goldequipeventtypes.PlayerGoldEquipStatusEventData)
	if !ok {
		return
	}

	subType := welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar
	level := goldequipData.GetUpstarLevel()
	err = handlePlayerGoldequipTakeOff(pl, subType, level)
	return
}

func handlePlayerGoldequipTakeOff(pl player.Player, subType welfaretypes.OpenActivityXiuxianBookSubType, level int32) (err error) {
	typ := welfaretypes.OpenActivityTypeXiuxianBook

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityData(typ, subType)
	if err != nil {
		return err
	}
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
		countLevelHandler := xiuxianbook.GetCountLevelHandler(obj.GetActivityType(), obj.GetActivitySubType())
		maxlevel, err := countLevelHandler.CountLevel(pl)
		if err != nil {
			return err
		}
		maxlevel = maxlevel + level
		if maxlevel > info.MaxLevel {
			info.MaxLevel = maxlevel
			welfareManager.UpdateObj(obj)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipStatusWhenTakeOff, event.EventListenerFunc(playerGoldequipTakeOffStrength))
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipStatusWhenTakeOff, event.EventListenerFunc(playerGoldequipTakeOffOpenLight))
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipStatusWhenTakeOff, event.EventListenerFunc(playerGoldequipTakeOffUpStar))
}
