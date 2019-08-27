package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/alliance/pbutil"
	playerAlliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeDepotPoint, command.CommandHandlerFunc(handleAllianceDepotPoint))
}

func handleAllianceDepotPoint(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置贡献值")

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	pointStr := c.Args[0]
	point, err := strconv.ParseInt(pointStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"point": pointStr,
				"error": err,
			}).Warn("gm:处理设置仙盟仓库积分,point不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = allianceDepotPoint(pl, point)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:仙盟仓库积分,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:仙盟仓库积分,完成")
	return
}

func allianceDepotPoint(p scene.Player, point int64) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playerAlliance.PlayerAllianceDataManager)

	curPoint := manager.GetDepotPoint()
	needPoint := int32(point) - curPoint
	if needPoint == 0 {
		return
	}
	if needPoint < 0 {
		manager.CostDepotPoint(-needPoint)
	} else {
		manager.AddDepotPoint(needPoint)
	}

	//发送仙盟个人信息
	scAlliancePlayerInfo := pbutil.BuildSCAlliancePlayerInfo(manager.GetPlayerAllianceObject(), manager.GetPlayerAllianceSkillMap())
	pl.SendMsg(scAlliancePlayerInfo)

	return

}
