package player

import (
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	rewpoolstypes "fgame/fgame/game/welfare/drew/rew_pools/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeRewPools, playerwelfare.ActivityObjInfoRefreshHandlerFunc(drewPoolsDataRefresh))
}

//充值抽奖-刷新
func drewPoolsDataRefresh(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	info := obj.GetActivityData().(*rewpoolstypes.RewPoolsInfo)
	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	if info.RecordTime == int64(0) {
		info.RecordTime = now
		welfareManager.UpdateObj(obj)
	}
	isSame, err := timeutils.IsSameDay(info.RecordTime, now)
	if err != nil {
		return err
	}
	// 跨天刷新
	if !isSame {
		info.Position = int32(0)
		info.BackTimes = int32(0)
		info.RecordTime = now
		welfareManager.UpdateObj(obj)
	}
	return nil

}
