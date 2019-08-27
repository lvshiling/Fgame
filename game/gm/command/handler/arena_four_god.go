package handler

import (
	"fgame/fgame/common/lang"
	arenalogic "fgame/fgame/game/arena/logic"
	arenatypes "fgame/fgame/game/arena/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeArenaFourGod, command.CommandHandlerFunc(handleArenaFourGod))
}

func handleArenaFourGod(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	log.Debug("gm:3v3竞技场选择四神")
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	fourGodTypeStr := c.Args[0]
	fourGodTypeInt, err := strconv.ParseInt(fourGodTypeStr, 10, 64)
	if err != nil {
		return
	}
	fourGodType := arenatypes.FourGodType(fourGodTypeInt)
	if !fourGodType.Valid() {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:3v3竞技场选择四神,参数无效")
		return
	}
	err = arenaSelectFourGod(pl, fourGodType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:3v3竞技场选择四神,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:3v3竞技场选择四神,完成")
	return
}

func arenaSelectFourGod(pl player.Player, fourGodType arenatypes.FourGodType) (err error) {
	//模仿跨服

	err = arenalogic.GMSelectFourGod(pl, fourGodType)
	if err != nil {
		return
	}
	return

}
