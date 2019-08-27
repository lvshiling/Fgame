package types

import (
	commonlog "fgame/fgame/common/log"
	alliancetypes "fgame/fgame/game/alliance/types"
)

type PlayerAllianceEventType string

const (
	EventTypePlayerAllianceJoin                 PlayerAllianceEventType = "PlayerAllianceJoin"                 //成员加入仙盟
	EventTypePlayerAllianceExit                                         = "PlayerAllianceExit"                 //成员退出仙盟
	EventTypePlayerAllianceMengZhuChanged                               = "PlayerAllianceMengZhuChanged"       //盟主变化
	EventTypePlayerAllianceChanged                                      = "PlayerAllianceChanged"              //仙盟变化
	EventTypePlayerAllianceSkillUpgrade                                 = "PlayerAllianceSkillUpgrade"         //成员仙术升级
	EventTypePlayerYaoPaiChanged                                        = "PlayerAllianceYaoPaiChanged"        //腰牌变化
	EventTypePlayerAllianceSkillChanged                                 = "PlayerAllianceSkillChanged"         //仙盟技能改变
	EventTypePlayerAllianceDepotPointChangedLog                         = "PlayerAllianceDepotPointChangedLog" //仓库积分改变日志
	EventTypePlayerAlliancePositionChanged                              = "PlayerAllianceMengZhuChanged"       //仙盟职位变化
)

type PlayerAllianceSkillUpgradeEventData struct {
	typ   alliancetypes.AllianceSkillType
	level int32
}

func CreatePlayerAllianceSkillUpgradeEventData(typ alliancetypes.AllianceSkillType, level int32) *PlayerAllianceSkillUpgradeEventData {
	data := &PlayerAllianceSkillUpgradeEventData{
		typ:   typ,
		level: level,
	}
	return data
}

func (d *PlayerAllianceSkillUpgradeEventData) GetSkillType() alliancetypes.AllianceSkillType {
	return d.typ
}

func (d *PlayerAllianceSkillUpgradeEventData) GetSkillLevel() int32 {
	return d.level
}

// 仓库积分日志
type PlayerAllianceDepotPointLogEventData struct {
	beforePoint  int32
	changedPoint int32
	reason       commonlog.AllianceLogReason
	reasonText   string
}

func CreatePlayerAllianceDepotPointLogEventData(beforePoint, changedPoint int32, reason commonlog.AllianceLogReason, reasonText string) *PlayerAllianceDepotPointLogEventData {
	d := &PlayerAllianceDepotPointLogEventData{
		beforePoint:  beforePoint,
		changedPoint: changedPoint,
		reason:       reason,
		reasonText:   reasonText,
	}
	return d
}

func (d *PlayerAllianceDepotPointLogEventData) GetBeforePoint() int32 {
	return d.beforePoint
}

func (d *PlayerAllianceDepotPointLogEventData) GetChangedPoint() int32 {
	return d.changedPoint
}

func (d *PlayerAllianceDepotPointLogEventData) GetReason() commonlog.AllianceLogReason {
	return d.reason
}

func (d *PlayerAllianceDepotPointLogEventData) GetReasonText() string {
	return d.reasonText
}
