package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	propertylogic "fgame/fgame/game/property/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_PLAYER_GET_DROP_ITEM_TYPE), dispatch.HandlerFunc(handlePlayerGetDropItem))
}

//处理跨服获取物品
func handlePlayerGetDropItem(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理跨服获取物品")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isPlayerGetDropItem := msg.(*crosspb.ISPlayerGetDropItem)
	itemId := isPlayerGetDropItem.GetItemId()
	itemNum := isPlayerGetDropItem.GetNum()
	level := isPlayerGetDropItem.GetLevel()
	bind := itemtypes.ItemBindTypeUnBind
	upstar := isPlayerGetDropItem.GetUpstar()
	attrList := isPlayerGetDropItem.GetAttrList()

	err = getDropItem(tpl, itemId, itemNum, level, upstar, attrList, bind)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:跨服获取物品,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("inventory:跨服获取物品,完成")
	return nil
}

//获取物品
func getDropItem(pl player.Player, itemId int32, num int32, level int32, upstar int32, attrList []int32, bind itemtypes.ItemBindType) (err error) {
	goldLog := commonlog.GoldLogReasonMonsterKilled
	goldReasonText := goldLog.String()
	silverLog := commonlog.SilverLogReasonMonsterKilled
	silverReasonText := silverLog.String()
	inventoryLog := commonlog.InventoryLogReasonMonsterKilled
	inventoryReasonText := inventoryLog.String()
	levelLog := commonlog.LevelLogReasonMonsterKilled
	levelReasonText := levelLog.String()

	flag, err := droplogic.AddItemWithProperty(pl,
		itemId,
		num,
		level,
		upstar,
		attrList,
		bind,
		goldLog,
		goldReasonText,
		silverLog,
		silverReasonText,
		inventoryLog,
		inventoryReasonText,
		levelLog,
		levelReasonText,
	)
	if err != nil {
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("inventory:跨服获取物品,添加物品失败")
	}

	//同步属性
	propertylogic.SnapChangedProperty(pl)

	inventorylogic.SnapInventoryChanged(pl)

	return
}
