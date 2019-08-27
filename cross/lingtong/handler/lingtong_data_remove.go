package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_REMOVE_TYPE), dispatch.HandlerFunc(handleLingTongDataRemove))
}

//战斗数据变化
func handleLingTongDataRemove(s session.Session, msg interface{}) error {
	log.Info("lingtong:处理跨服灵童数据移除")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)

	err := lingTongDataRemove(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
				"err":      err,
			}).Error("lingtong:处理跨服灵童数据移除,失败")
		return err
	}

	log.Info("lingtong:处理跨服灵童数据移除,完成")
	return nil
}

//灵童数据初始化
func lingTongDataRemove(pl *player.Player) (err error) {

	if pl.GetLingTong() == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingtong:处理跨服灵童数据,灵童不存在")
		return
	}

	pl.UpdateLingTong(nil)
	return nil
}
