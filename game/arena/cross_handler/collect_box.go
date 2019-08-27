package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arena/pbutil"
	collectlogic "fgame/fgame/game/collect/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_COLLECT_BOX_TYPE), dispatch.HandlerFunc(handleArenaCollectBox))
}

//处理获得宝箱
func handleArenaCollectBox(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理采集宝箱")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isArenaCollectBox := msg.(*crosspb.ISArenaCollectBox)
	err = arenaCollectBox(tpl, isArenaCollectBox.GetBoxId())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理采集宝箱,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理采集宝箱,完成")
	return nil

}

//采集宝箱
func arenaCollectBox(pl player.Player, boxId int32) (err error) {
	collectlogic.CollectDropToInventory(pl, boxId)
	siArenaCollectExpTree := pbutil.BuildSIArenaCollectExpTree()
	pl.SendMsg(siArenaCollectExpTree)
	return
}
