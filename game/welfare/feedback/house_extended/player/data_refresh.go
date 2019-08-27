package player

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackhouseextendedtemplate "fgame/fgame/game/welfare/feedback/house_extended/template"
	feedbackhouseextendedtypes "fgame/fgame/game/welfare/feedback/house_extended/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseExtended, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackHouseExtendedRefreshInfo))
}

//房产活动-刷新
func feedbackHouseExtendedRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//跨天
	info := obj.GetActivityData().(*feedbackhouseextendedtypes.FeedbackHouseExtendedInfo)
	groupId := obj.GetGroupId()
	pl := obj.GetPlayer()
	endTime := obj.GetEndTime()
	now := global.GetGame().GetTimeService().Now()
	isEnd := now > endTime

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackhouseextendedtemplate.GroupTemplateHouseExtended)

	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}
	if isSame {
		return
	}

	// 激活礼包未领取
	if !info.IsActivateGift {
		openTemp := groupTemp.GetActivateCanRewTemp(info.ActivateChargeNum)
		if openTemp != nil {
			title := coreutils.FormatNoticeStr(openTemp.Label)
			acName := chatlogic.FormatMailKeyWordNoticeStr(openTemp.Label)
			chargeStr := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", openTemp.Value2))
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackHouseExtendedEndContent), acName, chargeStr)
			rewItemList := openTemp.GetEmailRewItemDataList()
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(rewItemList, openTemp.GetExpireType(), openTemp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, now, newItemDataList)

			info.ReceiveActivateGift()
		}
	}

	// 激活礼包已领取
	if info.IsActivateGift {
		//装修礼包
		if !info.IsUplevelGift {
			openTemp := groupTemp.GetUplevelCanRewTemp(info.UplevelChargeNum, info.CurUplevelGiftLevel)
			if openTemp != nil {
				title := coreutils.FormatNoticeStr(openTemp.Label)
				acName := chatlogic.FormatMailKeyWordNoticeStr(openTemp.Label)
				chargeStr := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", openTemp.Value2))
				econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackHouseExtendedEndContent), acName, chargeStr)
				rewItemList := openTemp.GetEmailRewItemDataList()
				newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(rewItemList, openTemp.GetExpireType(), openTemp.GetExpireTime())
				emaillogic.AddEmailItemLevel(pl, title, econtent, now, newItemDataList)

				info.ReceiveUplevelGift()
			}
		}
		if !isEnd {
			info.CrossDayUplevelGift()
		}
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.UpdateObj(obj)

	return
}
