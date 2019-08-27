package types

import (
	commonlog "fgame/fgame/common/log"
)

type AllianceEventType string

const (
	//成员加入仙盟
	EventTypeAllianceMemberJoin AllianceEventType = "AllianceMemberJoin"
	//成员退出仙盟
	EventTypeAllianceMemberExit = "AlliacenMemberExit"
	//仙盟虎符改变
	EventTypeAllianceHuFuChanged = "AllianceHufuChanged"
	//斗神成员领域改变
	EventTypeAllianceDouShenMemberLingYuChanged = "AllianceDouShenMemberLingYuChanged"
	//斗神成员退出
	EventTypeAllianceDouShenMemberExit = "AllianceDouShenMemberExit"
	//斗神领域改变
	EventTypeAllianceDouShenLingYuChanged = "AllianceDouShenLingYuChanged"
	//仙盟升级
	EventTypeAllianceLevelChanged = "AllianceLevelChanged"
	//盟主变更
	EventTypeAllianceMengzhuChanged = "AllianceMengzhuChanged"
	//仙盟仓库变更
	EventTypeAllianceDepotChanged = "AllianceDepotChanged"
	//仙盟仓库整理
	EventTypeAllianceDepotMerge = "AllianceDepotMerge"
	//仙盟仓库设置改变
	EventTypeAllianceDepotSettingChanged = "AllianceDepotSettingChanged"
	//城战胜利
	EventTypeAllianceWinChengZhan = "AllianceWinChengZhan"
	//仙盟仓库设置改变日志
	EventTypeAllianceDepotSettingChangedLog = "AllianceDepotSettingChangedLog"
	//仙盟仓库物品变化日志
	EventTypeAllianceDepotItemChangedLog = "AllianceDepotItemChangedLog"
	//仙盟成员职位变更
	EventTypeAllianceMemberPositionChanged = "AlliacenMemberPositionChanged"
	//加载城战获胜仙盟盟主雕像
	EventTypeAllianceLoadWinnerModel = "AllianceLoadWinnerModel"
	//仙盟合并
	EventTypeAllianceMerge = "AllianceMerge"
	//仙盟名变更
	EventTypeAllianceNameChanged = "AllianceNameChanged"
	//仙盟合并日志
	EventTypeAllianceMergeLog = "AllianceMergeLog"
	//仙盟阵营类型改变
	EventTypeAllianceCampTypeChanged = "AllianceCampTypeChanged"
	//仙盟解散
	EventTypeAllianceDismiss = "AllianceDismiss"
)

//仙盟仓库设置改变日志
type AllianceDepotSettingChangedLogEventData struct {
	reason     commonlog.AllianceLogReason
	reasonText string
}

func CreateAllianceDepotSettingChangedLogEventData(reason commonlog.AllianceLogReason, reasonText string) *AllianceDepotSettingChangedLogEventData {
	d := &AllianceDepotSettingChangedLogEventData{
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *AllianceDepotSettingChangedLogEventData) GetReason() commonlog.AllianceLogReason {
	return d.reason
}

func (d *AllianceDepotSettingChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//仙盟仓库物品变化日志
type AllianceDepotItemChangedLogEventData struct {
	itemId     int32
	changedNum int32
	reason     commonlog.AllianceLogReason
	reasonText string
}

func CreateAllianceDepotItemChangedLogEventData(itemId, changedNum int32, reason commonlog.AllianceLogReason, reasonText string) *AllianceDepotItemChangedLogEventData {
	d := &AllianceDepotItemChangedLogEventData{
		itemId:     itemId,
		changedNum: changedNum,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *AllianceDepotItemChangedLogEventData) GetItemId() int32 {
	return d.itemId
}

func (d *AllianceDepotItemChangedLogEventData) GetChangedNum() int32 {
	return d.changedNum
}

func (d *AllianceDepotItemChangedLogEventData) GetReason() commonlog.AllianceLogReason {
	return d.reason
}

func (d *AllianceDepotItemChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//仙盟成员退出
type AllianceMemberExitEventData struct {
	memId             int64
	isClearPlayerData bool
}

func CreateAllianceMemberExitEventData(memId int64, isClearPlayerData bool) *AllianceMemberExitEventData {
	d := &AllianceMemberExitEventData{
		memId:             memId,
		isClearPlayerData: isClearPlayerData,
	}
	return d
}

func (d *AllianceMemberExitEventData) GetMemberId() int64 {
	return d.memId
}

func (d *AllianceMemberExitEventData) IsClearPlayerData() bool {
	return d.isClearPlayerData
}
