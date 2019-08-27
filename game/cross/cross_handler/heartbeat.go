package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/processor"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_HEARTBEAT_TYPE), dispatch.HandlerFunc(handleHeartBeat))
}

//处理跨服登录
func handleHeartBeat(s session.Session, msg interface{}) error {
	log.Debug("cross:处理跨服心跳")

	log.Debug("cross:处理跨服心跳完成")
	return nil
}
