package team

import (
	"fgame/fgame/core/utils"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/funcopen/funcopen"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	teamtypes "fgame/fgame/game/team/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"sort"
	"sync"
)

//组队接口处理
type TeamService interface {
	//获取玩家队伍列表
	GetTeam(teamId int64) (teamData *TeamObject)
	//获取队伍根据玩家
	GetTeamByPlayerId(playerId int64) (teamData *TeamObject)
	//玩家点击创建队伍
	CreateTeamByPlayer(pl player.Player, purpose teamtypes.TeamPurposeType) (teamData *TeamObject, err error)
	//邀请玩家
	InvitePlayer(pl player.Player, invitedId int64) (err error)
	//被邀请玩家决策
	InvitedPlayerChoose(pl player.Player, typ teamtypes.TeamInviteType, result teamtypes.TeamResultType, teamId int64) (teamData *TeamObject, err error)
	//队长对申请决策
	CaptainApplyChoose(pl player.Player, applyPlayer player.Player, result teamtypes.TeamResultType) (teamData *TeamObject, err error)
	//申请加入附近队伍
	JoinNearTeam(pl player.Player, teamId int64) (flag bool, err error)
	//获取附近队伍
	GetNearTeam(pl player.Player) (teamList []*TeamObject)
	//一键申请
	TeamApplyAll(pl player.Player) (err error)
	//清空列表
	ClearApplyList(pl player.Player) (err error)
	//获取申请列表
	GetApplyList(pl player.Player) (applyList []*TeamApplyData, err error)
	//转让队长
	TransferCaptain(pl player.Player, captainId int64) (teamData *TeamObject, err error)
	//离队
	LeaveTeam(pl player.Player) (teamData *TeamObject)
	//请离队伍
	BeLeavedTeam(pl player.Player, leavedId int64) (teamData *TeamObject, leaveMember *TeamMemberObject, err error)
	//修改组队自动审核
	AutoReviewChoose(pl player.Player, autoReview bool) (err error)
	//获取对应队伍标识的房间信息
	GetTeamsByPurpose(pl player.Player, purpose teamtypes.TeamPurposeType) (teamList []*TeamObject)
	//改变队伍标识
	TeamChangePurpose(pl player.Player, purpose teamtypes.TeamPurposeType) (teamData *TeamObject, err error)

	//更新战力
	UpdateMemberForce(playerId int64, force int64) (teamData *TeamObject)
	//玩家下线
	PlayerLogout(pl player.Player) (leaveStatus teamtypes.TeamLeaveStatusType)

	//一键邀请
	TeamInviteAll(pl player.Player, playerList []player.Player) (err error)

	//3v3匹配
	ArenaMatch(pl player.Player) (canMatch bool, err error)
	//3v3 停止匹配
	ArenaStopMatch(pl player.Player) (err error)
	//3v3匹配
	ArenaMatched(pl player.Player) (err error)
	//3v3匹配失败
	ArenaMatchFailed(pl player.Player) (err error)
	// ArenaMatchEnd(pl scene.Player) (err error)

	//组队副本
	TeamCopyStartBattle(pl player.Player) (err error)
	//开始战斗失败
	TeamCopyFailed(pl player.Player) (err error)
	//开始战斗成功
	TeamCopySucess(pl player.Player) (err error)

	TeamMatchCondtionFailedDeal(pl player.Player, result bool, memberIdList []int64) (err error)
	TeamMatchCondtionPrepareDeal(pl player.Player, result bool) (err error)
	TeamMatchRush(pl player.Player) (err error)

	//TODO zrc:添加玩家组队战斗中和退出战斗
}

type teamService struct {
	rwm sync.RWMutex
	//组队map信息
	teamMap map[int64]*TeamObject
	//玩家队伍
	playerTeamMap map[int64]*TeamObject
	//申请列表map
	applyMap map[int64]map[int64]*TeamApplyData
}

//初始化
func (ts *teamService) init() error {
	ts.teamMap = make(map[int64]*TeamObject)
	ts.playerTeamMap = make(map[int64]*TeamObject)
	ts.applyMap = make(map[int64]map[int64]*TeamApplyData)
	return nil
}

//获取玩家队伍
func (ts *teamService) GetTeam(teamId int64) (teamData *TeamObject) {
	ts.rwm.RLock()
	defer ts.rwm.RUnlock()
	return ts.getTeam(teamId)
}

//获取队伍
func (ts *teamService) getTeam(teamId int64) (t *TeamObject) {
	t, exist := ts.teamMap[teamId]
	if !exist {
		return nil
	}
	return
}

//获取玩家队伍
func (ts *teamService) GetTeamByPlayerId(playerId int64) (teamData *TeamObject) {
	ts.rwm.RLock()
	defer ts.rwm.RUnlock()
	return ts.getTeamByPlayerId(playerId)
}

//获取队伍根据玩家id
func (ts *teamService) getTeamByPlayerId(playerId int64) (t *TeamObject) {
	t, exist := ts.playerTeamMap[playerId]
	if !exist {
		return nil
	}
	return
}

