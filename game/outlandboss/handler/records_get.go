package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/outlandboss/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_DROP_RECORDS_GET_TYPE), dispatch.HandlerFunc(handleDropRecordGet))
}

//处理获取记录信息
func handleDropRecordGet(s session.Session, msg interface{}) (err error) {
	log.Debug("outlandboss:处理抢夺记录信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = dropRecordsGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("outlandboss:处理掉落记录信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("outlandboss:处理掉落记录信息完成")
	return nil
}

//处理抢夺记录界面信息逻辑
func dropRecordsGet(pl player.Player) (err error) {
	dropRecordsList := outlandboss.GetOutlandBossService().GetDropRecordsList()
	scDropRecordsGet := pbutil.BuildSCOutlandBossDropRecordsGet(dropRecordsList)
	pl.SendMsg(scDropRecordsGet)
	return
}
