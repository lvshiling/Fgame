package guaji_check

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	soullogic "fgame/fgame/game/soul/logic"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeSoul, guaji.GuaJiCheckHandlerFunc(soulGuaJiCheck))
}

//帝魂检查
func soulGuaJiCheck(pl player.Player) {

	//激活检查
	guaJiSoulActiveCheck(pl)
	//升级检查
	guaJiSoulSkillCheck(pl)

}

func guaJiSoulActiveCheck(pl player.Player) {
	soulManager := pl.GetPlayerDataManager(playertypes.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	for soulTag := soultypes.SoulTypeChongHua; soulTag <= soultypes.SoulTypeFuXi; soulTag++ {
		flag := soulManager.IfSoulTagExist(soulTag)
		if flag {
			return
		}

		soulActiveTemplate := soul.GetSoulService().GetSoulActiveTemplate(soulTag)
		//激活的前置帝魂条件
		preSoulCond := soulActiveTemplate.GetPreSoulCond()
		if preSoulCond != nil {
			preSoulTag := preSoulCond.GetSoulType()
			flag := soulManager.IfPreSoul(preSoulTag, preSoulCond.Level)
			if !flag {
				return
			}
		}

		//激活需要物品
		items := soulActiveTemplate.GetNeedItemMap()
		if len(items) != 0 {
			//判断物品是否足够
			flag := inventoryManager.HasEnoughItems(items)
			if !flag {
				return
			}
		}
		soullogic.HandleSoulActive(pl, soulTag)
		playerlogic.SendSystemMessage(pl, lang.GuaJiSoulActive)
	}
}

func guaJiSoulSkillCheck(pl player.Player) {
	//TODO
}