func memberObjFromPlayer(pl player.Player) *TeamMemberObject {
	playerId := pl.GetId()
	force := pl.GetForce()
	name := pl.GetName()
	level := pl.GetLevel()
	role := pl.GetRole()
	sex := pl.GetSex()
	fashionId := pl.GetFashionId()
	zhuanSheng := pl.GetZhuanSheng()
	serverId := pl.GetServerId()
	memberObj := NewTeamMemberObject(
		serverId,
		playerId,
		force,
		true,
		name,
		level,
		role,
		sex,
		fashionId,
		zhuanSheng)
	return memberObj
}

//玩家创建队伍map处理
func (ts *teamService) playerCreateTeam(playerId int64, teamData *TeamObject) {
	ts.playerTeamMap[playerId] = teamData
	ts.teamMap[teamData.GetTeamId()] = teamData
}

//玩家邀请他人组队成功map处理
func (ts *teamService) playerInviteCreateTeam(inviteId int64, invitedId int64, teamData *TeamObject) {
	ts.playerTeamMap[inviteId] = teamData
	ts.playerTeamMap[invitedId] = teamData
	ts.teamMap[teamData.GetTeamId()] = teamData
}

//玩家加入队伍
func (ts *teamService) playerJoinTeam(teamData *TeamObject, pl player.Player) {
	member := memberObjFromPlayer(pl)
	teamData.AddMember(member)
	ts.playerTeamMap[pl.GetId()] = teamData
}

//解散队伍
func (ts *teamService) dissolveTeam(teamData *TeamObject) {
	teamId := teamData.GetTeamId()
	defer delete(ts.teamMap, teamId)
	for _, memberObj := range teamData.GetMemberList() {
		delete(ts.playerTeamMap, memberObj.GetPlayerId())
	}
}

//玩家点击创建队伍(含创建房间)
func (ts *teamService) CreateTeamByPlayer(pl player.Player, purpose teamtypes.TeamPurposeType) (teamData *TeamObject, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData = ts.getTeamByPlayerId(playerId)
	if teamData != nil {
		if purpose == teamtypes.TeamPurposeTypeNormal {
			err = ErrorTeamPlayerInTeam
		} else {
			err = ErrorTeamCreateHouseInOther
		}
		return
	}
	teamId, _ := idutil.GetId()

	memberList := make([]*TeamMemberObject, 0, teamtypes.TeamMaxNum)
	member := memberObjFromPlayer(pl)
	memberList = append(memberList, member)
	teamData = CreateTeamObject(teamId, memberList, purpose)

	ts.playerCreateTeam(playerId, teamData)
	//TODO 发送创建事件
	return
}

//TODO 优化
func (ts *teamService) getNearTeam(pl player.Player) (teamList []*TeamObject) {
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	teamId := int64(0)
	if teamObj != nil {
		teamId = teamObj.GetTeamId()
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	for playerId, _ := range s.GetAllPlayers() {
		teamData := ts.getTeamByPlayerId(playerId)
		//未组队
		if teamData == nil {
			continue
		}
		//自己队伍
		if teamData.GetTeamId() == teamId {
			continue
		}
		_, pos := teamData.GetMember(playerId)
		//非队长
		if pos != 0 {
			continue
		}
		num := teamData.GetNum()
		//已满员
		if num == teamtypes.TeamMaxNum {
			continue
		}
		//玩家功能是否开启
		if !ts.isTeamCopyOpen(teamData, pl) {
			continue
		}
		teamList = append(teamList, teamData)
	}

	if len(teamList) > 1 {
		sort.Sort(sort.Reverse(TeamObjectList(teamList)))
	}
	return
}

//获取附近队伍
func (ts *teamService) GetNearTeam(pl player.Player) (teamList []*TeamObject) {
	ts.rwm.RLock()
	defer ts.rwm.RUnlock()
	return ts.getNearTeam(pl)
}

//TODO 优化
func (ts *teamService) getPurposeTeams(pl player.Player, purpose teamtypes.TeamPurposeType) (teamList []*TeamObject) {
	teamId := pl.GetTeamId()

	for _, teamData := range ts.teamMap {
		if !teamData.Equal(purpose) {
			continue
		}
		if teamData.teamId == teamId {
			continue
		}
		if teamData.IsFull() {
			continue
		}
		if teamData.IsBattling() {
			continue
		}
		hasCross := false
		for _, member := range teamData.GetMemberList() {
			//验证所有成员是否在副本中或跨服中
			spl := player.GetOnlinePlayerManager().GetPlayerById(member.playerId)
			if spl == nil {
				continue
			}

			//验证玩家是否处于跨服
			if spl.IsCross() {
				hasCross = true
				break
			}
		}
		if hasCross {
			continue
		}

		teamList = append(teamList, teamData)
	}
	if len(teamList) > 1 {
		sort.Sort(sort.Reverse(TeamObjectList(teamList)))
	}
	if len(teamList) > teamtypes.TeamPurposeMax {
		teamList = teamList[:teamtypes.TeamPurposeMax]
	}
	return
}

//获取队伍标识为purpose的所有队伍
func (ts *teamService) GetTeamsByPurpose(pl player.Player, purpose teamtypes.TeamPurposeType) (teamList []*TeamObject) {
	ts.rwm.RLock()
	defer ts.rwm.RUnlock()
	return ts.getPurposeTeams(pl, purpose)
}

//邀请玩家
func (ts *teamService) InvitePlayer(pl player.Player, invitedId int64) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	//被邀请者已组队状态
	inviteTeamObject := ts.getTeamByPlayerId(invitedId)
	if inviteTeamObject != nil {
		err = ErrorTeamPlayerInTeam
		return
	}
	selfTeamObject := ts.getTeamByPlayerId(playerId)
	if selfTeamObject != nil {
		//队伍已满员
		num := selfTeamObject.GetNum()
		if num == teamtypes.TeamMaxNum {
			err = ErrorTeamPlayerFull
			return
		}

		//3v3 正在匹配
		if selfTeamObject.IsMatch() {
			err = ErrorTeamInMatchInviteJion
			return
		}

		//正在组队副本
		if selfTeamObject.IsCopyBattle() {
			err = ErrorTeamInTeamCopyBattle
			return
		}

		if selfTeamObject.IsBattling() {
			err = ErrorTeamHouseIsBatting
			return
		}

		//队伍邀请玩家
		gameevent.Emit(teameventtypes.EventTypeTeamTeamInvitePlayer, selfTeamObject, invitedId)
		return
	}
	//玩家邀请玩家
	gameevent.Emit(teameventtypes.EventTypeTeamPlayerInvitePlayer, pl, invitedId)
	return
}

