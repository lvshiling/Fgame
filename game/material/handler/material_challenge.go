package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	materiallogic "fgame/fgame/game/material/logic"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MATERIAL_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerMaterialChallenge))
}

//材料副本挑战请求
func handlerMaterialChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("material:处理材料副本挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMaterialChallenge)
	typ := csMsg.GetMaterialType()

	materialType := materialtypes.MaterialType(typ)
	//验证参数
	if !materialType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"materialType": materialType,
			}).Warn("material:材料副本挑战请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = materialChallenge(tpl, materialType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     tpl.GetId(),
				"materialType": materialType,
				"err":          err,
			}).Error("material:处理材料副本挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     tpl.GetId(),
			"materialType": materialType,
		}).Debug("material：处理材料副本挑战请求完成")

	return
}

//仙府挑战逻辑
func materialChallenge(pl player.Player, materialType materialtypes.MaterialType) (err error) {
	return materiallogic.HandlePlayerMaterialChallenge(pl, materialType)
}
