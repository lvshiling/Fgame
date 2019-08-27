package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_ITEM_GET_TYPE), dispatch.HandlerFunc(handleItemGet))
}

//处理物品获取
func handleItemGet(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理物品获取")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csItemGet := msg.(*scenepb.CSItemGet)
	itemIds := csItemGet.GetItemId()
	//TODO 判断
	err = itemGet(tpl, itemIds)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemIds,
				"err":      err,
			}).Error("scene:处理物品获取,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemIds,
		}).Debug("scene:处理物品获取,完成")

	return nil
}

//玩家是否在场景
func itemGet(pl scene.Player, itemIds []int64) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemIds":  itemIds,
			}).Warn("scene:处理物品获取,玩家不在场景中")
		return
	}
	for _, itemId := range itemIds {
		dropItem := s.GetItem(itemId)
		if dropItem == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
				}).Warn("scene:处理物品获取,物品不存在")
			continue
		}
		if pl.GetScene() != dropItem.GetScene() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"ownerId":  dropItem.GetOwnerId(),
				}).Warn("scene:处理物品获取,不在同一个场景")
			continue
		}
		if !pl.IfCanGetDropItem(dropItem) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"ownerId":  dropItem.GetOwnerId(),
				}).Warn("scene:处理物品获取,物品不属于玩家")
			continue
		}

		flag, err := scenelogic.GetDropItem(pl, dropItem)
		if err != nil {
			return err
		}
		if !flag {
			continue
		}
	}
	// scItemGet := pbutil.BuildSCItemGet(itemIds)
	// err = pl.SendMsg(scItemGet)
	return
}
