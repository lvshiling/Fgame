package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	chargelogic "fgame/fgame/game/charge/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

//后台充值成功
func privilegeCharge(target event.EventTarget, data event.EventData) (err error) {
	playerId := target.(int64)

	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("charge:充值成功,玩家不存在")
		return
	}

	ctx := scene.WithPlayer(context.Background(), p)
	p.Post(message.NewScheduleMessage(onPrivilegeCharge, ctx, nil, nil))
	return
}

//充值
func onPrivilegeCharge(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	chargelogic.OnPlayerPrivilegeCharge(pl)
	return nil
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypePrivilegeCharge, event.EventListenerFunc(privilegeCharge))
}
