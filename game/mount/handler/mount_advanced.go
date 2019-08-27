package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	mountlogic "fgame/fgame/game/mount/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_ADVANCED_TYPE), dispatch.HandlerFunc(handleMountAdvanced))
}

//处理坐骑进阶信息
func handleMountAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMountAdvanced := msg.(*uipb.CSMountAdvanced)
	autoFlag := csMountAdvanced.GetAutoFlag()

	err = mountAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mount:处理坐骑进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mount:处理坐骑进阶完成")
	return nil
}

//坐骑进阶的逻辑
func mountAdvanced(pl player.Player, autoFlag bool) (err error) {
	return mountlogic.HandleMountAdvanced(pl, autoFlag)
}
