package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeGongDe, command.CommandHandlerFunc(handlePlayerGongDe))
}

func handlePlayerGongDe(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置贡献值")

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	gongDeStr := c.Args[0]
	gongDe, err := strconv.ParseInt(gongDeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"gongDe": gongDeStr,
				"error":  err,
			}).Warn("gm:处理设置功德,gongxian不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = addGongDe(pl, gongDe)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:设置功德,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:设置功德,完成")
	return
}

func addGongDe(p scene.Player, gongDe int64) (err error) {
	if gongDe < 0 {
		return
	}
	pl := p.(player.Player)
	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiManager.GMSetGongDe(gongDe)

	scMsg := pbutil.BuildSCFeiShengSanGong(gongDe)
	pl.SendMsg(scMsg)
	return

}
