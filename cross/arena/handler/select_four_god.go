package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/arena/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	arenatypes "fgame/fgame/game/arena/types"
	playerlogic "fgame/fgame/game/player/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_SELECT_FOUR_GOD_TYPE), dispatch.HandlerFunc(handleArenaSelectFourGod))
}

//处理选择四圣兽
func handleArenaSelectFourGod(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理选择四圣兽")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	csArenaSelectFourGod := msg.(*uipb.CSArenaSelectFourGod)
	fourGodTypeInt := csArenaSelectFourGod.GetFourGodType()
	fourGodType := arenatypes.FourGodType(fourGodTypeInt)
	if !fourGodType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理选择四圣兽,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = arenaSelectFourGod(tpl, fourGodType)
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

//3v3匹配
func arenaSelectFourGod(pl *player.Player, fourGodType arenatypes.FourGodType) (err error) {
	arena.GetArenaService().PlayerEnterFourGod(pl.GetId(), fourGodType)
	scArenaSelectFourGod := pbutil.BuildSCArenaSelectFourGod(fourGodType)
	pl.SendMsg(scArenaSelectFourGod)
	return
}
