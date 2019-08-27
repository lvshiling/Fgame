package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISCollectFinish(biologyId int32) *crosspb.ISCollectFinish {
	isCollectFinish := &crosspb.ISCollectFinish{}
	isCollectFinish.BiologyId = &biologyId
	return isCollectFinish
}

func BuildISCollectMiZang(npcId int64, biologyId int32, miZangId int32, openType int32) *crosspb.ISCollectMiZangFinish {
	isCollectFinish := &crosspb.ISCollectMiZangFinish{}
	isCollectFinish.NpcId = &npcId
	isCollectFinish.BiologyId = &biologyId
	isCollectFinish.MiZangId = &miZangId
	isCollectFinish.OpenType = &openType
	return isCollectFinish
}
