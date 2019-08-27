package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	tongtiantatemplate "fgame/fgame/game/welfare/tongtianta/template"
	tongtiantatypes "fgame/fgame/game/welfare/tongtianta/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeLingTong, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshTongTianTaInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeMingGe, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshTongTianTaInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeTuLong, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshTongTianTaInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeShengHen, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshTongTianTaInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeZhenFa, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshTongTianTaInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeBaby, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshTongTianTaInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeDianXing, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshTongTianTaInfo))

}

func refreshTongTianTaInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	sendEmailGroupEnd(obj)

	refreshTongTianTaInfoAboutForce(obj)
	refreshTongTianTaInfoAboutCharge(obj)
	return
}

// 活动结束奖励未领取发邮件
func sendEmailGroupEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}

	if now < endTime {
		return
	}

	if info.IsEmail {
		return
	}

	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
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

		title := temp.Label
		activityText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%s", title))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailTongTianTa), activityText)

		itemList := temp.GetEmailRewItemDataList()
		itemData := welfarelogic.ConvertToItemDataWithWelfareData(itemList, temp.GetExpireType(), temp.GetExpireTime())
		if len(itemData) != 0 {
			emaillogic.AddEmailItemLevel(pl, title, content, obj.GetUpdateTime(), itemData)
		}

		info.ReceiveSuccess(force)
	}

	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}

// 同步充值
func refreshTongTianTaInfoAboutCharge(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//同步第一天充值
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {

		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			pl := obj.GetPlayer()
			info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)
			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			if info.ChargeNum != int32(chargeManager.GetTodayChargeNum()) {
				info.ChargeNum = int32(chargeManager.GetTodayChargeNum())
				welfareManager.UpdateObj(obj)
			}
		}
	}

	return
}

// 同步战力
func refreshTongTianTaInfoAboutForce(obj *playerwelfare.PlayerOpenActivityObject) {
	if !welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		return
	}

	info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)
	if info.MinForce >= 0 {
		return
	}
	groupTempI := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(obj.GetGroupId())
	if groupTempI == nil {
		return
	}
	subType := groupTempI.GetSubType()
	effectTypeList, ok := tongtiantatypes.TongTianTaSubTypeTOPlayerPropertyEffectType(subType)
	if !ok {
		return
	}
	power := int32(0)
	pl := obj.GetPlayer()
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	for _, effType := range effectTypeList {
		power += int32(propertyManager.GetModuleForce(effType))
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

	info.MinForce = power
	info.MaxForce = power
	welfareManager.UpdateObj(obj)
	return
}
