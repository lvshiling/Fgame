package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	"fgame/fgame/game/global"
	propertypbutil "fgame/fgame/game/property/pbutil"
	propertytypes "fgame/fgame/game/property/types"
	playerweek "fgame/fgame/game/week/player"
	weektypes "fgame/fgame/game/week/types"
)

func BuildSCWeekInfo(weekInfoMap map[weektypes.WeekType]*playerweek.WeekData) *uipb.SCWeekInfo {
	scMsg := &uipb.SCWeekInfo{}
	scMsg.WeekInfo = buildWeekInfoList(weekInfoMap)
	return scMsg
}

func BuildSCWeekBuy(weekType weektypes.WeekType, expireTiem int64, itemMap map[int32]int32, rd *propertytypes.RewData, weekInfo *playerweek.WeekData) *uipb.SCWeekBuy {
	scMsg := &uipb.SCWeekBuy{}
	scMsg.ExpireTime = &expireTiem
	weekTypeInt := int32(weekType)
	scMsg.WeekType = &weekTypeInt
	scMsg.DropInfo = droppbutil.BuildSimpleDropInfoList(itemMap)
	scMsg.RewInfo = propertypbutil.BuildRewProperty(rd)
	scMsg.WeekInfo = buildWeekInfo(weekType, weekInfo)
	return scMsg
}

func BuildSCWeekReceiveRew(weekType weektypes.WeekType, weekInfo *playerweek.WeekData, rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCWeekReceiveRew {
	scMsg := &uipb.SCWeekReceiveRew{}
	scMsg.DropInfo = droppbutil.BuildSimpleDropInfoList(itemMap)
	scMsg.RewInfo = propertypbutil.BuildRewProperty(rd)
	scMsg.WeekInfo = buildWeekInfo(weekType, weekInfo)

	return scMsg
}

func buildWeekInfo(weekType weektypes.WeekType, weekInfo *playerweek.WeekData) *uipb.WeekInfo {
	info := &uipb.WeekInfo{}
	dayInt := weekInfo.GetCycleDay()
	weekTypeInt := int32(weekType)
	now := global.GetGame().GetTimeService().Now()
	expireTime := weekInfo.GetExpireTime()
	isReceive := weekInfo.IsReceiveRewards(now)

	info.IsReceive = &isReceive
	info.ExpireTime = &expireTime
	info.WeekType = &weekTypeInt
	info.DayInt = &dayInt

	return info
}

func buildWeekInfoList(weekInfoMap map[weektypes.WeekType]*playerweek.WeekData) (infoList []*uipb.WeekInfo) {
	for weekType, weekInfo := range weekInfoMap {
		infoList = append(infoList, buildWeekInfo(weekType, weekInfo))
	}

	return infoList
}
