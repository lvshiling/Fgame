package check

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	transpotationlogic "fgame/fgame/game/transportation/logic"
	transportationtemplate "fgame/fgame/game/transportation/template"
	"fgame/fgame/game/transportation/transpotation"
	transportationtypes "fgame/fgame/game/transportation/types"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeBiaoChe, guaji.GuaJiCheckHandlerFunc(biaocheGuaJiCheck))
}

func biaocheGuaJiCheck(pl player.Player) {

	//升级检查
	guaJiBiaoCheCheck(pl)

}

func guaJiBiaoCheCheck(pl player.Player) {

	biaoChe := transpotation.GetTransportService().GetTransportation(pl.GetId())
	if biaoChe == nil {
		return
	}

	obj := biaoChe.GetTransportationObject()
	tem := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(obj.GetTransportType())
	if tem == nil {
		return
	}

	//领取资源
	rewItemMap := make(map[int32]int32)

	switch obj.GetState() {
	case transportationtypes.TransportStateTypeFail:
		{
			rewItemMap = tem.GetLastItemMap()
			break
		}
	case transportationtypes.TransportStateTypeFinish:
		{
			rewItemMap = tem.GetFinishItemMap()
			break
		}
	default:
		return
	}

	//背包空间
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlots(rewItemMap) {
		return
	}
	transpotationlogic.HandleReceiveTransportationRew(pl)

	playerlogic.SendSystemMessage(pl, lang.GuaJiBiaoCheRewardGet)
}
