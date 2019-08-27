package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	playerxiantao "fgame/fgame/game/xiantao/player"
	xiantaotypes "fgame/fgame/game/xiantao/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXianTaoCount, command.CommandHandlerFunc(handleSetXianTaoCount))
}

func handleSetXianTaoCount(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typStr := c.Args[0]
	countStr := c.Args[1]
	typInt, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"typInt": typInt,
				"error":  err,
			}).Warn("gm:处理仙桃数量任务,类型typInt不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	count, err := strconv.ParseInt(countStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"count": count,
				"error": err,
			}).Warn("gm:处理仙桃数量任务,类型count不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//参数验证
	typ := xiantaotypes.XianTaoType(int32(typInt))
	if !typ.Valid() {
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typInt,
			}).Warn("gm:仙桃类型,错误")
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	switch typ {
	case xiantaotypes.XianTaoTypeBaiNian:
		manager.GmSetJuniorPeachCount(int32(count))
		break
	case xiantaotypes.XianTaoTypeQianNian:
		manager.GmSetHighPeachCount(int32(count))
		break
	}
	return
}
