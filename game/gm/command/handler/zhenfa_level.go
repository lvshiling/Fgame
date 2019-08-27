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
	command.Register(gmcommandtypes.CommandTypeZhenFaLevel, command.CommandHandlerFunc(handleZhenFaLevel))
}

func handleZhenFaLevel(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理设置阵法等级,wingLevel不是数字")
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
			}).Warn("gm:处理设置阵法等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	typ := zhenfatypes.ZhenFaType(typeInt64)
	if !typ.Vaild() {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"typeStr":  typeStr,
				"levelStr": levelStr,
				"error":    err,
			}).Warn("gm:处理设置阵法等级,wingLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaTempalte(typ, int32(levelInt64))
	if zhenFaTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"typeStr":  typeStr,
				"levelStr": levelStr,
			}).Warn("gm:处理设置阵法等级,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	obj := manager.GetZhenFaByType(typ)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"typeStr":  typeStr,
				"levelStr": levelStr,
				"error":    err,
			}).Warn("gm:处理设置阵法等级,该阵法还未激活")
		playerlogic.SendSystemMessage(pl, lang.ZhenFaShengJiNoActivate)
		err = nil
		return
	}

	manager.GmSetZhenFaLevel(typ, int32(levelInt64))
	//同步属性
	zhenfalogic.ZhenFaPropertyChanged(pl)

	scZhenFaShengJi := pbutil.BuildSCZhenFaShengJi(true, obj)
	pl.SendMsg(scZhenFaShengJi)
	return
}
