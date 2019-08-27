package transpotation

import (
	"fgame/fgame/common/lang"
	gamecommon "fgame/fgame/game/common/common"
)

var (
	//正在押镖
	errorTransportationOnDoing = gamecommon.CodeError(lang.TransportationOnDoing)
	//镖车不存在
	errorTransportationNotExsit = gamecommon.CodeError(lang.TransportationNotExist)
	//不是仙盟镖车
	errorTransportationNotAllianceTransportation = gamecommon.CodeError(lang.TransportationNotAllianceTransportation)
	//穿云箭CD
	errorTransportationDistressCD = gamecommon.CodeError(lang.TransportationDistressCD)
)