//一键邀请
func (ts *teamService) TeamInviteAll(pl player.Player, playerList []player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()

	teamData := ts.getTeamByPlayerId(pl.GetId())
	//队伍已满员
	if teamData != nil && teamData.GetNum() == teamtypes.TeamMaxNum {
		err = ErrorTeamPlayerFull
		return
	}

	for _, spl := range playerList {
		if spl == nil {
			continue
		}

		if teamData != nil {
			//对方功能是否开启
			if !ts.isTeamCopyOpen(teamData, pl) {
				continue
			}
			//队伍邀请玩家
			gameevent.Emit(teameventtypes.EventTypeTeamTeamInvitePlayer, teamData, spl.GetId())
		} else {
			//玩家邀请玩家组队
			gameevent.Emit(teameventtypes.EventTypeTeamPlayerInvitePlayer, pl, spl.GetId())
		}
	}
	return
}

//邀请创建队伍
func (ts *teamService) invitedPlayerChooseCreate(pl player.Player, result teamtypes.TeamResultType, invitePlayerId int64) (teamData *TeamObject, err error) {
	invitePlayer := player.GetOnlinePlayerManager().GetPlayerById(invitePlayerId)
	//邀请者已下线
	if invitePlayer == nil {
		//提示玩家已下线
		err = ErrorTeamPlayerOff
		return
	}
	memberData := memberObjFromPlayer(pl)
	inviteTeamObject := ts.getTeamByPlayerId(invitePlayerId)

	//TODO 改成邀请加入队伍
	if inviteTeamObject != nil {
		teamData, err = ts.invitedPlayerChooseJoin(pl, result, inviteTeamObject.GetTeamId())
		if err != nil {
			return
		}
		return
	}

	//玩家拒绝
	if result == teamtypes.TeamResultTypeNo {
		eventData := CreateTeamPlayerInviteDealCreateEventData(nil, invitePlayerId, result)
		gameevent.Emit(teameventtypes.EventTypeTeamPlayerInvitePlayerDealCreate, pl, eventData)
		return
	}

	teamId, _ := idutil.GetId()

	inviteMember := memberObjFromPlayer(invitePlayer)
	memberList := make([]*TeamMemberObject, 0, teamtypes.TeamMaxNum)
	memberList = append(memberList, inviteMember, memberData)
	teamData = CreateTeamObject(teamId, memberList, teamtypes.TeamPurposeTypeNormal)
	ts.playerInviteCreateTeam(invitePlayerId, memberData.GetPlayerId(), teamData)

	eventData := CreateTeamPlayerInviteDealCreateEventData(teamData, invitePlayerId, result)
	//发送玩家同意组队创建
	gameevent.Emit(teameventtypes.EventTypeTeamPlayerInvitePlayerDealCreate, pl, eventData)
	return
}

//邀请加入
func (ts *teamService) invitedPlayerChooseJoin(pl player.Player, result teamtypes.TeamResultType, inviteTeamId int64) (teamData *TeamObject, err error) {
	teamData = ts.getTeam(inviteTeamId)
	if teamData == nil {
		//队伍已解散
		err = ErrorTeamDissolve
		return
	}

	if result == teamtypes.TeamResultTypeOk {
		if teamData.IsMatch() {
			err = ErrorTeamInMatchDealAgreeFail
			return
		}

		num := int32(len(teamData.GetMemberList()))
		if num == teamtypes.TeamMaxNum {
			//队伍已满员
			err = ErrorTeamPlayerFull
			return
		}

		ts.playerJoinTeam(teamData, pl)
	}

	eventData := CreateTeamPlayerInviteDealJoinEventData(teamData, result)
	gameevent.Emit(teameventtypes.EventTypeTeamPlayerInvitePlayerDealJoin, pl, eventData)
	return
}

//被邀请玩家决策
func (ts *teamService) InvitedPlayerChoose(pl player.Player, typ teamtypes.TeamInviteType, result teamtypes.TeamResultType, id int64) (teamData *TeamObject, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	switch typ {
	case teamtypes.TeamInviteTypeCreate:
		return ts.invitedPlayerChooseCreate(pl, result, id)
	case teamtypes.TeamInviteTypeJoin:
		return ts.invitedPlayerChooseJoin(pl, result, id)
	}
	return
}

