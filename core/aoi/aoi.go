package aoi

import "fgame/fgame/core/types"

type AOI interface {
	OnEnterAOI(other AOI)
	OnLeaveAOI(other AOI, complete bool)
	GetId() int64
	GetPosition() types.Position
	SetPosition(types.Position)
}

type AOIManager interface {
	Enter(aoi AOI, pos types.Position)
	Leave(aoi AOI)
	Move(aoi AOI, pos types.Position)
}
