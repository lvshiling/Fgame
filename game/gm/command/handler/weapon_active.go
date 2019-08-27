package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"

	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeWeaponActive, command.CommandHandlerFunc(handleWeaponActive))
}

func handleWeaponActive(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理兵魂激活")
	if len(c.Args) < 1 {
		log.Warn("gm:兵魂激活,参数少于1")

		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	weaponIdStr := c.Args[0]
	weaponId, err := strconv.ParseInt(weaponIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"error":    err,
				"weaponId": weaponIdStr,
			}).Warn("gm:兵魂激活id,weaponId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//TODO 添加兵魂
	err = weaponActive(pl, int32(weaponId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"error":    err,
				"weaponId": weaponIdStr,
			}).Warn("gm:兵魂激活,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":       pl.GetId(),
			"weaponId": weaponIdStr,
		}).Debug("gm:处理兵魂激活完成")
	return
}

func weaponActive(p scene.Player, weaponId int32) error {
	pl := p.(player.Player)
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	flag := weaponManager.IsValid(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:无效参数")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return nil
	}
	flag = weaponManager.IfWeaponExist(weaponId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:兵魂已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.WeaponRepeatActive)
		return nil
	}

	weaponManager.GmWeaponActive(weaponId)
	scWeaponActive := pbutil.BuildSCWeaponActive(weaponId)
	pl.SendMsg(scWeaponActive)
	return nil
}
