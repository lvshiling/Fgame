package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/qixue/pbutil"
	playerqixue "fgame/fgame/game/qixue/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	//TODO:xzk 登陆不下发,功能开启也不下发?
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeQiXueQiang) {
		return
	}

	//泣血枪信息
	qixueManager := pl.GetPlayerDataManager(playertypes.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	qixueInfo := qixueManager.GetQiXueInfo()
	scMsg := pbutil.BuildSCQiXueGet(qixueInfo)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
