package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

func BuildSCWelfareSceneAttend(groupId int32) *uipb.SCWelfareSceneAttend {
	scMsg := &uipb.SCWelfareSceneAttend{}
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCWelfareSceneInfo(createTime int64, tempId int32, npcMap map[int64]scene.NPC, collectNum int32) *uipb.SCWelfareSceneInfo {
	scMsg := &uipb.SCWelfareSceneInfo{}
	scMsg.CreateTime = &createTime
	scMsg.TempId = &tempId
	scMsg.QiYuDaoInfo = buildQiYuDaoInfo(npcMap, collectNum)
	return scMsg
}

func BuildSCWelfareSceneDataChangedNotice(tempId int32, npcMap map[int64]scene.NPC, collectNum int32) *uipb.SCWelfareSceneDataChangedNotice {
	scMsg := &uipb.SCWelfareSceneDataChangedNotice{}
	scMsg.TempId = &tempId
	scMsg.QiYuDaoInfo = buildQiYuDaoInfo(npcMap, collectNum)

	return scMsg
}

func BuildSCWelfareSceneRefersh(groupId int32) *uipb.SCWelfareSceneRefersh {
	scMsg := &uipb.SCWelfareSceneRefersh{}
	scMsg.GroupId = &groupId
	return scMsg
}

func buildQiYuDaoInfo(npcMap map[int64]scene.NPC, collectNum int32) *uipb.QiYuDaoInfo {
	info := &uipb.QiYuDaoInfo{}
	info.BiologyInfo = scenepbutil.BuildGeneralCollectInfoList(npcMap)
	info.CollectNum = &collectNum
	return info
}
