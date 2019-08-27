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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_COMMON_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerReceive))
}

//处理领取运营活动
func handlerReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理运营活动领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityReceiveRew)
	rewId := csMsg.GetRewId()

	err = receive(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理运营活动领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理运营活动领取奖励请求完成")

	return
}

//运营活动领取奖励请求逻辑
func receive(pl player.Player, rewId int32) (err error) {
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取运营活动奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	groupId := openTemp.Group
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
				"groupId":  groupId,
			}).Warn("welfare:领取运营活动奖励请求，时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	h := welfare.GetReceiveHandler(timeTemp.GetOpenType(), timeTemp.GetOpenSubType())
	if h == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"type":     timeTemp.GetOpenType(),
				"subType":  timeTemp.GetOpenSubType(),
			}).Warn("welfare:处理运营活动领取请求，处理器不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonHandlerNotExist)
		return
	}

	err = h.ReceiveRew(pl, rewId)
	return
}
