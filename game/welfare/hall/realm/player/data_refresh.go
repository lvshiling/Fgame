package player

import (
	"fgame/fgame/common/lang"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	hallrealmtypes "fgame/fgame/game/welfare/hall/realm/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeRealm, playerwelfare.ActivityObjInfoRefreshHandlerFunc(welfareRealmRefreshInfo))
}

//天劫塔冲刺-刷新
func welfareRealmRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
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

	info := obj.GetActivityData().(*hallrealmtypes.WelfareRealmChallengeInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	realmEnd(obj)
	return
}

//天劫塔冲刺结束
func realmEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*hallrealmtypes.WelfareRealmChallengeInfo)
	realmLevel := info.Level
	recordList := info.RewRecord

	if realmLevel < 1 {
		return
	}

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needLevel := temp.Value1
		if needLevel > realmLevel {
			continue
		}

		isReceive := false
		for _, record := range recordList {
			if temp.Value1 == record {
				isReceive = true
			}
		}

		if !isReceive {
			title := lang.GetLangService().ReadLang(lang.EmailOpenActivityRealmTitle)
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityRealmContent))
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)
		}
	}

	info.IsEmail = true
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.UpdateObj(obj)

	return
}
