package types

type EventTypeRing string

const (
	EventTypeRingFuseChange         EventTypeRing = "玩家特戒融合等级变化"
	EventTypeRingLuckyPointsChange                = "宝库幸运值变化"
	EventTypeRingAttendPointsChange               = "宝库积分变化"
)

type PlayerRingFuseChangeEventData struct {
	lastItemId int32
	curItemId  int32
}

func CreatePlayerRingFuseChangeEventData(lastItemId int32, curItemId int32) *PlayerRingFuseChangeEventData {
	data := &PlayerRingFuseChangeEventData{
		lastItemId: lastItemId,
		curItemId:  curItemId,
	}
	return data
}

func (d *PlayerRingFuseChangeEventData) GetLastItemId() int32 {
	return d.lastItemId
}

func (d *PlayerRingFuseChangeEventData) GetCurItemId() int32 {
	return d.curItemId
}
