package player

import (
	"fgame/fgame/common/lang"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	advancedblessfeedbacktypes "fgame/fgame/game/welfare/advanced/bless_feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeBlessFeedback, playerwelfare.ActivityObjInfoRefreshHandlerFunc(blessAdvancedRefreshInfo))
}

//进阶祝福丹大放送-刷新（废弃）
func blessAdvancedRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		//发未领取奖励邮件
		advancedBlessChanged(obj)
	}
	return
}

//进阶祝福丹放送变更（废弃）
func advancedBlessChanged(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info := obj.GetActivityData().(*advancedblessfeedbacktypes.BlessAdvancedInfo)
	advancedNum := info.AdvancedNum
	blessDay := info.BlessDay
	recordList := info.RewRecord
	lastUpdateTime := obj.GetUpdateTime()

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		//对应的祝福日
		curDay := temp.Value1
		if curDay != blessDay {
			continue
		}

		needAdvancedNum := temp.Value2
		if needAdvancedNum > advancedNum {
			continue
		}

		isReceive := false
		for _, record := range recordList {
			if needAdvancedNum == record {
				isReceive = true
			}
		}

		if !isReceive {
			title := lang.GetLangService().ReadLang(lang.EmailOpenActivityAdvancedBlessTitle)
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityAdvancedBlessContent))
			itemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, itemDataList)
		}
	}

	//重置
	initHandler := playerwelfare.GetInfoInitHandler(obj.GetActivityType(), obj.GetActivitySubType())
	if initHandler == nil {
		return
	}
	initHandler.InitInfo(obj)
	welfareManager.UpdateObj(obj)
	return
}
