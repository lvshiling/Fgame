package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLYAER_MOUNT_SYNC_TYPE), dispatch.HandlerFunc(handlePlayerMountSync))
}

//处理坐骑同步
func handlePlayerMountSync(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑同步")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isPlayerMountSync := msg.(*crosspb.ISPlayerMountSync)
	hiddenFlag := isPlayerMountSync.GetHidden()

	err = mountSync(tpl, hiddenFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"hiddenFlag": hiddenFlag,
				"error":      err,
			}).Error("mount:处理坐骑同步,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"hiddenFlag": hiddenFlag,
		}).Debug("mount:处理坐骑同步")
	return nil

}

//坐骑隐藏展示的逻辑
func mountSync(pl player.Player, hiddenFlag bool) (err error) {
	pl.MountSync(hiddenFlag)
	return
}
