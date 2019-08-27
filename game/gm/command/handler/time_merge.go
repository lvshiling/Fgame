package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/charge/charge"
	chargelogic "fgame/fgame/game/charge/logic"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/merge/merge"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/timeutils"
	"time"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMergeTime, command.CommandHandlerFunc(handleSetMergeTime))
}

func handleSetMergeTime(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置合服时间")

	timeInt := int64(0)
	if len(c.Args) == 2 {
		timeStr := c.Args[0] + " " + c.Args[1]
		t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
		if err != nil {
			log.WithFields(
				log.Fields{
					"id":    pl.GetId(),
					"error": err,
				}).Warn("gm:设置合服时间,错误")
			playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
			return nil
		}
		timeInt = timeutils.TimeToMillisecond(t)
	} else if len(c.Args) != 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:设置合服时间,错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return nil
	}

	merge.GetMergeService().GMSetMergeTime(timeInt)
	flag := charge.GetChargeService().ResetFirstCharge()
	if !flag {
		return
	}
	chargelogic.BroadcastFirstCharge()
	return
}
