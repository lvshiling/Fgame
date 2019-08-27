package scene

import (
	arenatemplate "fgame/fgame/game/arena/template"
	arenatypes "fgame/fgame/game/arena/types"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
)

type MemberStatus int32

const (
	MemberStatusOffline MemberStatus = iota //离线
	MemberStatusOnline                      //在线
	MemberStatusFailed                      //失败
	MemberStatusGoAway                      //退出
)

type TeamMemberObject struct {
	serverId   int32                //服务器id
	playerId   int64                //玩家id
	force      int64                //玩家战力
	status     MemberStatus         //是否在线
	name       string               //玩家名字
	level      int32                //玩家等级
	role       playertypes.RoleType //玩家角色
	sex        playertypes.SexType  //性别
	fashionId  int32                //时装id
	reliveTime int32                //复活次数
	winCount   int32                //连胜次数
	robot      bool                 //机器人
}

func (o *TeamMemberObject) String() string {
	return fmt.Sprintf("id:%d,名字:%s", o.playerId, o.name)
}

func (o *TeamMemberObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *TeamMemberObject) GetForce() int64 {
	return o.force
}

func (o *TeamMemberObject) SetStatus(status MemberStatus) {
	o.status = status
}

func (o *TeamMemberObject) GetStatus() MemberStatus {
	return o.status
}

func (o *TeamMemberObject) GetName() string {
	return o.name
}

func (o *TeamMemberObject) GetLevel() int32 {
	return o.level
}

func (o *TeamMemberObject) GetRole() playertypes.RoleType {
	return o.role
}

func (o *TeamMemberObject) GetSex() playertypes.SexType {
	return o.sex
}

func (o *TeamMemberObject) GetFashionId() int32 {
	return o.fashionId
}

func (o *TeamMemberObject) SetReliveTime(reliveTime int32) {
	o.reliveTime = reliveTime
}

func (o *TeamMemberObject) GetReliveTime() int32 {
	return o.reliveTime
}

func (o *TeamMemberObject) GetRemainReliveTime() int32 {
	return arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RebornAmountMax - o.reliveTime
}

func (o *TeamMemberObject) GetServerId() int32 {
	return o.serverId
}

func (o *TeamMemberObject) IsRobot() bool {
	return o.robot
}

func CreateTeamMemberObject(
	serverId int32,
	playerId int64,
	force int64,
	name string,
	level int32,
	role playertypes.RoleType,
	sex playertypes.SexType,
	fashionId int32,
	robot bool,
	winCount int32) *TeamMemberObject {
	to := &TeamMemberObject{
		serverId:  serverId,
		playerId:  playerId,
		force:     force,
		status:    MemberStatusOffline,
		name:      name,
		level:     level,
		role:      role,
		sex:       sex,
		fashionId: fashionId,
		robot:     robot,
		winCount:  winCount,
	}
	return to
}

type TeamObject struct {
	teamId     int64
	force      int64
	memberList []*TeamMemberObject
}

func (o *TeamObject) String() string {
	return fmt.Sprintf("队伍:%d,战力:%d,成员:%s", o.teamId, o.force, o.memberList)
}

func (o *TeamObject) GetTeamId() int64 {
	return o.teamId
}

func (o *TeamObject) GetTeamName() string {
	captain := o.GetCaptain()
	if captain == nil {
		return ""
	}
	return captain.GetName()
}

func (o *TeamObject) GetMemberList() []*TeamMemberObject {
	return o.memberList
}

func (o *TeamObject) GetMember(playerId int64) (mem *TeamMemberObject, pos int32) {
	pos = -1
	for index, member := range o.memberList {
		if member.GetPlayerId() == playerId {
			pos = int32(index)
			mem = member
			break
		}
	}
	return
}

func (o *TeamObject) RemoveMember(memberId int64) {
	mem, pos := o.GetMember(memberId)
	if mem == nil {
		return
	}
	o.memberList = append(o.memberList[:pos], o.memberList[pos+1:]...)
}

func (o *TeamObject) GetCaptain() *TeamMemberObject {
	if len(o.memberList) == 0 {
		return nil
	}
	for _, mem := range o.memberList {
		if mem.status == MemberStatusOffline || mem.status == MemberStatusOnline {
			return mem
		}
	}
	return nil
}

func (o *TeamObject) GetForce() int64 {
	return o.force
}

func (o *TeamObject) IfAllLeave() bool {
	for _, mem := range o.memberList {
		if mem.status == MemberStatusOnline {
			return false
		}
	}
	return true
}

func (o *TeamObject) IfAllRobot() bool {
	for _, mem := range o.memberList {
		if !mem.robot {
			return false
		}
	}
	return true
}

func (o *TeamObject) GetTotalReliveTime() int32 {
	total := int32(0)
	for _, mem := range o.memberList {
		total += mem.GetRemainReliveTime()
	}
	return total
}

func CreateTeamObject(teamId int64, force int64, memberList []*TeamMemberObject) *TeamObject {
	to := &TeamObject{
		teamId:     teamId,
		force:      force,
		memberList: memberList,
	}
	return to
}

type ArenaTeamState int32

const (
	//初始化
	ArenaTeamStateInit ArenaTeamState = iota
	//匹配
	ArenaTeamStateMatch
	//比赛中
	ArenaTeamStateGame
	//比赛结束
	ArenaTeamStateGameEnd
	//四圣兽等待
	ArenaTeamStateFourGodInit
	//四圣兽进入中
	ArenaTeamStateFourGodEnter
	//四圣兽排队中
	ArenaTeamStateFourGodQueue
	//四圣兽中
	ArenaTeamStateFourGod
	//结束
	ArenaTeamStateEnd
)

