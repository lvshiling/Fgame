package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	playerfoe "fgame/fgame/game/foe/player"
	playercommon "fgame/fgame/game/player/common"
	playerpbutil "fgame/fgame/game/player/pbutil"
	playertypes "fgame/fgame/game/player/types"
	xiantaotypes "fgame/fgame/game/xiantao/types"
)

func BuildSCFoesGet(foeList []*uipb.FoeInfo) *uipb.SCFoesGet {
	scFoesGet := &uipb.SCFoesGet{}
	scFoesGet.FoeList = foeList
	return scFoesGet
}

func BuildFoe(attackId int64, killTime int64, info *playercommon.PlayerInfo) *uipb.FoeInfo {
	foeInfo := &uipb.FoeInfo{}
	foeInfo.FoeId = &attackId
	foeInfo.KillTime = &killTime
	foeInfo.PlayerBasicInfo = playerpbutil.BuildPlayerBasicInfo(info, false)
	return foeInfo
}

func BuildSCFoeRemove(foeId int64) *uipb.SCFoeRemove {
	scFoeRemove := &uipb.SCFoeRemove{}
	scFoeRemove.FoeId = &foeId
	return scFoeRemove
}

func BuildSCFoeAdd(foeInfo *uipb.FoeInfo) *uipb.SCFoeAdd {
	scFoeAdd := &uipb.SCFoeAdd{}
	scFoeAdd.FoeInfo = foeInfo
	return scFoeAdd
}

func BuildSCFoeViewPosNoOnline(foeId int64, foeName string) *uipb.SCFoeViewPos {
	scFoeViewPos := &uipb.SCFoeViewPos{}
	scFoeViewPos.FoeId = &foeId
	scFoeViewPos.FoeName = &foeName
	result := false
	scFoeViewPos.Result = &result
	return scFoeViewPos
}

func BuildSCFoeViewCross(foeId int64, foeName, name string) *uipb.SCFoeViewPos {
	scFoeViewPos := &uipb.SCFoeViewPos{}
	result := true
	isCross := true
	scFoeViewPos.Result = &result
	scFoeViewPos.IsCross = &isCross
	scFoeViewPos.FoeId = &foeId
	scFoeViewPos.Name = &name
	scFoeViewPos.FoeName = &foeName
	return scFoeViewPos
}

func BuildSCFoeViewPos(foeId int64, foeName, name string, mapId int32, pos coretypes.Position) *uipb.SCFoeViewPos {
	scFoeViewPos := &uipb.SCFoeViewPos{}
	result := true
	scFoeViewPos.Result = &result
	scFoeViewPos.FoeId = &foeId
	scFoeViewPos.FoeName = &foeName
	scFoeViewPos.Name = &name
	scFoeViewPos.MapId = &mapId
	scFoeViewPos.Pos = commonpbutil.BuildPos(pos)
	return scFoeViewPos
}

func BuildSCFoeTransfer(foeId int64) *uipb.SCFoeTransfer {
	scFoeTransfer := &uipb.SCFoeTransfer{}
	scFoeTransfer.FoeId = &foeId
	return scFoeTransfer
}

func BuildSCFoeNotice(foeId int64, foeName string, foeRole playertypes.RoleType, foeSex playertypes.SexType, sceneType int32) *uipb.SCFoeNotice {
	scMsg := &uipb.SCFoeNotice{}
	role := int32(foeRole)
	sex := int32(foeSex)

	scMsg.FoeId = &foeId
	scMsg.FoeName = &foeName
	scMsg.FoeRole = &role
	scMsg.FoeSex = &sex
	scMsg.SceneType = &sceneType
	return scMsg
}

func BuildSCFoeKillNotice(foeId int64, foeName string, foeRole playertypes.RoleType, foeSex playertypes.SexType, sceneType int32) *uipb.SCFoeKillNotice {
	scMsg := &uipb.SCFoeKillNotice{}
	role := int32(foeRole)
	sex := int32(foeSex)

	scMsg.DeadId = &foeId
	scMsg.DeadName = &foeName
	scMsg.DeadRole = &role
	scMsg.DeadSex = &sex
	scMsg.SceneType = &sceneType
	return scMsg
}

