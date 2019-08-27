package reddot

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/reddot/reddot"
	madeexptypes "fgame/fgame/game/welfare/made/exp/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMade, welfaretypes.OpenActivityMadeSubTypeResource, reddot.HandlerFunc(handleRedDotMadeRes))
}

//经验炼制红点
func handleRedDotMadeRes(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	madeTemp := welfaretemplate.GetWelfareTemplateService().GetMadeTemplate(groupId, pl.GetLevel())
	if madeTemp == nil {
		return
	}

	// 次数
	info := obj.GetActivityData().(*madeexptypes.MadeInfo)
	needGold := int64(madeTemp.GetNeedCost(info.Times))
	if needGold < 0 {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, false) {
		return
	}

	isNotice = true
	return
}
