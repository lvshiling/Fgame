package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	marryeventtypes "fgame/fgame/game/marry/event/types"
)

//玩家婚戒替换
func playerRingReplace(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryRingReplaceEventData)
	if !ok {
		return
	}
	ringType := eventData.GetRingType()
	spouseId := eventData.GetSpouseId()

	marry.GetMarryService().MarryReplace(pl.GetId(), ringType)
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		return
	}
	splCtx := scene.WithPlayer(context.Background(), spl)
	playerRingReplaceMsg := message.NewScheduleMessage(onPlayerRingReplace, splCtx, ringType, nil)
	spl.Post(playerRingReplaceMsg)
	return
}

func onPlayerRingReplace(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	ringType := result.(marrytypes.MarryRingType)

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	spouseId := marryInfo.SpouseId
	curRing := marryInfo.Ring
	if ringType <= curRing {
		return nil
	}

	manager.RingReplacedBySpouse(ringType)
	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())
	ringItem := int32(itemTemplate.TemplateId())
	marrylogic.MarryPropertyChanged(pl)
	scMarryRingReplace := pbuitl.BuildSCMarryRingReplace(ringItem)
	pl.SendMsg(scMarryRingReplace)

	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		return nil
	}
	scMarryRingChange := pbuitl.BuildSCMarryRingChange(pl.GetId(), ringItem)
	spl.SendMsg(scMarryRingChange)
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryRingReplace, event.EventListenerFunc(playerRingReplace))
}
