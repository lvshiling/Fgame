package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	"fgame/fgame/game/player"
	"fgame/fgame/game/player/pbutil"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PLAYER_OPEN_VIDEO_TYPE), dispatch.HandlerFunc(handlePlayerOpenVideo))
}

//处理玩家开场动画确认
func handlePlayerOpenVideo(s session.Session, msg interface{}) (err error) {
	log.Debug("player:处理玩家开场动画确认")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(player.Player)

	err = playerOpenVideo(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("player:处理玩家开场动画确认,创建失败")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("player:处理玩家开场动画确认,创建成功")
	return
}

func playerOpenVideo(pl player.Player) (err error) {
	isOpenVideo := pl.IsOpenVideo()
	if isOpenVideo {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:处理玩家开场动画确认,已经播放过")
		return
	}

	pl.OpenVideo()

	scMsg := pbutil.BuildscPlayerOpenVedio()
	pl.SendMsg(scMsg)
	return
}
