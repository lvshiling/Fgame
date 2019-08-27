package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	advancedrewexpendreturntypes "fgame/fgame/game/welfare/advancedrew/expend_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeExpendReturn, playerwelfare.ActivityObjInfoRefreshHandlerFunc(advancedExpendReturnRefreshInfo))
}

//升阶消耗返还
func advancedExpendReturnRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,活动未结束")
		return
	}
	info := obj.GetActivityData().(*advancedrewexpendreturntypes.AdvancedExpendReturnInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发未领取奖励邮件
	advancedExpendReturnEnd(obj)
	return
}

//进阶消耗返还结束
func advancedExpendReturnEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedrewexpendreturntypes.AdvancedExpendReturnInfo)
	lastUpdateTime := obj.GetUpdateTime()

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		//对应的进阶类型

		rewType := welfaretypes.AdvancedType(temp.Value1)
		if rewType != info.RewType {
			continue
		}

		needDanNum := temp.Value2
		if !info.IsCanReceiveRewards(needDanNum) {
			continue
		}

		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityAdvancedExpendReturnEndMailContent), acName)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)

		info.AddRecord(needDanNum)
	}

	//重置
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
