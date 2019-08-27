package info

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	qiyutemplate "fgame/fgame/game/welfare/huhu/qiyu/template"
	"fgame/fgame/game/welfare/pbutil"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	welfarescenescene "fgame/fgame/game/welfarescene/scene"
	"fgame/fgame/game/welfarescene/welfarescene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeQiYu, welfare.InfoGetHandlerFunc(handlerQiYuInfo))
}

//奇遇副本信息请求
func handlerQiYuInfo(pl player.Player, groupId int32) (err error) {
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

	qiYuGroupTemp, ok := groupInterface.(*qiyutemplate.GroupTemplateQiYuDao)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfarescene:运营活动副本挑战请求，不是奇遇岛活动id")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	wsTempId := qiYuGroupTemp.GetFirstOpenTemp().Value1

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	qiYuScene := welfarescene.GetWelfareSceneService().GetWelfareScene(groupId)
	if qiYuScene == nil {
		qiYuScene = welfarescene.GetWelfareSceneService().CreateWelfareScene(groupId, wsTempId, endTime)
	}
	if qiYuScene == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"wsTempId": wsTempId,
			}).Warn("welfarescene:运营活动副本挑战请求，场景不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	sd := qiYuScene.SceneDelegate().(welfarescenescene.WelfareQiYuBossSceneData)
	npcMap := sd.GetBossMap()

	scMsg := pbutil.BuildSCOpenActivityGetInfoQiYuDao(groupId, startTime, endTime, nil, npcMap)
	pl.SendMsg(scMsg)
	return
}
