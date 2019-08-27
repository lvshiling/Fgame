package sysskill_handler

import (
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingTongWeapon, systemskill.SystemSkillPropertyHandlerFunc(lingTongWeaponSystemSkillProperty))
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingTongMount, systemskill.SystemSkillPropertyHandlerFunc(lingTongMountSystemSkillProperty))
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingTongWing, systemskill.SystemSkillPropertyHandlerFunc(lingTongWingSystemSkillProperty))
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingTongShenFa, systemskill.SystemSkillPropertyHandlerFunc(lingTongShenFaSystemSkillProperty))
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingTongLingYu, systemskill.SystemSkillPropertyHandlerFunc(lingTongLingYuSystemSkillProperty))
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingTongFaBao, systemskill.SystemSkillPropertyHandlerFunc(lingTongFaBaoSystemSkillProperty))
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingTongXianTi, systemskill.SystemSkillPropertyHandlerFunc(lingTongXianTiSystemSkillProperty))
}

//获取灵兵系统属性更新
func lingTongWeaponSystemSkillProperty(pl player.Player) {
	lingtongdevlogic.LingTongDevPropertyChanged(pl, types.LingTongDevSysTypeLingBing)
}

//获取灵骑系统属性更新
func lingTongMountSystemSkillProperty(pl player.Player) {
	lingtongdevlogic.LingTongDevPropertyChanged(pl, types.LingTongDevSysTypeLingQi)
}

//获取灵翼系统属性更新
func lingTongWingSystemSkillProperty(pl player.Player) {
	lingtongdevlogic.LingTongDevPropertyChanged(pl, types.LingTongDevSysTypeLingYi)
}

//获取灵身系统属性更新
func lingTongShenFaSystemSkillProperty(pl player.Player) {
	lingtongdevlogic.LingTongDevPropertyChanged(pl, types.LingTongDevSysTypeLingShen)
}

//获取灵域系统属性更新
func lingTongLingYuSystemSkillProperty(pl player.Player) {
	lingtongdevlogic.LingTongDevPropertyChanged(pl, types.LingTongDevSysTypeLingYu)
}

//获取灵宝系统属性更新
func lingTongFaBaoSystemSkillProperty(pl player.Player) {
	lingtongdevlogic.LingTongDevPropertyChanged(pl, types.LingTongDevSysTypeLingBao)
}

//获取灵骑系统属性更新
func lingTongXianTiSystemSkillProperty(pl player.Player) {
	lingtongdevlogic.LingTongDevPropertyChanged(pl, types.LingTongDevSysTypeLingTi)
}
