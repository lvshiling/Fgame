package check

import (
	"fgame/fgame/game/guaji/guaji"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	unrealbosslogic "fgame/fgame/game/unrealboss/logic"
	"fgame/fgame/game/unrealboss/unrealboss"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeUnrealBoss, guaji.GuaJiEnterCheckHandlerFunc(unrealEnterCheck))
}

func unrealEnterCheck(pl player.Player) bool {
	//获取幻境boss
	unrealBossList := unrealboss.GetUnrealBossService().GetGuaiJiUnrealBossList(pl.GetForce())
	lenOfUnrealBossList := len(unrealBossList)
	if lenOfUnrealBossList <= 0 {
		return false
	}
	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	for _, boss := range unrealBossList {
		s := boss.GetScene()
		if s == nil {
			continue
		}
		currentPiLao := pl.GetPilao()
		if currentPiLao < boss.GetBiologyTemplate().NeedPilao {
			guaJiData := guaJiManager.GetGuaJiType(guajitypes.GuaJiTypeUnrealBoss)
			if guaJiData.GetOptionValue(guajitypes.GuaJiTypeUnrealBossOptionTypeAutoBuy) == 0 {
				continue
			}
			needBuyNum := int32(1)
			flag := unrealbosslogic.CheckIfPlayerUnrealbossBuy(pl, needBuyNum)
			if !flag {
				continue
			}
			unrealbosslogic.HandleUnrealbossBuy(pl, needBuyNum)
		}
		return true
	}
	return false
}
