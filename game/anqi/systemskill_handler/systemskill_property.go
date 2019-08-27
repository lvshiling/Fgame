package sysskill_handler

import (
	anqilogic "fgame/fgame/game/anqi/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeAnQi, systemskill.SystemSkillPropertyHandlerFunc(anQiSystemSkillProperty))
}

//获取暗器系统属性更新
func anQiSystemSkillProperty(pl player.Player) {
	anqilogic.AnqiPropertyChanged(pl)
}
