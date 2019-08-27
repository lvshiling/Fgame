package handler

import (
	"fgame/fgame/client/inventory"
	"fgame/fgame/client/player"
	clientsession "fgame/fgame/client/session"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

//获取物品
func handleInventoryGet(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:获取物品")
	cs := clientsession.SessionInContext(s.Context())

	pl := cs.Player().(*player.Player)
	scInventoryGet := msg.(*uipb.SCInventoryGet)
	page := scInventoryGet.GetPage()
	items := scInventoryGet.GetItemList()

	inventory.OnInventoryGet(pl, page, convertFromSlotItems(items))
	log.Debug("inventory:获取物品完成")
	return
}
