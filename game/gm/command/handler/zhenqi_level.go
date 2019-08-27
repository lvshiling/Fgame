package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	zhenfalogic "fgame/fgame/game/zhenfa/logic"
	"fgame/fgame/game/zhenfa/pbutil"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
	zhenfatemplate "fgame/fgame/game/zhenfa/template"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeZhenQiLevel, command.CommandHandlerFunc(handleZhenQiLevel))
}

func handleZhenQiLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typeStr := c.Args[0]
	posTagStr := c.Args[1]
	levelStr := c.Args[2]
	typeInt64, err := strconv.ParseInt(typeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"typeStr":   typeStr,
				"posTagStr": posTagStr,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置阵旗等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	posTagInt64, err := strconv.ParseInt(posTagStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"typeStr":   typeStr,
				"posTagStr": posTagStr,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置阵旗等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	levelInt64, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"typeStr":   typeStr,
				"posTagStr": posTagStr,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置阵旗等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	typ := zhenfatypes.ZhenFaType(typeInt64)
	if !typ.Vaild() {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"typeStr":   typeStr,
				"posTagStr": posTagStr,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置阵旗等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	posTagType := zhenfatypes.ZhenQiType(posTagInt64)
	if !posTagType.Vaild() {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"typeStr":   typeStr,
				"posTagStr": posTagStr,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置阵旗等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaZhenQiTemplate(typ, posTagType, int32(levelInt64))
	if zhenFaTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"typeStr":   typeStr,
				"posTagStr": posTagStr,
				"levelStr":  levelStr,
			}).Warn("gm:处理设置阵旗等级,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	obj := manager.GetZhenQi(typ, posTagType)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"typeStr":   typeStr,
				"posTagStr": posTagStr,
				"levelStr":  levelStr,
				"error":     err,
			}).Warn("gm:处理设置阵旗等级,该阵旗还未激活")
		playerlogic.SendSystemMessage(pl, lang.ZhenQiAdvancedNoActivate)
		err = nil
		return
	}

	manager.GmSetZhenQiLevel(typ, posTagType, int32(levelInt64))
	//同步属性
	zhenfalogic.ZhenFaPropertyChanged(pl)

	scZhenQiAdvanced := pbutil.BuildSCZhenQiAdvanced(true, obj)
	pl.SendMsg(scZhenQiAdvanced)
	return
}
