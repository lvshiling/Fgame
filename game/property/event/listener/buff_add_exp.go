package listener

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	"fgame/fgame/game/buff/buff"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fmt"
	"math"
)

//玩家玩家进入场景
func buffAddExp(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	s := p.GetScene()
	if s == nil {
		return
	}
	eventData := data.(*buff.BuffExpEventData)
	buffId := eventData.GetBuffId()
	exp := int64(eventData.GetExp())
	expPoint := int64(eventData.GetExpPoint())
	propertyManager := p.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	reason := commonlog.LevelLogReasonBuff
	reasonText := fmt.Sprintf(reason.String(), buffId)

	expAddPercent := s.GetExpAddPercent(p)
	if expAddPercent > 0 {
		expPoint = int64(math.Ceil(float64(expPoint) * (1 + float64(expAddPercent)/float64(common.MAX_RATE))))
		exp = int64(math.Ceil(float64(exp) * (1 + float64(expAddPercent)/float64(common.MAX_RATE))))
	}

	if expPoint > 0 {
		propertyManager.AddExpPoint(expPoint, reason, reasonText)
	}
	if exp > 0 {
		propertyManager.AddExp(exp, reason, reasonText)
	}
	propertylogic.SnapChangedProperty(p)
	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffAddExp, event.EventListenerFunc(buffAddExp))
}
