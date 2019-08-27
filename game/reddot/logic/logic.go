package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/reddot/pbutil"
	"fgame/fgame/game/reddot/reddot"
	reddottypes "fgame/fgame/game/reddot/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func CheckReddotByType(pl player.Player, checkMap map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]int32) {
	//检查抽奖红点
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	var objList []*playerwelfare.PlayerOpenActivityObject
	allTimeList := welfaretemplate.GetWelfareTemplateService().GetAllActivityTimeTemplate()
	for _, tiemTemp := range allTimeList {
		typ := tiemTemp.GetOpenType()
		subType := tiemTemp.GetOpenSubType()
		groupId := tiemTemp.Group

		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		subTypeMap, ok := checkMap[typ]
		if !ok {
			continue
		}
		_, ok = subTypeMap[subType]
		if !ok {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		objList = append(objList, obj)
	}

	ReddotNoticeChanged(pl, objList)
}

func ReddotNoticeChanged(pl player.Player, objList []*playerwelfare.PlayerOpenActivityObject) {
	var redInfoList []*reddottypes.RedDotInfo
	for _, obj := range objList {
		isReddot, isChanged := reddot.Handle(pl, obj)
		if isChanged {
			redData := reddottypes.NewRedDotInfo(obj.GetGroupId(), isReddot)
			redInfoList = append(redInfoList, redData)
		}
	}

	if len(redInfoList) > 0 {
		scMsg := pbutil.BuildSCActivityNoticeChanged(redInfoList)
		pl.SendMsg(scMsg)
	}
}

func CheckReddotOnTimeAll(pl player.Player) {
	//检查所有在活动时间内红点
	var groupIdList []int32
	allTimeList := welfaretemplate.GetWelfareTemplateService().GetAllActivityTimeTemplate()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	for _, timeTemp := range allTimeList {
		groupId := timeTemp.Group
		typ := timeTemp.GetOpenType()
		subType := timeTemp.GetOpenSubType()
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		isReddot, _ := reddot.Handle(pl, obj)
		if isReddot {
			groupIdList = append(groupIdList, groupId)
		}
	}

	if len(groupIdList) > 0 {
		scMsg := pbutil.BuildSCActivityNoticeOnTimeAll(groupIdList)
		pl.SendMsg(scMsg)
	}
	return
}
