package types

type MingGeEventType string

const (
	//命格命理
	EventTypeMingGeMingLi MingGeEventType = "MingGeEventTypeMingLi"
	//命格祭炼
	EventTypeMingGeJiLian MingGeEventType = "MingGeEventTypeJiLian"
)

type PlayerMingGeMingLiEventData struct {
	useItemMap map[int32]int32
}

func CreatePlayerMingGeMingLiEventData(needItemMap map[int32]int32) *PlayerMingGeMingLiEventData {
	d := &PlayerMingGeMingLiEventData{
		useItemMap: needItemMap,
	}
	return d
}

func (data *PlayerMingGeMingLiEventData) GetUseItemMap() map[int32]int32 {
	return data.useItemMap
}

type PlayerMingGeJiLianEventData struct {
	useItemMap map[int32]int32
}

func CreatePlayerMingGeJiLianEventData(needItemMap map[int32]int32) *PlayerMingGeJiLianEventData {
	d := &PlayerMingGeJiLianEventData{
		useItemMap: needItemMap,
	}
	return d
}

func (data *PlayerMingGeJiLianEventData) GetUseItemMap() map[int32]int32 {
	return data.useItemMap
}
