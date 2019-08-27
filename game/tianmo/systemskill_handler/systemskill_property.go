package sysskill_handler

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	tianmologic "fgame/fgame/game/tianmo/logic"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeTianMo, systemskill.SystemSkillPropertyHandlerFunc(tianMoSystemSkillProperty))
}

//获取天魔系统属性更新
func tianMoSystemSkillProperty(pl player.Player) {
	tianmologic.TianMoPropertyChanged(pl)
}
