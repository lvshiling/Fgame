package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/onearena/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_GET_TYPE), dispatch.HandlerFunc(handleOneArenaGet))
}

//处理灵池争夺界面信息
func handleOneArenaGet(s session.Session, msg interface{}) (err error) {
	log.Debug("onearena:处理获取灵池争夺界面消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = oneArenaGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("onearena:处理获取灵池争夺界面消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("onearena:处理获取灵池争夺界面消息完成")
	return nil
}

//处理灵池争夺界面界面信息逻辑
func oneArenaGet(pl player.Player) (err error) {
	oneArenaList := onearena.GetOneArenaService().GetOneArenaList()
	scOneArenaGet := pbutil.BuildSCOneArenaGet(pl, oneArenaList)
	pl.SendMsg(scOneArenaGet)
	return
}
