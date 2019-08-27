package player

import (
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	discountdiscounttypes "fgame/fgame/game/welfare/discount/discount/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeCommon, playerwelfare.ActivityObjInfoRefreshHandlerFunc(discountRefreshInfo))
}

//限时礼包-刷新
func discountRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime != 0 && now > endTime {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动跨天刷新,不是活动时间")
		return
	}

	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		pl := obj.GetPlayer()
		info := obj.GetActivityData().(*discountdiscounttypes.DiscountInfo)
		info.BuyRecord = map[int32]int32{}
		info.DiscountDay = welfarelogic.CountCurActivityDay(obj.GetGroupId())
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		welfareManager.UpdateObj(obj)
	}
	return
}