//离队
func (ts *teamService) LeaveTeam(pl player.Player) (teamData *TeamObject) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData = ts.getTeamByPlayerId(playerId)
	if teamData == nil {
		return
	}

	//停止匹配
	ts.teamStopMatch(pl, teamData)
	//停止开始战斗
	ts.teamCopyStopStartBattle(pl, teamData)

	_, pos := teamData.GetMember(playerId)
	num := teamData.GetNum()

	ts.removeTeamMember(playerId)
	eventData := CreateTeamPlayerLeaveEventData(teamData, pos)
	gameevent.Emit(teameventtypes.EventTypeTeamPlayerLeave, pl, eventData)

	offLen := teamData.GetOfflineNum()
	if num == offLen+1 {
		//解散组队
		ts.dissolveTeam(teamData)
		return
	}
	return
}

//转让队长
func (ts *teamService) TransferCaptain(pl player.Player, captainId int64) (teamData *TeamObject, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData = ts.getTeamByPlayerId(playerId)
	if teamData == nil {
		err = ErrorTeamPlayerNotInTeam
		return
	}
	//正在匹配中
	if teamData.IsMatch() {
		err = ErrorTeamInMatch
		return
	}

	_, pos := teamData.GetMember(playerId)
	if pos != 0 {
		err = ErrorTeamCaptainIsOther
		return
	}

	_, captainPos := teamData.GetMember(captainId)
	if captainPos == -1 {
		err = ErrorTeamPlayerNoMember
		return
	}
	flag := teamData.TransferCaptain(captainPos)
	if !flag {
		panic(fmt.Errorf("team:转移队长应该成功"))
	}
	gameevent.Emit(teameventtypes.EventTypeTeamCaptainTransfer, pl, teamData)

	return
}

//请离队伍
func (ts *teamService) BeLeavedTeam(pl player.Player, leavedId int64) (teamData *TeamObject, leaveMember *TeamMemberObject, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData = ts.getTeamByPlayerId(playerId)
	if teamData == nil {
		err = ErrorTeamPlayerNotInTeam
		return
	}
	_, pos := teamData.GetMember(playerId)
	if pos != 0 {
		err = ErrorTeamCaptainIsOther
		return
	}

	leaveMember, leavePos := teamData.GetMember(leavedId)
	if leavePos == -1 {
		err = ErrorTeamPlayerNoMember
		return
	}

	//停止匹配
	ts.teamStopMatch(pl, teamData)
	//TODO
	ts.removeTeamMember(leavedId)
	//被请离队事件
	gameevent.Emit(teameventtypes.EventTypeTeamPlayerBeLeaved, teamData, leaveMember)
	return
}

//申请加入附近队伍
func (ts *teamService) JoinNearTeam(pl player.Player, teamId int64) (flag bool, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	playerTeam := ts.getTeamByPlayerId(playerId)
	//玩家还在队伍中
	if playerTeam != nil {
		err = ErrorTeamPlayerInTeam
		return
	}

	teamData := ts.getTeam(teamId)
	if teamData == nil {
		//队伍已解散
		err = ErrorTeamDissolve
		return
	}
	num := teamData.GetNum()
	if num == teamtypes.TeamMaxNum {
		//队伍已满员
		err = ErrorTeamPlayerFull
		return
	}

	playerIdList := teamData.GetKickTimeInCd()
	if utils.ContainInt64(playerIdList, pl.GetId()) {
		err = ErrorTeamJionByLeavedInCd
		return
	}

	//判断3v3 是否在匹配
	if teamData.IsMatch() {
		err = ErrorTeamInMatchJionFail
		return
	}

	//判断队伍是否在组队副本
	if teamData.IsCopyBattle() {
		err = ErrorTeamInTeamCopyJionFail
		return
	}

	if teamData.IsBattling() {
		err = ErrorTeamHouseIsBatting
		return
	}

	//功能是否开启
	if !ts.isTeamCopyOpen(teamData, pl) {
		err = ErrorTeamApplyJoinFuncNoOpen
		return
	}

	//是否是自动审核
	if ts.isAutoReview(teamData, pl) {
		flag = true
		return
	}
	teamData.Apply(playerId)
	//发送事件
	gameevent.Emit(teameventtypes.EventTypeTeamNearApplyJoin, pl, teamData)

	return
}

//队长对玩家申请决策
func (ts *teamService) CaptainApplyChoose(pl player.Player, applyPlayer player.Player, result teamtypes.TeamResultType) (teamData *TeamObject, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData = ts.getTeamByPlayerId(playerId)
	if teamData == nil {
		err = ErrorTeamPlayerNotInTeam
	}

	_, pos := teamData.GetMember(playerId)
	//你不是队长
	if pos != 0 {
		err = ErrorTeamCaptainIsOther
		return
	}

	//正在匹配中
	if teamData.IsMatch() {
		err = ErrorTeamInMatch
		return
	}

	applyPlayerId := applyPlayer.GetId()
	apply := teamData.GetApply(applyPlayerId)
	//从申请列表移除
	defer teamData.RemoveApply(applyPlayerId)

	if apply == nil {
		err = ErrorTeamApplyNoExist
		return
	}
	if result == teamtypes.TeamResultTypeOk {
		applyTeamObject := ts.getTeamByPlayerId(applyPlayerId)
		//当前已有队伍,无法加入他人队伍
		if applyTeamObject != nil {
			err = ErrorTeamPlayerInTeam
			return
		}

		len := teamData.GetNum()
		if len == teamtypes.TeamMaxNum {
			//当前队伍人数已达上限
			err = ErrorTeamPlayerFull
			return
		}
		ts.playerJoinTeam(teamData, applyPlayer)
	}

	eventData := teameventtypes.CreateTeamApplyDealEventData(applyPlayerId, result)
	//发送处理成功事件
	gameevent.Emit(teameventtypes.EventTypeTeamApplyDeal, teamData, eventData)
	return
}

