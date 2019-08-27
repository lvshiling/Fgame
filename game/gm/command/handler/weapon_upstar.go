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
	"fgame/fgame/game/weapon/weapon"

	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeWeaponUpstar, command.CommandHandlerFunc(handleWeaponUpstar))
}

func handleWeaponUpstar(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理兵魂升星")
	if len(c.Args) < 2 {
		log.Warn("gm:兵魂升星,参数少于2")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	weaponIdStr := c.Args[0]
	upStarLevelStr := c.Args[1]
	weaponId, err := strconv.ParseInt(weaponIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"error":    err,
				"weaponId": weaponIdStr,
			}).Warn("gm:兵魂升星,weaponId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	upStarLevel, err := strconv.ParseInt(upStarLevelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":          pl.GetId(),
				"error":       err,
				"upStarLevel": upStarLevel,
			}).Warn("gm:兵魂升星,upStarLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//TODO 添加兵魂
	err = weaponUpstar(pl, int32(weaponId), int32(upStarLevel))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"error":    err,
				"weaponId": weaponIdStr,
			}).Warn("gm:兵魂升星,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":       pl.GetId(),
			"weaponId": weaponIdStr,
		}).Debug("gm:处理兵魂升星完成")
	return
}

func weaponUpstar(p scene.Player, weaponId int32, upStarLevel int32) error {
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
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:未激活的兵魂")
		playerlogic.SendSystemMessage(pl, lang.WeaponNotActiveNotEat)
		return nil
	}

	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	tempTemplateObject := to.GetWeaponUpstarByLevel(upStarLevel)
	//修改等级
	if tempTemplateObject == nil {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return nil
	}

	weaponManager.GmWeaponUpstar(weaponId, upStarLevel)
	scWeaponUpstar := pbutil.BuildSCWeaponUpstar(weaponId, upStarLevel, 0)
	pl.SendMsg(scWeaponUpstar)
	return nil
}
