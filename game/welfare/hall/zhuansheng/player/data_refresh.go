package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	hallzhuanshengtemplate "fgame/fgame/game/welfare/hall/zhuansheng/template"
	hallzhuanshengtypes "fgame/fgame/game/welfare/hall/zhuansheng/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeZhaunSheng, playerwelfare.ActivityObjInfoRefreshHandlerFunc(zhuanShengRefreshInfo))
}

//转生冲刺-刷新
func zhuanShengRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {

	//发未领取奖励邮件
	advancedRewEnd(obj)
	return
}

//进阶奖励结束
func advancedRewEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		return
	}
	info := obj.GetActivityData().(*hallzhuanshengtypes.ZhuanShengInfo)
	if info.IsMail {
		return
	}

	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	lastUpdateTime := obj.GetUpdateTime()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	groupTemp := groupInterface.(*hallzhuanshengtemplate.GroupTemplateZhuanSheng)
	rewTempList := groupTemp.GetCanRewTempList(info.ZhuanSheng, info.RewRecord)
	for _, temp := range rewTempList {
		rewZhuanSheng := temp.Value1

		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		zhuanShengStr := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", rewZhuanSheng))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityHallZhuanShengEndContent), acName, zhuanShengStr)
		itemDataList := welfarelogic.ConvertToItemData(temp.GetEmailRewItemMap(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, itemDataList)
		info.AddRecord(rewZhuanSheng)
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsMail = true
	welfareManager.UpdateObj(obj)

	return
}
