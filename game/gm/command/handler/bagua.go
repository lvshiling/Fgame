package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/bagua/pbutil"
	playerbagua "fgame/fgame/game/bagua/player"
	baguatemplate "fgame/fgame/game/bagua/template"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeBaGuaLevel, command.CommandHandlerFunc(handleBaGuaLevel))

}

func handleBaGuaLevel(p scene.Player, c *command.Command) (err error) {
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
			}).Warn("gm:处理设置八卦秘境,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := baguatemplate.GetBaGuaTemplateService().GetBaGuaTemplateByLevel(int32(level))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": level,
				"error": err,
			}).Warn("gm:处理设置八卦秘境,八卦秘境模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	manager.GmSetLevel(int32(level))

	scBaGuaLevel := pbutil.BuildSCBaGuaLevel(int32(level))
	pl.SendMsg(scBaGuaLevel)
	return
}
