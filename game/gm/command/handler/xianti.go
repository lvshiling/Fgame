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
	xiantilogic "fgame/fgame/game/xianti/logic"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXianTi, command.CommandHandlerFunc(handleXianTiLevel))

}

func handleXianTiLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	xianTiLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"xianTiLevel": xianTiLevel,
				"error":       err,
			}).Warn("gm:处理设置仙体阶别,xianTiLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := xianti.GetXianTiService().GetXianTiNumber(int32(xianTiLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"xianTiLevel": xianTiLevel,
				"error":       err,
			}).Warn("gm:处理设置仙体阶别,xianTiLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	manager.GmSetXianTiAdvanced(int(xianTiLevel))

	//同步属性
	xiantilogic.XianTiPropertyChanged(pl)

	xianTiId := manager.GetXianTiInfo().XianTiId
	scXianTiAdvanced := pbutil.BuildSCXianTiAdavancedFinshed(int32(xianTiLevel), xianTiId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scXianTiAdvanced)
	return
}
