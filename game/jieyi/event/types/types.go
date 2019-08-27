package types

import (
	commonlog "fgame/fgame/common/log"
	jieyitypes "fgame/fgame/game/jieyi/types"
)

// 结义事件类型
type JieYiEventType string

const (
	JieYiEventTypeDaoJuTypeChange   JieYiEventType = "结义道具类型改变"
	JieYiEventTypeJieYiSuccess                     = "玩家结义成功"
	JieYiEventTypeJieYiChange                      = "结义改变"
	JieYiEventTypeJieYiInviteFail                  = "结义邀请失败"
	JieYiEventTypeTokenLevelChange                 = "信物等级改变"
	JieYiEventTypeNameUpLev                        = "威名升级"
	JieYiEventTypeShengWeiZhiChange                = "声威值改变"

	JieYiEventTypeJionJieYiLog         = "加入结义"
	JieYiEventTypeLeaveJieYiLog        = "离开结义"
	JieYiEventTypeDaoJuTypeChangeLog   = "道具类型发生变化"
	JieYiEventTypeTokenTypeChangeLog   = "信物类型发生变化"
	JieYiEventTypeTokenLevelChangeLog  = "信物等级发生变化"
	JieYiEventTypeNameLevelChangeLog   = "威名等级发生变化"
	JieYiEventTypeShengWeiZhiChangeLog = "声威值发生变化"
)

// 结义道具变化日志
type PlayerJieYiDaoJuChangeLogEventData struct {
	beformType jieyitypes.JieYiDaoJuType
	curType    jieyitypes.JieYiDaoJuType
	reason     commonlog.JieYiLogReason
	reasonText string
}

func CreatePlayerJieYiDaoJuChangeLogEventData(beformType jieyitypes.JieYiDaoJuType, curType jieyitypes.JieYiDaoJuType, reason commonlog.JieYiLogReason, reasonText string) *PlayerJieYiDaoJuChangeLogEventData {
	d := &PlayerJieYiDaoJuChangeLogEventData{
		beformType: beformType,
		curType:    curType,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerJieYiDaoJuChangeLogEventData) GetBeformType() jieyitypes.JieYiDaoJuType {
	return d.beformType
}

func (d *PlayerJieYiDaoJuChangeLogEventData) GetCurType() jieyitypes.JieYiDaoJuType {
	return d.curType
}

func (d *PlayerJieYiDaoJuChangeLogEventData) GetReason() commonlog.JieYiLogReason {
	return d.reason
}

func (d *PlayerJieYiDaoJuChangeLogEventData) GetReasonText() string {
	return d.reasonText
}

// 结义信物类型变化日志
type PlayerJieYiTokenTypeChangeLogEventData struct {
	beformType jieyitypes.JieYiTokenType
	curType    jieyitypes.JieYiTokenType
	reason     commonlog.JieYiLogReason
	reasonText string
}

func CreatePlayerJieYiTokenTypeChangeLogEventData(beformType jieyitypes.JieYiTokenType, curType jieyitypes.JieYiTokenType, reason commonlog.JieYiLogReason, reasonText string) *PlayerJieYiTokenTypeChangeLogEventData {
	d := &PlayerJieYiTokenTypeChangeLogEventData{
		beformType: beformType,
		curType:    curType,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerJieYiTokenTypeChangeLogEventData) GetBeformType() jieyitypes.JieYiTokenType {
	return d.beformType
}

func (d *PlayerJieYiTokenTypeChangeLogEventData) GetCurType() jieyitypes.JieYiTokenType {
	return d.curType
}

func (d *PlayerJieYiTokenTypeChangeLogEventData) GetReason() commonlog.JieYiLogReason {
	return d.reason
}

func (d *PlayerJieYiTokenTypeChangeLogEventData) GetReasonText() string {
	return d.reasonText
}

// 结义信物等级变化日志
type PlayerJieYiTokenLevelChangeLogEventData struct {
	beformLevel int32
	curLevel    int32
	reason      commonlog.JieYiLogReason
	reasonText  string
}

func CreatePlayerJieYiTokenLevelChangeLogEventData(beformLevel int32, curLevel int32, reason commonlog.JieYiLogReason, reasonText string) *PlayerJieYiTokenLevelChangeLogEventData {
	d := &PlayerJieYiTokenLevelChangeLogEventData{
		beformLevel: beformLevel,
		curLevel:    curLevel,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerJieYiTokenLevelChangeLogEventData) GetBeformLevel() int32 {
	return d.beformLevel
}

func (d *PlayerJieYiTokenLevelChangeLogEventData) GetCurLevel() int32 {
	return d.curLevel
}

func (d *PlayerJieYiTokenLevelChangeLogEventData) GetReason() commonlog.JieYiLogReason {
	return d.reason
}

func (d *PlayerJieYiTokenLevelChangeLogEventData) GetReasonText() string {
	return d.reasonText
}

// 结义威名等级变化日志
type PlayerJieYiNameLevelChangeLogEventData struct {
	beformLevel int32
	curLevel    int32
	reason      commonlog.JieYiLogReason
	reasonText  string
}

func CreatePlayerJieYiNameLevelChangeLogEventData(beformLevel int32, curLevel int32, reason commonlog.JieYiLogReason, reasonText string) *PlayerJieYiNameLevelChangeLogEventData {
	d := &PlayerJieYiNameLevelChangeLogEventData{
		beformLevel: beformLevel,
		curLevel:    curLevel,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerJieYiNameLevelChangeLogEventData) GetBeformLevel() int32 {
	return d.beformLevel
}

func (d *PlayerJieYiNameLevelChangeLogEventData) GetCurLevel() int32 {
	return d.curLevel
}

func (d *PlayerJieYiNameLevelChangeLogEventData) GetReason() commonlog.JieYiLogReason {
	return d.reason
}

func (d *PlayerJieYiNameLevelChangeLogEventData) GetReasonText() string {
	return d.reasonText
}

// 声威值变化日志
type PlayerJieYiShengWeiZhiChangeLogEventData struct {
	beformNum  int32
	curNum     int32
	reason     commonlog.JieYiLogReason
	reasonText string
}

func CreatePlayerJieYiShengWeiZhiChangeLogEventData(beformNum int32, curNum int32, reason commonlog.JieYiLogReason, reasonText string) *PlayerJieYiShengWeiZhiChangeLogEventData {
	d := &PlayerJieYiShengWeiZhiChangeLogEventData{
		beformNum:  beformNum,
		curNum:     curNum,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerJieYiShengWeiZhiChangeLogEventData) GetBeformNum() int32 {
	return d.beformNum
}

func (d *PlayerJieYiShengWeiZhiChangeLogEventData) GetCurNum() int32 {
	return d.curNum
}

func (d *PlayerJieYiShengWeiZhiChangeLogEventData) GetReason() commonlog.JieYiLogReason {
	return d.reason
}

func (d *PlayerJieYiShengWeiZhiChangeLogEventData) GetReasonText() string {
	return d.reasonText
}
