package gm

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func buildCSGMCommand(cmdStr string) *uipb.CSGMCommand {
	cmd := &uipb.CSGMCommand{}
	cmd.Command = &cmdStr
	return cmd
}
