package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	discountyunyintemplate "fgame/fgame/game/welfare/discount/yunyin/template"
	discountyunyintypes "fgame/fgame/game/welfare/discount/yunyin/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeYunYin, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshDiscountYunYinData))
}

func refreshDiscountYunYinData(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	info := obj.GetActivityData().(*discountyunyintypes.YunYinInfo)
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
	groupTemp := groupTempI.(*discountyunyintemplate.GroupTemplateDiscountYunYinShop)

	openTempList := groupTemp.GetCanReceiveRewardTemplateList(info.GoldNum)
	for _, temp := range openTempList {
		goldNum := temp.Value1

		if !info.IsCanReceive(goldNum) {
			continue
		}
		if info.IsAlreadyReceive(goldNum) {
			continue
		}

		title := temp.Label
		activityText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%s", title))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailYunYinShop), activityText)

		itemList := temp.GetEmailRewItemDataList()
		itemData := welfarelogic.ConvertToItemDataWithWelfareData(itemList, temp.GetExpireType(), temp.GetExpireTime())
		if len(itemData) != 0 {
			emaillogic.AddEmailItemLevel(pl, title, content, obj.GetUpdateTime(), itemData)
		}
		info.AddReceiveRecord(goldNum)
	}

	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
