package handler

import (
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/transportation/pbutil"
	playertransportation "fgame/fgame/game/transportation/player"
	"fgame/fgame/game/transportation/transpotation"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeTransportReset, command.CommandHandlerFunc(handleTransportReset))
}

func handleTransportReset(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:重置镖车")
	err = transportReset(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:重置镖车,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:重置镖车,完成")
	return
}

func transportReset(p scene.Player) (err error) {
	pl := p.(player.Player)
	transManager := pl.GetPlayerDataManager(types.PlayerTransportationType).(*playertransportation.PlayerTransportationDataManager)
	transManager.GMResetTimes()
	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
	if al != nil {
		al.GMResetTimes()
	}

	transpotation.GetTransportService().RemoveTransportation(pl.GetId())

	personalTimes := transManager.GetTranspotTimes()
	allianceTimes := alliance.GetAllianceService().GetAllianceTransportTimes(pl.GetId())
	scPlayerTransportationBriefInfo := pbutil.BuildSCPlayerTransportationBriefInfo(personalTimes, allianceTimes)
	err = pl.SendMsg(scPlayerTransportationBriefInfo)
	return
}
