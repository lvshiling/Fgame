package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/equipbaoku/equipbaoku"
	"fgame/fgame/game/equipbaoku/pbutil"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_LOG_INCR_TYPE), dispatch.HandlerFunc(handleEquipBaoKuLogIncr))
}

//处理装备宝库日志请求
func handleEquipBaoKuLogIncr(s session.Session, msg interface{}) (err error) {
	log.Debug("equipbaoku:处理获取宝库日志请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csEquipBaoKuLogIncr := msg.(*uipb.CSEquipbaokuLogIncr)
	logTime := csEquipBaoKuLogIncr.GetLogTime()
	typ := equipbaokutypes.BaoKuType(csEquipBaoKuLogIncr.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Warn("equipbaoku:处理获取宝库日志请求,宝库类型不合法")
		return
	}

	err = equipBaoKuLogIncr(tpl, logTime, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
				"error":    err,
			}).Error("equipbaoku:处理获取宝库日志请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Debug("equipbaoku:处理获取宝库日志请求,完成")
	return nil

}

//获取装备宝库界面信息逻辑
func equipBaoKuLogIncr(pl player.Player, logTime int64, typ equipbaokutypes.BaoKuType) (err error) {
	logList := equipbaoku.GetEquipBaoKuService().GetLogByTime(logTime, typ)
	if len(logList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
			}).Warn("equipbaoku:处理获取装备宝库日志请求,日志增量列表为空")
		return
	}

	scEquipBaoKuLogIncr := pbutil.BuildSCEquipBaoKuLogIncr(logList, int32(typ))
	pl.SendMsg(scEquipBaoKuLogIncr)
	return
}
