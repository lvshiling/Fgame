package guaji

import (
	inventoryguaji "fgame/fgame/game/inventory/guaji/guaji"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
)

func init() {
	inventoryguaji.RegisterGuaJiItemUseHandler(itemtypes.ItemTypeGem, inventoryguaji.GuaJiItemUseHandlerFunc(handleGemUse))
}

//TODO
//宝石镶嵌
func handleGemUse(pl player.Player, index int32, num int32) {

}
