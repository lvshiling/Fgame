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
	command.Register(gmcommandtypes.CommandTypeLingTongUnrealDan, command.CommandHandlerFunc(handleLingTongUnrealDan))

}

func handleLingTongUnrealDan(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	classTypeStr := c.Args[0]
	unrealDanStr := c.Args[1]

	classTypeInt64, err := strconv.ParseInt(classTypeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"classTypeInt64": classTypeInt64,
				"error":          err,
			}).Warn("gm:处理设置灵童养成类食幻化丹等级,classTypeInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	unrealDanLevel, err := strconv.ParseInt(unrealDanStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"classTypeInt64": classTypeInt64,
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置灵童养成类食幻化丹等级,unrealDanLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	classType := lingtongdevtypes.LingTongDevSysType(classTypeInt64)
	if !classType.Vaild() {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"classTypeInt64": classTypeInt64,
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置灵童养成类食幻化丹等级,classTypeInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	lingTongManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := lingTongManager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId":       pl.GetId(),
			"classTypeStr":   classTypeStr,
			"unrealDanLevel": unrealDanLevel,
		}).Warn("gm:请先激活灵童时装激活系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	lingTongDevHuanHuaTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevHuanHuaTemplate(classType, int32(unrealDanLevel))

	//修改等级
	if lingTongDevHuanHuaTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"classTypeStr":   classTypeStr,
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置灵童养成类食幻化丹等级,unrealDanLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	if lingTongDevInfo == nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"classTypeStr":   classTypeStr,
				"unrealDanLevel": unrealDanLevel,
				"error":          err,
			}).Warn("gm:处理设置灵童养成类幻化,classTypeInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		err = nil
		return
	}

	manager.GmSetLingTongDevUnrealDanLevel(classType, int32(unrealDanLevel))

	//同步属性
	lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)

	scLingTongDevUnrealDan := pbutil.BuildSCLingTongDevUnrealDan(int32(classType), int32(unrealDanLevel), 0)
	pl.SendMsg(scLingTongDevUnrealDan)
	return
}
