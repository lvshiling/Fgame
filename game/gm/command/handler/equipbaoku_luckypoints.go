package handler

import (
	"fgame/fgame/common/lang"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
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
	command.Register(gmcommandtypes.CommandTypeEquipBaoKuLuckyPoints, command.CommandHandlerFunc(handleSetEquipBaoKuLuckyPoints))
}

func handleSetEquipBaoKuLuckyPoints(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理装备宝库幸运值任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	typStr := c.Args[1]
	typ, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理装备宝库积分任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	baoKuTyp := equipbaokutypes.BaoKuType(typ)
	if !baoKuTyp.Valid() {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理装备宝库积分任务,宝库类型不合法")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	manager.GMSetLuckyPoints(int32(num), baoKuTyp)

	// obj := manager.GetEquipBaoKuObj(baoKuTyp)
	// logList := equipbaoku.GetEquipBaoKuService().GetLogByTime(0, baoKuTyp)
	// shopBuyCountMap := manager.GetEquipBaoKuShopBuyAll()

	// scEquipBaoKuInfoGet := pbutil.BuildSCEquipBaoKuInfoGet(obj, logList, shopBuyCountMap, int32(baoKuTyp))
	// pl.SendMsg(scEquipBaoKuInfoGet)
	return
}
