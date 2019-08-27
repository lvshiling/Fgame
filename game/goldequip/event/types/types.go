package types

import (
	commonlog "fgame/fgame/common/log"
	inventorytypes "fgame/fgame/game/inventory/types"
)

type GoldEquipEventType string

const (
	EventTypeGoldEquipPutOn               GoldEquipEventType = "GoldEquipPutOn"               // 元神金装穿戴
	EventTypeGoldEquipTakeOff                                = "GoldEquipTakeOff"             // 元神金装卸下
	EventTypeGoldEquipStatusWhenTakeOff                      = "GoldEquipStatusWhenTakeOff"   // 元神金装卸下时候的信息
	EventTypeGoldEquipStrengSuccess                          = "GoldEquipStrengSuccess"       // 元神金装强化成功
	EventTypeGoldEquipStrengUpstarSuccess                    = "GoldEquipStrengUpstarSuccess" // 元神金装强化升星成功
	EventTypeGoldEquipResolve                                = "GoldEquipResolve"             // 元神金装分解
	EventTypeGoldEquipEmbedGem                               = "GoldEquipEmbedGem"            // 元神金装镶嵌宝石
	EventTypeGoldEquipTakeOffGem                             = "GoldEquipTakeOffGem"          // 元神金装脱下宝石
	EventTypeGoldEquipUseGemAll                              = "GoldEquipUseGemAll"           // 元神金装一键镶嵌
	EventTypeGoldEquipExtendLog                              = "GoldEquipExtendLog"           // 元神金装继承日志

)

//继承日志
type PlayerGoldEquipExtendLogEventData struct {
	beforeLevel int32
	afterLevel  int32
	reason      commonlog.GoldEquipLogReason
	reasonText  string
}

func CreatePlayerGoldEquipExtendLogEventData(beforeLevel, afterLevel int32, reason commonlog.GoldEquipLogReason, reasonText string) *PlayerGoldEquipExtendLogEventData {
	d := &PlayerGoldEquipExtendLogEventData{
		beforeLevel: beforeLevel,
		afterLevel:  afterLevel,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerGoldEquipExtendLogEventData) GetBeforeLevel() int32 {
	return d.beforeLevel
}

func (d *PlayerGoldEquipExtendLogEventData) GetAfterLevel() int32 {
	return d.afterLevel
}

func (d *PlayerGoldEquipExtendLogEventData) GetReason() commonlog.GoldEquipLogReason {
	return d.reason
}

func (d *PlayerGoldEquipExtendLogEventData) GetReasonText() string {
	return d.reasonText
}

//元神金装信息
type PlayerGoldEquipStatusEventData struct {
	bodyPos         inventorytypes.BodyPositionType
	openlightLevel  int32
	strengthenLevel int32
	upstarLevel     int32
}

func CreatePlayerGoldEquipStatusEventData(bodyPos inventorytypes.BodyPositionType, openlightLevel, strengthenLevel, upstarLevel int32) *PlayerGoldEquipStatusEventData {
	d := &PlayerGoldEquipStatusEventData{
		bodyPos:         bodyPos,
		openlightLevel:  openlightLevel,
		strengthenLevel: strengthenLevel,
		upstarLevel:     upstarLevel,
	}
	return d
}

func (d *PlayerGoldEquipStatusEventData) GetBodyPos() inventorytypes.BodyPositionType {
	return d.bodyPos
}

func (d *PlayerGoldEquipStatusEventData) GetOpenlightLevel() int32 {
	return d.openlightLevel
}

func (d *PlayerGoldEquipStatusEventData) GetStrengthenLevel() int32 {
	return d.strengthenLevel
}

func (d *PlayerGoldEquipStatusEventData) GetUpstarLevel() int32 {
	return d.upstarLevel
}
