package handler

import (
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	// processor.Register(codec.MessageType(uipb.MessageType_CS_WORLD_BOSS_LIST_TYPE), dispatch.HandlerFunc(handlerWorldBossList))
}

//世界BOSS列表请求
func handlerWorldBossList(s session.Session, msg interface{}) (err error) {
	log.Debug("worldBoss:处理世界BOSS列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = worldBossList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("worldBoss:处理世界BOSS列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("worldBoss:处理世界BOSS列表请求完成")

	return
}

func worldBossList(pl player.Player) (err error) {
	// //功能开启判断
	// flag := pl.IsFuncOpen(funcopentypes.FuncOpenTypeWorldBoss)
	// if !flag {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 		}).Warn("transport:世界BOSS列表错误，功能未开启")

	// 	playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
	// 	return
	// }

	// bossList := worldboss.GetWorldBossService().GetWorldBossList()
	// scWorldBossList := pbutil.BuildSCWorldBossList(bossList)
	// pl.SendMsg(scWorldBossList)

	return
}
