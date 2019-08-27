package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	relivelogic "fgame/fgame/game/relive/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_PLAYER_RELIVE_TYPE), dispatch.HandlerFunc(handleRelive))
}

//处理对象复活
func handleRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理对象复活消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csPlayerRelive := msg.(*scenepb.CSPlayerRelive)
	reliveTypeInt := csPlayerRelive.GetReliveType()
	reliveType := scenetypes.ReliveType(reliveTypeInt)
	autoBuy := csPlayerRelive.GetAutoBuy()
	if !reliveType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"reliveType": reliveTypeInt,
			}).Warn("scene:处理对象复活消息,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	//玩家复活
	err = playerRelive(tpl, reliveType, autoBuy)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("scene:处理对象复活消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("scene:处理对象复活消息,完成")
	return
}

func playerRelive(pl scene.Player, reliveType scenetypes.ReliveType, autoBuy bool) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:处理对象复活消息,场景为空")
		return
	}

	if !pl.IsDead() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:处理对象复活消息,还没死亡")
		return
	}
	//判断场景是否可以复活
	mapTemplate := s.MapTemplate()
	if reliveType.Mask()&mapTemplate.LimitReborn == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"reliveType":  reliveType,
				"limitReborn": mapTemplate.LimitReborn,
			}).Warn("scene:处理对象复活消息,限制复活方式")
		return
	}

	switch reliveType {
	case scenetypes.ReliveTypeImmediate:
		reliveHandler := scene.GetReliveHandler(s.MapTemplate().GetMapType())
		if reliveHandler == nil {
			//处理复活
			relivelogic.Relive(pl, autoBuy)
		} else {
			reliveHandler.Relive(pl, autoBuy)
		}
		break
	case scenetypes.ReliveTypeBack:
		scenelogic.RebornBack(pl)
		break
	case scenetypes.ReliveTypeEnterPoint:
		pl.Reborn(mapTemplate.GetBornPos())
		break
	case scenetypes.ReliveTypeRelivePoint:
		//特殊处理
		scenelogic.EnterRelivePoint(pl)
		break
	}

	return
}
