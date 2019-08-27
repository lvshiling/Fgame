package use

import (
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeSilver, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeGold, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeBindGold, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeYaoPai, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeKey, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeExp, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeStorage, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeNormalFireworks, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeSeniorFireworks, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeShaQi, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeBloodZhi, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeZhuoQi, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeXingChen, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeGongDe, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeEquipBaoKuAttendPoints, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeLingQi, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeQianNianXianTao, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeBaiNianXianTao, playerinventory.ItemUseHandleFunc(handleResource))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeMaterialBaoKuAttendPoints, playerinventory.ItemUseHandleFunc(handleResource))

}

//资源使用
func handleResource(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	if num <= 0 {
		return
	}
	itemId := it.ItemId

	goldReason := commonlog.GoldLogReasonInventoryResourceUse
	silverReason := commonlog.SilverLogReasonInventoryResourceUse
	levelReason := commonlog.LevelLogReasonInventoryResourceUse
	itemData := droptemplate.CreateItemData(itemId, num, 0, itemtypes.ItemBindTypeUnBind)
	flag, err = droplogic.AddItem(pl, itemData, goldReason, goldReason.String(), silverReason, silverReason.String(), 0, "", levelReason, levelReason.String())
	if err != nil {
		return
	}
	if !flag {
		return
	}

	propertylogic.SnapChangedProperty(pl)

	flag = true
	return
}
