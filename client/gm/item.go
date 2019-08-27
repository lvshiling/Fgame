package gm

import (
	"fgame/fgame/client/player"
	"fgame/fgame/game/gm/command/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func GmChangeItem(pl *player.Player, itemId int32, num int32) (err error) {
	log.WithFields(
		log.Fields{
			"id":     pl.Id(),
			"itemId": itemId,
			"num":    num,
		}).Infoln("gm:设置物品数量")
	cmdStr := fmt.Sprintf("%s %d %d", types.CommandTypeItem, itemId, num)
	cmd := buildCSGMCommand(cmdStr)
	pl.SendMessage(cmd)
	return nil
}

func GmItemClear(pl *player.Player) (err error) {
	log.WithFields(
		log.Fields{
			"id": pl.Id(),
		}).Infoln("gm:清空物品")
	cmdStr := fmt.Sprintf("%s", types.CommandTypeItemClear)
	cmd := buildCSGMCommand(cmdStr)
	pl.SendMessage(cmd)
	return nil
}
