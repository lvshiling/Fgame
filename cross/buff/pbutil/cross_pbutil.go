package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	buffcommon "fgame/fgame/game/buff/common"
)

func ConvertFromBuffObject(buffData *crosspb.BuffData) buffcommon.BuffObject {
	ownerId := buffData.GetOwnerId()
	buffId := buffData.GetBuffId()
	groupId := buffData.GetGroupId()
	startTime := buffData.GetStartTime()
	useTime := buffData.GetUseTime()
	culTime := buffData.GetCulTime()
	lastTouchTime := buffData.GetLastTouchTime()
	duration := buffData.GetDuration()
	bo := buffcommon.NewBuffObject(ownerId, buffId, groupId, startTime, useTime, culTime, lastTouchTime, duration, nil)
	return bo
}
