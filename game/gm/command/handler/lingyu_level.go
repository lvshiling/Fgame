package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLingyuLevel, command.CommandHandlerFunc(handleLingyuLevel))

}

func handleLingyuLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	lingyuLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"lingyuLevel": lingyuLevel,
				"error":       err,
			}).Warn("gm:处理设置领域阶别,lingyuLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(int32(lingyuLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"lingyuLevel": lingyuLevel,
				"error":       err,
			}).Warn("gm:处理设置领域阶别,lingyuLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	manager.GmSetLingyuAdvanced(int(lingyuLevel))

	//同步属性
	lingyulogic.LingyuPropertyChanged(pl)
	lingyuId := manager.GetLingyuInfo().LingyuId
	scLingyuAdvanced := pbutil.BuildSCLingyuAdavancedFinshed(int32(lingyuLevel), lingyuId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scLingyuAdvanced)
	return
}
