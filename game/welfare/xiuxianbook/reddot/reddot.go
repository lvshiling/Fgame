package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiuxianbooklogic "fgame/fgame/game/welfare/xiuxianbook/logic"
	xiuxianbooktemplate "fgame/fgame/game/welfare/xiuxianbook/template"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength, reddot.HandlerFunc(handleRedDotxiuxianbook))
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, reddot.HandlerFunc(handleRedDotxiuxianbook))
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar, reddot.HandlerFunc(handleRedDotxiuxianbook))
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, reddot.HandlerFunc(handleRedDotxiuxianbook))
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, reddot.HandlerFunc(handleRedDotxiuxianbook))
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, reddot.HandlerFunc(handleRedDotxiuxianbook))
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, reddot.HandlerFunc(handleRedDotxiuxianbook))
	reddot.Register(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, reddot.HandlerFunc(handleRedDotxiuxianbook))
}

func handleRedDotxiuxianbook(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	welfareTemplateService := welfaretemplate.GetWelfareTemplateService()
	groupInterface := welfareTemplateService.GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*xiuxianbooktemplate.GroupTemplateXiuxianBook)
	for _, opentemp := range groupTemp.GetOpenTempMap() {
		can, _, _ := xiuxianbooklogic.IsCanReceiceRew(obj, opentemp)
		if can {
			return true
		}
	}
	return false
}
