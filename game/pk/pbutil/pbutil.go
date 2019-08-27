package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	pktypes "fgame/fgame/game/pk/types"
)

func BuildSCPkStateSwitch(state pktypes.PkState) *uipb.SCPkStateSwitch {
	scPkStateSwitch := &uipb.SCPkStateSwitch{}
	stateInt := int32(state)
	scPkStateSwitch.PkState = &stateInt
	return scPkStateSwitch
}

func BuildSCPKValueChanged(pkValue int32, onlineTime, loginTime int64) *uipb.SCPKValueChanged {
	scPKValueChanged := &uipb.SCPKValueChanged{}
	scPKValueChanged.PkValue = &pkValue
	scPKValueChanged.OnlineTime = &onlineTime
	scPKValueChanged.LoginTime = &loginTime
	return scPKValueChanged
}
