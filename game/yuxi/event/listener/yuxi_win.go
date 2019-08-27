package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	yuxieventtypes "fgame/fgame/game/yuxi/event/types"
	yuxilogic "fgame/fgame/game/yuxi/logic"
	"fgame/fgame/game/yuxi/pbutil"
	yuxiscene "fgame/fgame/game/yuxi/scene"
)

//玉玺之战结束
func yuXiFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(yuxiscene.YuXiSceneData)
	if !ok {
		return
	}
	allianceId, ok := data.(int64)
	if !ok {
		return
	}

	//广播
	scMsg := pbutil.BuildSCYuXiWinnerBroadcast(allianceId)
	sd.GetScene().BroadcastMsg(scMsg)

	// 获胜
	yuxilogic.WinYuXiWar(allianceId)
	return
}

func init() {
	gameevent.AddEventListener(yuxieventtypes.EventTypeYuXiWin, event.EventListenerFunc(yuXiFinish))
}
