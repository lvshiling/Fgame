package sysskill_handler

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	winglogic "fgame/fgame/game/wing/logic"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeWing, systemskill.SystemSkillPropertyHandlerFunc(wingSystemSkillProperty))
}

//获取战翼系统属性更新
func wingSystemSkillProperty(pl player.Player) {
	winglogic.WingPropertyChanged(pl)
}
