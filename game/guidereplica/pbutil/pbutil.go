package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
)

func BuildSCGuideReplicaChallenge() *uipb.SCGuideReplicaChallenge {
	scMsg := &uipb.SCGuideReplicaChallenge{}
	return scMsg
}

func BuildSCGuideReplicaChallengeResult(isSuccess bool, itemList []*droptemplate.DropItemData) *uipb.SCGuideReplicaChallengeResult {
	scMsg := &uipb.SCGuideReplicaChallengeResult{}
	scMsg.IsSuccess = &isSuccess
	scMsg.DropList = droppbutil.BuildDropInfoList(itemList)
	return scMsg
}

func BuildSCGuideReplicaSceneInfo(createTime int64, mapId int32, guideType int32, questId int32) *uipb.SCGuideReplicaSceneInfo {
	scMsg := &uipb.SCGuideReplicaSceneInfo{}
	scMsg.CreateTime = &createTime
	scMsg.MapId = &mapId
	scMsg.GuideType = &guideType
	scMsg.QuestId = &questId
	return scMsg
}

func BuildSCGuideReplicaSceneInfoWithCatDog(createTime int64, mapId int32, guideType int32, questId int32, killMap map[guidereplicatypes.CatDogKillType]int32) *uipb.SCGuideReplicaSceneInfo {
	scMsg := BuildSCGuideReplicaSceneInfo(createTime, mapId, guideType, questId)
	scMsg.CatDogInfo = buildCatDogInfo(killMap)
	return scMsg
}

func BuildSCGuideReplicaSceneDataChangedNoticeWithCatDog(guideType int32, killMap map[guidereplicatypes.CatDogKillType]int32) *uipb.SCGuideReplicaSceneDataChangedNotice {
	scMsg := &uipb.SCGuideReplicaSceneDataChangedNotice{}
	scMsg.GuideType = &guideType
	scMsg.CatDogInfo = buildCatDogInfo(killMap)
	return scMsg
}

func BuildSCGuideReplicaSceneDataChangedNoticeWithRescure(guideType int32, herbsFlag bool) *uipb.SCGuideReplicaSceneDataChangedNotice {
	scMsg := &uipb.SCGuideReplicaSceneDataChangedNotice{}
	scMsg.GuideType = &guideType
	scMsg.HerbsFlag = &herbsFlag
	return scMsg
}

func buildCatDogInfo(killMap map[guidereplicatypes.CatDogKillType]int32) *uipb.CatDogInfo {
	info := &uipb.CatDogInfo{}
	info.KillInfo = buildCatDogKillInfo(killMap)
	return info
}
func buildCatDogKillInfo(killMap map[guidereplicatypes.CatDogKillType]int32) (infoList []*uipb.CatDogKillInfo) {
	for typ, num := range killMap {
		info := &uipb.CatDogKillInfo{}
		typInt := int32(typ)
		killNum := num

		info.Type = &typInt
		info.Num = &killNum
		infoList = append(infoList, info)
	}
	return infoList
}

func BuildSCGuideReplicaRescureCommitHerbs(guideType int32) *uipb.SCGuideReplicaPlayerCommonOperate {
	scMsg := &uipb.SCGuideReplicaPlayerCommonOperate{}
	scMsg.GuideType = &guideType
	return scMsg
}