//更新战力
func (ts *teamService) UpdateMemberForce(playerId int64, force int64) (teamData *TeamObject) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	teamData = ts.getTeamByPlayerId(playerId)
	if teamData == nil {
		return
	}
	teamData.UpdateMemberForce(playerId, force)
	return
}

//玩家下线
func (ts *teamService) PlayerLogout(pl player.Player) (leaveStatus teamtypes.TeamLeaveStatusType) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData := ts.getTeamByPlayerId(playerId)

	if teamData == nil {
		leaveStatus = teamtypes.TeamLeaveStatusTypeDissolve
		return
	}

	ts.teamLogoutStopMatch(pl, teamData)

	offLen := teamData.GetOfflineNum()
	len := teamData.GetNum()

	if len == (offLen + 1) {
		leaveStatus = teamtypes.TeamLeaveStatusTypeDissolve
		//解散队伍
		ts.dissolveTeam(teamData)
	} else {
		memberObj, _ := teamData.GetMember(playerId)
		memberObj.online = false
	}

	return
}

//一键申请
func (ts *teamService) TeamApplyAll(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData := ts.getTeamByPlayerId(playerId)
	if teamData != nil {
		err = ErrorTeamPlayerInTeam
		return
	}

	teamList := ts.getNearTeam(pl)

	for _, teamData := range teamList {
		if !ts.isTeamCopyOpen(teamData, pl) {
			continue
		}

		playerIdList := teamData.GetKickTimeInCd()
		if utils.ContainInt64(playerIdList, pl.GetId()) {
			continue
		}

		//自动审核
		if ts.isAutoReview(teamData, pl) {
			continue
		}
		teamData.Apply(playerId)
		//发送事件
		gameevent.Emit(teameventtypes.EventTypeTeamNearApplyJoin, pl, teamData)
	}
	return
}

//获取申请列表
func (ts *teamService) GetApplyList(pl player.Player) (applyList []*TeamApplyData, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData := ts.getTeamByPlayerId(playerId)
	if teamData == nil {
		err = ErrorTeamPlayerNotInTeam
		return
	}
	_, pos := teamData.GetMember(playerId)
	if pos != 0 {
		err = ErrorTeamCaptainIsOther
		return
	}
	applyList = teamData.GetAllApplyList()
	return
}

//清空列表
func (ts *teamService) ClearApplyList(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamData := ts.getTeamByPlayerId(playerId)
	if teamData == nil {
		err = ErrorTeamPlayerNotInTeam
		return
	}
	_, pos := teamData.GetMember(playerId)
	if pos != 0 {
		err = ErrorTeamCaptainIsOther
		return
	}
	teamData.ClearAllApply()

	return
}

func (ts *teamService) addTeamMember(teamData *TeamObject, member *TeamMemberObject) {
	teamData.AddMember(member)
	ts.playerTeamMap[member.GetPlayerId()] = teamData
}

func (ts *teamService) removeTeamMember(memberId int64) {
	teamData := ts.getTeamByPlayerId(memberId)
	if teamData == nil {
		return
	}
	teamData.RemoveMember(memberId)
	delete(ts.playerTeamMap, memberId)
}

//竞技场匹配
func (ts *teamService) ArenaMatch(pl player.Player) (canMatch bool, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamPlayerNotInTeam
		return
	}
	captainMember := teamObj.GetCaptain()
	//不是队长
	if captainMember.GetPlayerId() != playerId {
		err = ErrorTeamCaptainIsOther
		return
	}

	//匹配条件
	memberIdList, err := ts.arenaMatchCondtion(teamObj, pl)
	if err != nil {
		return
	}

	//匹配条件不足
	if len(memberIdList) != 0 {
		gameevent.Emit(teameventtypes.EventTypeTeamMatchNoEough, teamObj, memberIdList)
		return
	}

	//判断队伍是不是正在匹配
	flag := teamObj.Matching()
	if !flag {
		err = ErrorTeamInMatch
		return
	}
	gameevent.Emit(teameventtypes.EventTypeTeamArenaMatch, pl, teamObj)
	canMatch = true
	return
}

//3v3条件判断
func (ts *teamService) arenaMatchCondtion(teamObj *TeamObject, pl player.Player) (memberIdList []int64, err error) {
	funcOpenTemplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(funcopentypes.FuncOpenTypeArena)
	if funcOpenTemplate == nil {
		return
	}
	return ts.memberConditionArean(teamObj, funcOpenTemplate, pl)
}

