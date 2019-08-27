package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	realmlogic "fgame/fgame/game/realm/logic"
	playerrealm "fgame/fgame/game/realm/player"
	realmtemplate "fgame/fgame/game/realm/template"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTianJieTa, command.CommandHandlerFunc(handleTianJieTa))
}

func handleTianJieTa(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理进入天劫塔")
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
				"level": level,
				"error": err,
			}).Warn("gm:处理跳转天劫塔,天劫塔id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tjtTemplate := realmtemplate.GetRealmTemplateService().GetTianJieTaTemplateByLevel(int32(level))
	if tjtTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
		}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}
	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	if level >= 1 {
		manager.GmSetLevel(int32(level) - 1)
	}

	err = enterTianJieTa(pl, int32(level))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": level,
				"error": err,
			}).Warn("gm:处理跳转天劫塔,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":    pl.GetId(),
			"level": level,
		}).Debug("gm:处理跳转天劫塔,完成")
	return
}

func enterTianJieTa(pl player.Player, level int32) (err error) {
	realmlogic.PlayerEnterTianJieTa(pl, 0, level)
	return
}
