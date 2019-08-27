package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	guidereplica "fgame/fgame/game/guidereplica/guidereplica"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GUIDE_REPLICA_PLAYER_COMMON_OPERATE_TYPE), dispatch.HandlerFunc(handlerCommonOperate))
}

//处理通用接口操作
func handlerCommonOperate(s session.Session, msg interface{}) (err error) {
	log.Debug("guidereplica:引导副本请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGuideReplicaPlayerCommonOperate)
	guideTypeInt := csMsg.GetGuideType()

	//参数验证
	guideType := guidereplicatypes.GuideReplicaType(guideTypeInt)
	if !guideType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"guideType": guideType,
			}).Warn("gm:引导副本类型,错误")
		return
	}

	err = commonOperate(tpl, guideType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("guidereplica:引导副本请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("guidereplica:引导副本请求完成")

	return
}

//处理通用接口操作请求逻辑
func commonOperate(pl player.Player, typ guidereplicatypes.GuideReplicaType) (err error) {
	h := guidereplica.GetGuideCommonOperate(pl, typ)
	if h == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"type":     typ,
			}).Warn("guidereplica:引导副本请求，处理器不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonHandlerNotExist)
		return
	}

	err = h.GuideCommonOperate(pl)
	return
}
