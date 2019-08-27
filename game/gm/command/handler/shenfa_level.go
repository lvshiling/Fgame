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
	shenfalogic "fgame/fgame/game/shenfa/logic"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShenfaLevel, command.CommandHandlerFunc(handleShenfaLevel))

}

func handleShenfaLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	shenfaLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"shenfaLevel": shenfaLevel,
				"error":       err,
			}).Warn("gm:处理设置身法阶别,shenfaLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(shenfaLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"shenfaLevel": shenfaLevel,
				"error":       err,
			}).Warn("gm:处理设置身法阶别,shenfaLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	manager.GmSetShenfaAdvanced(int(shenfaLevel))

	//同步属性
	shenfalogic.ShenfaPropertyChanged(pl)
	shenfaId := manager.GetShenfaInfo().ShenfaId
	scShenfaAdvanced := pbutil.BuildSCShenfaAdavancedFinshed(int32(shenfaLevel), shenfaId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scShenfaAdvanced)
	return
}
