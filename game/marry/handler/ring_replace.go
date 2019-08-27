package handler

import (
	//"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	//"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	//"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	//processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_RING_REPLACE_TYPE), dispatch.HandlerFunc(handleMarryRingReplace))
}

//处理婚戒替换信息
//物品替换
func handleMarryRingReplace(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚戒替换消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryRingReplace := msg.(*uipb.CSMarryRingReplace)
	ring := csMarryRingReplace.GetRingType()
	err = marryRingReplace(tpl, marrytypes.MarryRingType(ring))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"ring":     ring,
				"error":    err,
			}).Error("marry:处理婚戒替换消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理婚戒替换消息完成")
	return nil
}

//处理婚戒替换信息逻辑
func marryRingReplace(pl player.Player, ringType marrytypes.MarryRingType) (err error) {
	if !ringType.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ringType": ringType,
		}).Warn("marry:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	spouseId := marryInfo.SpouseId
	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried ||
		marryInfo.Status == marrytypes.MarryStatusTypeDivorce {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ringType": ringType,
		}).Warn("marry:未婚,无法替换")
		playerlogic.SendSystemMessage(pl, lang.MarryRingReplaceNoMarried)
		return
	}

	curRingType := marryInfo.Ring
	if ringType <= curRingType {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"ringType": ringType,
		}).Warn("marry:婚戒替换需要更高级")
		playerlogic.SendSystemMessage(pl, lang.MarryRingReplaceNeedSenior)
		return
	}

	//获取婚烟对戒
	itemTempalte := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())
	ringItem := int32(itemTempalte.TemplateId())

	//判断物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag := inventoryManager.HasEnoughItem(ringItem, 1)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"ringType": ringType,
			}).Warn("marry:婚戒不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//扣除物品
	inventoryReason := commonlog.InventoryLogReasonMarryRingReplace
	flag = inventoryManager.UseItem(ringItem, 1, inventoryReason, inventoryReason.String())
	if !flag {
		panic(fmt.Errorf("marry: marryRingReplace UseItem should be ok"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	//婚戒替换
	manager.RingReplace(ringType)
	marrylogic.MarryPropertyChanged(pl)
	scMarryRingReplace := pbuitl.BuildSCMarryRingReplace(ringItem)
	pl.SendMsg(scMarryRingReplace)
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		return
	}
	scMarryRingChange := pbuitl.BuildSCMarryRingChange(pl.GetId(), ringItem)
	spl.SendMsg(scMarryRingChange)
	return
}
