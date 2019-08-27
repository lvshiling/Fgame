package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/charge/charge"
	chargelogic "fgame/fgame/game/charge/logic"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/timeutils"
	"time"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeOpenTime, command.CommandHandlerFunc(handleSetOpenTime))
}

func handleSetOpenTime(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置合服时间")
	if len(c.Args) < 2 {
		log.Warn("gm:设置合服时间,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	timeStr := c.Args[0] + " " + c.Args[1]
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:设置合服时间,错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	time := timeutils.TimeToMillisecond(t)
	center.GetCenterService().SetStartTime(time)
	flag := charge.GetChargeService().ResetFirstCharge()
	if !flag {
		return
	}
	chargelogic.BroadcastFirstCharge()
	return
}
