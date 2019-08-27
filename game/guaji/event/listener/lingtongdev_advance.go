package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
)

var (
	guaJiTypeMap = map[lingtongdevtypes.LingTongDevSysType]guajitypes.GuaJiAdvanceType{
		lingtongdevtypes.LingTongDevSysTypeLingBing: guajitypes.GuaJiAdvanceTypeLingTongWeapon,
		lingtongdevtypes.LingTongDevSysTypeLingQi:   guajitypes.GuaJiAdvanceTypeLingTongMount,
		lingtongdevtypes.LingTongDevSysTypeLingYi:   guajitypes.GuaJiAdvanceTypeLingTongWing,
		lingtongdevtypes.LingTongDevSysTypeLingShen: guajitypes.GuaJiAdvanceTypeLingTongShenFa,
		lingtongdevtypes.LingTongDevSysTypeLingYu:   guajitypes.GuaJiAdvanceTypeLingTongLingYu,
		lingtongdevtypes.LingTongDevSysTypeLingBao:  guajitypes.GuaJiAdvanceTypeLingTongFaBao,
		lingtongdevtypes.LingTongDevSysTypeLingTi:   guajitypes.GuaJiAdvanceTypeLingTongXianTi,
	}
)

func getLingTongDevGuaJiType(typ lingtongdevtypes.LingTongDevSysType) guajitypes.GuaJiAdvanceType {
	return guaJiTypeMap[typ]
}

//坐骑进阶
func lingTongDevAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTongObj, ok := data.(*playerlingtongdev.PlayerLingTongDevObject)
	if !ok {
		return
	}
	classType := lingTongObj.GetClassType()
	guaJiType := getLingTongDevGuaJiType(classType)
	advanceId := lingTongObj.GetAdvancedId()
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, advanceId)
	if lingTongDevTemplate == nil {
		return
	}

	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guaJiType, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevAdvanced, event.EventListenerFunc(lingTongDevAdvanced))
}
