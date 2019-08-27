package listener

import (
	"fgame/fgame/core/event"
	consttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	playerfourgod "fgame/fgame/game/fourgod/player"
	fourgodtemplate "fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/mathutils"
)

//玩家死亡钥匙减半
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	pl, ok := bo.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	attackId := data.(int64)
	sp := s.GetPlayer(attackId)
	spl, ok := sp.(player.Player)
	if !ok {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	itemNum := manager.KeyHalve()
	if itemNum <= 0 {
		return
	}

	itemId := int32(consttypes.FourGodKey)
	fourGodTemplate := fourgodtemplate.GetFourGodTemplateService().GetFourGodConstTemplate()
	minStack := fourGodTemplate.MinStack
	maxStack := fourGodTemplate.MaxStack + 1
	protectedTime := fourGodTemplate.ProtectedTime
	existTime := fourGodTemplate.ExistTime

	stack := int32(mathutils.RandomRange(int(minStack), int(maxStack)))
	scenelogic.CustomItemDrop(s, pl.GetPosition(), spl.GetId(), itemId, itemNum, stack, protectedTime, existTime)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
