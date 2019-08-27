package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	advancedrewrewmaxtemplate "fgame/fgame/game/welfare/advancedrew/rew_max/template"
	advancedrewrewmaxtypes "fgame/fgame/game/welfare/advancedrew/rew_max/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewMax, playerwelfare.ActivityObjInfoRefreshHandlerFunc(rewMaxAdvancedRefreshInfo))
}

//进阶奖励-刷新
func rewMaxAdvancedRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	groupId := obj.GetGroupId()
	pl := obj.GetPlayer()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,活动未结束")
		return
	}
	info := obj.GetActivityData().(*advancedrewrewmaxtypes.AdvancedRewMaxInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发未领取奖励邮件
	lastUpdateTime := obj.GetUpdateTime()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewrewmaxtemplate.GroupTemplateRewMax)
	for _, temp := range groupTemp.GetRewTempDescList() {
		needAdvancedNum := temp.Value2
		needChargeNum := temp.Value3
		if info.IsCanReceiveRewards(needAdvancedNum, needChargeNum) {
			title := temp.Label
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			moduleName := chatlogic.FormatMailKeyWordNoticeStr(info.RewType.String())
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityAdvancedRewContent), acName, moduleName, needAdvancedNum)
			itemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, itemDataList)
			info.AddRecord(needAdvancedNum)
		}

		// 初始条件最近的档次
		if needAdvancedNum <= info.InitAdvancedNum {
			break
		}
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