var (
	arenaTeamStateMap = map[ArenaTeamState]string{
		ArenaTeamStateInit:         "初始化",
		ArenaTeamStateMatch:        "匹配中",
		ArenaTeamStateGame:         "比赛中",
		ArenaTeamStateGameEnd:      "比赛完",
		ArenaTeamStateFourGodInit:  "四圣兽初始化",
		ArenaTeamStateFourGodEnter: "四圣兽进入中",
		ArenaTeamStateFourGodQueue: "四圣兽排队中",
		ArenaTeamStateFourGod:      "四圣兽中",
		ArenaTeamStateEnd:          "结束",
	}
)

func (s ArenaTeamState) String() string {
	return arenaTeamStateMap[s]
}

//3v3队伍
type ArenaTeam struct {
	//队伍
	tm *TeamObject
	//当前轮数
	current int32
	//状态
	state ArenaTeamState
	//上一次时间
	lastTime int64
	//当前圣兽场景
	fourGodType arenatypes.FourGodType
}

func (t ArenaTeam) String() string {
	return fmt.Sprintf("队伍:%s,当前轮数:%d,状态:%s", t.tm.String(), t.current, t.state.String())
}

func (t *ArenaTeam) GetTeam() *TeamObject {
	return t.tm
}

func (t *ArenaTeam) GetCurrent() int32 {
	return t.current
}

func (t *ArenaTeam) GetMemberMaxWinCount() int32 {
	memMaxCount := int32(0)
	for _, temMem := range t.tm.memberList {
		if memMaxCount > temMem.winCount {
			continue
		}

		memMaxCount = temMem.winCount
	}
	return memMaxCount
}

func (t *ArenaTeam) Win() {
	t.current += 1
}

func (t *ArenaTeam) GetState() ArenaTeamState {
	return t.state
}

func (t *ArenaTeam) Match() bool {
	if t.state != ArenaTeamStateInit && t.state != ArenaTeamStateGameEnd {
		return false
	}

	//踢掉失败的玩家
	t.KickFailed()
	t.state = ArenaTeamStateMatch
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) Game() bool {
	// if t.state != ArenaTeamStateMatch {
	// 	return false
	// }
	t.state = ArenaTeamStateGame
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) GameEnd() bool {
	if t.state != ArenaTeamStateGame {
		return false
	}
	t.state = ArenaTeamStateGameEnd
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) EnterFourGod() bool {
	if t.state != ArenaTeamStateGameEnd && t.state != ArenaTeamStateFourGod {
		return false
	}
	t.KickFailed()
	t.state = ArenaTeamStateFourGodInit
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) KickFailed() {
	var kickMemberList []*TeamMemberObject
	//需要踢掉的玩家
	for _, mem := range t.GetTeam().GetMemberList() {
		if mem.GetStatus() == MemberStatusFailed || mem.GetStatus() == MemberStatusGoAway {
			kickMemberList = append(kickMemberList, mem)
		}
	}
	for _, kickMem := range kickMemberList {
		t.GetTeam().RemoveMember(kickMem.GetPlayerId())
	}
}

func (t *ArenaTeam) EnterFourGodGame(fourGodType arenatypes.FourGodType) bool {
	if t.state != ArenaTeamStateFourGodInit {
		return false
	}

	t.state = ArenaTeamStateFourGodEnter
	t.fourGodType = fourGodType
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) FourGodGameQueue() bool {
	if t.state != ArenaTeamStateFourGodEnter {
		return false
	}
	t.state = ArenaTeamStateFourGodQueue
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) CancelFourGodGameQueue() bool {
	if t.state != ArenaTeamStateFourGodQueue {
		return false
	}
	t.state = ArenaTeamStateFourGodInit
	t.fourGodType = arenatypes.FourGodTypeQingLong
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) FourGodGame() bool {
	if t.state != ArenaTeamStateFourGodEnter && t.state != ArenaTeamStateFourGodQueue {
		return false
	}
	t.KickFailed()
	t.state = ArenaTeamStateFourGod
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) End() bool {

	t.state = ArenaTeamStateEnd
	now := global.GetGame().GetTimeService().Now()
	t.lastTime = now
	return true
}

func (t *ArenaTeam) GetFourGodType() arenatypes.FourGodType {
	return t.fourGodType
}

func (t *ArenaTeam) GetLastTime() int64 {
	return t.lastTime
}

func CreateArenaTeam(tm *TeamObject) *ArenaTeam {
	at := &ArenaTeam{
		tm:      tm,
		current: 0,
		state:   ArenaTeamStateInit,
	}
	return at
}

func CreateArenaTeamWithMembers(teamId int64, memList []*TeamMemberObject) *ArenaTeam {
	t := CreateTeamObject(teamId, 0, memList)
	at := CreateArenaTeam(t)
	return at
}

func CreateArenaTeamWithMemberLists(teamId int64, memListOfList ...[]*TeamMemberObject) *ArenaTeam {
	teamMemList := make([]*TeamMemberObject, 0, 3)
	for _, memList := range memListOfList {
		teamMemList = append(teamMemList, memList...)
	}
	t := CreateTeamObject(teamId, 0, teamMemList)
	at := CreateArenaTeam(t)
	return at
}
