package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tower/pbutil"
	"fgame/fgame/game/tower/tower"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TOWER_LOG_INCR_TYPE), dispatch.HandlerFunc(handlerTowerLogList))
}

//打宝塔日志列表请求
func handlerTowerLogList(s session.Session, msg interface{}) (err error) {
	log.Debug("towerBoss:处理打宝塔日志列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTowerLogIncr)
	lastLogTime := csMsg.GetLogTime()

	err = towerLogList(tpl, lastLogTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("towerBoss:处理打宝塔日志列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("towerBoss:处理打宝塔日志列表请求完成")

	return
}

func towerLogList(pl player.Player, lastLogTime int64) (err error) {

	s := pl.GetScene()
	if s == nil {
		return
	}
	if !s.MapTemplate().IsTower() {
		return
	}

	sd := s.SceneDelegate().(tower.TowerSceneData)
	logList := sd.GetLogByTime(lastLogTime)
	if len(logList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"lastLogTime": lastLogTime,
			}).Warn("chess:处理获取打宝塔日志请求,日志增量列表为空")
		return
	}

	scMsg := pbutil.BuildSCTowerLogIncr(logList)
	pl.SendMsg(scMsg)
	return
}
