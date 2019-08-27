package player

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	hallupleveltypes "fgame/fgame/game/welfare/hall/uplevel/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeUpLevel, playerwelfare.ActivityObjInfoRefreshHandlerFunc(welfareLevelRefreshInfo))
}

//福利大厅升级-刷新
func welfareLevelRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime != 0 && now > endTime {
		//发送未领取奖励邮件
		welfareUplevelEnd(obj)
	}

	return
}

//升级活动结束
func welfareUplevelEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*hallupleveltypes.WelfareUplevelInfo)

	if info.IsEmail {
		return
	}

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needLevel := temp.Value1
		if needLevel > pl.GetLevel() {
			continue
		}

		if !info.IsReceive(needLevel) {
			title := coreutils.FormatNoticeStr(temp.Label)
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityWelfareUplevelMailContent), acName, needLevel)
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)

			info.AddRecord(needLevel)
		}
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
