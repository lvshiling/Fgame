package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLingTongUnreal, command.CommandHandlerFunc(handleLingTongUnreal))

}

func handleLingTongUnreal(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	classTypeStr := c.Args[0]
	seqIdStr := c.Args[1]
	classTypeInt64, err := strconv.ParseInt(classTypeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"classTypeStr": classTypeStr,
				"seqIdStr":     seqIdStr,
				"error":        err,
			}).Warn("gm:处理设置灵童养成类幻化,classTypeInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	seqIdInt64, err := strconv.ParseInt(seqIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"classTypeStr": classTypeStr,
				"seqIdStr":     seqIdStr,
				"error":        err,
			}).Warn("gm:处理设置灵童养成类幻化,seqIdInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	lingTongManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := lingTongManager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId":     pl.GetId(),
			"classTypeStr": classTypeStr,
			"seqIdStr":     seqIdStr,
		}).Warn("gm:请先激活灵童时装激活系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	classType := lingtongdevtypes.LingTongDevSysType(classTypeInt64)
	if !classType.Vaild() {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"classTypeStr": classTypeStr,
				"seqIdStr":     seqIdStr,
				"error":        err,
			}).Warn("gm:处理设置灵童养成类幻化,classTypeInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	if lingTongDevInfo == nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"classTypeStr": classTypeStr,
				"seqIdStr":     seqIdStr,
				"error":        err,
			}).Warn("gm:处理设置灵童养成类幻化,classTypeInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		err = nil
		return
	}

	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqIdInt64))
	if lingTongDevTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"classTypeStr": classTypeStr,
				"seqIdStr":     seqIdStr,
				"error":        err,
			}).Warn("gm:处理设置灵童养成类幻化,幻化模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager.GmSetLingTongDevUnreal(classType, int(seqIdInt64))
	//同步属性
	lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)

	scLingTongDevUnreal := pbutil.BuildSCLingTongDevUnreal(int32(classType), int32(seqIdInt64))
	pl.SendMsg(scLingTongDevUnreal)
	return
}
