package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	accounttypes "fgame/fgame/login/types"
)

var (
	scGetIdentifyCode = &uipb.SCGetIdentifyCode{}
)

func BuildSCGetIdentifyCode() *uipb.SCGetIdentifyCode {
	return scGetIdentifyCode
}

var (
	scRealNameAuth = &uipb.SCRealNameAuth{}
)

func BuildSCRealNameAuth(state accounttypes.RealNameState) *uipb.SCRealNameAuth {
	scRealNameAuth = &uipb.SCRealNameAuth{}
	stateInt := int32(state)
	scRealNameAuth.RealNameState = &stateInt
	return scRealNameAuth
}

var (
	scExitKaSi = &uipb.SCExitKaSi{}
)

func BuildSCExitKaSi() *uipb.SCExitKaSi {

	return scExitKaSi
}
