package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeHp, command.CommandHandlerFunc(handleHp))
}

func handleHp(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:设置血量")
	if len(c.Args[0]) <= 0 {
		log.Warn("gm:设置血量,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	hpStr := c.Args[0]
	hp, err := strconv.ParseInt(hpStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"hp":    hpStr,
			}).Warn("gm:设置血量,hp不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	if hp < 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	//TODO 修改物品数量
	err = setHp(pl, hp)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"hp":    hp,
			}).Warn("gm:设置血量,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
			"hp": hp,
		}).Debug("gm:设置血量,完成")
	return
}

func setHp(pl scene.Player, hp int64) (err error) {

	currentHp := pl.GetHP()
	needAdd := hp - currentHp
	if needAdd == 0 {
		return
	}
	if needAdd < 0 {
		// pl.CostHP(-needAdd, 0)
		// if pl.GetScene() != nil {
		// 	scObjectDamage := pbutil.BuildSCObjectDamage(pl, scenetypes.DamageTypeAttack, -needAdd, 0, 0)
		// 	scenelogic.BroadcastNeighborIncludeSelf(pl, scObjectDamage)
		// }
		scenelogic.CostHP(pl, -needAdd, 0, 0, scenetypes.DamageTypeAttack)
	} else {
		pl.AddHP(needAdd)
		if pl.GetScene() != nil {
			scObjectDamage := pbutil.BuildSCObjectDamage(pl, scenetypes.DamageTypeRecovery, needAdd, 0, 0)
			scenelogic.BroadcastNeighborIncludeSelf(pl, scObjectDamage)
		}
	}

	return
}
