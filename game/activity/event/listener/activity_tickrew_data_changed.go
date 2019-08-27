package listener

// import (
// 	commonlog "fgame/fgame/common/log"
// 	"fgame/fgame/core/event"
// 	"fgame/fgame/game/battle/battle"
// 	battleeventtypes "fgame/fgame/game/battle/event/types"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	droplogic "fgame/fgame/game/drop/logic"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fmt"
// )

// //采集数据变化
// func battlePlayerActivityTickRewDataChanged(target event.EventTarget, data event.EventData) (err error) {
// 	pl, ok := target.(player.Player)
// 	if !ok {
// 		return
// 	}
// 	eventData, ok := data.(*battle.BattlePlayerActivityTickRewDataChangedEventData)
// 	if !ok {
// 		return
// 	}

// 	activityType := eventData.GetActivityType()
// 	addResMap := eventData.GetAddResMap()
// 	addJifen := eventData.GetAddJifen()

// 	//
// 	_, resMap := droplogic.SeperateItems(addResMap)
// 	if len(resMap) > 0 {
// 		reasonGold := commonlog.GoldLogReasonActivityTickRew
// 		reasonGoldText := fmt.Sprintf(reasonGold.String(), activityType)
// 		reasonSilver := commonlog.SilverLogReasonActivityTickRew
// 		reasonSilverText := fmt.Sprintf(reasonSilver.String(), activityType)
// 		reasonLevel := commonlog.LevelLogReasonActivityTickRew
// 		reasonLevelText := fmt.Sprintf(reasonLevel.String(), activityType)
// 		droplogic.AddRes(pl, resMap, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
// 	}

// 	//
// 	if addJifen > 0 {
// 		chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 		chuangShiManager.AddJiFen(addJifen)
// 	}

// 	return
// }

// func init() {
// 	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityTickRewDataChanged, event.EventListenerFunc(battlePlayerActivityTickRewDataChanged))
// }
