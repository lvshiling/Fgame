package sysskill_handler

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	xiantilogic "fgame/fgame/game/xianti/logic"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeXianTi, systemskill.SystemSkillPropertyHandlerFunc(xianTiSystemSkillProperty))
}

//获取仙体系统属性更新
func xianTiSystemSkillProperty(pl player.Player) {
	xiantilogic.XianTiPropertyChanged(pl)
}
