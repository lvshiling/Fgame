package types

import (
	commonlog "fgame/fgame/common/log"
	housetypes "fgame/fgame/game/house/types"
)

type HouseEventType string

const (
	EventTypeHouseUplevel        HouseEventType = "HouseUplevel"        //房子升级事件
	EventTypeHouseActivate                      = "HouseActivate"       //房子激活事件
	EventTypeHouseSell                          = "HouseSell"           //房子出售事件
	EventTypeHouseBeforeCrossDay                = "HouseBeforeCrossDay" //房子跨天前
	EventTypeHouseUplevelLog                    = "HouseUplevelLog"     //房子升级日志
)

//
type PlayerHouseUplevelLogEventData struct {
	beforeUplevelNum int32
	advancedNum      int32
	reason           commonlog.HouseLogReason
	reasonText       string
}

func CreatePlayerHouseUplevelLogEventData(beforeUplevelNum, advancedNum int32, reason commonlog.HouseLogReason, reasonText string) *PlayerHouseUplevelLogEventData {
	d := &PlayerHouseUplevelLogEventData{
		beforeUplevelNum: beforeUplevelNum,
		advancedNum:      advancedNum,
		reason:           reason,
		reasonText:       reasonText,
	}
	return d
}

func (d *PlayerHouseUplevelLogEventData) GetBeforeUplevelNum() int32 {
	return d.beforeUplevelNum
}

func (d *PlayerHouseUplevelLogEventData) GetUplevelNum() int32 {
	return d.advancedNum
}

func (d *PlayerHouseUplevelLogEventData) GetReason() commonlog.HouseLogReason {
	return d.reason
}

func (d *PlayerHouseUplevelLogEventData) GetReasonText() string {
	return d.reasonText
}

//
type PlayerHouseSellEventData struct {
	houseType  housetypes.HouseType
	houseLevel int32
	houseIndex int32
}

func CreatePlayerHouseSellEventData(houseType housetypes.HouseType, houseLevel, houseIndex int32) *PlayerHouseSellEventData {
	d := &PlayerHouseSellEventData{
		houseType:  houseType,
		houseLevel: houseLevel,
		houseIndex: houseIndex,
	}
	return d
}

func (d *PlayerHouseSellEventData) GetHouseType() housetypes.HouseType {
	return d.houseType
}

func (d *PlayerHouseSellEventData) GetHouseLevel() int32 {
	return d.houseLevel
}

func (d *PlayerHouseSellEventData) GetHouseIndex() int32 {
	return d.houseIndex
}
