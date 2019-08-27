package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/xuechi/pbutil"
	playerxuechi "fgame/fgame/game/xuechi/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXueChi, command.CommandHandlerFunc(handleXueChi))

}

//设置血池剩余血量
func handleXueChi(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	bloodStr := c.Args[0]
	blood, err := strconv.ParseInt(bloodStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"blood": blood,
				"error": err,
			}).Warn("gm:处理设置血池血量,blood不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if blood < 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"blood": blood,
				"error": err,
			}).Warn("gm:处理设置血池血量,level小于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXueChiDataManagerType).(*playerxuechi.PlayerXueChiDataManager)
	manager.GmSetBlood(blood)
	scXueChiBlood := pbutil.BuildSCXueChiBlood(blood)
	pl.SendMsg(scXueChiBlood)
	return
}