//3v3
func (ts *teamService) memberConditionArean(teamObj *TeamObject, funcOpenTemplate *gametemplate.ModuleOpenedTemplate, pl player.Player) (memberIdList []int64, err error) {
	if teamObj.IsMatch() {
		err = ErrorTeamInMatch
		return
	}

	if teamObj.IsCopyBattle() {
		err = ErrorTeamInTeamCopyBattle
		return
	}

	if teamObj.IsBattling() {
		err = ErrorTeamHouseIsBatting
		return
	}

	isReturn, err := ts.selfCondition(pl, funcOpenTemplate)
	if isReturn {
		return nil, err
	}

	for _, member := range teamObj.GetMemberList() {
		//验证所有成员是否在副本中或跨服中
		spl := player.GetOnlinePlayerManager().GetPlayerById(member.playerId)
		if spl == nil {
			err = ErrorTeamMemberPlayerOffline
			return
		}
		if pl.GetId() == member.GetPlayerId() {
			continue
		}

		//验证所有成员等级
		if member.level < funcOpenTemplate.OpenedLevel {
			err = ErrorTeamInMatchLevelLow
			return
		}

		//验证玩家是否处于跨服
		if spl.IsCross() {
			//err = ErrorTeamInMatchInCross
			//return
			memberIdList = append(memberIdList, member.GetPlayerId())
			continue
		}

		//是否在无间炼狱排队
		if spl.IsLianYuLineUp() {
			//err = ErrorTeamInMatchInLianYuLineUp
			//return
			memberIdList = append(memberIdList, member.GetPlayerId())
			continue
		}

		//是否在神兽攻城排队
		if spl.IsGodSiegeLineUp() {
			//err = ErrorTeamInMatchInGodSiegeLineUp
			//return
			memberIdList = append(memberIdList, member.GetPlayerId())
			continue
		}

		//神魔战场排队
		if spl.IsShenMoLineUp() {
			memberIdList = append(memberIdList, member.GetPlayerId())
			continue
		}

		//验证玩家是否处于副本
		pls := spl.GetScene()
		if pls == nil {
			memberIdList = append(memberIdList, member.GetPlayerId())
			continue
		}

		if pls.MapTemplate().IsFuBen() {
			// err = ErrorTeamInMatchInFuBen
			// return
			memberIdList = append(memberIdList, member.GetPlayerId())
			continue
		}

		//验证玩家是否处于战斗
		if spl.IsPvpBattle() {
			err = ErrorTeamInMatchInPvp
			return
		}
	}
	return
}

func (ts *teamService) selfCondition(pl player.Player, funcOpenTemplate *gametemplate.ModuleOpenedTemplate) (isReturn bool, err error) {
	//验证玩家是否处于跨服
	if pl.IsCross() {
		err = ErrorTeamInMatchSelfInCross
		isReturn = true
		return
	}

	//是否在无间炼狱排队
	if pl.IsLianYuLineUp() {
		err = ErrorTeamInMatchSelfInLianYuLineUp
		isReturn = true
		return
	}

	//是否在神兽攻城排队
	if pl.IsGodSiegeLineUp() {
		err = ErrorTeamInMatchInSelfGodSiegeLineUp
		isReturn = true
		return
	}

	//是否在神魔战场排队
	if pl.IsShenMoLineUp() {
		err = ErrorTeamInMatchInSelfShenMoLineUp
		isReturn = true
		return
	}

	//验证玩家是否处于副本
	pls := pl.GetScene()
	if pls.MapTemplate().IsFuBen() {
		err = ErrorTeamInMatchSelfInFuBen
		isReturn = true
		return
	}

	//验证玩家是否处于战斗
	if pl.IsPvpBattle() {
		err = ErrorTeamInMatchSelfInPvp
		isReturn = true
		return
	}
	return
}

func (ts *teamService) memberCondition(teamObj *TeamObject, funcOpenTemplate *gametemplate.ModuleOpenedTemplate) (err error) {
	if teamObj.IsMatch() {
		err = ErrorTeamInMatch
		return
	}

	if teamObj.IsCopyBattle() {
		err = ErrorTeamInTeamCopyBattle
		return
	}

	if teamObj.IsBattling() {
		err = ErrorTeamHouseIsBatting
		return
	}
	for _, member := range teamObj.GetMemberList() {
		//验证所有成员是否在副本中或跨服中
		spl := player.GetOnlinePlayerManager().GetPlayerById(member.playerId)
		if spl == nil {
			err = ErrorTeamMemberPlayerOffline
			return
		}
		//验证所有成员等级
		if member.level < funcOpenTemplate.OpenedLevel {
			err = ErrorTeamInMatchLevelLow
			return
		}

		//验证玩家是否处于跨服
		if spl.IsCross() {
			err = ErrorTeamInMatchInCross
			return
		}

		//是否在无间炼狱排队
		if spl.IsLianYuLineUp() {
			err = ErrorTeamInMatchInLianYuLineUp
			return
		}

		//是否在神兽攻城排队
		if spl.IsGodSiegeLineUp() {
			err = ErrorTeamInMatchInGodSiegeLineUp
			return
		}

		if spl.IsShenMoLineUp() {
			err = ErrorTeamInMatchInSelfShenMoLineUp
			return
		}

		//验证玩家是否处于副本
		pls := spl.GetScene()
		if pls != nil && pls.MapTemplate().IsFuBen() {
			err = ErrorTeamInMatchInFuBen
			return
		}

		//验证玩家是否处于战斗
		if spl.IsPvpBattle() {
			err = ErrorTeamInMatchInPvp
			return
		}
	}
	return
}

