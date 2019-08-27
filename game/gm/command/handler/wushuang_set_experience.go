package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeWushuangSetExp, command.CommandHandlerFunc(handleWushuangSetExp))
}

func handleWushuangSetExp(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	slotId, err := strconv.ParseInt(c.Args[0], 10, 32)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"slotId": c.Args[0],
				"error":  err,
			}).Warn("gm:处理部位序号错误,格式不正确")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	exp, err := strconv.ParseInt(c.Args[1], 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"exp":   c.Args[1],
				"error": err,
			}).Warn("gm:处理经验值数值错误,格式不正确")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	wushuangManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	bodyPos := wushuangweapontypes.WushuangWeaponPart(slotId)
	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"slotId": c.Args[0],
				"error":  err,
			}).Warn("gm:处理部位序号错误,无双神器类型不正确")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	wushuangManager.GmSetSlotExperience(bodyPos, int64(exp))
	// obj := wushuangManager.GetSlotObjectFromBodyPos(bodyPos)
	// if obj.IsEquip() {
	// 	wushuangweaponlogic.PutOnEquipmentChangeLevel(obj)
	// }
	// obj := wushuangManager.GetSlotObjectFromBodyPos(bodyPos)
	// if obj.IsEquip() {
	// 	itemTemp := item.GetItemService().GetItem(int(obj.GetItemId()))
	// 	level := itemTemp.GetWushuangBaseTemplate().CalculateLevel(obj.GetExperience())
	// 	obj.ChangeLevel(level)
	// 	wushuangweaponlogic.WushuangWeaponPropertyChanged(pl)
	// 	propertylogic.SnapChangedProperty(pl)
	// }
	return
}
