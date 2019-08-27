package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSoulUpgrade, command.CommandHandlerFunc(handleSoulUpgrade))

}

func handleSoulUpgrade(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	tagStr := c.Args[0]
	orderStr := c.Args[1]
	tag, err := strconv.ParseInt(tagStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"soulTag": tag,
				"error":   err,
			}).Warn("gm:处理帝魂升级,tag不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	order, err := strconv.ParseInt(orderStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"order": order,
				"error": err,
			}).Warn("gm:处理帝魂升级,order不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	soulTag := soultypes.SoulType(tag)
	soulManager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := soulTag.Valid()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = soulManager.IfSoulTagExist(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("gm:未激活的帝魂,无法升级")
		playerlogic.SendSystemMessage(pl, lang.SoulNotActiveNotUpgrade)
		return
	}

	to := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, int32(order))
	if to == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	soulManager.GmSoulAwakenOrder(soulTag, int32(order))

	scSoulUpgrade := pbutil.BuildSCSoulUpgrade(int32(soulTag), int32(order))
	pl.SendMsg(scSoulUpgrade)
	return
}
