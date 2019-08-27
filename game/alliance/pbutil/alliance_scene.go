package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	playertypes "fgame/fgame/game/player/types"
	propertypbutil "fgame/fgame/game/property/pbutil"
	propertytypes "fgame/fgame/game/property/types"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

func BuildSCAllianceSceneInfo(defendAllianceId int64,
	currentDefendAllianceId int64,
	currentDefendAllianceName string,
	currentDefendAllianceHuFu int64,
	currentDoor int32,
	endTime int64,
	doorRewardList []int32,
	currentReliveAllianceId int64,
	collectReliveAllianceId int64,
	collectReliveStartTime int64,
	warPoint int32,
	yuXi scene.NPC) *uipb.SCAllianceSceneInfo {

	scMsg := &uipb.SCAllianceSceneInfo{}
	scMsg.DefendAllianceId = &defendAllianceId
	scMsg.CurrentDefendAllianceId = &currentDefendAllianceId
	scMsg.CurrentDefendAllianceName = &currentDefendAllianceName
	scMsg.CurrentDefendAllianceHuFu = &currentDefendAllianceHuFu
	scMsg.CurrentDoor = &currentDoor
	scMsg.EndTime = &endTime
	scMsg.DoorRewardList = doorRewardList
	scMsg.CuurentReliveAllianceId = &currentReliveAllianceId
	scMsg.CurrentCollectReliveAllianceId = &collectReliveAllianceId
	scMsg.CollectReliveStartTime = &collectReliveStartTime
	scMsg.WarPoint = &warPoint
	if yuXi != nil {
		scMsg.YuXi = scenepbutil.BuildGeneralCollectInfo(yuXi)
	}
	return scMsg
}

func BuildSCAllianceSceneCall(guardId int32) *uipb.SCAllianceSceneCall {
	scAllianceSceneCall := &uipb.SCAllianceSceneCall{}
	scAllianceSceneCall.GuardId = &guardId
	return scAllianceSceneCall
}

func BuildSCAllianceSceneCalledGuardList(calledGuardList []int32) *uipb.SCAllianceSceneCalledGuardList {
	scAllianceSceneCalledGuardList := &uipb.SCAllianceSceneCalledGuardList{}
	scAllianceSceneCalledGuardList.CalledGuardList = calledGuardList
	return scAllianceSceneCalledGuardList
}

func BuildSCAllianceSceneDoorBroke(door int32) *uipb.SCAllianceSceneDoorBroke {
	scAllianceSceneDoorBroke := &uipb.SCAllianceSceneDoorBroke{}
	scAllianceSceneDoorBroke.Door = &door
	return scAllianceSceneDoorBroke
}

// func BuildSCAllianceSceneOccupying(alliaceId int64) *uipb.SCAllianceSceneOccupying {
// 	scAllianceSceneOccupying := &uipb.SCAllianceSceneOccupying{}
// 	scAllianceSceneOccupying.AllianceId = &alliaceId
// 	return scAllianceSceneOccupying
// }

// func BuildSCAllianceSceneOccupyStop(alliaceId int64) *uipb.SCAllianceSceneOccupyStop {
// 	scAllianceSceneOccupyStop := &uipb.SCAllianceSceneOccupyStop{}
// 	scAllianceSceneOccupyStop.AllianceId = &alliaceId
// 	return scAllianceSceneOccupyStop
// }

func BuildSCAllianceSceneOccupyFinish(alliaceId int64, name string, hufu int64) *uipb.SCAllianceSceneOccupy {
	scAllianceSceneOccupy := &uipb.SCAllianceSceneOccupy{}
	scAllianceSceneOccupy.AllianceId = &alliaceId
	scAllianceSceneOccupy.AllianceName = &name
	scAllianceSceneOccupy.HuFu = &hufu
	return scAllianceSceneOccupy
}

func BuildSCAllianceSceneFinish(allianceId int64, allianceName string) *uipb.SCAllianceSceneFinish {
	scAllianceSceneFinish := &uipb.SCAllianceSceneFinish{}
	scAllianceSceneFinish.AllianceId = &allianceId
	scAllianceSceneFinish.AllianceName = &allianceName
	return scAllianceSceneFinish
}

