package use

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/lucky/pbutil"
	playerlucky "fgame/fgame/game/lucky/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeAttackLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeHpLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeDefenceLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeMountLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeWingLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeAnqiLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeBodyShieldLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeShenfaLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingyuLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeFeatherLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeShieldLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeFaBaoLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeXianTiLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeShiHunFanLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeTianMoTiLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeResouceCard, itemtypes.ItemResourceCardSubTypeExp, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeResouceCard, itemtypes.ItemResourceCardSubTypeDrop, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))

	//灵童
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingBingLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingQiLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingYiLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingShenLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingTongYuLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingBaoLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLuckyRate, itemtypes.ItemLuckySubTypeLingTiLucky, playerinventory.ItemUseHandleFunc(handleSynthesisLucky))
}

//幸运符
func handleSynthesisLucky(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	luckyManager := pl.GetPlayerDataManager(playertypes.PlayerLuckyDataManagerType).(*playerlucky.PlayerLuckyDataManager)
	err = luckyManager.AddLuckyType(itemId, num)
	if err != nil {
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	typ := itemTemplate.GetItemType()
	subType := itemTemplate.GetItemSubType()
	expire := luckyManager.GetLuckyExpireTime(typ, subType)
	scMsg := pbutil.BuildSCLuckyInfoChanged(expire, itemId)
	pl.SendMsg(scMsg)

	flag = true
	return
}