//竞技场停止匹配
func (ts *teamService) ArenaStopMatch(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamPlayerNotInTeam
		return
	}
	captainMember := teamObj.GetCaptain()
	//不是队长
	if captainMember.GetPlayerId() != playerId {
		err = ErrorTeamCaptainIsOther
		return
	}
	flag := ts.teamStopMatch(pl, teamObj)
	if !flag {
		err = ErrorTeamNotInMatch
		return
	}

	return
}

func (ts *teamService) teamStopMatch(pl player.Player, teamObj *TeamObject) bool {
	//判断队伍是不是正在匹配
	flag := teamObj.StopMatching()
	if !flag {
		return false
	}
	gameevent.Emit(teameventtypes.EventTypeTeamArenaStopMatch, pl, teamObj)
	return true
}

func (ts *teamService) teamLogoutStopMatch(pl player.Player, teamObj *TeamObject) bool {
	//判断队伍是不是正在匹配
	flag := teamObj.StopMatching()
	if !flag {
		return false
	}
	gameevent.Emit(teameventtypes.EventTypeTeamArenaStopMatchOther, pl, teamObj)
	return true
}

//竞技场匹配到了
func (ts *teamService) ArenaMatched(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamNoExist
		return
	}

	//TODO 细化错误
	flag := teamObj.StopMatching()
	if !flag {
		err = ErrorTeamNotInMatch
		return
	}

	gameevent.Emit(teameventtypes.EventTypeTeamArenaMatched, pl, teamObj)
	return
}

//竞技场匹配失败
func (ts *teamService) ArenaMatchFailed(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamNoExist
		return
	}

	flag := teamObj.StopMatching()
	if !flag {
		err = ErrorTeamNotInMatch
		return
	}

	gameevent.Emit(teameventtypes.EventTypeTeamArenaMatchFailed, pl, teamObj)
	return
}

func (ts *teamService) isTeamCopyOpen(teamObj *TeamObject, pl player.Player) (flag bool) {
	teamPurpose := teamObj.GetTeamPurpose()
	if !teamPurpose.Vaild() {
		flag = true
		return
	}
	openType := teamPurpose.GetFuncOpenType()
	funcOpenTeamplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(openType)
	if funcOpenTeamplate == nil {
		return
	}
	if pl.GetLevel() < funcOpenTeamplate.OpenedLevel {
		return
	}
	flag = true
	return
}

func (ts *teamService) isAutoReview(teamData *TeamObject, pl player.Player) (flag bool) {
	if !teamData.IsAutoReview() {
		return
	}
	if teamData.IsMatch() {
		return
	}
	num := teamData.GetNum()
	//已满员
	if num == teamtypes.TeamMaxNum {
		return
	}
	ts.playerJoinTeam(teamData, pl)
	eventData := teameventtypes.CreateTeamApplyDealEventData(pl.GetId(), teamtypes.TeamResultTypeOk)
	//发送处理成功事件
	gameevent.Emit(teameventtypes.EventTypeTeamApplyDeal, teamData, eventData)
	flag = true
	return
}

//修改组队自动审核
func (ts *teamService) AutoReviewChoose(pl player.Player, autoReview bool) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()

	teamData := ts.getTeamByPlayerId(pl.GetId())
	if teamData == nil {
		err = ErrorTeamNoExist
		return
	}

	_, pos := teamData.GetMember(pl.GetId())
	//你不是队长
	if pos != 0 {
		err = ErrorTeamCaptainIsOther
		return
	}

	teamData.autoReview = autoReview
	return
}

//改变队伍标识
func (ts *teamService) TeamChangePurpose(pl player.Player, purpose teamtypes.TeamPurposeType) (teamData *TeamObject, err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()

	openType := purpose.GetFuncOpenType()
	funcOpenTemplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(openType)
	if funcOpenTemplate == nil {
		return
	}
	teamData = ts.getTeamByPlayerId(pl.GetId())
	if teamData == nil {
		if pl.GetLevel() < funcOpenTemplate.OpenedLevel {
			err = ErrorTeamCreateHouseSelfLevelLow
			return
		}
		teamId, _ := idutil.GetId()
		memberList := make([]*TeamMemberObject, 0, teamtypes.TeamMaxNum)
		member := memberObjFromPlayer(pl)
		memberList = append(memberList, member)
		teamData = CreateTeamObject(teamId, memberList, purpose)
		ts.playerCreateTeam(pl.GetId(), teamData)
	} else {
		_, pos := teamData.GetMember(pl.GetId())
		//非队长
		if pos != 0 {
			err = ErrorTeamCreateHouseInOther
			return
		}
		if teamData.Equal(purpose) {
			return
		}

		for _, mem := range teamData.GetMemberList() {
			if mem.GetLevel() < funcOpenTemplate.OpenedLevel {
				err = ErrorTeamCreateHouseLevelLow
				return
			}
		}
		teamData.purpose = purpose
	}
	gameevent.Emit(teameventtypes.EventTypeTeamCaptainChangePurpose, pl, teamData)
	return
}

