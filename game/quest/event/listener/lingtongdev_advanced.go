package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

var lingTongDevSysTypeSystemXMap = map[lingtongdevtypes.LingTongDevSysType]questtypes.SystemReachXType{
	lingtongdevtypes.LingTongDevSysTypeLingBao:  questtypes.SystemReachXTypeLingTongFaBao,
	lingtongdevtypes.LingTongDevSysTypeLingBing: questtypes.SystemReachXTypeLingTongWeapon,
	lingtongdevtypes.LingTongDevSysTypeLingQi:   questtypes.SystemReachXTypeLingTongMount,
	lingtongdevtypes.LingTongDevSysTypeLingShen: questtypes.SystemReachXTypeLingTongShenFa,
	lingtongdevtypes.LingTongDevSysTypeLingTi:   questtypes.SystemReachXTypeLingTongXianTi,
	lingtongdevtypes.LingTongDevSysTypeLingYi:   questtypes.SystemReachXTypeLingTongWing,
	lingtongdevtypes.LingTongDevSysTypeLingYu:   questtypes.SystemReachXTypeLingTongLingYu,
}

//玩家灵童养成类进阶
func playerLingTongDevAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTongObj, ok := data.(*playerlingtongdev.PlayerLingTongDevObject)
	if !ok {
		return
	}
	classType := lingTongObj.GetClassType()
	advanceId := lingTongObj.GetAdvancedId()
	systemReachXType, ok := lingTongDevSysTypeSystemXMap[classType]
	if !ok {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(systemReachXType), advanceId)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevAdvanced, event.EventListenerFunc(playerLingTongDevAdavanced))
}
