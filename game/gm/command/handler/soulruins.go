package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	soulruinislogic "fgame/fgame/game/soulruins/logic"
	playersoulruins "fgame/fgame/game/soulruins/player"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSoulRuins, command.CommandHandlerFunc(handleSoulRuins))
}

func handleSoulRuins(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理进入帝陵遗迹")
	pl := p.(player.Player)
	if len(c.Args) < 3 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	chapterStr := c.Args[0]
	chapter, err := strconv.ParseInt(chapterStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"chapter": chapter,
				"error":   err,
			}).Warn("gm:处理跳转帝陵遗迹,chapter不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	typStr := c.Args[1]
	typ, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"typ":   typ,
				"error": err,
			}).Warn("gm:处理跳转帝陵遗迹,typ不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	levelStr := c.Args[2]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": level,
				"error": err,
			}).Warn("gm:处理跳转帝陵遗迹,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	flag := manager.IsValid(int32(chapter), soulruinstypes.SoulRuinsType(typ), int32(level))
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
		}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}

	if level > 1 {
		manager.GMSetLevel(int32(chapter), soulruinstypes.SoulRuinsType(typ), int32(level)-1)
	}

	err = enterSoulRuins(pl, int32(chapter), soulruinstypes.SoulRuinsType(typ), int32(level))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": level,
				"error": err,
			}).Warn("gm:处理跳转帝陵遗迹,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":    pl.GetId(),
			"level": level,
		}).Debug("gm:处理跳转帝陵遗迹,完成")
	return
}

func enterSoulRuins(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32) (err error) {
	soulruinislogic.PlayerEnterSoulRuins(pl, chapter, typ, level)
	return
}
