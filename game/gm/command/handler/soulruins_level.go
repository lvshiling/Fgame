package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeSoulRuinsLevel, command.CommandHandlerFunc(handleSoulRuinsLevel))
}

func handleSoulRuinsLevel(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置帝陵遗迹关卡等级")
	if len(c.Args) < 3 {
		log.Warn("gm:处理设置帝陵遗迹关卡等级,参数少于3")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	chapterStr := c.Args[0]
	typStr := c.Args[1]
	levelStr := c.Args[2]
	chapter, err := strconv.ParseInt(chapterStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"chapter": chapter,
			}).Warn("gm:处理设置帝陵遗迹关卡等级,chapter不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	typ, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"typ":   typ,
			}).Warn("gm:处理设置帝陵遗迹关卡等级,typ不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"level": level,
			}).Warn("gm:处理设置帝陵遗迹关卡等级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	if level < 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	//TODO 修改帝陵遗迹关卡等级
	err = setSoulRuinsLevel(pl, int32(chapter), soulruinstypes.SoulRuinsType(typ), int32(level))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"chapter": chapter,
				"typ":     chapter,
				"level":   level,
			}).Warn("gm:设置帝陵遗迹关卡等级,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"chapter": chapter,
			"typ":     chapter,
			"level":   level,
		}).Debug("gm:设置帝陵遗迹关卡等级,完成")
	return
}

func setSoulRuinsLevel(p scene.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	flag := manager.IsValid(chapter, typ, level)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"chapter":  chapter,
			"typ":      typ,
			"level":    level,
		}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}

	manager.GMSetLevel(chapter, typ, level)

	scSoulRuinsGet := pbutil.BuildSCSoulRuinsGet(pl)
	pl.SendMsg(scSoulRuinsGet)
	return
}
