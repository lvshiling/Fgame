package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeZhuanSheng, command.CommandHandlerFunc(handleZhuanSheng))

}

func handleZhuanSheng(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
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
			}).Warn("gm:处理设置转生,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if level <= -1 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置转生,level小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if level == 0 {
		reasonZhuanSheng := commonlog.ZhuanShengLogReasonGoldEquip
		manager.SetZhuanSheng(0, reasonZhuanSheng, reasonZhuanSheng.String())
		scGoldEquipZhuanSheng := pbutil.BuildSCGoldEquipZhuanSheng(0)
		pl.SendMsg(scGoldEquipZhuanSheng)
		goldequiplogic.ZhuanShengPropertyChanged(pl)
		return
	}

	tempTemplateObject := goldequiptemplate.GetGoldEquipTemplateService().GetZhuanShengTemplate(int32(level))
	if tempTemplateObject == nil {
		return
	}

	reasonZhuanSheng := commonlog.ZhuanShengLogReasonGoldEquip
	manager.SetZhuanSheng(int32(level), reasonZhuanSheng, reasonZhuanSheng.String())
	scGoldEquipZhuanSheng := pbutil.BuildSCGoldEquipZhuanSheng(int32(level))
	pl.SendMsg(scGoldEquipZhuanSheng)
	goldequiplogic.ZhuanShengPropertyChanged(pl)
	return
}