func BuildSCFoeNoticeShenYu(foeId int64, foeName string, foeRole playertypes.RoleType, foeSex playertypes.SexType, sceneType, dropKey int32) *uipb.SCFoeNotice {
	scMsg := BuildSCFoeNotice(foeId, foeName, foeRole, foeSex, sceneType)

	info := &uipb.ShenYuFoeNoticeInfo{}
	info.DropKeyNum = &dropKey

	scMsg.ShenYuFoeInfo = info
	return scMsg
}

func BuildSCFoeNoticeXianTao(foeId int64, foeName string, foeRole playertypes.RoleType, foeSex playertypes.SexType, sceneType int32, dropNumMap map[xiantaotypes.XianTaoType]int32) *uipb.SCFoeNotice {
	scMsg := BuildSCFoeNotice(foeId, foeName, foeRole, foeSex, sceneType)

	xianTaoFoeInfo := &uipb.XianTaoFoeNoticeInfo{}
	for typ, count := range dropNumMap {
		xianTaoFoeInfo.Info = append(xianTaoFoeInfo.Info, buildFoeXianTaoDrop(typ, count))
	}

	scMsg.XianTaoFoeInfo = xianTaoFoeInfo
	return scMsg
}

func BuildSCFoeKillNoticeXianTao(foeId int64, foeName string, foeRole playertypes.RoleType, foeSex playertypes.SexType, sceneType int32, dropNumMap map[xiantaotypes.XianTaoType]int32) *uipb.SCFoeKillNotice {
	scMsg := BuildSCFoeKillNotice(foeId, foeName, foeRole, foeSex, sceneType)

	xianTaoFoeInfo := &uipb.XianTaoFoeNoticeInfo{}
	for typ, count := range dropNumMap {
		xianTaoFoeInfo.Info = append(xianTaoFoeInfo.Info, buildFoeXianTaoDrop(typ, count))
	}

	scMsg.XianTaoFoeInfo = xianTaoFoeInfo
	return scMsg
}

func BuildSCFoeFeedback(isProtected bool, foeName string, lostSilver int32, args string, foeSex playertypes.SexType) *uipb.SCFoeFeedback {
	scMsg := &uipb.SCFoeFeedback{}
	sexInt := int32(foeSex)
	scMsg.IsProtected = &isProtected
	scMsg.FoeName = &foeName
	scMsg.LostSilver = &lostSilver
	scMsg.Args = &args
	scMsg.Sex = &sexInt
	return scMsg
}

func BuildSCFoeFeedbackRead(feedbackList []*playerfoe.PlayerFoeFeedbackObject) *uipb.SCFoeFeedbackRead {
	scMsg := &uipb.SCFoeFeedbackRead{}
	for _, obj := range feedbackList {
		scMsg.FeedbackList = append(scMsg.FeedbackList, buildFoeFeedback(obj))
	}
	return scMsg
}

func BuildSCFoeFeedbackNotice(plId int64, plName string, isProtected bool, needSilver int64) *uipb.SCFoeFeedbackNotice {
	scMsg := &uipb.SCFoeFeedbackNotice{}
	scMsg.PlayerId = &plId
	scMsg.PlayerName = &plName
	scMsg.IsProtected = &isProtected
	scMsg.LostSilver = &needSilver
	return scMsg
}

func BuildSCFoeFeedbackBuyProtect(expireTime int64) *uipb.SCFoeFeedbackBuyProtect {
	scMsg := &uipb.SCFoeFeedbackBuyProtect{}
	scMsg.ExpireTime = &expireTime
	return scMsg
}

func BuildSCFoeFeedbackInfo(expireTime int64, feedbackList []*playerfoe.PlayerFoeFeedbackObject) *uipb.SCFoeFeedbackInfo {
	scMsg := &uipb.SCFoeFeedbackInfo{}
	scMsg.ProtectExpireTime = &expireTime

	for _, obj := range feedbackList {
		scMsg.FeedbackList = append(scMsg.FeedbackList, buildFoeFeedback(obj))
	}
	return scMsg
}

func buildFoeFeedback(obj *playerfoe.PlayerFoeFeedbackObject) *uipb.FoeFeedback {
	info := &uipb.FoeFeedback{}
	name := obj.GetFeedbackName()
	isProtect := obj.GetIsProtect()
	info.IsProtected = &isProtect
	info.PlayerName = &name

	return info
}

func buildFoeXianTaoDrop(typ xiantaotypes.XianTaoType, count int32) *uipb.XianTaoInfo {
	info := &uipb.XianTaoInfo{}
	typInt := int32(typ)
	countInt := count
	info.Typ = &typInt
	info.Num = &countInt

	return info
}
