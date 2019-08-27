package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCReliveInfo(culTime int32, lastReliveTime int64) *uipb.SCReliveInfo {
	reliveInfo := &uipb.SCReliveInfo{}
	reliveInfo.CulTime = &culTime
	reliveInfo.LastReliveTime = &lastReliveTime
	return reliveInfo
}
