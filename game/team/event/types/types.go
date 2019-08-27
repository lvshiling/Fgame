package types

import (
	"fgame/fgame/game/player"
	teamtypes "fgame/fgame/game/team/types"
)

const (
	//队伍邀请玩家
	EventTypeTeamTeamInvitePlayer = "TeamTeamInvitePlayer"
	//玩家邀请玩家组队
	EventTypeTeamPlayerInvitePlayer = "TeamPlayerInvitePlayer"
	//被邀请玩家决策玩家组队创建
	EventTypeTeamPlayerInvitePlayerDealCreate = "TeamPlayerInvitePlayerDealCreate"
	//被邀请玩家决策玩家组队加入
	EventTypeTeamPlayerInvitePlayerDealJoin = "TeamPlayerInvitePlayerDealJoin"
	//玩家申请加入队伍
	EventTypeTeamNearApplyJoin = "TeamNearPlayerApplyJoin"
	//玩家被请离队
	EventTypeTeamPlayerBeLeaved = "TeamPlayerBeLeaved"
	//队长对申请决策
	EventTypeTeamApplyDeal = "TeamApplyDeal"
	//队长转让
	EventTypeTeamCaptainTransfer = "TeamCaptainTransfer"
	//玩家离队
	EventTypeTeamPlayerLeave = "TeamPlayerLeave"
	//队长改变组队标识
	EventTypeTeamCaptainChangePurpose = "TeamCaptainChangePurpose"
	//匹配条件不足
	EventTypeTeamMatchNoEough = "TeamMatchNoEough"
	//队长匹配条件不足决策
	EventTypeTeamMatchCondtionFailedDeal = "TeamMatchCondtionFailedDeal"
	//队员准备决策
	EventTypeTeamMatchCondtionPrepareDeal = "TeamMatchCondtionPrepareDeal"
	//玩家催处开始战斗
	EventTypeTeamMatchRushStart = "TeamRushStart"
)

const (
	//玩家队伍改变事件
	EventTypePlayerTeamChange = "PlayerTeamChange"
)

const (
	//队伍匹配中
	EventTypeTeamArenaMatch = "TeamArenaMatch"
	//停止匹配
	EventTypeTeamArenaStopMatch = "TeamArenaStopMatch"
	//匹配成功
	EventTypeTeamArenaMatched = "TeamArenaMatched"
	//匹配失败
	EventTypeTeamArenaMatchFailed = "TeamArenaMatchFailed"
	//由于队伍原因停止匹配
	EventTypeTeamArenaStopMatchOther = "TeamArenaStopMatchOther"
)

const (
	//组队副本开始战斗
	EventTypeTeamCopyStartBattle = "TeamCopyStartBattle"
	//组队副本开始战斗请求失败
	EventTypeTeamCopyStartBattleFailed = "TeamCopyStartBattleFailed"
	//组队副本开始战斗请求成功
	EventTypeTeamCopyStartBattleSucess = "TeamCopyStartBattleSucess"
)

type TeamInviteEventData struct {
	playerId  int64
	inviteTyp teamtypes.TeamInviteType
	id        int64
}

func CreateTeamInviteData(invitedId int64, typ teamtypes.TeamInviteType, id int64) *TeamInviteEventData {
	tted := &TeamInviteEventData{
		playerId:  invitedId,
		inviteTyp: typ,
		id:        id,
	}
	return tted
}

func (d *TeamInviteEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *TeamInviteEventData) GetId() int64 {
	return d.id
}

func (d *TeamInviteEventData) GetInviteTyp() teamtypes.TeamInviteType {
	return d.inviteTyp
}

type TeamInviteDealEventData struct {
	playerId  int64
	inviteTyp teamtypes.TeamInviteType
	id        int64
	teamId    int64
	teamName  string
	agree     bool
}

func CreateTeamInviteDealData(playerId int64, typ teamtypes.TeamInviteType, id int64, teamId int64, teamName string, agree bool) *TeamInviteDealEventData {
	tted := &TeamInviteDealEventData{
		playerId:  playerId,
		inviteTyp: typ,
		id:        id,
		teamId:    teamId,
		teamName:  teamName,
		agree:     agree,
	}
	return tted
}

func (d *TeamInviteDealEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *TeamInviteDealEventData) GetInviteTyp() teamtypes.TeamInviteType {
	return d.inviteTyp
}

func (d *TeamInviteDealEventData) GetId() int64 {
	return d.id
}

func (d *TeamInviteDealEventData) GetTeamId() int64 {
	return d.teamId
}

func (d *TeamInviteDealEventData) GetTeamName() string {
	return d.teamName
}

func (d *TeamInviteDealEventData) GetAgree() bool {
	return d.agree
}

type TeamPlayerIdEventData struct {
	playerId int64
}

func CreateTeamPlayerIdData(playerId int64) *TeamPlayerIdEventData {
	tted := &TeamPlayerIdEventData{
		playerId: playerId,
	}
	return tted
}

func (d *TeamPlayerIdEventData) GetPlayerId() int64 {
	return d.playerId
}

type TeamApplyDealEventData struct {
	applyId int64
	result  teamtypes.TeamResultType
}

func (d *TeamApplyDealEventData) GetApplyId() int64 {
	return d.applyId
}

func (d *TeamApplyDealEventData) GetResult() teamtypes.TeamResultType {
	return d.result
}

func CreateTeamApplyDealEventData(applyId int64, result teamtypes.TeamResultType) *TeamApplyDealEventData {
	d := &TeamApplyDealEventData{
		applyId: applyId,
		result:  result,
	}
	return d
}

type TeamMatchCondtionFailedDealEventData struct {
	result       bool
	memberIdList []int64
}

func (d *TeamMatchCondtionFailedDealEventData) GetMemberIdList() []int64 {
	return d.memberIdList
}

func (d *TeamMatchCondtionFailedDealEventData) GetResult() bool {
	return d.result
}

func CreateTeamMatchCondtionFailedDealEventData(result bool, memberIdList []int64) *TeamMatchCondtionFailedDealEventData {
	d := &TeamMatchCondtionFailedDealEventData{
		result:       result,
		memberIdList: memberIdList,
	}
	return d
}

type TeamMatchCondtionPrepareDealEventData struct {
	pl     player.Player
	result bool
}

func (d *TeamMatchCondtionPrepareDealEventData) GetPlayer() player.Player {
	return d.pl
}

func (d *TeamMatchCondtionPrepareDealEventData) GetResult() bool {
	return d.result
}

func CreateTeamMatchCondtionPrepareDealEventData(pl player.Player, result bool) *TeamMatchCondtionPrepareDealEventData {
	d := &TeamMatchCondtionPrepareDealEventData{
		result: result,
		pl:     pl,
	}
	return d
}
