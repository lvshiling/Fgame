package outlandboss

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	outlandbossentity "fgame/fgame/game/outlandboss/entity"

	"github.com/pkg/errors"
)

//外域boss掉落记录列表对象
type OutlandBossDropRecordsObject struct {
	id         int64
	serverId   int32
	killerName string
	biologyId  int32
	mapId      int32
	dropTime   int64
	ItemMap    map[int32]int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewOutlandBossDropRecordsObject() *OutlandBossDropRecordsObject {
	o := &OutlandBossDropRecordsObject{}
	return o
}

func convertNewOutlandBossDropRecordsObjectToEntity(o *OutlandBossDropRecordsObject) (*outlandbossentity.OutlandBossDropRecordsEntity, error) {
	e := &outlandbossentity.OutlandBossDropRecordsEntity{}
	itemInfoBytes, err := json.Marshal(o.ItemMap)
	if err != nil {
		return nil, err
	}
	e.Id = o.id
	e.ServerId = o.serverId
	e.KillerName = o.killerName
	e.BiologyId = o.biologyId
	e.MapId = o.mapId
	e.DropTime = o.dropTime
	e.ItemInfo = string(itemInfoBytes)
	e.UpdateTime = o.updateTime
	e.CreateTime = o.createTime
	e.DeleteTime = o.deleteTime
	return e, nil
}

func (o *OutlandBossDropRecordsObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *OutlandBossDropRecordsObject) GetKillerName() string {
	return o.killerName
}

func (o *OutlandBossDropRecordsObject) GetBiologyId() int32 {
	return o.biologyId
}

func (o *OutlandBossDropRecordsObject) GetMapId() int32 {
	return o.mapId
}

func (o *OutlandBossDropRecordsObject) GetDropTime() int64 {
	return o.dropTime
}

func (o *OutlandBossDropRecordsObject) GetItemInfo() map[int32]int32 {
	return o.ItemMap
}

func (o *OutlandBossDropRecordsObject) GetDBId() int64 {
	return o.id
}

func (o *OutlandBossDropRecordsObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewOutlandBossDropRecordsObjectToEntity(o)
	return e, err
}

func (o *OutlandBossDropRecordsObject) FromEntity(e storage.Entity) error {

	pse, _ := e.(*outlandbossentity.OutlandBossDropRecordsEntity)
	itemMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(pse.ItemInfo), &itemMap); err != nil {
		return err
	}
	o.id = pse.Id
	o.serverId = pse.ServerId
	o.killerName = pse.KillerName
	o.biologyId = pse.BiologyId
	o.mapId = pse.MapId
	o.dropTime = pse.DropTime
	o.ItemMap = itemMap
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *OutlandBossDropRecordsObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OutlandBossDropRecords"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
