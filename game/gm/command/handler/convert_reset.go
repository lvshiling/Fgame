package handler

import (
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeYaoPaiConvertReset, command.CommandHandlerFunc(handleConvertReset))
}

func handleConvertReset(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理兑换次数重置")

	err = reset(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理兑换次数重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理兑换次数重置完成")
	return
}

func reset(p scene.Player) (err error) {
	pl := p.(player.Player)
	allianceManager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.GMResetConvertTimes()

	//发送仙盟个人信息
	scAlliancePlayerInfo := pbutil.BuildSCAlliancePlayerInfo(allianceManager.GetPlayerAllianceObject(), allianceManager.GetPlayerAllianceSkillMap())
	pl.SendMsg(scAlliancePlayerInfo)

	return
}
