package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	moonloveeventtypes "fgame/fgame/game/moonlove/event/types"
	"fgame/fgame/game/moonlove/pbutil"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
)

//月下情缘组成双人赏月
func combine(target event.EventTarget, data event.EventData) (err error) {
	eventData := data.(*moonloveeventtypes.MoonloveDoubleCombineEventData)
	player1Id := eventData.GetPlayer1()
	player2Id := eventData.GetPlayer2()
	player1 := player.GetOnlinePlayerManager().GetPlayerById(player1Id)
	player2 := player.GetOnlinePlayerManager().GetPlayerById(player2Id)

	player1MoonloveViewDoubleCombine := pbutil.BuildMoonloveViewDoubleState(player2Id, true)
	player1.SendMsg(player1MoonloveViewDoubleCombine)

	player2MoonloveViewDoubleCombine := pbutil.BuildMoonloveViewDoubleState(player1Id, true)
	player2.SendMsg(player2MoonloveViewDoubleCombine)

	// 双月光效
	buffId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMoonloveBuff))
	scenelogic.AddBuff(player1, buffId, player1.GetId(), common.MAX_RATE)
	scenelogic.AddBuff(player2, buffId, player2.GetId(), common.MAX_RATE)

	return
}

func init() {
	gameevent.AddEventListener(moonloveeventtypes.EventTypeMoonloveDoubleCombine, event.EventListenerFunc(combine))
}
