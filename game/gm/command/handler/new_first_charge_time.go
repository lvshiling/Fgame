package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/charge/charge"
	chargelogic "fgame/fgame/game/charge/logic"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeNewFirstChargeTime, command.CommandHandlerFunc(handleNewFirstChargeTime))

}

// const (
// 	duration = 7
// )

func handleNewFirstChargeTime(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置新首充活动时间")

	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	// startServerTime := welfare.GetWelfareService().GetServerStartTime()
	// mergeServerTime := welfare.GetWelfareService().GetServerMergeTime()
	// dur := int64(duration) * int64(common.DAY)

	// startTimeStr := c.Args[0] + " " + c.Args[1]
	// startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, time.Local)
	// if err != nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"error":    err,
	// 		}).Warn("gm:设置新首充活动时间,错误")
	// 	playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
	// 	err = nil
	// 	return
	// }
	// startTimeMill := timeutils.TimeToMillisecond(startTime)
	// if startTimeMill > mergeServerTime && startTimeMill < mergeServerTime+dur {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"error":    err,
	// 		}).Warn("gm:设置新首充活动时间,错误")
	// 	playerlogic.SendSystemMessage(pl, lang.GMCanNotUse)
	// 	return
	// }
	// if startTimeMill > startServerTime && startTimeMill < startServerTime+dur {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"error":    err,
	// 		}).Warn("gm:设置新首充活动时间,错误")
	// 	playerlogic.SendSystemMessage(pl, lang.GMCanNotUse)
	// 	return
	// }
	now := global.GetGame().GetTimeService().Now()
	flag := charge.GetChargeService().GmSetNewFirstChargeTime(now)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Warn("gm:设置新首充活动时间,首冲时间还没完")
		playerlogic.SendSystemMessage(pl, lang.ChargeFirstChargeNoDone)
		return
	}

	chargelogic.BroadcastFirstCharge()

	return nil
}
