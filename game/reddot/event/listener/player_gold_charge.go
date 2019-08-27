package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	reddotlogic "fgame/fgame/game/reddot/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

var (
	goldChangedMap = map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]int32{
		welfaretypes.OpenActivityTypeInvest: map[welfaretypes.OpenActivitySubType]int32{
			welfaretypes.OpenActivityInvestSubTypeServenDay: 1,
		},
		welfaretypes.OpenActivityTypeMergeDrew: map[welfaretypes.OpenActivitySubType]int32{
			welfaretypes.OpenActivityDrewSubTypeBombOre:    1,
			welfaretypes.OpenActivityDrewSubTypeTray:       1,
			welfaretypes.OpenActivityDrewSubTypeChargeDrew: 1,
			welfaretypes.OpenActivityDrewSubTypeSmashEgg:   1,
		},
		welfaretypes.OpenActivityTypeDiscount: map[welfaretypes.OpenActivitySubType]int32{
			welfaretypes.OpenActivityDiscountSubTypeCommon:     1,
			welfaretypes.OpenActivityDiscountSubTypeZhuanSheng: 1,
		},
		welfaretypes.OpenActivityTypeFeedback: map[welfaretypes.OpenActivitySubType]int32{
			welfaretypes.OpenActivityFeedbackSubTypeCycleCharge: 1,
			welfaretypes.OpenActivityFeedbackSubTypeGoldLaBa:    1,
		},
		welfaretypes.OpenActivityTypeMade: map[welfaretypes.OpenActivitySubType]int32{
			welfaretypes.OpenActivityMadeSubTypeResource: 1,
		},
	}
)

//玩家充值元宝
func playerChargeGold(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	reddotlogic.CheckReddotByType(pl, goldChangedMap)
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeGold))
}
