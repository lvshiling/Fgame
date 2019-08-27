package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISTuLongKillBoss(biologyId int32) *crosspb.ISTuLongKillBoss {
	tuLongKillBoss := &crosspb.ISTuLongKillBoss{}
	tuLongKillBoss.BossId = &biologyId
	return tuLongKillBoss
}

func BuildISTuLongAttend(isLineUp bool, sceneId int64) *crosspb.ISTuLongAttend {
	isTuLongMsg := &crosspb.ISTuLongAttend{}
	isTuLongMsg.IsLineUp = &isLineUp
	isTuLongMsg.SceneId = &sceneId
	return isTuLongMsg
}
