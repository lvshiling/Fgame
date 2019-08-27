package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	equipbaokulogic "fgame/fgame/game/equipbaoku/logic"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_ATTEND_TYPE), dispatch.HandlerFunc(handleEquipBaoKuAttend))

}

//探索宝库
func handleEquipBaoKuAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("equipbaoku:探索宝库")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csEquipBaoKuAttend := msg.(*uipb.CSEquipbaokuAttend)
	logTime := csEquipBaoKuAttend.GetLogTime()
	autoFlag := csEquipBaoKuAttend.GetAutoFlag()
	typ := equipbaokutypes.BaoKuType(csEquipBaoKuAttend.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"typ":      typ,
				"error":    err,
			}).Warn("equipbaoku:处理探索宝库,宝库类型不合法")
		return
	}

	err = equipBaoKuAttend(tpl, logTime, autoFlag, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("equipbaoku:处理探索宝库,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("equipbaoku:处理探索宝库完成")
	return nil

}

//探索装备宝库逻辑
func equipBaoKuAttend(pl player.Player, logTime int64, autoFlag bool, typ equipbaokutypes.BaoKuType) (err error) {
	return equipbaokulogic.EquipBaoKuAttend(pl, logTime, autoFlag, typ)
}
