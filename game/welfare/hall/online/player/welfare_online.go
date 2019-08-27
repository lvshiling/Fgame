package player

import (
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	hallonlinetypes "fgame/fgame/game/welfare/hall/online/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeOnline, playerwelfare.ActivityObjInfoRefreshHandlerFunc(welfareOnlineRefreshInfo))
}

//在线抽奖-刷新
func welfareOnlineRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		pl := obj.GetPlayer()
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

		info := obj.GetActivityData().(*hallonlinetypes.WelfareOnlineInfo)
		info.DrawTimes = 0
		welfareManager.UpdateObj(obj)
	}
	return
}
