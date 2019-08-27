package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dan/dan"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYRECEIVE_TYPE), dispatch.HandlerFunc(handleAlchemyReceive))
}

//处理练完丹领取信息
func handleAlchemyReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("dan:处理练完丹领取信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAlchemyReceive := msg.(*uipb.CSAlchemyReceive)
	kindId := csAlchemyReceive.GetKindId()

	err = alchemyReceive(tpl, int(kindId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"kindId":   kindId,
				"error":    err,
			}).Error("dan:处理练完丹领取信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
		}).Error("dan:处理练完丹领取信息完成")
	return nil
}

//处理炼丹完成,领取丹药的逻辑
func alchemyReceive(pl player.Player, kindId int) (err error) {
	danManager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	alchemyTemplate := dan.GetDanService().GetAlchemy(kindId)
	if alchemyTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
		}).Warn("dan:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	flag := danManager.IsAlchemyFinish(kindId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
		}).Warn("dan:炼丹未完成,无法领取丹药")
		playerlogic.SendSystemMessage(pl, lang.DanNotFinished)
		return
	}

	alchemyObj := danManager.GetAlchemy(kindId)
	num := int32(alchemyObj.Num)
	itemId := alchemyTemplate.SynthetiseId

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag = inventoryManager.HasEnoughSlot(int32(itemId), num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
		}).Warn("dan:背包空间不足，请清理后再领取")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//丹药放入玩家背包
	reasonText := commonlog.InventoryLogReasonAchemyRec.String()
	flag = inventoryManager.AddItem(int32(itemId), num, commonlog.InventoryLogReasonAchemyRec, reasonText)
	if !flag {
		panic(fmt.Errorf("dan: alchemyReceive add item should be ok"))
	}
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)

	danManager.ClearAlchemyReceive(kindId)
	scAlchemyReceive := pbuitl.BuildSCAlchemyReceive(num, kindId)
	pl.SendMsg(scAlchemyReceive)
	return
}
