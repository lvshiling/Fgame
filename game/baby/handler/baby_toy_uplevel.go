package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	babylogic "fgame/fgame/game/baby/logic"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	babytypes "fgame/fgame/game/baby/types"
	"fgame/fgame/game/common/common"
	commontypes "fgame/fgame/game/common/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_TOY_UPLEVEL_TYPE), dispatch.HandlerFunc(handleToyUplevel))
}

//处理玩具升级
func handleToyUplevel(s session.Session, msg interface{}) (err error) {
	log.Debug("baby:处理玩具升级")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyToyUplevel)
	slotId := csMsg.GetSlotId()
	suitInt := csMsg.GetSuitType()

	slotPosition := inventorytypes.BodyPositionType(slotId)
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition,
				"suitType":     suitInt,
				"error":        err,
			}).Warn("baby:处理玩具升级,参数无效")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	suitType := babytypes.ToySuitType(suitInt)
	if !suitType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition,
				"suitType":     suitType,
				"error":        err,
			}).Warn("baby:处理玩具升级,参数无效")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = babyToyUplevel(tpl, suitType, slotPosition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition,
				"suitType":     suitType,
				"error":        err,
			}).Error("baby:处理玩具升级,错误")

		return err
	}
	log.Debug("baby:处理玩具升级,完成")
	return nil
}

//宝宝升级
func babyToyUplevel(pl player.Player, suitType babytypes.ToySuitType, posType inventorytypes.BodyPositionType) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	toyBag := babyManager.GetBabyToyBag(suitType)
	if toyBag == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitType": suitType,
			}).Warn("inventory:升级玩具,玩具套装背包不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//玩具不存在
	targetIt := toyBag.GetByPosition(posType)
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitType": suitType,
				"posType":  posType,
			}).Warn("baby:升级玩具失败,玩具不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//能否被升级
	nextLevel := targetIt.GetLevel() + 1
	nextToyUplevelTemp := babytemplate.GetBabyTemplateService().GetBabyToyUplevelTemplate(suitType, posType, nextLevel)
	if nextToyUplevelTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"suitType":  suitType,
				"posType":   posType,
				"nextLevel": nextLevel,
			}).Warn("baby:升级玩具失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.BabyToyFullLevel)
		return
	}

	needItemId := nextToyUplevelTemp.NeedItem
	needItemNum := nextToyUplevelTemp.ItemCount
	if !inventoryManager.HasEnoughItem(needItemId, needItemNum) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"nextLevel":   nextLevel,
				"needItemId":  needItemId,
				"needItemNum": needItemNum,
			}).Warn("baby:升级玩具失败,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗材料
	if needItemNum > 0 {
		reason := commonlog.InventoryLogReasonBabyToyUplevelUse
		reasonText := fmt.Sprintf(reason.String(), suitType, posType.String(), nextLevel)
		flag := inventoryManager.UseItem(needItemId, needItemNum, reason, reasonText)
		if !flag {
			panic(fmt.Errorf("baby:背包升级玩具移除材料应该成功"))
		}
	}

	//计算成功
	var result commontypes.ResultType
	isReturn := false
	isSuccess := mathutils.RandomHit(common.MAX_RATE, int(nextToyUplevelTemp.Rate))
	if isSuccess {
		flag := toyBag.ToyUplevel(posType)
		if !flag {
			panic(fmt.Errorf("baby: 升级玩具应该成功"))
		}

		result = commontypes.ResultTypeSuccess

	} else {
		// 回退计算
		isReturn = mathutils.RandomHit(common.MAX_RATE, int(nextToyUplevelTemp.ReturnRate))
		if isReturn {
			returnLevel := nextToyUplevelTemp.GetFaildReturnTemplate().Level
			toyBag.ToyReturn(posType, returnLevel)
			result = commontypes.ResultTypeBack
		} else {
			result = commontypes.ResultTypeFailed
		}
	}

	if isSuccess || isReturn {
		babylogic.SnapBabyToyChanged(pl)
		babylogic.BabyPropertyChanged(pl)
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCBabyToyUplevel(int32(suitType), int32(posType), targetIt.GetLevel(), int32(result))
	pl.SendMsg(scMsg)

	return
}
