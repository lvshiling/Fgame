package pbutil

import (
	scenepb "fgame/fgame/common/codec/pb/scene"
	coretypes "fgame/fgame/core/types"
	scenetypes "fgame/fgame/game/scene/types"
)

func BuildCSObjectMove(uid int64, pos coretypes.Position, angle float32, moveType scenetypes.MoveType) *scenepb.CSObjectMove {
	csObjectMove := &scenepb.CSObjectMove{}
	csObjectMove.Pos = BuildPosition(pos)
	csObjectMove.Angle = &angle
	moveSpeed := float32(0)
	csObjectMove.MoveSpeed = &moveSpeed
	moveTypeInt := int32(moveType)
	csObjectMove.MoveType = &moveTypeInt

	return csObjectMove
}

func BuildPosition(pos coretypes.Position) *scenepb.Position {
	tPos := &scenepb.Position{}
	x := float32(pos.X)
	tPos.PosX = &x
	y := float32(pos.Y)
	tPos.PosY = &y
	z := float32(pos.Z)
	tPos.PosZ = &z
	return tPos
}

func BuildPing() *scenepb.CSPing {
	p := &scenepb.CSPing{}
	return p
}