func BuildSCAllianceSceneDefendHuFuChanged(hufu int64) *uipb.SCAllianceSceneDefendHuFuChanged {
	scAllianceSceneDefendHuFuChanged := &uipb.SCAllianceSceneDefendHuFuChanged{}
	scAllianceSceneDefendHuFuChanged.HuFu = &hufu

	return scAllianceSceneDefendHuFuChanged
}

func BuildSCAllianceSceneGetReward(door int32, itemMap map[int32]int32, rd *propertytypes.RewData) *uipb.SCAllianceSceneGetReward {
	scAllianceSceneGetReward := &uipb.SCAllianceSceneGetReward{}
	scAllianceSceneGetReward.Door = &door
	scAllianceSceneGetReward.DropInfo = droppbutil.BuildSimpleDropInfoList(itemMap)
	scAllianceSceneGetReward.RewProperty = propertypbutil.BuildRewProperty(rd)
	return scAllianceSceneGetReward
}

func BuildSCAllianceHegemonInfo(allianceId int64, allianceName string, mengZhuId int64, mengZhuName string, mengzhuSex playertypes.SexType, totalForce int64, winNum int32) *uipb.SCAllianceHegemonInfo {
	scAllianceHegemonInfo := &uipb.SCAllianceHegemonInfo{}
	scAllianceHegemonInfo.AllianceId = &allianceId
	scAllianceHegemonInfo.AllianceName = &allianceName
	scAllianceHegemonInfo.MengZhuId = &mengZhuId
	scAllianceHegemonInfo.MengZhuName = &mengZhuName
	sexInt := int32(mengzhuSex)
	scAllianceHegemonInfo.MengZhuSex = &sexInt
	scAllianceHegemonInfo.TotalForce = &totalForce
	scAllianceHegemonInfo.WinNum = &winNum
	return scAllianceHegemonInfo
}

func BuildSCAllianceSceneReliveOccupying(alliaceId int64, playerId int64) *uipb.SCAllianceSceneReliveOccupying {
	scAllianceSceneReliveOccupying := &uipb.SCAllianceSceneReliveOccupying{}
	scAllianceSceneReliveOccupying.AllianceId = &alliaceId
	scAllianceSceneReliveOccupying.PlayerId = &playerId
	return scAllianceSceneReliveOccupying
}

func BuildSCAllianceSceneReliveOccupyStop(alliaceId int64) *uipb.SCAllianceSceneReliveOccupyStop {
	scAllianceSceneReliveOccupyStop := &uipb.SCAllianceSceneReliveOccupyStop{}
	scAllianceSceneReliveOccupyStop.AllianceId = &alliaceId
	return scAllianceSceneReliveOccupyStop
}

func BuildSCAllianceSceneReliveOccupyFinish(alliaceId int64, allianceName string, playerName string) *uipb.SCAllianceSceneReliveOccupy {
	scAllianceSceneReliveOccupy := &uipb.SCAllianceSceneReliveOccupy{}
	scAllianceSceneReliveOccupy.AllianceId = &alliaceId
	scAllianceSceneReliveOccupy.AllianceName = &allianceName
	scAllianceSceneReliveOccupy.PlayerName = &playerName
	return scAllianceSceneReliveOccupy
}

func BuildSCAllianceSceneReliveTimeChange(reliveTime int32) *uipb.SCAllianceSceneReliveTimeChange {
	scAllianceSceneReliveTimeChange := &uipb.SCAllianceSceneReliveTimeChange{}
	scAllianceSceneReliveTimeChange.ReliveTime = &reliveTime
	return scAllianceSceneReliveTimeChange
}

func BuildSCAllianceSceneWarPointChanged(warPoint int32) *uipb.SCAllianceSceneWarPointChanged {
	scMsg := &uipb.SCAllianceSceneWarPointChanged{}
	scMsg.WarPoint = &warPoint
	return scMsg
}

func BuildSCAllianceSceneYuXiBroadcast(npc scene.NPC) *uipb.SCAllianceSceneYuXiBroadcast {
	scMsg := &uipb.SCAllianceSceneYuXiBroadcast{}
	scMsg.YuXi = scenepbutil.BuildGeneralCollectInfo(npc)
	return scMsg
}
