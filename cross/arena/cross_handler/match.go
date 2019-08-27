package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENA_MATCH_TYPE), dispatch.HandlerFunc(handleArenaMatch))
}

//处理3v3匹配
func handleArenaMatch(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理3v3匹配")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siArenaMatch := msg.(*crosspb.SIArenaMatch)
	playerList := pbutil.ConvertFromArenaPlayerList(siArenaMatch.GetPlayerList())

	err = arenaMatch(tpl, playerList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理3v3匹配,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理3v3匹配,完成")
	return nil

}

//3v3匹配
func arenaMatch(pl *player.Player, playerList []*arena.MatchTeamMember) (err error) {
	flag := arena.GetArenaService().ArenaMatch(playerList)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:处理3v3匹配,匹配失败")
		isArenaMatchResult := pbutil.BuildISArenaMatchResult(false)
		pl.SendMsg(isArenaMatchResult)
		return
	}

	isArenaMatch := pbutil.BuildISArenaMatch()
	pl.SendMsg(isArenaMatch)
	return
}
