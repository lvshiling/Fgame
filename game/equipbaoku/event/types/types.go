package types

import (
	commonlog "fgame/fgame/common/log"
	droptemplate "fgame/fgame/game/drop/template"
	equiptypes "fgame/fgame/game/equipbaoku/types"
	"fmt"
)

type EquipBaoKuEventType string

const (
	EventTypeEquipBaoKuLuckyBox        EquipBaoKuEventType = "EquipBaoKuLuckyBox"         //0点发幸运宝箱事件
	EventTypeEquipBaoKuAttend                              = "EquipBaoKuAttend"           //宝库探索
	EventTypeEquipBaoKuBuy                                 = "EquipBaoKuBuy"              //宝库兑换
	EventTypeEquipBaoKuAttendPointsLog                     = "EquipBaoKuAttendPointsLog " //宝库积分变化日志
	EventTypeEquipBaoKuLuckyPointsLog                      = "EquipBaoKuLuckyPointsLog"   //宝宝库幸运值变化日志
)

//宝库积分变化日志
type PlayerEquipBaoKuAttendPointsLogEventData struct {
	typ        equiptypes.BaoKuType
	beforeNum  int32
	curNum     int32
	itemId     int32
	itemCount  int32
	reason     commonlog.EquipBaoKuLogReason
	reasonText string
}

func CreatePlayerEquipBaoKuAttendPointsLogEventData(beforeNum int32, curNum int32, itemId int32, itemCount int32, reason commonlog.EquipBaoKuLogReason, reasonText string, typ equiptypes.BaoKuType) *PlayerEquipBaoKuAttendPointsLogEventData {
	d := &PlayerEquipBaoKuAttendPointsLogEventData{
		beforeNum:  beforeNum,
		curNum:     curNum,
		itemId:     itemId,
		itemCount:  itemCount,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerEquipBaoKuAttendPointsLogEventData) GetBaoKuType() equiptypes.BaoKuType {
	return d.typ
}

func (d *PlayerEquipBaoKuAttendPointsLogEventData) GetBeforeNum() int32 {
	return d.beforeNum
}

func (d *PlayerEquipBaoKuAttendPointsLogEventData) GetCurNum() int32 {
	return d.curNum
}

func (d *PlayerEquipBaoKuAttendPointsLogEventData) GetItemId() int32 {
	return d.itemId
}

func (d *PlayerEquipBaoKuAttendPointsLogEventData) GetItemCount() int32 {
	return d.itemCount
}

func (d *PlayerEquipBaoKuAttendPointsLogEventData) GetReason() commonlog.EquipBaoKuLogReason {
	return d.reason
}

func (d *PlayerEquipBaoKuAttendPointsLogEventData) GetReasonText() string {
	return d.reasonText
}

//宝库幸运值变化日志
type PlayerEquipBaoKuLuckyPointsLogEventData struct {
	typ        equiptypes.BaoKuType
	beforeNum  int32
	curNum     int32
	withItems  []*droptemplate.DropItemData
	reason     commonlog.EquipBaoKuLogReason
	reasonText string
}

func CreatePlayerEquipBaoKuLuckyPointsLogEventData(beforeNum int32, curNum int32, itemDataArr []*droptemplate.DropItemData, reason commonlog.EquipBaoKuLogReason, reasonText string, typ equiptypes.BaoKuType) *PlayerEquipBaoKuLuckyPointsLogEventData {
	d := &PlayerEquipBaoKuLuckyPointsLogEventData{
		typ:        typ,
		beforeNum:  beforeNum,
		curNum:     curNum,
		withItems:  itemDataArr,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerEquipBaoKuLuckyPointsLogEventData) GetBaoKuType() equiptypes.BaoKuType {
	return d.typ
}

func (d *PlayerEquipBaoKuLuckyPointsLogEventData) GetBeforeNum() int32 {
	return d.beforeNum
}

func (d *PlayerEquipBaoKuLuckyPointsLogEventData) GetCurNum() int32 {
	return d.curNum
}

func (d *PlayerEquipBaoKuLuckyPointsLogEventData) GetWithItems() string {
	rewStr := ""
	for _, item := range d.withItems {
		tempStr := fmt.Sprintf(commonlog.EquipBaoKuLogLuckyPointRewContent.String(), item.ItemId, item.Num)
		rewStr = fmt.Sprint(rewStr, tempStr)
	}
	return rewStr
}

func (d *PlayerEquipBaoKuLuckyPointsLogEventData) GetReason() commonlog.EquipBaoKuLogReason {
	return d.reason
}

func (d *PlayerEquipBaoKuLuckyPointsLogEventData) GetReasonText() string {
	return d.reasonText
}
