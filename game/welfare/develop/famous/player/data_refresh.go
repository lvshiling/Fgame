package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	developfamoustypes "fgame/fgame/game/welfare/develop/famous/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeDevelop, welfaretypes.OpenActivityDefaultSubTypeDefault, playerwelfare.ActivityObjInfoRefreshHandlerFunc(developFameEndRefreshInfo))
}

//名人普-刷新
func developFameEndRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//跨天
	now := global.GetGame().GetTimeService().Now()
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*developfamoustypes.DevelopFameInfo)
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}
	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

	if !isSame {
		info.DayFavorableNum = 0
		welfareManager.UpdateObj(obj)
	}

	//结束
	endTime := obj.GetEndTime()
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

	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	developFamousFeedEnd(obj)
	return
}

//名人培养结束
func developFamousFeedEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*developfamoustypes.DevelopFameInfo)
	endTime := obj.GetEndTime()

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needFavorableNum := temp.Value1
		if !info.IsCanReceiveRewards(needFavorableNum) {
			continue
		}
		famousTemp := welfaretemplate.GetWelfareTemplateService().GetFamousTemplate(groupId)
		title := temp.Label

		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		famousName := chatlogic.FormatMailKeyWordNoticeStr(famousTemp.Name)
		favorable := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", info.FavorableNum))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackDevelopFameEndMailContent), acName, famousName, favorable)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)
		info.AddRecord(needFavorableNum)
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)

	return
}
