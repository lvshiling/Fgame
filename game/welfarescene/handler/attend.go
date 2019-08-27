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
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	qiyutemplate "fgame/fgame/game/welfare/huhu/qiyu/template"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"
	"fgame/fgame/game/welfarescene/pbutil"
	welfarescenetemplate "fgame/fgame/game/welfarescene/template"
	"fgame/fgame/game/welfarescene/welfarescene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WELFARE_SCENE_ATTEND_TYPE), dispatch.HandlerFunc(handlerWelfareSceneChallenge))
}

//运营活动副本挑战
func handlerWelfareSceneChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("welfarescene:处理运营活动副本挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWelfareSceneAttend)
	groupId := csMsg.GetGroupId()

	err = welfaresceneChallenge(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"groupId":  groupId,
				"err":      err,
			}).Error("welfarescene:处理运营活动副本挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"groupId":  groupId,
		}).Debug("welfarescene:处理运营活动副本挑战请求完成")

	return
}

//运营活动副本挑战逻辑
func welfaresceneChallenge(pl player.Player, groupId int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	//  当前是否处于活动时间内
	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfarescene:运营活动副本挑战请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfarescene:运营活动副本挑战请求，活动模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	qiYuGroupTemp := groupInterface.(*qiyutemplate.GroupTemplateQiYuDao)
	wsTempId := qiYuGroupTemp.GetFirstOpenTemp().Value1
	wsTemp := welfarescenetemplate.GetWelfareSceneTemplateService().GetWelfareSceneTemplate(wsTempId)
	if wsTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfarescene:运营活动副本挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//进入场景
	s := welfarescene.GetWelfareSceneService().GetWelfareScene(groupId)
	if s == nil {
		_, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		s = welfarescene.GetWelfareSceneService().CreateWelfareScene(groupId, wsTempId, endTime)
	}
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfarescene:运营活动副本挑战请求，场景不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if !scenelogic.PlayerEnterSingleFuBenScene(pl, s) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfarescene:运营活动副本挑战请求，进入场景失败")
		return
	}

	scMsg := pbutil.BuildSCWelfareSceneAttend(groupId)
	pl.SendMsg(scMsg)
	return
}
