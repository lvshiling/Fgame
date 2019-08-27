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
	xianfulogic "fgame/fgame/game/xianfu/logic"

	gamesession "fgame/fgame/game/session"
	xianfutypes "fgame/fgame/game/xianfu/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANFU_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerXianfuChallenge))
}

//秘境仙府挑战请求
func handlerXianfuChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("xianfu:处理秘境仙府挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csXianfuChallenge := msg.(*uipb.CSXianfuChallenge)
	typ := csXianfuChallenge.GetXianfuType()
	xianfuType := xianfutypes.XianfuType(typ)
	//验证参数
	if !xianfuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府挑战请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.XianfuArgumentInvalid)
		return
	}

	err = xianfuChallenge(tpl, xianfuType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   tpl.GetId(),
				"xianfuType": xianfuType,
				"err":        err,
			}).Error("xianfu:处理秘境仙府挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   tpl.GetId(),
			"xianfuType": xianfuType,
		}).Debug("xianfu：处理秘境仙府挑战请求完成")

	return
}

//仙府挑战逻辑
func xianfuChallenge(pl player.Player, xianfuType xianfutypes.XianfuType) (err error) {

	return xianfulogic.HandleXianfuChallenge(pl, xianfuType)
}
