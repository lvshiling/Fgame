package sysskill_handler

import (
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeFaBao, systemskill.SystemSkillPropertyHandlerFunc(faBaoSystemSkillProperty))
}

//获取法宝系统属性更新
func faBaoSystemSkillProperty(pl player.Player) {
	fabaologic.FaBaoPropertyChanged(pl)
}
