package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	wardrobelogic "fgame/fgame/game/wardrobe/logic"
	"fgame/fgame/game/wardrobe/pbutil"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetemplate "fgame/fgame/game/wardrobe/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeYiChuLevel, command.CommandHandlerFunc(handleYiChuLevel))
}

func handleYiChuLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typeStr := c.Args[0]
	levelStr := c.Args[1]
	typeInt64, err := strconv.ParseInt(typeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"typeStr":  typeStr,
				"levelStr": levelStr,
				"error":    err,
			}).Warn("gm:处理设置衣橱培养等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	levelInt64, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"typeStr":  typeStr,
				"levelStr": levelStr,
				"error":    err,
			}).Warn("gm:处理设置衣橱培养等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	typ := int32(typeInt64)
	// if !typ.Valid() {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"id":       pl.GetId(),
	// 			"typeStr":  typeStr,
	// 			"levelStr": levelStr,
	// 			"error":    err,
	// 		}).Warn("gm:处理设置衣橱培养等级,wingLevel不是数字")
	// 	playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
	// 	err = nil
	// 	return
	// }
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	_, flag := manager.IfCanPeiYang(typ)
	if !flag {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"typeStr":  typeStr,
				"levelStr": levelStr,
				"error":    err,
			}).Warn("gm:处理设置衣橱培养等级,无法培养")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	suitTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSuitTemplate(typ)
	peiYangTemplate := suitTemplate.GetPeiYangByLevel(int32(levelInt64))
	if peiYangTemplate == nil {
		return
	}

	manager.GmSetPeiYangLevel(typ, int32(levelInt64))
	//同步属性
	wardrobelogic.WardrobePropertyChanged(pl)

	scWardrobePeiYang := pbutil.BuildSCWardrobePeiYang(int32(typ), int32(levelInt64))
	pl.SendMsg(scWardrobePeiYang)
	return
}
