package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	xiantilogic "fgame/fgame/game/xianti/logic"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXianTiUnreal, command.CommandHandlerFunc(handleXianTiUnreal))

}

func handleXianTiUnreal(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	xianTiIdStr := c.Args[0]
	xianTiId, err := strconv.ParseInt(xianTiIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"xianTiId": xianTiId,
				"error":    err,
			}).Warn("gm:处理设置仙体幻化,xianTiId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := xianti.GetXianTiService().GetXianTi(int(xianTiId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"xianTiId": xianTiId,
				"error":    err,
			}).Warn("gm:处理设置仙体幻化,xianTiId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	manager.GmSetXianTiUnreal(int(xianTiId))

	//同步属性
	xiantilogic.XianTiPropertyChanged(pl)

	scXianTiUnreal := pbutil.BuildSCXianTiUnreal(int32(xianTiId))
	pl.SendMsg(scXianTiUnreal)
	return
}
