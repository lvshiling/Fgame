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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_LIST_TYPE), dispatch.HandlerFunc(handleArenaFourGodList))
}

//处理获取四圣兽信息
func handleArenaFourGodList(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理获取四圣兽信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = arenaFourGodList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理获取四圣兽信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理选择四圣兽")
	return nil

}

//四圣兽信息
func arenaFourGodList(pl *player.Player) (err error) {

	fourGodSceneList := arena.GetArenaService().GetFourGodSceneList()
	scArenaFourGodList := pbutil.BuildSCArenaFourGodList(fourGodSceneList)
	pl.SendMsg(scArenaFourGodList)
	return
}
