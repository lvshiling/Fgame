package use

// func init() {
// 	playerinventory.RegisterUseHandler(itemtypes.ItemTypeVigorousPill, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleVigorousPill))
// }

// func handleVigorousPill(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
// 	itemId := it.ItemId
// 	itemTemplate := item.GetItemService().GetItem(int(itemId))
// 	buffId := itemTemplate.TypeFlag1
// 	scenelogic.AddBuffs(pl, buffId, pl.GetId(), num, common.MAX_RATE)
// 	flag = true
// 	return
// }
