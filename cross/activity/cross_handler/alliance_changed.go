package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerActivityPkDataChanged))
}

//玩家仙盟变化
func handlePlayerActivityPkDataChanged(s session.Session, msg interface{}) (err error) {

	return nil
}
