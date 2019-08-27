package handler

import (
	"fgame/fgame/client/inventory"
	"fgame/fgame/client/player"
	clientsession "fgame/fgame/client/session"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

//物品改变了
func handleInventoryChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:物品改变了")
	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)
	scInventoryChanged := msg.(*uipb.SCInventoryChanged)
	itemList := convertFromSlotItems(scInventoryChanged.GetItemList())
	inventory.OnInventoryChanged(pl, itemList)
	log.Debug("inventory:物品改变了,成功")
	return
}
