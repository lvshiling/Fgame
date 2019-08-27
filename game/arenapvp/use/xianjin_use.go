package use

import (
	commonlog "fgame/fgame/common/log"
	feedbackfeelogic "fgame/fgame/game/feedbackfee/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fmt"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeXianJinYuanBaoKa, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleXianJinUse))
}

//现金卡使用
func handleXianJinUse(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))

	money := itemTemplate.TypeFlag1 * num
	reason := commonlog.FeedbackLogReasonUseItemRew
	reasonText := fmt.Sprintf(reason.String(), num)
	feedbackfeelogic.AddMoney(pl, money, reason, reasonText)

	flag = true
	return
}
