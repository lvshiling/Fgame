package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	babylogic "fgame/fgame/game/baby/logic"
	playerbaby "fgame/fgame/game/baby/player"
	babytypes "fgame/fgame/game/baby/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

// 配偶宝宝变化
func coupleBabyChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	coupleBabyList, ok := data.([]*babytypes.CoupleBabyData)
	if !ok {
		return
	}

	//同步配偶宝宝
	spl := player.GetOnlinePlayerManager().GetPlayerById(pl.GetSpouseId())
	if spl != nil {
		ctx := scene.WithPlayer(context.Background(), spl)
		msg := message.NewScheduleMessage(spouseBabyChanged, ctx, coupleBabyList, nil)
		spl.Post(msg)
	}

	return
}

// 配偶宝宝变化
func spouseBabyChanged(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	babyDataList := result.([]*babytypes.CoupleBabyData)

	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	babyManager.LoadAllCoupleBaby(babyDataList)
	babylogic.BabyPropertyChanged(pl)

	return nil
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeCoupleBabyChanged, event.EventListenerFunc(coupleBabyChanged))
}
