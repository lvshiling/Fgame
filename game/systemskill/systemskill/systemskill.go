package systemskill

import (
	"fgame/fgame/game/player"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	"fmt"
)

const (
	defaultAdvancedNum = 0
)

type SystemSkillHandler interface {
	SystemSkill(player.Player) (number int32)
}

type SystemSkillHandlerFunc func(player.Player) (number int32)

func (f SystemSkillHandlerFunc) SystemSkill(pl player.Player) (number int32) {
	return f(pl)
}

var (
	systemSkillHandlerMap = make(map[sysskilltypes.SystemSkillType]SystemSkillHandler)
)

func RegisterSystemSkillHandler(tag sysskilltypes.SystemSkillType, h SystemSkillHandler) {
	_, ok := systemSkillHandlerMap[tag]
	if ok {
		panic(fmt.Errorf("systemskill:repeat register %s", tag.String()))
	}
	systemSkillHandlerMap[tag] = h
}

func GetSystemAdvancedNum(pl player.Player, tag sysskilltypes.SystemSkillType) int32 {
	h, ok := systemSkillHandlerMap[tag]
	if !ok {
		return defaultAdvancedNum
	}

	return h.SystemSkill(pl)
}
