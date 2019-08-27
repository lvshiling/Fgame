package emperor

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	emperorentity "fgame/fgame/game/emperor/entity"
	"fgame/fgame/game/global"
)

//抢龙椅记录数据
type EmperorRecordsObject struct {
	Id          int64
	ServerId    int32
	Type        int32
	EmperorName string
	RobbedName  string
	RobTime     int64
	ItemMap     map[int32]int32
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

//抢夺记录排序
type EmperorRecordsObjectList []*EmperorRecordsObject

func (erol EmperorRecordsObjectList) Len() int {
	return len(erol)
}

func (erol EmperorRecordsObjectList) Less(i, j int) bool {
	return erol[i].RobTime < erol[j].RobTime
}

func (erol EmperorRecordsObjectList) Swap(i, j int) {
	erol[i], erol[j] = erol[j], erol[i]
}

func NewEmperorRecordsObject() *EmperorRecordsObject {
	pso := &EmperorRecordsObject{}
	return pso
}

func (ero *EmperorRecordsObject) GetDBId() int64 {
	return ero.Id
}

func (ero *EmperorRecordsObject) ToEntity() (e storage.Entity, err error) {
	pe := &emperorentity.EmperorRecordsEntity{}
	itemInfoBytes, err := json.Marshal(ero.ItemMap)
	if err != nil {
		return nil, err
	}
	pe.Id = ero.Id
	pe.ServerId = ero.ServerId
	pe.Type = ero.Type
	pe.EmperorName = ero.EmperorName
	pe.RobbedName = ero.RobbedName
	pe.RobTime = ero.RobTime
	pe.ItemInfo = string(itemInfoBytes)
	pe.UpdateTime = ero.UpdateTime
	pe.CreateTime = ero.CreateTime
	pe.DeleteTime = ero.DeleteTime
	e = pe
	return
}

func (ero *EmperorRecordsObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*emperorentity.EmperorRecordsEntity)
	itemMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(pe.ItemInfo), &itemMap); err != nil {
		return err
	}
	ero.Id = pe.Id
	ero.ServerId = pe.ServerId
	ero.Type = pe.Type
	ero.EmperorName = pe.EmperorName
	ero.RobbedName = pe.RobbedName
	ero.RobTime = pe.RobTime
	ero.ItemMap = itemMap
	ero.UpdateTime = pe.UpdateTime
	ero.CreateTime = pe.CreateTime
	ero.DeleteTime = pe.DeleteTime
	return
}

func (ero *EmperorRecordsObject) SetModified() {
	e, err := ero.ToEntity()
	if err != nil {
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
