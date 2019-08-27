package player

import (
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	madeexptypes "fgame/fgame/game/welfare/made/exp/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeMade, welfaretypes.OpenActivityMadeSubTypeResource, playerwelfare.ActivityObjInfoRefreshHandlerFunc(madeResRefreshInfo))
}

//炼制资源-刷新
func madeResRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		info := obj.GetActivityData().(*madeexptypes.MadeInfo)
		info.Times = 0
		pl := obj.GetPlayer()
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		welfareManager.UpdateObj(obj)
	}
	return
}
