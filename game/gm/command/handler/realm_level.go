package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/realm/pbutil"
	playerrealm "fgame/fgame/game/realm/player"
	realmtemplate "fgame/fgame/game/realm/template"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeRealmChangeLevel, command.CommandHandlerFunc(handleRealmChangeLevel))
}

func handleRealmChangeLevel(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置境界")
	if len(c.Args) < 1 {
		log.Warn("gm:处理设置境界,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"level": level,
			}).Warn("gm:处理设置境界,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	if level < 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	//TODO 修改境界
	err = setLevel(pl, int32(level))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"level": level,
			}).Warn("gm:设置境界,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":    pl.GetId(),
			"level": level,
		}).Debug("gm:设置境界,完成")
	return
}

func setLevel(p scene.Player, level int32) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)

	tjtTemplate := realmtemplate.GetRealmTemplateService().GetTianJieTaTemplateByLevel(level)
	if tjtTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
		}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}

	manager.GmSetLevel(level)

	scRealmLevel := pbutil.BuildSCRealmLevel(level)
	pl.SendMsg(scRealmLevel)

	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeRealm.Mask())
	return
}
