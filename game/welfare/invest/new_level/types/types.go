package types

import (
	commonlog "fgame/fgame/common/log"
)

//新等级投资计划类型
type InvestNewLevelType int32

const (
	InvesetNewLevelTypeJunior  InvestNewLevelType = iota //初级投资计划
	InvesetNewLevelTypeMiddle                            //中级投资计划
	InvesetNewLevelTypeHigher                            //高级投资计划
	InvesetNewLevelTypeHighest                           //顶级投资计划
)

const (
	MinType = InvesetNewLevelTypeJunior
	MaxType = InvesetNewLevelTypeHighest
)

func (t InvestNewLevelType) Valid() bool {
	switch t {
	case InvesetNewLevelTypeJunior,
		InvesetNewLevelTypeMiddle,
		InvesetNewLevelTypeHigher,
		InvesetNewLevelTypeHighest:
		return true
	default:
		return false
	}
}

// 购买新等级计划消耗日志类型
var invesetNewLevelCostLogTypeMap = map[InvestNewLevelType]commonlog.GoldLogReason{
	InvesetNewLevelTypeJunior:  commonlog.GoldLogReasonBuyJunior,
	InvesetNewLevelTypeMiddle:  commonlog.GoldLogReasonBuyMiddle,
	InvesetNewLevelTypeHigher:  commonlog.GoldLogReasonBuyHigher,
	InvesetNewLevelTypeHighest: commonlog.GoldLogReasonBuyHighest,
}

func (t InvestNewLevelType) ConvertToInvesetNewLevelCostLogType() (commonlog.GoldLogReason, bool) {
	typ, ok := invesetNewLevelCostLogTypeMap[t]
	return typ, ok
}

// 购买新等级计划消耗日志类型
var invesetNewLevelUpgradeCostLogTypeMap = map[InvestNewLevelType]commonlog.GoldLogReason{

	InvesetNewLevelTypeMiddle:  commonlog.GoldLogReasonInvestUpLevelMiddle,
	InvesetNewLevelTypeHigher:  commonlog.GoldLogReasonInvestUpLevelHigher,
	InvesetNewLevelTypeHighest: commonlog.GoldLogReasonInvestUpLevelHighest,
}

func (t InvestNewLevelType) ConvertToInvesetNewLevelUpgradeCostLogType() (commonlog.GoldLogReason, bool) {
	typ, ok := invesetNewLevelUpgradeCostLogTypeMap[t]
	return typ, ok
}

// 投资计划字符串类型，邮件用
type InvestNewLevelStringType string

const (
	InvesetStringTypeJunior  InvestNewLevelStringType = "初级投资"
	InvesetStringTypeMiddle                           = "中级投资"
	InvesetStringTypeHigher                           = "高级投资"
	InvesetStringTypeHighest                          = "顶级投资"
)

// 新等级计划类型，邮件用
var invesetNewLevelEmailTypeMap = map[InvestNewLevelType]InvestNewLevelStringType{
	InvesetNewLevelTypeJunior:  InvesetStringTypeJunior,
	InvesetNewLevelTypeMiddle:  InvesetStringTypeMiddle,
	InvesetNewLevelTypeHigher:  InvesetStringTypeHigher,
	InvesetNewLevelTypeHighest: InvesetStringTypeHighest,
}

func (t InvestNewLevelType) ConvertToInvesetNewLevelEmailType() (string, bool) {
	typ, ok := invesetNewLevelEmailTypeMap[t]
	return string(typ), ok
}
