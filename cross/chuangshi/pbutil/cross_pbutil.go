package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISChuangShiEnterCity(isLineup bool) *crosspb.ISChuangShiEnterCity {
	isMsg := &crosspb.ISChuangShiEnterCity{}
	isMsg.IsLineUp = &isLineup
	return isMsg
}
