package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAllianceJianShe, command.CommandHandlerFunc(handleAllianceLevel))

}

func handleAllianceLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   numStr,
				"error": err,
			}).Warn("gm:处理设置仙盟建设度,num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   numStr,
				"error": err,
			}).Warn("gm:处理设置仙盟建设度,num小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	member := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	al := member.GetAlliance()
	curJianShe := al.GetAllianceObject().GetJianShe()
	dif := num - curJianShe
	if dif < 0 {
		al.CostJianShe(-dif)
	} else {
		al.AddJianShe(dif)
	}

	propertylogic.SnapChangedProperty(pl)
	return
}
