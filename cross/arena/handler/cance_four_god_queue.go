package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/player/player"

	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_CANCEL_QUEUE_TYPE), dispatch.HandlerFunc(handleArenaCancelFourGodQueue))
}

//处理取消四圣兽
func handleArenaCancelFourGodQueue(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理选择四圣兽取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = arenaCancelFourGodQueue(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理选择四圣兽,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理选择四圣兽")
	return nil

}

//四神取消
func arenaCancelFourGodQueue(pl *player.Player) (err error) {
	//判断是否有组队

	arena.GetArenaService().PlayerCancelFourGodQueue(pl.GetId())

	return
}
