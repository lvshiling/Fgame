package advance

import (
	"fgame/fgame/common/lang"
	guajitypes "fgame/fgame/game/guaji/types"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
)

var (
	guaJiTypeMap = map[guajitypes.GuaJiAdvanceType]lingtongdevtypes.LingTongDevSysType{
		guajitypes.GuaJiAdvanceTypeLingTongWeapon: lingtongdevtypes.LingTongDevSysTypeLingBing,
		guajitypes.GuaJiAdvanceTypeLingTongMount:  lingtongdevtypes.LingTongDevSysTypeLingQi,
		guajitypes.GuaJiAdvanceTypeLingTongWing:   lingtongdevtypes.LingTongDevSysTypeLingYi,
		guajitypes.GuaJiAdvanceTypeLingTongShenFa: lingtongdevtypes.LingTongDevSysTypeLingShen,
		guajitypes.GuaJiAdvanceTypeLingTongLingYu: lingtongdevtypes.LingTongDevSysTypeLingYu,
		guajitypes.GuaJiAdvanceTypeLingTongFaBao:  lingtongdevtypes.LingTongDevSysTypeLingBao,
		guajitypes.GuaJiAdvanceTypeLingTongXianTi: lingtongdevtypes.LingTongDevSysTypeLingTi,
	}
)

func getLingTongDevType(typ guajitypes.GuaJiAdvanceType) lingtongdevtypes.LingTongDevSysType {
	return guaJiTypeMap[typ]
}

var (
	guaJiTypeLangMap = map[guajitypes.GuaJiAdvanceType]lang.LangCode{
		guajitypes.GuaJiAdvanceTypeLingTongWeapon: lang.GuaJiLingTongWeaponAdvanced,
		guajitypes.GuaJiAdvanceTypeLingTongMount:  lang.GuaJiLingTongMountAdvanced,
		guajitypes.GuaJiAdvanceTypeLingTongWing:   lang.GuaJiLingTongWingAdvanced,
		guajitypes.GuaJiAdvanceTypeLingTongShenFa: lang.GuaJiLingTongShenFaAdvanced,
		guajitypes.GuaJiAdvanceTypeLingTongLingYu: lang.GuaJiLingTongLingYuAdvanced,
		guajitypes.GuaJiAdvanceTypeLingTongFaBao:  lang.GuaJiLingTongFaBaoAdvanced,
		guajitypes.GuaJiAdvanceTypeLingTongXianTi: lang.GuaJiLingTongXianTiAdvanced,
	}
)

func getLingTongDevLangCode(typ guajitypes.GuaJiAdvanceType) lang.LangCode {
	return guaJiTypeLangMap[typ]
}
