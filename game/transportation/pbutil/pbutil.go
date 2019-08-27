package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/scene/scene"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	transportationtypes "fgame/fgame/game/transportation/types"
)

func BuildSCAllianceTransportation(startTime int64) *uipb.SCAllianceTransportation {
	scAllianceTransportation := &uipb.SCAllianceTransportation{}
	scAllianceTransportation.StartTime = &startTime

	return scAllianceTransportation
}

func BuildSCPersonalTransportation(startTime int64) *uipb.SCPersonalTransportation {
	scPersonalTransportation := &uipb.SCPersonalTransportation{}
	scPersonalTransportation.StartTime = &startTime
	return scPersonalTransportation
}

func BuildSCPlayerTransportationBriefInfo(personalTimes, allianceTimes int32) *uipb.SCPlayerTransportInfo {
	scPlayerTransportInfo := &uipb.SCPlayerTransportInfo{}
	scPlayerTransportInfo.PersonalTransportTimes = &personalTimes
	scPlayerTransportInfo.AllianceTransportTimes = &allianceTimes

	return scPlayerTransportInfo
}

func BuildSCDistressSignalBroadcast(targetId int64) *uipb.SCDistressSignalBroadcast {
	scDistressSignalBroadcast := &uipb.SCDistressSignalBroadcast{}
	scDistressSignalBroadcast.TargetId = &targetId
	return scDistressSignalBroadcast
}

func BuildSCDistressSignal() *uipb.SCDistressSignal {
	scDistressSignal := &uipb.SCDistressSignal{}
	return scDistressSignal
}

func BuildSCAgreeDistressSignal(pos coretypes.Position) *uipb.SCAgreeDistressSignal {
	scAgreeDistressSignal := &uipb.SCAgreeDistressSignal{}
	targetPos := commonpbutil.BuildPos(pos)
	scAgreeDistressSignal.TargetPosition = targetPos
	return scAgreeDistressSignal
}

func BuildSCTransportBriefInfoNotice(obj *biaochenpc.TransportationObject, n scene.NPC) *uipb.SCTransportBriefInfoNotice {
	scTransportBriefInfoNotice := &uipb.SCTransportBriefInfoNotice{}
	robName := obj.GetRobName()
	scTransportBriefInfoNotice.RobName = &robName
	stateInt := int32(obj.GetState())
	scTransportBriefInfoNotice.State = &stateInt
	typeInt := int32(obj.GetTransportType())
	scTransportBriefInfoNotice.TransportType = &typeInt
	pos := commonpbutil.BuildPos(n.GetPosition())
	scTransportBriefInfoNotice.TargetPosition = pos
	moveId := obj.GetTransportMoveId()
	scTransportBriefInfoNotice.MoveId = &moveId
	startTime := obj.GetCreateTime()
	scTransportBriefInfoNotice.StartTime = &startTime

	return scTransportBriefInfoNotice
}

func BuildSCTransportationProtect() *uipb.SCTransportationProtect {
	scTransportationProtect := &uipb.SCTransportationProtect{}
	return scTransportationProtect
}

func BuildSCReceiveTransportRew(rewGold, rewSilver int64, typ transportationtypes.TransportationType) *uipb.SCReceiveTransportRew {
	scReceiveTransportRew := &uipb.SCReceiveTransportRew{}

	if typ == transportationtypes.TransportationTypeGold {
		silver := int32(0)
		gold := int32(rewGold)
		scReceiveTransportRew.RewGold = &gold
		scReceiveTransportRew.RewSilver = &silver
	}
	if typ == transportationtypes.TransportationTypeSilver {
		silver := int32(rewSilver)
		gold := int32(0)
		scReceiveTransportRew.RewGold = &gold
		scReceiveTransportRew.RewSilver = &silver
	}
	if typ == transportationtypes.TransportationTypeAlliance {
		silver := int32(rewSilver)
		gold := int32(rewGold)
		scReceiveTransportRew.RewGold = &gold
		scReceiveTransportRew.RewSilver = &silver
	}

	ttyp := int32(typ)
	scReceiveTransportRew.TransportType = &ttyp

	return scReceiveTransportRew
}

func BuildSCTransportationProtectNotice(pos coretypes.Position) *uipb.SCTransportationProtectNotice {
	scTransportationProtectNotice := &uipb.SCTransportationProtectNotice{}
	targetPos := commonpbutil.BuildPos(pos)
	scTransportationProtectNotice.TargetPosition = targetPos

	return scTransportationProtectNotice
}

func BuildSCRobSuccessNotice(robNum int64, typ transportationtypes.TransportationType) *uipb.SCRobSuccessNotice {
	scRobSuccessNotice := &uipb.SCRobSuccessNotice{}

	if typ == transportationtypes.TransportationTypeGold {
		silver := int32(0)
		gold := int32(robNum)
		scRobSuccessNotice.RewGold = &gold
		scRobSuccessNotice.RewSilver = &silver
	} else {
		silver := int32(robNum)
		gold := int32(0)
		scRobSuccessNotice.RewGold = &gold
		scRobSuccessNotice.RewSilver = &silver
	}

	ttyp := int32(typ)
	scRobSuccessNotice.TransportType = &ttyp

	return scRobSuccessNotice
}
