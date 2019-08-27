package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCSynthesisStart(synthesisId int32, num int32, resultArr []bool) *uipb.SCSynthesisStart {
	synthesisStart := &uipb.SCSynthesisStart{
		SynthesisId: &synthesisId,
		Num:         &num,
	}
	synthesisStart.IsSuccess = resultArr

	return synthesisStart
}
