package types

type WushuangEventType string

const (
	//无双神器吞噬强化
	EventTypeWushuangDevouring WushuangEventType = "WushuangEventTypeDevouring"
)

type PlayerWushuangDevouringEventData struct {
	useItemId  int32
	useItemNum int32
}

func (data *PlayerWushuangDevouringEventData) GetUseItemMap() map[int32]int32 {
	d := map[int32]int32{
		data.useItemId: data.useItemNum,
	}
	return d
}

func CreatePlayerWushuangDevouringEventData(useItemId, useItemNum int32) *PlayerWushuangDevouringEventData {
	d := &PlayerWushuangDevouringEventData{
		useItemId:  useItemId,
		useItemNum: useItemNum,
	}
	return d
}
