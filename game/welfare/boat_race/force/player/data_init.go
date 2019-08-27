package player

import (
	playertypes "fgame/fgame/game/player/types"
	boatraceforcetypes "fgame/fgame/game/welfare/boat_race/force/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 赛龙舟战力
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeBoatRace, welfaretypes.OpenActivityDefaultSubTypeDefault, playerwelfare.ActivityObjInfoInitFunc(boatRaceForceInitInfo))
}

func boatRaceForceInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*boatraceforcetypes.BoatRaceForceInfo)
	pl := obj.GetPlayer()
	//初始化战力，当前战力和上次活动的最高历史战力相比
	initForce := pl.GetForce()
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(obj.GetGroupId())
	relationGroupList := timeTemp.GetRelationToGroupList()
	if len(relationGroupList) == 1 {
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		lastObj := welfareManager.GetOpenActivity(relationGroupList[0])
		if lastObj != nil {
			lastInfo := lastObj.GetActivityData().(*boatraceforcetypes.BoatRaceForceInfo)
			if initForce < lastInfo.MaxForce {
				initForce = lastInfo.MaxForce
			}
		}
	}
	info.StartForce = initForce
	info.MaxForce = initForce
}
