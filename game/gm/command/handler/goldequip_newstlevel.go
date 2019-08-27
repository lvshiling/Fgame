package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeGoldEquipNewStLevel, command.CommandHandlerFunc(handleGoldEquipNewStLevel))
}

func handleGoldEquipNewStLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	posStr := c.Args[0]
	levStr := c.Args[1]
	posInt, err := strconv.ParseInt(posStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"pos":   posInt,
				"error": err,
			}).Warn("gm:处理金装新强化等级任务,类型pos不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	lev, err := strconv.ParseInt(levStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"error": err,
			}).Warn("gm:处理金装新强化等级任务,类型lev不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//参数验证
	posType := inventorytypes.BodyPositionType(posInt)

	if !posType.Valid() {
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"posType":  posType.String(),
			}).Warn("gm:部位类型,错误")
		return
	}

	tempTemplatet := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenBuWeiTemplate(posType, int32(lev))
	if tempTemplatet == nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"posType": posType,
				"lev":     lev,
			}).Warn("gm:处理金装新强化等级任务,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//修改等级
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()
	flag := goldequipBag.GmSetStrengthBuWeiLevel(posType, int32(lev))
	if !flag {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"posType": posType,
				"lev":     lev,
			}).Warn("gm:处理金装新强化等级任务,错误")
		return
	}

	//同步改变
	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	propertylogic.SnapChangedProperty(pl)
	return
}
