package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobeentity "fgame/fgame/game/wardrobe/entity"
	"fmt"
)

//玩家衣橱对象
type PlayerWardrobeObject struct {
	player     player.Player
	id         int64
	playerId   int64
	typ        int32
	subType    int32
	activeFlag int32
	permanent  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerWardrobeObject(pl player.Player) *PlayerWardrobeObject {
	pmo := &PlayerWardrobeObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerObjectToEntity(o *PlayerWardrobeObject) (*wardrobeentity.PlayerWardrobeEntity, error) {
	e := &wardrobeentity.PlayerWardrobeEntity{
		Id:         o.id,
		PlayerId:   o.playerId,
		Type:       o.typ,
		SubType:    o.subType,
		ActiveFlag: o.activeFlag,
		Permanent:  o.permanent,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerWardrobeObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *PlayerWardrobeObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerWardrobeObject) GetType() int32 {
	return o.typ
}

func (o *PlayerWardrobeObject) GetSubType() int32 {
	return o.subType
}

func (o *PlayerWardrobeObject) GetIsActive() bool {
	return o.activeFlag == 1
}

func (o *PlayerWardrobeObject) GetIsPermanent() bool {
	return o.permanent == 1
}

func (o *PlayerWardrobeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerObjectToEntity(o)
	return e, err
}

func (o *PlayerWardrobeObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*wardrobeentity.PlayerWardrobeEntity)

	o.id = pse.Id
	o.playerId = pse.PlayerId
	o.typ = pse.Type
	subType := pse.SubType
	o.subType = subType
	o.activeFlag = pse.ActiveFlag
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerWardrobeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("wardrobe: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
