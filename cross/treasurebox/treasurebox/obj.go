package treasurebox

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	treasureboxentity "fgame/fgame/cross/treasurebox/entity"
	"fgame/fgame/game/global"
)

//跨服宝箱日志数据
type TreasureBoxLogObject struct {
	Id         int64
	AreaId     int32
	ServerId   int32
	PlayerName string
	ItemMap    map[int32]int32
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

//记录排序
type TreasureBoxLogObjectList []*TreasureBoxLogObject

func (adl TreasureBoxLogObjectList) Len() int {
	return len(adl)
}

func (adl TreasureBoxLogObjectList) Less(i, j int) bool {
	return adl[i].LastTime < adl[j].LastTime
}

func (adl TreasureBoxLogObjectList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

func NewTreasureBoxLogObject() *TreasureBoxLogObject {
	poo := &TreasureBoxLogObject{}
	return poo
}

func (tlro *TreasureBoxLogObject) GetDBId() int64 {
	return tlro.Id
}

func (oo *TreasureBoxLogObject) ToEntity() (e storage.Entity, err error) {

	itemInfoBytes, err := json.Marshal(oo.ItemMap)
	if err != nil {
		return nil, err
	}

	oe := &treasureboxentity.TreasureBoxLogEntity{}
	oe.Id = oo.Id
	oe.AreaId = oo.AreaId
	oe.ServerId = oo.ServerId
	oe.PlayerName = oo.PlayerName
	oe.ItemInfo = string(itemInfoBytes)
	oe.LastTime = oo.LastTime
	oe.UpdateTime = oo.UpdateTime
	oe.CreateTime = oo.CreateTime
	oe.DeleteTime = oo.DeleteTime
	e = oe
	return
}

func (oo *TreasureBoxLogObject) FromEntity(e storage.Entity) (err error) {
	oe, _ := e.(*treasureboxentity.TreasureBoxLogEntity)
	itemMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(oe.ItemInfo), &itemMap); err != nil {
		return err
	}

	oo.Id = oe.Id
	oo.AreaId = oe.AreaId
	oo.ServerId = oe.ServerId
	oo.PlayerName = oe.PlayerName
	oo.ItemMap = itemMap
	oo.LastTime = oe.LastTime
	oo.UpdateTime = oe.UpdateTime
	oo.CreateTime = oe.CreateTime
	oo.DeleteTime = oe.DeleteTime
	return
}

func (oo *TreasureBoxLogObject) SetModified() {
	e, err := oo.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
