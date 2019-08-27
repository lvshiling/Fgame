package types

import (
	commonlog "fgame/fgame/common/log"
	shenqitypes "fgame/fgame/game/shenqi/types"
)

type ShenQiEventType string

const (
	EventTypeShenQiUpLevel           ShenQiEventType = "ShenQiUpLevel"           //神器升级事件
	EventTypeShenQiRelatedUpLevelLog                 = "ShenQiRelatedUpLevelLog" //神器相关升级事件日志
	EventTypeShenQiUseItem                           = "ShenQiUseItem"           //神器消耗物品
	EventTypeShenQiLingQiNumChanged                  = "ShenQiLingQiNumChanged"  //灵气值变化
)

//神器升级事件信息
type PlayerShenQiUpLevelEventData struct {
	shenQiType shenqitypes.ShenQiType
	oldLevel   int32
	newLevel   int32
}

func CreatePlayerShenQiUpLevelEventData(typ shenqitypes.ShenQiType, oldlev int32, newlev int32) *PlayerShenQiUpLevelEventData {
	d := &PlayerShenQiUpLevelEventData{
		shenQiType: typ,
		oldLevel:   oldlev,
		newLevel:   newlev,
	}
	return d
}

func (d *PlayerShenQiUpLevelEventData) GetShenQiType() shenqitypes.ShenQiType {
	return d.shenQiType
}

func (d *PlayerShenQiUpLevelEventData) GetShenQiOldLevel() int32 {
	return d.oldLevel
}

func (d *PlayerShenQiUpLevelEventData) GetShenQiNewLevel() int32 {
	return d.newLevel
}

//神器相关升级事件日志信息
type PlayerShenQiRelatedUpLevelLogEventData struct {
	befLev     int32
	curLev     int32
	reason     commonlog.ShenQiLogReason
	reasonText string
}

func CreatePlayerShenQiRelatedUpLevelLogEventData(oldlev int32, newlev int32, reason commonlog.ShenQiLogReason, reasonText string) *PlayerShenQiRelatedUpLevelLogEventData {
	d := &PlayerShenQiRelatedUpLevelLogEventData{
		befLev:     oldlev,
		curLev:     newlev,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerShenQiRelatedUpLevelLogEventData) GetBefLev() int32 {
	return d.befLev
}

func (d *PlayerShenQiRelatedUpLevelLogEventData) GetCurLev() int32 {
	return d.curLev
}

func (d *PlayerShenQiRelatedUpLevelLogEventData) GetReason() commonlog.ShenQiLogReason {
	return d.reason
}

func (d *PlayerShenQiRelatedUpLevelLogEventData) GetReasonText() string {
	return d.reasonText
}

type PlayerShenQiUseItemEventData struct {
	itemId  int32
	itemNum int32
}

func CreatePlayerShenQiUseItemEventData(itemId, itemNum int32) *PlayerShenQiUseItemEventData {
	d := &PlayerShenQiUseItemEventData{
		itemId:  itemId,
		itemNum: itemNum,
	}
	return d
}

func (data *PlayerShenQiUseItemEventData) GetUseItemMap() map[int32]int32 {
	d := map[int32]int32{
		data.itemId: data.itemNum,
	}
	return d
}
