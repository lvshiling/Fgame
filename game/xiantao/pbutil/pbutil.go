package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	collectnpc "fgame/fgame/game/collect/npc"
	propertytypes "fgame/fgame/game/property/types"
	playerxiantao "fgame/fgame/game/xiantao/player"
)

func BuildSCXiantaoGet(obj *playerxiantao.PlayerXianTaoObject, cpnList []*collectnpc.CollectPointNPC, collectCount int32) *uipb.SCXiantaoGet {
	scMsg := &uipb.SCXiantaoGet{}
	scMsg.AInfo = buildAttendInfo(obj, collectCount)

	for _, cpn := range cpnList {
		scMsg.PInfo = append(scMsg.PInfo, buildPeachPointInfo(cpn))
	}
	return scMsg
}

func BuildSCXiantaoPeachCommit(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCXiantaoPeachCommit {
	scMsg := &uipb.SCXiantaoPeachCommit{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	return scMsg
}

func BuildSCXiantaoPlayerAttendChange(obj *playerxiantao.PlayerXianTaoObject, collectCount int32) *uipb.SCXiantaoPlayerAttendChange {
	scMsg := &uipb.SCXiantaoPlayerAttendChange{}
	scMsg.AInfo = buildAttendInfo(obj, collectCount)
	return scMsg
}

func BuildSCXiantaoPeachPointChange(cpn *collectnpc.CollectPointNPC) *uipb.SCXiantaoPeachPointChange {
	scMsg := &uipb.SCXiantaoPeachPointChange{}
	scMsg.PInfo = buildPeachPointInfo(cpn)

	return scMsg
}

func BuildSCXiantaoResult() *uipb.SCXiantaoResult {
	scMsg := &uipb.SCXiantaoResult{}
	return scMsg
}

func buildAttendInfo(obj *playerxiantao.PlayerXianTaoObject, collectCount int32) *uipb.AttendInfo {
	info := &uipb.AttendInfo{}
	juniorPeachCount := obj.JuniorPeachCount
	highPeachCount := obj.HighPeachCount
	robCount := obj.RobCount
	beRobCount := obj.BeRobCount

	info.CollectCount = &collectCount
	info.JuniorPeachCount = &juniorPeachCount
	info.HighPeachCount = &highPeachCount
	info.RobCount = &robCount
	info.BeRobCount = &beRobCount
	return info
}

func buildPeachPointInfo(cpn *collectnpc.CollectPointNPC) *uipb.PeachPointInfo {
	info := &uipb.PeachPointInfo{}
	cpnObj := cpn.GetCollect()
	biologyId := int32(cpn.GetBiologyTemplate().TemplateId())
	totalCount := cpnObj.GetTotalCount()
	useCount := cpnObj.GetUseCount()
	lastRecoverTime := cpnObj.GetLastRecoverTime()

	info.BiologyId = &biologyId
	info.TotalCount = &totalCount
	info.UseCount = &useCount
	info.LastRecoverTime = &lastRecoverTime
	return info
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

func buildRewProperty(rd *propertytypes.RewData) *uipb.RewProperty {
	rewProperty := &uipb.RewProperty{}
	rewExp := rd.GetRewExp()
	rewExpPoint := rd.GetRewExpPoint()
	rewGold := rd.GetRewGold()
	rewBindGold := rd.GetRewBindGold()
	rewSilver := rd.GetRewSilver()

	rewProperty.Exp = &rewExp
	rewProperty.ExpPoint = &rewExpPoint
	rewProperty.Silver = &rewSilver
	rewProperty.Gold = &rewGold
	rewProperty.BindGold = &rewBindGold

	return rewProperty
}
