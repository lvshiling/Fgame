package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	tongtiantatemplate "fgame/fgame/game/welfare/tongtianta/template"
	tongtiantatypes "fgame/fgame/game/welfare/tongtianta/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeLingTong, reddot.HandlerFunc(tongTianTaReddot))
	reddot.Register(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeMingGe, reddot.HandlerFunc(tongTianTaReddot))
	reddot.Register(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeTuLong, reddot.HandlerFunc(tongTianTaReddot))
	reddot.Register(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeShengHen, reddot.HandlerFunc(tongTianTaReddot))
	reddot.Register(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeZhenFa, reddot.HandlerFunc(tongTianTaReddot))
	reddot.Register(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeBaby, reddot.HandlerFunc(tongTianTaReddot))
	reddot.Register(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeDianXing, reddot.HandlerFunc(tongTianTaReddot))
}

func tongTianTaReddot(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)

	groupTempI := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(obj.GetGroupId())
	if groupTempI == nil {
		return
	}
	groupTemp := groupTempI.(*tongtiantatemplate.GroupTemplateTongTianTa)
	nearForceTemp := groupTemp.GetTongTianTaTemplateByNearForce(info.MinForce)
	nearForce := int32(0)
	if nearForceTemp != nil {
		nearForce = nearForceTemp.Value1
	}

	openTempList := groupTemp.GetTongTianTaForceTemplateListByForce(nearForce, info.MaxForce)
	for _, temp := range openTempList {
		force := temp.Value1
		goldNum := temp.Value2

		if info.IsAlreadyReceiveByForce(force) {
			continue
		}
		if !info.IsEnoughChargeNum(goldNum) {
			continue
		}
		isNotice = true
		return
	}

	return
}
