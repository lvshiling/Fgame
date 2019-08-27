package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
	viplogic "fgame/fgame/game/vip/logic"
	playervip "fgame/fgame/game/vip/player"
	viptemplate "fgame/fgame/game/vip/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeVipLevel, command.CommandHandlerFunc(handleVipLevel))
}

func handleVipLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置VIP等级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if level < 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置VIP等级,level小于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	starStr := c.Args[1]
	satr, err := strconv.ParseInt(starStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"satr":  starStr,
				"error": err,
			}).Warn("gm:处理设置VIP等级,satr不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	if satr < 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"satr":  satr,
				"error": err,
			}).Warn("gm:处理设置VIP等级,star小于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	temp := viptemplate.GetVipTemplateService().GetVipTemplate(int32(level), int32(satr))
	//修改VIP等级
	if temp == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置VIP等级,level模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	vipManager := pl.GetPlayerDataManager(types.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	curLevel, curStar := vipManager.GetVipLevel()
	curTemp := viptemplate.GetVipTemplateService().GetVipTemplate(curLevel, curStar)
	if curTemp == nil {
		return
	}

	tempLevelTemplate := viptemplate.GetVipTemplateService().GetVipTemplateById(int32(temp.Id))
	needNum := int64(tempLevelTemplate.NeedValue)
	vipManager.GMSetChargeNum(needNum)

	viplogic.VipInfoNotice(pl)
	propertylogic.SnapChangedProperty(pl)
	return
}
