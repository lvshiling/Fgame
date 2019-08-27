package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	cyclechargesingletemplate "fgame/fgame/game/welfare/cycle/charge_single/template"
	cyclechargesingletypes "fgame/fgame/game/welfare/cycle/charge_single/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleCharge, playerwelfare.ActivityObjInfoRefreshHandlerFunc(cycleSingleChargeRefreshInfo))
}

//每日单笔充值-刷新
func cycleSingleChargeRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//重置
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		info := obj.GetActivityData().(*cyclechargesingletypes.CycleSingleChargeInfo)
		if info.IsEmail {
			log.WithFields(
				log.Fields{
					"groupId": obj.GetGroupId(),
				}).Debugf("活动结束发放邮件,邮件已发送")
			return
		}

		//发未领取奖励邮件
		playerCycleSingleChanged(obj)
	}

	return
}

//每日单笔充值变更
func playerCycleSingleChanged(obj *playerwelfare.PlayerOpenActivityObject) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	lastUpdateTime := obj.GetUpdateTime()
	info := obj.GetActivityData().(*cyclechargesingletypes.CycleSingleChargeInfo)
	rewCycDay := info.CycleDay

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp, ok := groupInterface.(*cyclechargesingletemplate.GroupTemplateCycleSingle)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare：类型错误，转换失败")
	}

	tempList := groupTemp.GetCurDayTempList(rewCycDay)
	for _, temp := range tempList {
		needNum := temp.Value2
		if info.IsCanReceiveRewards(needNum) {
			title := temp.Label
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", needNum))
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleSingleChargeMailContent), acName, chargeText)
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)
		}
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.MaxSingleChargeNum = 0
	info.RewRecord = []int32{}
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	welfareManager.UpdateObj(obj)
	return
}
