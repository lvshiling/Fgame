package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiuxianbooklogic "fgame/fgame/game/welfare/xiuxianbook/logic"
	xiuxianbooktemplate "fgame/fgame/game/welfare/xiuxianbook/template"
	xiuxianbooktypes "fgame/fgame/game/welfare/xiuxianbook/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, playerwelfare.ActivityObjInfoRefreshHandlerFunc(xiuxianBookRefreshInfo))
}

func xiuxianBookRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
	now := global.GetGame().GetTimeService().Now()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:修仙典籍刷新错误，模板不存在")
		return
	}
	groupTemp, ok := groupInterface.(*xiuxianbooktemplate.GroupTemplateXiuxianBook)
	if !ok {
		return
	}
	tempList := groupTemp.GetGroupXiuxianBookList()
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		//同步第一天充值
		if diff == 0 {
			pl := obj.GetPlayer()

			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			if info.ChargeNum != int32(chargeManager.GetTodayChargeNum()) {
				info.ChargeNum = int32(chargeManager.GetTodayChargeNum())
				welfareManager.UpdateObj(obj)
			}
		}

		if info.FirstTimeRewRecord == -1 {
			//计算等级
			needLevel, err := xiuxianbooklogic.CountRecentCanReceiveRewLevel(obj)
			if err != nil {
				return err
			}
			info.FirstTimeRewRecord = needLevel
			welfareManager.UpdateObj(obj)
		}
	} else {
		//邮件补发
		for _, temp := range tempList {
			can, _, needLevel := xiuxianbooklogic.IsCanReceiceRew(obj, temp)
			if can {
				title := temp.Label
				acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
				econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityXiuXianBook), acName)
				lastUpdateTime := obj.GetUpdateTime()
				newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
				emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)
				//增加记录
				info.AddReceiveRecord(needLevel)
				welfareManager.UpdateObj(obj)
			}

		}

	}

	return
}
