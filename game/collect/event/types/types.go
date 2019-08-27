package types

import (
	collecttypes "fgame/fgame/game/collect/types"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/scene/scene"
)

type CollectEventType string

const (
	//采集完成
	EventTypeCollectFinish CollectEventType = "CollectFinish"
	//采集完成关联
	EventTypeCollectFinishWith CollectEventType = "CollectFinishWith"
	//采集点信息变化
	EventTypeCollectPointChange CollectEventType = "CollectPointChange"
	//采集物选择完成
	EventTypeCollectChooseFinish CollectEventType = "CollectChooseFinish"
	//密藏采集物完成
	EventTypeCollectMiZangFinish CollectEventType = "CollectMiZangFinish"
)

//采集物选择完成事件信息
type CollectChooseFinishEventData struct {
	npc              scene.NPC
	chooseFinishType collecttypes.CollectChooseFinishType
}

func CreateCollectChooseFinishEventData(n scene.NPC, typ collecttypes.CollectChooseFinishType) *CollectChooseFinishEventData {
	d := &CollectChooseFinishEventData{
		npc:              n,
		chooseFinishType: typ,
	}
	return d
}

func (d *CollectChooseFinishEventData) GetCollectNpc() scene.NPC {
	return d.npc
}

func (d *CollectChooseFinishEventData) GetChooseFinishType() collecttypes.CollectChooseFinishType {
	return d.chooseFinishType
}

//采集物完成关联事件信息
type CollectFinishWithEventData struct {
	npc          scene.NPC
	itemDataList []*droptemplate.DropItemData
}

func CreateCollectFinishWithEventData(n scene.NPC, itemDataList []*droptemplate.DropItemData) *CollectFinishWithEventData {
	d := &CollectFinishWithEventData{
		npc:          n,
		itemDataList: itemDataList,
	}
	return d
}

func (d *CollectFinishWithEventData) GetCollectNpc() scene.NPC {
	return d.npc
}

func (d *CollectFinishWithEventData) GetItemDataList() []*droptemplate.DropItemData {
	return d.itemDataList
}
