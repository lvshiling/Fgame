package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	drewcrazyboxtemplate "fgame/fgame/game/welfare/drew/crazy_box/template"
	drewcrazyboxtypes "fgame/fgame/game/welfare/drew/crazy_box/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_CRAZY_BOX_INFO_TYPE), dispatch.HandlerFunc(handlerCrazyBoxGetInfo))
}

//处理获取疯狂宝箱信息
func handlerCrazyBoxGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取疯狂宝箱请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityCrazyBoxInfo)
	groupId := csMsg.GetGroupId()

	err = getCrazyBoxInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取疯狂宝箱请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取疯狂宝箱请求完成")

	return
}

//获取疯狂宝箱请求逻辑
func getCrazyBoxInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeCrazyBox
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参数无效,活动时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*drewcrazyboxtemplate.GroupTemplateCrazyBox)

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivityIfNotCreate(timeTemp.GetOpenType(), timeTemp.GetOpenSubType(), groupId)
	logList := welfare.GetWelfareService().GetCrazyBoxLogByTime(groupId, 0)

	info := obj.GetActivityData().(*drewcrazyboxtypes.CrazyBoxInfo)
	curBoxLevel, boxLeftTimes := groupTemp.GetCrazyBoxArg(info.AttendTimes)
	openActivityTemp := groupTemp.GetOpenActivityCrazyBox(curBoxLevel)
	if openActivityTemp == nil {
		return
	}
	curBoxTimes := groupTemp.GetOpenActivityCrazyBoxUpTimes(curBoxLevel) - boxLeftTimes
	totalTimes := groupTemp.GetCrazyBoxTotalTimes(info.GoldNum)
	drewTimes := totalTimes - info.AttendTimes
	scMsg := pbutil.BuildSCOpenActivityCrazyBoxInfo(obj, logList, groupId, startTime, endTime, drewTimes, curBoxLevel, curBoxTimes)
	pl.SendMsg(scMsg)
	return
}
