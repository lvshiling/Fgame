package use

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/global"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/dao"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeRenameCard, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handlerRename))
}

// 改名卡
func handlerRename(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {

	if pl.IsCross() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:处理改名请求,玩家处于跨服")
		playerlogic.SendSystemMessage(pl, lang.InventoryInCross)
		return
	}

	newName := args
	if len(newName) <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("player:处理改名请求,名字无效")
		playerlogic.SendSystemMessage(pl, lang.NameInvalid)
		return
	}

	// 重名
	// TODO 优化
	serverId := global.GetGame().GetServerIndex()
	pe, err := dao.GetPlayerDao().QueryByName(serverId, newName)
	if err != nil {
		return
	}
	if pe != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("player:处理改名请求，名字已经存在")
		playerlogic.SendSystemMessage(pl, lang.NameAlreadyExist)
		return
	}

	pl.ChangeName(newName)
	scMsg := pbutil.BuildSCPlayerNameChanged(newName)
	pl.SendMsg(scMsg)

	flag = true
	return
}
