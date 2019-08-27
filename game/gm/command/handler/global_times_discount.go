package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	"fgame/fgame/game/welfare/welfare"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeDiscountGlobalTimes, command.CommandHandlerFunc(handleDiscountGlobalTimes))
}

func handleDiscountGlobalTimes(p scene.Player, c *command.Command) (err error) {
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
	welfare.GetWelfareService().GMSetGlobalDiscountTimes(int32(groupId), int32(key), int32(tiems))

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timesList := welfare.GetWelfareService().GetDiscountTimes(groupId)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	discountObj := welfareManager.GetOpenActivity(groupId)
	dicountDay := welfarelogic.CountCurActivityDay(groupId)
	scMsg := pbutil.BuildSCMergeActivityDiscountInfo(discountObj, groupId, dicountDay, timesList, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
