package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_COMMON_INFO_TYPE), dispatch.HandlerFunc(handlerCommonInfo))
}

//处理获取运营活动信息
func handlerCommonInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取运营活动请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityGetInfo)
	groupId := csMsg.GetGroupId()

	err = commonInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理获取运营活动请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理获取运营活动请求完成")

	return
}

//获取运营活动请求逻辑
func commonInfo(pl player.Player, groupId int32) (err error) {
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

	h := welfare.GetInfoGetHandler(timeTemp.GetOpenType(), timeTemp.GetOpenSubType())
	if h == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"type":     timeTemp.GetOpenType(),
				"subType":  timeTemp.GetOpenSubType(),
			}).Warn("welfare:处理运营活动信息请求，处理器不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonHandlerNotExist)
		return
	}

	err = h.GetInfo(pl, groupId)
	return
}
