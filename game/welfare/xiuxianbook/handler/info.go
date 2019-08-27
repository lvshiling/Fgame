package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	xiuxianbooklogic "fgame/fgame/game/welfare/xiuxianbook/logic"
	xiuxianbooktypes "fgame/fgame/game/welfare/xiuxianbook/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, welfare.InfoGetHandlerFunc(handleXiuxianBookInfo))
}

func handleXiuxianBookInfo(pl player.Player, groupId int32) (err error) {

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:修仙典籍信息请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	//刷新一下数据先
	welfareManager.RefreshActivityDataByGroupId(groupId)

	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:处理运营活动信息请求，时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	obj := welfareManager.GetOpenActivityIfNotCreate(timeTemp.GetOpenType(), timeTemp.GetOpenSubType(), groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Error("welfare:修仙典籍信息请求，活动对象不存在")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityObjectNotExist)
		return
	}
	info := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
	hadRewRecord := info.HasReceiveRecord
	level, maxLevel, err := xiuxianbooklogic.GetCanRewList(obj)
	if err != nil {
		return err
	}
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityGetInfoXiuxianBook(groupId, startTime, endTime, hadRewRecord, level, maxLevel)
	pl.SendMsg(scMsg)

	chargeNum := info.ChargeNum
	scMsg2 := pbutil.BuildSCOpenActivityFeedbackChargeNotice(groupId, int64(chargeNum))
	pl.SendMsg(scMsg2)
	return
}
