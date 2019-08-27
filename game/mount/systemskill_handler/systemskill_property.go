package sysskill_handler

import (
	mountlogic "fgame/fgame/game/mount/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeMount, systemskill.SystemSkillPropertyHandlerFunc(mountSystemSkillProperty))
}

//获取坐骑系统属性更新
func mountSystemSkillProperty(pl player.Player) {
	mountlogic.MountPropertyChanged(pl)
}
