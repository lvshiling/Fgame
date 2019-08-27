package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_GET_DROP_ITEM_TYPE), dispatch.HandlerFunc(handlePlayerGetDropItem))
}

//TODO 是否需要回复后才删除物品
//处理获取掉落回复
func handlePlayerGetDropItem(s session.Session, msg interface{}) error {
	log.Debugln("inventory:处理跨服获取掉落回复")

	log.WithFields(
		log.Fields{}).Debug("scene:处理跨服获取掉落回复,完成")
	return nil
}
