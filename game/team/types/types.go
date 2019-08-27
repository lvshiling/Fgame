package types

const (
	//组队人数最大3人
	TeamMaxNum = 3
	//一键cd时间(单位ms) 30s
	TeamCdTime     = 27 * 1000
	TeamOperCd     = "30s"
	TeamPurposeMax = 20
)

type TeamResultType int32

const (
	//决策同意
	TeamResultTypeOk TeamResultType = 1 + iota
	//决策拒绝
	TeamResultTypeNo
)

func (t TeamResultType) Valid() bool {
	switch t {
	case TeamResultTypeOk,
		TeamResultTypeNo:
		return true
	}
	return false
}

type TeamApplyJoinCodeType int32

const (
	//发送申请加入
	TeamApplyJoinCodeTypeSend TeamApplyJoinCodeType = iota
	//当前队伍人数已达上限
	TeamApplyJoinCodeTypeFull
	//当前队伍已解散
	TeamApplyJoinCodeTypeDissolve
	//入队成功
	TeamApplyJoinCodeTypeSucess
)

type TeamInviteType int32

const (
	//邀请组队
	TeamInviteTypeCreate TeamInviteType = 1 + iota
	//邀请加入
	TeamInviteTypeJoin
)

func (tit TeamInviteType) Valid() bool {
	switch tit {
	case TeamInviteTypeCreate,
		TeamInviteTypeJoin:
		return true
	}
	return false
}

//离队队伍状态
type TeamLeaveStatusType int32

const (
	//解散队伍
	TeamLeaveStatusTypeDissolve TeamLeaveStatusType = 1 + iota
	//转让队长
	TeamLeaveStatusTypeTransfer
)

//广播类型
type TeamBroadcastType int32

const (
	//转让队长
	TeamBroadcastTypeTransfer TeamBroadcastType = 1 + iota
	//离队
	TeamBroadcastTypeleave
	//请离队
	TeamBroadcastTypeleaved
	//申请者加入成功
	TeamBroadcastTypeApplyJoin
	//被邀请者同意加入
	TeamBroadCastTypeInviteAgree
)
