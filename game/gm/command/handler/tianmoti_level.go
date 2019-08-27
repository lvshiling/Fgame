package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	tianmologic "fgame/fgame/game/tianmo/logic"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTianMoTiLevel, command.CommandHandlerFunc(handleTianMoLevel))

}

func handleTianMoLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	tianMoLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"tianMoLevel": tianMoLevel,
				"error":       err,
			}).Warn("gm:处理设置天魔体阶别,tianMoLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := tianmotemplate.GetTianMoTemplateService().GetTianMo(int(tianMoLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"tianMoLevel": tianMoLevel,
				"error":       err,
			}).Warn("gm:处理设置天魔体阶别,天魔体模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	manager.GmSetTianMoAdvanced(int32(tianMoLevel))

	//同步属性
	tianmologic.TianMoPropertyChanged(pl)

	scTianMoAdvanced := pbutil.BuildSCTianMoAdavancedFinshed(int32(tianMoLevel), commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scTianMoAdvanced)
	return
}
