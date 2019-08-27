package pbutil

import uipb "fgame/fgame/common/codec/pb/ui"

var (
	csGMCommand = &uipb.CSGMCommand{}
)

func BuildCSGMCommand(cmd string) *uipb.CSGMCommand {
	csGMCommand = &uipb.CSGMCommand{}
	csGMCommand.Command = &cmd
	return csGMCommand
}
