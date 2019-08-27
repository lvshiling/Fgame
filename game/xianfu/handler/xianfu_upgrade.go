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
	xianfulogic "fgame/fgame/game/xianfu/logic"
	xianfutypes "fgame/fgame/game/xianfu/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANFU_UPGRADE_TYPE), dispatch.HandlerFunc(handlerXianfuUpgrade))
}

//秘境仙府升级请求
func handlerXianfuUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("xianfu:处理秘境仙府升级请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csXianfuUpgrade := msg.(*uipb.CSXianfuUpgrade)
	typ := csXianfuUpgrade.GetXianfuType()

	//验证参数
	xianfuType := xianfutypes.XianfuType(typ)
	if !xianfuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府升级请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = xianfuUpgrade(tpl, xianfuType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   tpl.GetId(),
				"xianfuType": xianfuType,
				"err":        err,
			}).Error("xianfu:处理秘境仙府升级请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   tpl.GetId(),
			"xianfuType": xianfuType,
		}).Debug("xianfu：处理秘境仙府升级请求完成")

	return
}

//升级请求逻辑
func xianfuUpgrade(pl player.Player, xianfuType xianfutypes.XianfuType) (err error) {
	return xianfulogic.HandleXianfuUpgrade(pl, xianfuType)
}
