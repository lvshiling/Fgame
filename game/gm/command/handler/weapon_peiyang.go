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
	command.Register(gmcommandtypes.CommandTypeWeaponPeiYang, command.CommandHandlerFunc(handleWeaponPeiYang))
}

func handleWeaponPeiYang(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理兵魂培养")
	if len(c.Args) < 2 {
		log.Warn("gm:兵魂培养,参数少于2")

		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	weaponIdStr := c.Args[0]
	peiYangLevelStr := c.Args[1]
	weaponId, err := strconv.ParseInt(weaponIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"error":    err,
				"weaponId": weaponIdStr,
			}).Warn("gm:兵魂培养,weaponId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	peiYangLevel, err := strconv.ParseInt(peiYangLevelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"error":        err,
				"peiYangLevel": peiYangLevel,
			}).Warn("gm:兵魂培养,peiYangLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//TODO 添加兵魂
	err = weaponPeiYang(pl, int32(weaponId), int32(peiYangLevel))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"error":    err,
				"weaponId": weaponIdStr,
			}).Warn("gm:兵魂培养,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":       pl.GetId(),
			"weaponId": weaponIdStr,
		}).Debug("gm:处理兵魂培养完成")
	return
}

func weaponPeiYang(p scene.Player, weaponId int32, peiYangLevel int32) error {
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

	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	tempTemplateObject := to.GetWeaponPeiYangByLevel(peiYangLevel)

	//修改等级
	if tempTemplateObject == nil {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
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

	weaponManager.GmWeaponPeiYang(weaponId, peiYangLevel)
	scWeaponEatDan := pbutil.BuildSCWeaponEatDan(weaponId, peiYangLevel, 0)
	pl.SendMsg(scWeaponEatDan)
	return nil
}
