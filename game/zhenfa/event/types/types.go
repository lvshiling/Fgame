package types

type ZhenFaEventType string

const (
	EventTypeZhenFaXianHuoShengJiUseItem ZhenFaEventType = "ZhenFaXianHuoShengJiUseItem" //玉玺之战结束
)

type PlayerZhenFaXianHuoShengJiUseItemEventData struct {
	itemMap map[int32]int32
}

func CreatePlayerZhenFaXianHuoShengJiUseItemEventData(itemMap map[int32]int32) *PlayerZhenFaXianHuoShengJiUseItemEventData {
	d := &PlayerZhenFaXianHuoShengJiUseItemEventData{
		itemMap: itemMap,
	}
	return d
}

func (data *PlayerZhenFaXianHuoShengJiUseItemEventData) GetItemMap() map[int32]int32 {
	return data.itemMap
}
