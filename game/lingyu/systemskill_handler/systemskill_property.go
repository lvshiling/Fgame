package sysskill_handler

import (
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeLingYu, systemskill.SystemSkillPropertyHandlerFunc(lingYuSystemSkillProperty))
}

//获取领域系统属性更新
func lingYuSystemSkillProperty(pl player.Player) {
	lingyulogic.LingyuPropertyChanged(pl)
}
