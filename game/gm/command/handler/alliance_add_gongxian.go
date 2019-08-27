package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/alliance/pbutil"
	playerAlliance "fgame/fgame/game/alliance/player"
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

	command.Register(gmcommandtypes.CommandTypeAllianceAddGongXian, command.CommandHandlerFunc(handleAllianceAddGongxian))
}

func handleAllianceAddGongxian(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置贡献值")

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	gongxianStr := c.Args[0]
	gongxian, err := strconv.ParseInt(gongxianStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"gongxian": gongxianStr,
				"error":    err,
			}).Warn("gm:处理设置元宝数量,gongxian不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	err = allianceAddGongxian(pl, gongxian)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:仙盟捐献,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:仙盟捐献,完成")
	return
}

func allianceAddGongxian(p scene.Player, gongxian int64) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playerAlliance.PlayerAllianceDataManager)

	curGongxian := manager.GetPlayerAllianceObject().GetCurrentGongXian()
	needGongxian := gongxian - curGongxian
	if needGongxian == 0 {
		return
	}
	if needGongxian < 0 {
		flag := manager.UseGongXian(-needGongxian)
		if !flag {
			panic("gm:使用贡献应该成功")
		}
	} else {
		manager.AddGongXian(needGongxian)
	}

	//发送仙盟个人信息
	scAlliancePlayerInfo := pbutil.BuildSCAlliancePlayerInfo(manager.GetPlayerAllianceObject(), manager.GetPlayerAllianceSkillMap())
	pl.SendMsg(scAlliancePlayerInfo)

	return

}
