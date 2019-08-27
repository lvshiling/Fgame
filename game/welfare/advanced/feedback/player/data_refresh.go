package player

import (
	"fgame/fgame/common/lang"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	advancedfeedbacktypes "fgame/fgame/game/welfare/advanced/feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeFeedback, playerwelfare.ActivityObjInfoRefreshHandlerFunc(advancedRefreshInfo))
}

//进阶返利-刷新(废弃)
func advancedRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		//发未领取奖励邮件
		advancedChanged(obj)
	}
	return
}

//进阶返利变更（废弃）
func advancedChanged(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedfeedbacktypes.AdvancedInfo)
	danNum := info.DanNum
	advancedDay := info.AdvancedDay
	recordList := info.RewRecord
	lastUpdateTime := obj.GetUpdateTime()

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		//对应的充值日
		cycDay := temp.Value1
		if cycDay != advancedDay {
			continue
		}

		needDanNum := temp.Value2
		if needDanNum > danNum {
			continue
		}

		isReceive := false
		for _, record := range recordList {
			if needDanNum == record {
				isReceive = true
			}
		}

		if !isReceive {
			title := lang.GetLangService().ReadLang(lang.EmailOpenActivityAdvancedTitle)
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityAdvancedContent))
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)
		}
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.DanNum = 0
	info.RewRecord = []int32{}
	info.AdvancedDay = welfarelogic.CountCurActivityDay(groupId)
	welfareManager.UpdateObj(obj)

	return
}
