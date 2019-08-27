package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenqilogic "fgame/fgame/game/shenqi/logic"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitemplate "fgame/fgame/game/shenqi/template"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENQI_USE_QILING_TYPE), dispatch.HandlerFunc(handleShenQiUseQiLing))
}

//使用器灵
func handleShenQiUseQiLing(s session.Session, msg interface{}) (err error) {
	log.Debug("useqiling:处理使用装备")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csShenqiUseQiling := msg.(*uipb.CSShenqiUseQiling)
	typInt := csShenqiUseQiling.GetShenQiType()
	subTypeInt := csShenqiUseQiling.GetSubType()
	slotIdInt := csShenqiUseQiling.GetSlotId()
	index := csShenqiUseQiling.GetIndex()
	typ := shenqitypes.ShenQiType(typInt)
	subType := shenqitypes.QiLingType(subTypeInt)
	//参数不对
	if !typ.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typInt":   typInt,
			}).Warn("useqiling:器灵类型,错误")
		return
	}
	if !subType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"subTypeInt": subTypeInt,
			}).Warn("useqiling:器灵类型,错误")
		return
	}
	slotId := shenqitypes.CreateQiLingSubType(subType, slotIdInt)
	if !slotId.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"subTypeInt": subTypeInt,
				"slotIdInt":  slotIdInt,
			}).Warn("useqiling:器灵类型,错误")
		return
	}
	if index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("useqiling:物品背包索引,错误")
		return
	}

	err = useQiLing(tpl, typ, subType, slotId, index)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("useqiling:处理使用装备,错误")

		return err
	}
	log.Debug("useqiling:处理使用装备,完成")
	return nil
}

//使用装备
func useQiLing(pl player.Player, typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, slotId shenqitypes.QiLingSubType, index int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenQiQiLing) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("useqiling:装备器灵,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	shenQiManager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)

	//判断神器等级
	qiLingEffectTemp := shenqitemplate.GetShenQiTemplateService().GetShenQiQiLingEffectByArg(typ, subType, slotId)
	shenQiLevel := shenQiManager.GetShenQiDebrisMinLevelByShenQi(typ)
	if shenQiLevel < qiLingEffectTemp.NeedLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"slotId":   slotId.String(),
			}).Warn("useqiling:装备器灵,神器等级不足")
		playerlogic.SendSystemMessage(pl, lang.ShenQiLevelNotEnough)
		return
	}

	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypeQiLing, index)
	//物品不存在
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("useqiling:使用装备,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	itemId := it.ItemId
	bind := it.BindType
	//判断物品是否可以装备
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	qiLingTemplate := itemTemplate.GetShenQiQiLingTemplate()
	if qiLingTemplate == nil {
		log.WithFields(
			log.Fields{
				"index":  index,
				"itemId": itemId,
			}).Warn("useqiling:器灵模板错误")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if !qiLingTemplate.IsTypeByArg(subType, slotId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"slotId":   slotId.String(),
			}).Warn("useqiling:使用装备,此物品槽位不匹配")
		playerlogic.SendSystemMessage(pl, lang.ShenQiQiLingSlotNotMatch)
		return
	}

	slotObj := shenQiManager.GetShenQiQiLingOrInitByArg(typ, subType, slotId)
	if !slotObj.IsEmpty() {
		flag := shenqilogic.TakeOffQiLing(pl, typ, subType, slotId)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ.String(),
					"subType":  int32(subType),
					"slotId":   slotId.String(),
				}).Warn("useqiling:使用器灵,卸下器灵失败")
			playerlogic.SendSystemMessage(pl, lang.ShenQiQiLingSlotReplaceFail)
			return
		}
	}
	//移除物品
	reasonText := commonlog.InventoryLogReasonPutOn.String()
	flag, _ := inventoryManager.RemoveIndex(inventorytypes.BagTypeQiLing, index, 1, commonlog.InventoryLogReasonPutOn, reasonText)
	if !flag {
		panic("useqiling:移除物品应该是可以的")
	}

	flag = shenQiManager.QiLingPutOn(typ, subType, slotId, itemId, bind)
	if !flag {
		panic(fmt.Errorf("useqiling:穿上位置 [%s]应该是可以的", slotId.String()))
	}

	shenQiManager.RefreshShenQiQiLingTaoZhuang(typ)
	//更新属性
	shenqilogic.ShenQiPropertyChanged(pl)

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCShenQiUseQiling(slotObj)
	pl.SendMsg(scMsg)
	return
}
