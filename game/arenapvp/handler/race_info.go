package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/pbutil"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENAPVP_RACE_INFO_TYPE), dispatch.HandlerFunc(handleArenapvpRace))
}

//处理跨服1v1赛程
func handleArenapvpRace(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理跨服1v1赛程")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSArenapvpRaceInfo)
	typeInt := csMsg.GetType()

	pvpType := arenapvptypes.ArenapvpType(typeInt)

	err = arenapvpRace(tpl, pvpType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理跨服1v1赛程,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理跨服1v1赛程,完成")
	return nil

}

func arenapvpRace(pl player.Player, pvpType arenapvptypes.ArenapvpType) (err error) {
	raceInfoList := arenapvp.GetArenapvpService().GetArenapvpPlayerDataList()
	scMsg := pbutil.BuildSCArenapvpRaceInfo(pvpType, raceInfoList)
	pl.SendMsg(scMsg)
	return
}
