package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/onearena/pbutil"
	playeronearena "fgame/fgame/game/onearena/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_RECORED_TYPE), dispatch.HandlerFunc(handleOneArenaRecord))
}

//处理灵池被抢记录信息
func handleOneArenaRecord(s session.Session, msg interface{}) (err error) {
	log.Debug("onearena:处理获取灵池被抢记录消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csOneArenaRecord := msg.(*uipb.CSOneArenaRecord)
	logTime := csOneArenaRecord.GetLogTime()
	err = oneArenaRecord(tpl, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
				"error":    err,
			}).Error("onearena:处理获取灵池被抢记录消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("onearena:处理获取灵池被抢记录消息完成")
	return nil
}

//处理灵池被抢记录界面信息逻辑
func oneArenaRecord(pl player.Player, logTime int64) (err error) {
	if logTime < 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Warn("onearena:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	recordList := manager.GetRecordByLogTime(logTime)
	scOneArenaRecord := pbutil.BuilSCOneArenaRecord(recordList)
	pl.SendMsg(scOneArenaRecord)
	return
}
