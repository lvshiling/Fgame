package sysskill_handler

import (
	"fgame/fgame/game/player"
	shenfalogic "fgame/fgame/game/shenfa/logic"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeShenFa, systemskill.SystemSkillPropertyHandlerFunc(shenFaSystemSkillProperty))
}

//获取身法系统属性更新
func shenFaSystemSkillProperty(pl player.Player) {
	shenfalogic.ShenfaPropertyChanged(pl)
}
