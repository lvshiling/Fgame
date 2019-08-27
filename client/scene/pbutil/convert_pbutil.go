package pbutil

import (
	scenepb "fgame/fgame/common/codec/pb/scene"
	coretypes "fgame/fgame/core/types"
)

func ConvertFromPos(pos *scenepb.Position) coretypes.Position {
	tPos := coretypes.Position{
		X: float64(pos.GetPosX()),
		Y: float64(pos.GetPosY()),
		Z: float64(pos.GetPosZ()),
	}
	return tPos
}
