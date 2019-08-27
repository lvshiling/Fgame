package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_ADVANCED_TYPE), dispatch.HandlerFunc(handleLingTongDevAdvanced))
}

//处理灵童养成类进阶信息
func handleLingTongDevAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成类进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevAdvanced := msg.(*uipb.CSLingTongDevAdvanced)
	classType := csLingTongDevAdvanced.GetClassType()
	autoFlag := csLingTongDevAdvanced.GetAutoFlag()

	err = lingtongdevlogic.HandleLingTongDevAdvanced(tpl, types.LingTongDevSysType(classType), autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"autoFlag":  autoFlag,
				"error":     err,
			}).Error("lingtongdev:处理灵童养成类进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingtongdev:处理灵童养成类进阶完成")
	return nil

}
