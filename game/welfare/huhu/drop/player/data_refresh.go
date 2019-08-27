package player

import (
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	huhudroptypes "fgame/fgame/game/welfare/huhu/drop/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeDrop, playerwelfare.ActivityObjInfoRefreshHandlerFunc(huhuRefreshInfo))
}

//虎虎生风-刷新
func huhuRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime != 0 && now > endTime {
		return
	}

	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		info := obj.GetActivityData().(*huhudroptypes.HuHuInfo)
		info.CurDayDropNum = 0
		pl := obj.GetPlayer()
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		welfareManager.UpdateObj(obj)

	}

	return
}
