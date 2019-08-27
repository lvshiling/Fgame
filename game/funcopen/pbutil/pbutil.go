package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	funcopentypes "fgame/fgame/game/funcopen/types"
)

func BuildSCFuncOpenList(funcOpenList []funcopentypes.FuncOpenType) *uipb.SCFuncOpenList {
	scFuncOpenList := &uipb.SCFuncOpenList{}
	for _, funcOpen := range funcOpenList {
		scFuncOpenList.FuncOpenList = append(scFuncOpenList.FuncOpenList, int32(funcOpen))
	}
	return scFuncOpenList
}

func BuildSCFuncOpenUpdateList(funcOpenList []funcopentypes.FuncOpenType) *uipb.SCFuncOpenUpdateList {
	scFuncOpenUpdateList := &uipb.SCFuncOpenUpdateList{}
	for _, funcOpen := range funcOpenList {
		scFuncOpenUpdateList.FuncOpenList = append(scFuncOpenUpdateList.FuncOpenList, int32(funcOpen))
	}
	return scFuncOpenUpdateList
}

func BuildSCFuncOpenManualActive(funcOpenList []funcopentypes.FuncOpenType, result int32) *uipb.SCFuncOpenManualActive {
	scFuncOpenManualActive := &uipb.SCFuncOpenManualActive{}
	for _, funcOpen := range funcOpenList {
		scFuncOpenManualActive.FuncOpenList = append(scFuncOpenManualActive.FuncOpenList, int32(funcOpen))
	}
	scFuncOpenManualActive.Result = &result
	return scFuncOpenManualActive
}
