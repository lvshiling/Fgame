package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	"fgame/fgame/game/welfare/welfare"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeReceiveGlobalTimes, command.CommandHandlerFunc(handleReceiveGlobalTimes))
}

func handleReceiveGlobalTimes(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 3 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	groupIdStr := c.Args[0]
	keyStr := c.Args[1]
	timesStr := c.Args[2]
	groupId64, err := strconv.ParseInt(groupIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"groupIdStr": groupIdStr,
				"error":      err,
			}).Warn("gm:处理全服次数修改,groupId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	key, err := strconv.ParseInt(keyStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"keyStr": keyStr,
				"error":  err,
			}).Warn("gm:处理全服次数修改,key不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tiems, err := strconv.ParseInt(timesStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"timesStr": timesStr,
				"error":    err,
			}).Warn("gm:处理全服次数修改,timesStr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	groupId := int32(groupId64)
	welfare.GetWelfareService().GMSetGlobalReceiveTimes(int32(groupId), int32(key), int32(tiems))
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	realmObj := welfareManager.GetOpenActivity(groupId)
	timesList := welfare.GetWelfareService().GetReceiveTimesList(groupId)
	scOpenActivityRealmInfo := pbutil.BuildSCOpenActivityRealmInfo(realmObj, groupId, timesList, startTime, endTime)
	pl.SendMsg(scOpenActivityRealmInfo)
	return
}
