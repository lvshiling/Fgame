package handler

import "fgame/fgame/game/gm/command"
import commandtypes "fgame/fgame/game/gm/command/types"

func init() {
	initCross()
}
func initCross() {
	command.RegisterCross(commandtypes.CommandTypeScene)
	command.RegisterCross(commandtypes.CommandTypeTime)
	command.RegisterCross(commandtypes.CommandTypeBuffAdd)
	command.RegisterCross(commandtypes.CommandTypeHp)
	command.RegisterCross(commandtypes.CommandTypeBuffRemove)
	command.RegisterCross(commandtypes.CommandTypeAddMonster)
	command.RegisterCross(commandtypes.CommandTypeRobot)
	command.RegisterCross(commandtypes.CommandTypeRobotClear)
}
