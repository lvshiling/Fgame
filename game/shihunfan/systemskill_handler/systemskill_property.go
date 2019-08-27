package sysskill_handler

import (
	"fgame/fgame/game/player"
	shihunfanlogic "fgame/fgame/game/shihunfan/logic"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func init() {
	systemskill.RegisterSystemSkillPropertyHandler(sysskilltypes.SystemSkillTypeShiHunFan, systemskill.SystemSkillPropertyHandlerFunc(shiHunFanSystemSkillProperty))
}

//获取噬魂幡系统属性更新
func shiHunFanSystemSkillProperty(pl player.Player) {
	shihunfanlogic.ShiHunFanPropertyChanged(pl)
}
