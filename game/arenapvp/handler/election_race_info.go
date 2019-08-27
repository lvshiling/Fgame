package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENAPVP_ELECTION_RACE_INFO_TYPE), dispatch.HandlerFunc(handleArenapvpElectionRace))
}

//处理跨服海选赛程
func handleArenapvpElectionRace(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理跨服海选赛程")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenapvpElectionRace(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理跨服海选赛程,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理跨服海选赛程,完成")
	return nil
}

func arenapvpElectionRace(pl player.Player) (err error) {
	raceInfoList := arenapvp.GetArenapvpService().GetArenapvpElectionList()
	scMsg := pbutil.BuildSCArenapvpElectionRaceInfo(raceInfoList)
	pl.SendMsg(scMsg)
	return
}
