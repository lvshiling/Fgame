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
	tongtiantatemplate "fgame/fgame/game/welfare/tongtianta/template"
	tongtiantatypes "fgame/fgame/game/welfare/tongtianta/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeLingTong, welfare.InfoGetHandlerFunc(handleTongTianTaInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeMingGe, welfare.InfoGetHandlerFunc(handleTongTianTaInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeTuLong, welfare.InfoGetHandlerFunc(handleTongTianTaInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeShengHen, welfare.InfoGetHandlerFunc(handleTongTianTaInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeZhenFa, welfare.InfoGetHandlerFunc(handleTongTianTaInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeBaby, welfare.InfoGetHandlerFunc(handleTongTianTaInfo))
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeDianXing, welfare.InfoGetHandlerFunc(handleTongTianTaInfo))
}

func handleTongTianTaInfo(pl player.Player, groupId int32) (err error) {

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
	typ := timeTemp.GetOpenType()
	subType := timeTemp.GetOpenSubType()
	flag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !flag {
		return
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)

	groupTempI := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupTempI == nil {
		return
	}
	groupTemp := groupTempI.(*tongtiantatemplate.GroupTemplateTongTianTa)
	nearForceTemp := groupTemp.GetTongTianTaTemplateByNearForce(info.MinForce)
	nearForce := int32(0)
	if nearForceTemp != nil {
		nearForce = nearForceTemp.Value1
	}

	scOpenActivityTongTianTaInfo := pbutil.BuildSCOpenActivityTongTianTaInfo(nearForce, info.MaxForce, groupId, startTime, endTime, info.Record)
	pl.SendMsg(scOpenActivityTongTianTaInfo)

	scMsg := pbutil.BuildSCOpenActivityFeedbackChargeNotice(groupId, int64(info.ChargeNum))
	pl.SendMsg(scMsg)
	return
}
