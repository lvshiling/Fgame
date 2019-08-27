package systemskill

import (
	"fgame/fgame/game/player"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	"fmt"
)

type SystemSkillPropertyHandler interface {
	SystemSkillProperty(player.Player)
}

type SystemSkillPropertyHandlerFunc func(player.Player)

func (f SystemSkillPropertyHandlerFunc) SystemSkillProperty(pl player.Player) {
	f(pl)
}

var (
	systemSkillPropertyHandlerMap = make(map[sysskilltypes.SystemSkillType]SystemSkillPropertyHandler)
)

func RegisterSystemSkillPropertyHandler(tag sysskilltypes.SystemSkillType, h SystemSkillPropertyHandler) {
	_, ok := systemSkillPropertyHandlerMap[tag]
	if ok {
		panic(fmt.Errorf("systemskillproperty:repeat register %s", tag.String()))
	}
	systemSkillPropertyHandlerMap[tag] = h
}

func SystemSkillPropertyChange(pl player.Player, tag sysskilltypes.SystemSkillType) {
	h, ok := systemSkillPropertyHandlerMap[tag]
	if !ok {
		return
	}
	h.SystemSkillProperty(pl)
}
