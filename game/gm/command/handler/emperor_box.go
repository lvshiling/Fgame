package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/emperor/pbutil"
	playeremperor "fgame/fgame/game/emperor/player"
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
	command.Register(gmcommandtypes.CommandTypeEmperorBox, command.CommandHandlerFunc(handleEmperorBox))

}

func handleEmperorBox(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	boxNumStr := c.Args[0]
	boxNum, err := strconv.ParseInt(boxNumStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"boxNum": boxNum,
				"error":  err,
			}).Warn("gm:boxNum不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = emperorBox(pl, boxNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理抢龙椅宝箱数量,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理抢龙椅宝箱数量完成")
	return
}

func emperorBox(pl player.Player, boxNum int64) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerEmperorDataManagerType).(*playeremperor.PlayerEmperorDataManager)
	emperor.GetEmperorService().GmSetEmperorBox(boxNum)
	emperorObj := emperor.GetEmperorService().GetEmperorInfo()
	worshipNum := manager.GetWorshipNum()
	scEmperorGet := pbuitl.BuildSCEmperorGet(emperorObj, worshipNum)
	pl.SendMsg(scEmperorGet)
	return
}
