package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dragon/dragon"
	dragonlogic "fgame/fgame/game/dragon/logic"
	"fgame/fgame/game/dragon/pbutil"
	playerdragon "fgame/fgame/game/dragon/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DRAGON_FEED_TYPE), dispatch.HandlerFunc(handleDragonFeed))
}

//处理神龙喂养
func handleDragonFeed(s session.Session, msg interface{}) (err error) {
	log.Debug("dragon:处理神龙喂养消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csDragonFeed := msg.(*uipb.CSDragonFeed)
	itemId := csDragonFeed.GetItemId()

	err = dragonFeed(tpl, itemId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"error":    err,
			}).Error("dragon:处理神龙喂养消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dragon:处理神龙喂养消息完成")
	return nil
}

//处理神龙喂养逻辑
func dragonFeed(pl player.Player, itemId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	dragonInfo := manager.GetDragon()
	beforeStage := dragonInfo.StageId
	flag := manager.IsValid(itemId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
		}).Warn("dragon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	num, flag := manager.CanEatNum(itemId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
		}).Warn("dragon:当前食用已达上限")
		playerlogic.SendSystemMessage(pl, lang.DragonEatReachLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	curNum := inventoryManager.NumOfItems(int32(itemId))
	if curNum < num {
		num = curNum
	}
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
		}).Warn("dragon:道具不足,无法喂食")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//喂养扣除物品
	reasonText := commonlog.InventoryLogReasonDragonFeed.String()
	flag = inventoryManager.UseItem(itemId, num, commonlog.InventoryLogReasonDragonFeed, reasonText)
	if !flag {
		panic(fmt.Errorf("dragon: dragonFeed UseItem should be ok"))
	}

	eatFull := manager.DragonFeed(itemId, num)
	rewItemId := int32(0)
	if eatFull {
		dragonTemplate := dragon.GetDragonService().GetDragonTemplate(beforeStage)
		rewItemId = dragonTemplate.ItemId
		dragonlogic.GiveDragonFeedReward(pl, rewItemId, 1)
	}

	//同步物品
	inventorylogic.SnapInventoryChanged(pl)
	//同步属性
	dragonlogic.DragonPropertyChanged(pl)
	scDragonFeed := pbuitl.BuildSCDragonFeed(dragonInfo)
	pl.SendMsg(scDragonFeed)
	return
}
