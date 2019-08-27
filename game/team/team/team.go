package team

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	teamtypes "fgame/fgame/game/team/types"
	"sort"
)

type TeamObject struct {
	teamId       int64
	force        int64
	memberList   []*TeamMemberObject
	applyDataMap map[int64]*TeamApplyData
	purpose      teamtypes.TeamPurposeType
	match        bool
	autoReview   bool
	copyBattle   bool
	isBattling   bool
	rushTime     int64
	kickTimeMap  map[int64]int64
}

func (td *TeamObject) GetTeamId() int64 {
	return td.teamId
}

func (td *TeamObject) GetForce() int64 {
	return td.force
}

func (td *TeamObject) GetMemberList() []*TeamMemberObject {
	return td.memberList
}

func (td *TeamObject) RemoveApply(playerId int64) {
	delete(td.applyDataMap, playerId)
}

func (td *TeamObject) ClearAllApply() {
	//TODO 优化
	td.applyDataMap = make(map[int64]*TeamApplyData)
}
func (td *TeamObject) GetAllApplyList() []*TeamApplyData {
	applyList := make([]*TeamApplyData, 0, len(td.applyDataMap))

	for _, applyData := range td.applyDataMap {
		applyList = append(applyList, applyData)
	}
	if len(applyList) > 1 {
		sort.Sort(TeamApplyDataList(applyList))
	}

	return applyList
}

func (td *TeamObject) GetApply(playerId int64) *TeamApplyData {
	applyData, ok := td.applyDataMap[playerId]
	if !ok {
		return nil
	}
	return applyData
}

func (td *TeamObject) IfCanApply(playerId int64) bool {
	applyData := td.GetApply(playerId)
	if applyData != nil {
		//TODO 判断过期时间
		return false
	}

	return true
}

func (td *TeamObject) Apply(playerId int64) {
	applyData := td.GetApply(playerId)
	now := global.GetGame().GetTimeService().Now()
	if applyData == nil {
		applyData = NewTeamApplyData(playerId, now)
		td.applyDataMap[playerId] = applyData
	}
	applyData.applyTime = now

	return
}

func (td *TeamObject) GetNum() int32 {
	return int32(len(td.memberList))
}

func (td *TeamObject) GetOfflineNum() int32 {
	num := int32(0)
	for _, member := range td.memberList {
		if !member.online {
			num += 1
		}
	}
	return num
}

func (td *TeamObject) GetCaptain() *TeamMemberObject {
	return td.memberList[0]
}

func (td *TeamObject) GetMember(playerId int64) (mem *TeamMemberObject, pos int32) {
	pos = -1
	for index, member := range td.memberList {
		if member.GetPlayerId() == playerId {
			pos = int32(index)
			mem = member
			break
		}
	}
	return
}

func (td *TeamObject) TransferCaptain(captainPos int32) bool {
	if captainPos <= 0 {
		return false
	}
	num := td.GetNum()
	if num <= captainPos {
		return false
	}
	td.memberList[0], td.memberList[captainPos] = td.memberList[captainPos], td.memberList[0]
	td.autoReview = false
	return true
}

func (td *TeamObject) UpdateForce() {
	force := int64(0)
	for _, member := range td.memberList {
		force += member.GetForce()
	}
	td.force = force
	return
}

func (td *TeamObject) GetTeamName() string {
	return td.memberList[0].GetName()
}

func (td *TeamObject) AddMember(member *TeamMemberObject) {
	td.memberList = append(td.memberList, member)
	td.UpdateForce()
}

func (td *TeamObject) RemoveMember(memberId int64) {
	mem, pos := td.GetMember(memberId)
	if mem == nil {
		return
	}
	td.memberList = append(td.memberList[:pos], td.memberList[pos+1:]...)
}

func (td *TeamObject) UpdateMemberForce(memberId int64, force int64) {
	mem, _ := td.GetMember(memberId)
	if mem == nil {
		return
	}
	mem.UpdateForce(force)
	td.UpdateForce()
}

func (td *TeamObject) Matching() bool {
	if td.match {
		return false
	}
	td.match = true
	return true
}

func (td *TeamObject) StopMatching() bool {
	if !td.match {
		return false
	}
	td.match = false
	return true
}

func (td *TeamObject) IsMatch() bool {
	return td.match
}

func (td *TeamObject) IsAutoReview() bool {
	return td.autoReview
}

func (td *TeamObject) GetTeamPurpose() teamtypes.TeamPurposeType {
	return td.purpose
}

func (td *TeamObject) Equal(purpose teamtypes.TeamPurposeType) bool {
	return td.purpose == purpose
}

func (td *TeamObject) IsFull() bool {
	return len(td.memberList) == teamtypes.TeamMaxNum
}

func (td *TeamObject) IsCopyBattle() bool {
	return td.copyBattle
}

func (td *TeamObject) TeamCopyStartBattling() bool {
	if td.copyBattle {
		return false
	}
	td.copyBattle = true
	return true
}

func (td *TeamObject) TeamCopyStopStartBattle() bool {
	if !td.copyBattle {
		return false
	}
	td.copyBattle = false
	return true
}

func (td *TeamObject) IsBattling() bool {
	return td.isBattling
}

func (td *TeamObject) SetBattling() {
	if td.isBattling {
		return
	}
	td.isBattling = true
	return
}

func (td *TeamObject) GetRushTime() int64 {
	return td.rushTime
}

func (td *TeamObject) IsRushTimeInCd() bool {
	now := global.GetGame().GetTimeService().Now()
	if td.rushTime == 0 {
		return false
	}
	rushCdTime3v3 := constant.GetConstantService().GetConstant(constanttypes.ConstantType3V3RushCdTime) - 3000
	return (now - td.rushTime) < int64(rushCdTime3v3)
}

func (td *TeamObject) SetRushTime() {
	now := global.GetGame().GetTimeService().Now()
	td.rushTime = now
}

func (td *TeamObject) SetKickTime(playerId int64) {
	now := global.GetGame().GetTimeService().Now()
	td.kickTimeMap[playerId] = now
}

func (td *TeamObject) GetKickTimeInCd() (playerIdList []int64) {
	now := global.GetGame().GetTimeService().Now()
	for playerId, kickTime := range td.kickTimeMap {
		diffTime := now - kickTime
		if diffTime >= 30*1000 {
			delete(td.kickTimeMap, playerId)
		}
		playerIdList = append(playerIdList, playerId)
	}
	return
}

func CreateTeamObject(teamId int64, memberList []*TeamMemberObject, purpose teamtypes.TeamPurposeType) *TeamObject {
	data := &TeamObject{
		teamId:       teamId,
		memberList:   memberList,
		applyDataMap: make(map[int64]*TeamApplyData),
		autoReview:   true,
		purpose:      purpose,
		kickTimeMap:  make(map[int64]int64),
	}
	data.UpdateForce()
	return data
}

func NewTeamObject(teamId int64, force int64, memberList []*TeamMemberObject) *TeamObject {
	data := &TeamObject{
		teamId:     teamId,
		force:      force,
		memberList: memberList,
	}
	return data
}

//组队列表排序
type TeamObjectList []*TeamObject

func (tdl TeamObjectList) Len() int {
	return len(tdl)
}

func (tdl TeamObjectList) Less(i, j int) bool {
	if tdl[i].force < tdl[j].force {
		return true
	}
	if len(tdl[i].memberList) < len(tdl[j].memberList) {
		return true
	}
	return false
}

func (tdl TeamObjectList) Swap(i, j int) {
	tdl[i], tdl[j] = tdl[j], tdl[i]
}
