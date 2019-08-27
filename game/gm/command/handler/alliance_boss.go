package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAllianceBoss, command.CommandHandlerFunc(handleAllianceBoss))
}

func handleAllianceBoss(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:仙盟boss")

	err = allianceBoss(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:仙盟boss,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:仙盟boss,完成")
	return
}

func allianceBoss(p scene.Player) (err error) {
	pl := p.(player.Player)
	allianceId := pl.GetAllianceId()
	if allianceId == 0 || pl.GetMengZhuId() != pl.GetId() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您不是盟主,无法召唤仙盟boss")
		playerlogic.SendSystemMessage(pl, lang.AllianceBossSummonNoMengZhu)
		return
	}

	err = alliance.GetAllianceService().AllianceSummonBoss(pl)
	if err != nil && err == alliance.ErrorAllianceBossSummonedBoss {
		flag := alliance.GetAllianceService().GMAllianceBossReset(allianceId)
		if !flag {
			return
		}
	}

	alliance.GetAllianceService().AllianceSummonBoss(pl)
	return
}
