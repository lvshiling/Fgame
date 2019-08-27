package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	"fgame/fgame/game/welfare/welfare"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeOpenActivityXunHuan, command.CommandHandlerFunc(handleXunHuan))
}

func handleXunHuan(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置循环活动组")
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	arrStr := c.Args[0]
	arrGroup, err := strconv.ParseInt(arrStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"arrGroup": arrGroup,
				"error":    err,
			}).Warn("gm:处理设置循环活动组,arrGroup不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if arrGroup <= 0 {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"arrGroup": arrGroup,
				"error":    err,
			}).Warn("gm:处理设置循环活动组,arrGroup小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = setXunHuanGroup(pl, int32(arrGroup))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"arrGroup": arrGroup,
				"error":    err,
			}).Warn("gm:处理设置循环活动组,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":       pl.GetId(),
			"arrGroup": arrGroup,
		}).Debug("gm:处理设置循环活动组,完成")
	return
}

func setXunHuanGroup(pl player.Player, arrGroup int32) (err error) {
	welfare.GetWelfareService().GMSetXunHuanGroupArr(arrGroup)
	groupIdList := welfarelogic.GetXunHuanGroupList()
	alPlList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range alPlList {
		scMsg := pbutil.BuildSCOpenActivityXunHuanInfo(groupIdList)
		pl.SendMsg(scMsg)
	}
	return
}
