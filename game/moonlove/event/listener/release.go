package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	moonloveeventtypes "fgame/fgame/game/moonlove/event/types"
	"fgame/fgame/game/moonlove/pbutil"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
)

//月下情缘移动，解除双人赏月
func release(target event.EventTarget, data event.EventData) (err error) {
	eventData := data.(*moonloveeventtypes.MoonloveDoubleReleaseEventData)
	player1Id := eventData.GetPlayer1()
	player2Id := eventData.GetPlayer2()
	player1 := player.GetOnlinePlayerManager().GetPlayerById(player1Id)
	player2 := player.GetOnlinePlayerManager().GetPlayerById(player2Id)
	// 移除光效
	buffId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDaBaoBuff))

	if player1 != nil {
		player1MoonloveViewDoubleRelease := pbutil.BuildMoonloveViewDoubleRelease(player2Id)
		player1.SendMsg(player1MoonloveViewDoubleRelease)
		scenelogic.RemoveBuff(player1, buffId)
	}

	if player2 != nil {
		player2MoonloveViewDoubleRelease := pbutil.BuildMoonloveViewDoubleRelease(player1Id)
		player2.SendMsg(player2MoonloveViewDoubleRelease)
		scenelogic.RemoveBuff(player2, buffId)
	}

	return
}

func init() {
	gameevent.AddEventListener(moonloveeventtypes.EventTypeMoonloveDoubleRelease, event.EventListenerFunc(release))
}
