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
	command.Register(gmcommandtypes.CommandTypeSoulStrengthen, command.CommandHandlerFunc(handleSoulStrengthen))

}

func handleSoulStrengthen(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	tagStr := c.Args[0]
	strengthenStr := c.Args[1]
	tag, err := strconv.ParseInt(tagStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"soulTag": tag,
				"error":   err,
			}).Warn("gm:处理帝魂强化,tag不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	strengthenLevel, err := strconv.ParseInt(strengthenStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":              pl.GetId(),
				"strengthenLevel": strengthenLevel,
				"error":           err,
			}).Warn("gm:处理帝魂强化,strengthenLevel不是数字")
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
		}).Warn("gm:未激活的帝魂,无法强化")
		playerlogic.SendSystemMessage(pl, lang.SoulNotActiveNotStrengthen)
		return
	}

	to := soul.GetSoulService().GetSoulStrengthenTemplateByLevel(soulTag, int32(strengthenLevel))
	if to == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	soulManager.GmSoulStrengthenLevel(soulTag, int32(strengthenLevel))

	scSoulStrengthen := pbutil.BuildSCSoulStrengthen(int32(soulTag), int32(strengthenLevel), 0)
	pl.SendMsg(scSoulStrengthen)
	return
}