//组队副本
func (ts *teamService) TeamCopyStartBattle(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamPlayerNotInTeam
		return
	}
	captainMember := teamObj.GetCaptain()
	//不是队长
	if captainMember.GetPlayerId() != playerId {
		err = ErrorTeamCaptainIsOther
		return
	}

	teamPurose := teamObj.GetTeamPurpose()
	if !teamPurose.Vaild() {
		err = ErrorTeamBattlePuroseIsNormal
		return
	}

	//匹配条件
	err = ts.teamCopyStartCondtion(teamObj)
	if err != nil {
		return
	}

	flag := teamObj.TeamCopyStartBattling()
	if !flag {
		err = ErrorTeamInTeamCopyBattle
		return
	}
	gameevent.Emit(teameventtypes.EventTypeTeamCopyStartBattle, pl, teamObj)
	return
}

func (ts *teamService) getTeamCopyOpenType(teamObj *TeamObject) (funcOpenTeamplate *gametemplate.ModuleOpenedTemplate) {
	teamPurpose := teamObj.GetTeamPurpose()
	if !teamPurpose.Vaild() {
		return
	}
	openType := teamPurpose.GetFuncOpenType()
	funcOpenTeamplate = funcopen.GetFuncOpenService().GetFuncOpenTemplate(openType)
	if funcOpenTeamplate == nil {
		return
	}
	return
}

//组队副本条件判断
func (ts *teamService) teamCopyStartCondtion(teamObj *TeamObject) (err error) {
	funcOpenTemplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(funcopentypes.FuncOpenTypeTeamCopy)
	if funcOpenTemplate == nil {
		return
	}
	funcOpenTemplate = ts.getTeamCopyOpenType(teamObj)
	if funcOpenTemplate == nil {
		err = ErrorTeamInMatchLevelLow
		return
	}
	return ts.memberCondition(teamObj, funcOpenTemplate)
}

//组队副本开始战斗失败
func (ts *teamService) TeamCopyFailed(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamNoExist
		return
	}

	flag := teamObj.TeamCopyStopStartBattle()
	if !flag {
		err = ErrorTeamNotInTeamCopyBattle
		return
	}
	gameevent.Emit(teameventtypes.EventTypeTeamCopyStartBattleFailed, pl, teamObj)
	return
}

//组队副本开始战斗成功
func (ts *teamService) TeamCopySucess(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()
	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamNoExist
		return
	}

	flag := teamObj.IsCopyBattle()
	if !flag {
		err = ErrorTeamNotInTeamCopyBattle
		return
	}
	teamObj.TeamCopyStopStartBattle()
	//zrc: 临时注释掉
	// teamObj.SetBattling()
	gameevent.Emit(teameventtypes.EventTypeTeamCopyStartBattleSucess, pl, teamObj)
	return
}

func (ts *teamService) teamCopyStopStartBattle(pl player.Player, teamObj *TeamObject) bool {
	//判断队伍是不是正在战斗
	flag := teamObj.TeamCopyStopStartBattle()
	if !flag {
		return false
	}
	gameevent.Emit(teameventtypes.EventTypeTeamCopyStartBattleFailed, pl, teamObj)
	return true
}

func (ts *teamService) TeamMatchCondtionFailedDeal(pl player.Player, result bool, memberIdList []int64) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()

	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamNoExist
		return
	}

	//决策期间有队员离队
	var tempMemberIdList []int64
	for _, playerId := range memberIdList {
		_, pos := teamObj.GetMember(playerId)
		if pos == -1 {
			continue
		}
		tempMemberIdList = append(tempMemberIdList, playerId)
	}

	if len(tempMemberIdList) == 0 {
		return
	}
	eventData := teameventtypes.CreateTeamMatchCondtionFailedDealEventData(result, tempMemberIdList)
	gameevent.Emit(teameventtypes.EventTypeTeamMatchCondtionFailedDeal, teamObj, eventData)
	return
}

func (ts *teamService) TeamMatchCondtionPrepareDeal(pl player.Player, result bool) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()

	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamNoExist
		return
	}
	eventData := teameventtypes.CreateTeamMatchCondtionPrepareDealEventData(pl, result)
	gameevent.Emit(teameventtypes.EventTypeTeamMatchCondtionPrepareDeal, teamObj, eventData)
	return
}

func (ts *teamService) TeamMatchRush(pl player.Player) (err error) {
	ts.rwm.Lock()
	defer ts.rwm.Unlock()

	playerId := pl.GetId()
	teamObj := ts.getTeamByPlayerId(playerId)
	//队伍不存在
	if teamObj == nil {
		err = ErrorTeamNoExist
		return
	}

	captainId := teamObj.GetCaptain().GetPlayerId()

	captainPl := player.GetOnlinePlayerManager().GetPlayerById(captainId)
	if captainPl == nil {
		err = ErrorTeamRushMatchCaptainOffline
		return
	}

	flag := teamObj.IsRushTimeInCd()
	if flag {
		err = ErrorTeamRushMatchIsExist
		return
	}

	teamObj.SetRushTime()
	gameevent.Emit(teameventtypes.EventTypeTeamMatchRushStart, pl, captainPl)
	return
}

var (
	once sync.Once
	cs   *teamService
)

func Init() (err error) {
	once.Do(func() {
		cs = &teamService{}
		err = cs.init()
	})
	return err
}

func GetTeamService() TeamService {
	return cs
}
