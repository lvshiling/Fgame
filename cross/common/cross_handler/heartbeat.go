package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_HEARTBEAT_TYPE), dispatch.HandlerFunc(handleHeartbeat))
}

//获取时间
func handleHeartbeat(s session.Session, msg interface{}) error {
	log.Debug("common:处理心跳")

	gcs := gamesession.SessionInContext(s.Context())
	gcs.Ping()
	log.Debug("common:处理心跳,完成")
	return nil
}
