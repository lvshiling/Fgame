package types

type FourGodEventType string

const (
	//玩家进入四神遗迹
	EventTypeFourGodPlayerEnter FourGodEventType = "FourGodPlayerEnter"
	//玩家钥匙改变
	EventTypeFourGodKeyChange FourGodEventType = "FourGodPlayerKeyChange"
	//玩家退出四神遗迹
	EventTypeFourGodPlayerExit FourGodEventType = "FourGodPlayerExit"
	//副本生物改变
	EventTypeFourGodBioChange FourGodEventType = "FourGodBioChange"
	//玩家四神遗迹获得物品
	EventTypeFourGodGetItem FourGodEventType = "FourGodPlayerGetItem"
	//四神遗迹活动结束
	EventTypeFourGodSceneFinish FourGodEventType = "FourGodSceneFinish"
	//四神遗迹宝箱采集完成
	EventTypeFourGodCollectBoxFinish FourGodEventType = "FourGodCollectBoxFinish"
)

type FourGodItemGetEventData struct {
	itemId int32
	num    int32
	level  int32
}

func (d *FourGodItemGetEventData) GetItemId() int32 {
	return d.itemId
}

func (d *FourGodItemGetEventData) GetNum() int32 {
	return d.num
}

func (d *FourGodItemGetEventData) GetLevel() int32 {
	return d.level
}

func CreateFourGodItemGetEventData(itemId int32, num int32, level int32) *FourGodItemGetEventData {
	d := &FourGodItemGetEventData{
		itemId: itemId,
		num:    num,
		level:  level,
	}
	return d
}
