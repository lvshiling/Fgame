package handler

import (
	"context"
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/common/lang"
	settimeclient "fgame/fgame/cross/settime/client"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/timeutils"
	"time"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeRegionTime, command.CommandHandlerFunc(handleRegionSetTime))
}

func handleRegionSetTime(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置组服服务器时间")
	if len(c.Args) < 2 {
		log.Warn("gm:设置组服 服务器时间,参数少于1")
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
			}).Warn("gm:设置服务器时间,错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	now := time.Now()
	offTime := int64(t.Sub(now)) / int64(time.Millisecond)

	global.GetGame().GetTimeService().SetOffTime(offTime)
	nowInt := global.GetGame().GetTimeService().Now()
	scGetTime := pbutil.BuildSCGetTime(nowInt)

	player.GetOnlinePlayerManager().BroadcastMsg(scGetTime)
	log.WithFields(
		log.Fields{
			"id":   pl.GetId(),
			"time": timeStr,
			"now":  timeutils.MillisecondToTime(global.GetGame().GetTimeService().Now()).String(),
		}).Debug("gm:设置服务器时间")

	//同步所有跨服
	conn := center.GetCenterService().GetCross(centertypes.GameServerTypeRegion)
	if conn == nil {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:设置服务器时间,跨服连接不存在")
		playerlogic.SendSystemMessage(pl, lang.CrossFailed)
		return
	}

	setTimeClient := settimeclient.NewSetTimeClient(conn)
	setTimeClient.SetTime(context.Background(), nowInt)

	return
}
