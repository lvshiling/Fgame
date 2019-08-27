package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	"time"
)

const (
	guaJiGetItemTaskTime = 5 * time.Second
)

type GuaJiGetItemTask struct {
	p scene.Player
}

func (t *GuaJiGetItemTask) Run() {

	//TODO:zrc 限制个数
	for _, nei := range t.p.GetNeighbors() {
		switch obj := nei.(type) {
		case scene.DropItem:
			if !t.p.IfCanGetDropItem(obj) {
				continue
			}
			gameevent.Emit(battleeventtypes.EventTypeBattlePlayerGuaJiGetItem, t.p, obj)
			break
		}
	}

}

func (t *GuaJiGetItemTask) ElapseTime() time.Duration {
	return guaJiGetItemTaskTime
}

func CreateGuaJiGetItemTask(p scene.Player) *GuaJiGetItemTask {
	t := &GuaJiGetItemTask{
		p: p,
	}
	return t
}
