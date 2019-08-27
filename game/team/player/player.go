package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	teamtypes "fgame/fgame/game/team/types"
)

//玩家组队管理器
type PlayerTeamDataManager struct {
	p player.Player
	// //我的队伍id
	// teamId int64
	// //我的队伍名称
	// teamName string
	//TODO 限制最大邀请人数
	//邀请我的
	inviteMap map[teamtypes.TeamInviteType]map[int64]int64
	//一键申请时间
	applyAllTime int64
	//一键邀请时间
	inviteAllTime int64
	//是否匹配条件不足队员
	matchFailedMemberList []int64
	//自己是否收到匹配条件不足
	isMatchCondtionFailed bool
	//开始催处时间
	rushTime int64
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (ptdm *PlayerTeamDataManager) Player() player.Player {
	return ptdm.p
}

//加载
func (ptdm *PlayerTeamDataManager) Load() (err error) {
	return nil
}

//加载
func (ptdm *PlayerTeamDataManager) AfterLoad() (err error) {

	return nil
}

//心跳
func (ptdm *PlayerTeamDataManager) Heartbeat() {

}

// //获取队伍id
// func (ptdm *PlayerTeamDataManager) GetTeamId() int64 {
// 	return ptdm.teamId
// }

// //获取队伍名称
// func (ptdm *PlayerTeamDataManager) GetTeamName() string {
// 	return ptdm.teamName
// }

//获取一键申请时间
func (ptdm *PlayerTeamDataManager) GetApplyAllTime() int64 {
	return ptdm.applyAllTime
}

//获取一键邀请时间
func (ptdm *PlayerTeamDataManager) GetInviteAllTime() int64 {
	return ptdm.inviteAllTime
}

//判断cd
func (ptdm *PlayerTeamDataManager) IfCanApplyAllTime() (sucess bool) {
	now := global.GetGame().GetTimeService().Now()
	diff := now - ptdm.applyAllTime
	if diff < teamtypes.TeamCdTime {
		return
	}
	sucess = true
	return
}

func (ptdm *PlayerTeamDataManager) TeamApplyAllTime() {
	now := global.GetGame().GetTimeService().Now()
	ptdm.applyAllTime = now
}

func (ptdm *PlayerTeamDataManager) TeamInviteAllTime() {
	now := global.GetGame().GetTimeService().Now()
	ptdm.inviteAllTime = now
}

//判断cd
func (ptdm *PlayerTeamDataManager) IfCanInviteAllTime() (sucess bool) {
	now := global.GetGame().GetTimeService().Now()
	diff := now - ptdm.inviteAllTime
	if diff < teamtypes.TeamCdTime {
		return
	}
	sucess = true
	return
}

//是否存在的邀请
func (ptdm *PlayerTeamDataManager) IfExistInvite(typ teamtypes.TeamInviteType, id int64) (flag bool) {
	inviteTypMap, exist := ptdm.inviteMap[typ]
	if !exist {
		return
	}
	_, exist = inviteTypMap[id]
	if !exist {
		return
	}
	flag = true
	return
}

//存储申请
func (ptdm *PlayerTeamDataManager) TeamInvite(typ teamtypes.TeamInviteType, id int64) {
	inviteTypMap, exist := ptdm.inviteMap[typ]
	if !exist {
		inviteTypMap = make(map[int64]int64)
		ptdm.inviteMap[typ] = inviteTypMap
	}
	now := global.GetGame().GetTimeService().Now()
	inviteTypMap[id] = now
}

//被邀请玩家决策
func (ptdm *PlayerTeamDataManager) TeamInvitedDeal(typ teamtypes.TeamInviteType, id int64) {
	inviteTypMap, exist := ptdm.inviteMap[typ]
	if !exist {
		return
	}
	delete(inviteTypMap, id)
}

// //设置队伍id
// func (ptdm *PlayerTeamDataManager) SetTeam(teamId int64, teamName string) {
// 	if teamId <= 0 {
// 		return
// 	}
// 	if ptdm.teamId == teamId {
// 		return
// 	}
// 	ptdm.teamId = teamId
// 	ptdm.teamName = teamName
// 	gameevent.Emit(teameventtypes.EventTypePlayerTeamChange, ptdm.p, nil)
// }

// func (ptdm *PlayerTeamDataManager) SetTeamName(teamName string) {
// 	if ptdm.teamName == teamName {
// 		return
// 	}
// 	ptdm.teamName = teamName
// 	gameevent.Emit(teameventtypes.EventTypePlayerTeamChange, ptdm.p, nil)
// }

// //离队或被请离队伍id置0
// func (ptdm *PlayerTeamDataManager) ResetTeam() {
// 	ptdm.teamId = 0
// 	ptdm.teamName = ""
// 	gameevent.Emit(teameventtypes.EventTypePlayerTeamChange, ptdm.p, nil)
// }

func (ptdm *PlayerTeamDataManager) SetMatchCondtionFailedList(memberIdList []int64) {
	ptdm.matchFailedMemberList = memberIdList
}

func (ptdm *PlayerTeamDataManager) GetMatchCondtionFailedList() ([]int64, bool) {
	flag := false
	if len(ptdm.matchFailedMemberList) != 0 {
		flag = true
	}
	return ptdm.matchFailedMemberList, flag
}

func (ptdm *PlayerTeamDataManager) ResetMatchCondtionFailedList() {
	ptdm.matchFailedMemberList = nil
}

func (ptdm *PlayerTeamDataManager) SetMatchCondtionFailed() {
	ptdm.isMatchCondtionFailed = true
}

func (ptdm *PlayerTeamDataManager) IsExistMatchCondtionFailed() bool {
	return ptdm.isMatchCondtionFailed
}

func (ptdm *PlayerTeamDataManager) ResetMatchCondtionFailed() {
	ptdm.isMatchCondtionFailed = false
}

func (ptdm *PlayerTeamDataManager) GetRushTime() int64 {
	return ptdm.rushTime
}

func (ptdm *PlayerTeamDataManager) RushTimeLeftTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - ptdm.rushTime
	rushCdTime3v3 := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantType3V3RushCdTime) - 3000)
	return rushCdTime3v3 - elapse
}

func (ptdm *PlayerTeamDataManager) IsRushTimeInCd() (sucess bool) {
	now := global.GetGame().GetTimeService().Now()
	if ptdm.rushTime == 0 {
		return
	}
	diff := now - ptdm.rushTime
	rushCdTime3v3 := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantType3V3RushCdTime) - 3000)
	if diff >= rushCdTime3v3 {
		return
	}
	sucess = true
	return
}

func (ptdm *PlayerTeamDataManager) SetRushTime() {
	now := global.GetGame().GetTimeService().Now()
	ptdm.rushTime = now
}

func CreatePlayerTeamDataManager(p player.Player) player.PlayerDataManager {
	ptdm := &PlayerTeamDataManager{}
	ptdm.p = p
	// ptdm.teamId = 0
	ptdm.applyAllTime = 0
	ptdm.inviteAllTime = 0
	ptdm.isMatchCondtionFailed = false
	ptdm.inviteMap = make(map[teamtypes.TeamInviteType]map[int64]int64)
	ptdm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return ptdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerTeamDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTeamDataManager))
}
