package pbuitl

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

const (
	//生物存在
	bioLive = int32(1)
	//生物死亡
	bioDead = int32(2)
)

func BuildSCFourGodGet(keyNum int32, ncpMap map[int64]scene.NPC) *uipb.SCFourGodGet {
	fourGodGet := &uipb.SCFourGodGet{}
	fourGodGet.KeyNum = &keyNum

	for ncpId, ncp := range ncpMap {
		status := bioLive
		statusTime := int64(0)
		typ := int32(ncp.GetBiologyTemplate().GetBiologyType())

		biologyId := int32(ncp.GetBiologyTemplate().TemplateId())
		if ncp.IsDead() {
			status = bioDead
			statusTime = ncp.GetDeadTime()
		}
		if typ == int32(scenetypes.BiologyScriptTypeFourGodBoss) && ncp.IsDead() {
			fourGodGet.BioList = append(fourGodGet.BioList, buildBio(ncpId, typ, status, statusTime, biologyId))
		} else {
			pos := ncp.GetPosition()
			fourGodGet.BioList = append(fourGodGet.BioList, buildBioWithPos(ncpId, typ, status, statusTime, pos, biologyId))
		}
	}
	return fourGodGet
}

func BuildSCFourGodKeyNumChange(keyNum int32) *uipb.SCFourGodKeyNumChange {
	fourGodKeyNumChange := &uipb.SCFourGodKeyNumChange{}
	fourGodKeyNumChange.KeyNum = &keyNum
	return fourGodKeyNumChange
}

func BuildSCFourGodBioBroadcast(npcId int64, npc scene.NPC) *uipb.SCFourGodBioBroadcast {
	fourGodBioBroadcast := &uipb.SCFourGodBioBroadcast{}
	status := bioLive
	statusTime := int64(0)
	typ := int32(npc.GetBiologyTemplate().GetBiologyType())
	if npc.IsDead() {
		status = bioDead
		statusTime = npc.GetDeadTime()
	}

	biologyId := int32(npc.GetBiologyTemplate().TemplateId())
	if typ == int32(scenetypes.BiologyScriptTypeFourGodBoss) && npc.IsDead() {
		fourGodBioBroadcast.Bio = buildBio(npcId, typ, status, statusTime, biologyId)
	} else {
		pos := npc.GetPosition()
		fourGodBioBroadcast.Bio = buildBioWithPos(npcId, typ, status, statusTime, pos, biologyId)
	}

	return fourGodBioBroadcast
}

func BuildSCFourGodTotal(exp int64, itemMap map[int32]int32) *uipb.SCFourGodTotal {
	fourGodTotal := &uipb.SCFourGodTotal{}
	fourGodTotal.Exp = &exp
	for itemId, num := range itemMap {
		fourGodTotal.ItemList = append(fourGodTotal.ItemList, buildItem(itemId, num))
	}
	return fourGodTotal
}

func BuildSCFourGodUseMarsked(result bool) *uipb.SCFourGodUseMasked {
	fourGodUseMasked := &uipb.SCFourGodUseMasked{}
	fourGodUseMasked.Result = &result
	return fourGodUseMasked
}

func BuildSCFourGodOpenBox(npcId int64) *uipb.SCFourGodOpenBox {
	fourGodOpenBox := &uipb.SCFourGodOpenBox{}
	fourGodOpenBox.NpcId = &npcId
	return fourGodOpenBox
}

func BuildSCFourGodOpenBoxStop(npcId int64) *uipb.SCFourGodOpenBoxStop {
	fourGodOpenBoxStop := &uipb.SCFourGodOpenBoxStop{}
	fourGodOpenBoxStop.NpcId = &npcId
	return fourGodOpenBoxStop
}

func BuildSCFourGodOpenBoxFinish(npcId int64) *uipb.SCFourGodOpenBoxFinish {
	fourGodOpenBoxFinish := &uipb.SCFourGodOpenBoxFinish{}
	fourGodOpenBoxFinish.NpcId = &npcId
	return fourGodOpenBoxFinish
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}

func buildBioWithPos(ncpId int64, typ int32, status int32, statusTime int64, pos types.Position, biologyId int32) *uipb.FourGodBio {
	bio := &uipb.FourGodBio{}
	bio.NcpId = &ncpId
	bio.Typ = &typ
	bio.Status = &status
	bio.StatusTime = &statusTime
	bio.Pos = commonpbutil.BuildPos(pos)
	bio.BiologyId = &biologyId
	return bio
}

func buildBio(ncpId int64, typ int32, status int32, statusTime int64, biologyId int32) *uipb.FourGodBio {
	bio := &uipb.FourGodBio{}
	bio.NcpId = &ncpId
	bio.Typ = &typ
	bio.Status = &status
	bio.StatusTime = &statusTime
	bio.BiologyId = &biologyId
	return bio
}
