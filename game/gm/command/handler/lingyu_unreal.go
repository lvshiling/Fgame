package handler

import (
	"fgame/fgame/common/lang"
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
	command.Register(gmcommandtypes.CommandTypeLingYuUnreal, command.CommandHandlerFunc(handleLingYuUnreal))

}

func handleLingYuUnreal(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	lingYuIdStr := c.Args[0]
	lingYuId, err := strconv.ParseInt(lingYuIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"lingYuId": lingYuId,
				"error":    err,
			}).Warn("gm:处理设置领域幻化,lingYuId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingYuId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"lingYuId": lingYuId,
				"error":    err,
			}).Warn("gm:处理设置领域幻化,lingYuId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	manager.GmSetLingYuUnreal(int(lingYuId))

	//同步属性
	lingyulogic.LingyuPropertyChanged(pl)

	scLingyuUnreal := pbutil.BuildSCLingyuUnreal(int32(lingYuId))
	pl.SendMsg(scLingyuUnreal)
	return
}
