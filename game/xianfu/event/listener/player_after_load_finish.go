package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/xianfu/pbutil"
	playerxianfu "fgame/fgame/game/xianfu/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	xianfuManager := p.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*playerxianfu.PlayerXinafuDataManager)
	now := global.GetGame().GetTimeService().Now()

	//刷新数据
	err = xianfuManager.RefreshData(now)
	if err != nil {
		return
	}
	//修正经验仙府波数
	xianfuManager.FixXianExpGroupRecord()

	//获取信息List
	xianfuArr := xianfuManager.GetPlayerXianfuInfoList()
	scXianfuGet := pbutil.BuildSCXianfuGet(xianfuArr)
	p.SendMsg(scXianfuGet)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
