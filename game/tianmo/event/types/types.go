package types

import (
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
)

type TianMoEventType string

const (
	EventTypeTianMoAdvanced     TianMoEventType = "TianMoAdvanced"     //天魔进阶事件
	EventTypeTianMoAdvancedCost                 = "TianMoAdvancedCost" //天魔进阶消耗
	EventTypeTianMoAdvancedLog                  = "TianMoAdvancedLog"  //天魔进阶日志
	EventTypeTianMoPowerChanged                 = "TianMoPowerChanged" //天魔战力变化
	EventTypeTianMoUnitePiFu                    = "TianMoUnitePiFu"    //天魔关联皮肤事件
)

//
type PlayerTianMoAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.TianMoLogReason
	reasonText        string
}

func CreatePlayerTianMoAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.TianMoLogReason, reasonText string) *PlayerTianMoAdvancedLogEventData {
	d := &PlayerTianMoAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerTianMoAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerTianMoAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerTianMoAdvancedLogEventData) GetReason() commonlog.TianMoLogReason {
	return d.reason
}

func (d *PlayerTianMoAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}

//天魔体关联皮肤事件信息
type PlayerTianMoTiUnitePiFuEventData struct {
	piFuType commontypes.AdvancedUnitePiFuType
	piFuId   int32
}

func CreatePlayerTianMoTiUnitePiFuEventData(piFuType commontypes.AdvancedUnitePiFuType, piFuId int32) *PlayerTianMoTiUnitePiFuEventData {
	d := &PlayerTianMoTiUnitePiFuEventData{
		piFuType: piFuType,
		piFuId:   piFuId,
	}
	return d
}

func (d *PlayerTianMoTiUnitePiFuEventData) GetPiFuType() commontypes.AdvancedUnitePiFuType {
	return d.piFuType
}

func (d *PlayerTianMoTiUnitePiFuEventData) GetPiFuId() int32 {
	return d.piFuId
}
