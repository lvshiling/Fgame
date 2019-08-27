package types

import (
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
)

type ShiHunFanEventType string

const (
	EventTypeShiHunFanAdvanced     ShiHunFanEventType = "ShiHunFanAdvanced"     //噬魂幡进阶事件
	EventTypeShiHunFanUnitePiFu                       = "ShiHunFanUnitePiFu"    //噬魂幡关联皮肤事件
	EventTypeShiHunFanAdvancedLog                     = "ShiHunFanAdvancedLog"  //噬魂幡进阶日志
	EventTypeShiHunFanPowerChanged                    = "ShiHunFanPowerChanged" //噬魂幡战力变化
	EventTypeShiHunFanAdvancedCost                    = "ShiHunFanAdvancedCost" //噬魂幡进阶消耗
)

//噬魂幡关联皮肤事件信息
type PlayerShiHunFanUnitePiFuEventData struct {
	piFuType commontypes.AdvancedUnitePiFuType
	piFuId   int32
}

func CreatePlayerShiHunFanUnitePiFuEventData(piFuType commontypes.AdvancedUnitePiFuType, piFuId int32) *PlayerShiHunFanUnitePiFuEventData {
	d := &PlayerShiHunFanUnitePiFuEventData{
		piFuType: piFuType,
		piFuId:   piFuId,
	}
	return d
}

func (d *PlayerShiHunFanUnitePiFuEventData) GetPiFuType() commontypes.AdvancedUnitePiFuType {
	return d.piFuType
}

func (d *PlayerShiHunFanUnitePiFuEventData) GetPiFuId() int32 {
	return d.piFuId
}

//进阶日志信息
type PlayerShiHunFanAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.ShiHunFanLogReason
	reasonText        string
}

func CreatePlayerShiHunFanAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.ShiHunFanLogReason, reasonText string) *PlayerShiHunFanAdvancedLogEventData {
	d := &PlayerShiHunFanAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerShiHunFanAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerShiHunFanAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerShiHunFanAdvancedLogEventData) GetReason() commonlog.ShiHunFanLogReason {
	return d.reason
}

func (d *PlayerShiHunFanAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
