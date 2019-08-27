package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/arena/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_NEXT_MATCH_TYPE), dispatch.HandlerFunc(handleNext))
}

//下一场
func handleNext(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理选择下一场")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = arenaNext(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理选择下一场,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理选择四圣兽")
	return nil

}

//3v3匹配
func arenaNext(pl *player.Player) (err error) {
	//判断是否有组队
	flag := arena.GetArenaService().ArenaNext(pl.GetId())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:下一个匹配失败")
		return
	}

	scArenaNextMatch := pbutil.BuildSCArenaNextMatch()
	pl.SendMsg(scArenaNextMatch)
	return
}
