package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	wardrobeentity "fgame/fgame/game/wardrobe/entity"
	"fmt"
)

//玩家衣橱培养对象
type PlayerWardrobePeiYangObject struct {
	player     player.Player
	id         int64
	playerId   int64
	typ        int32
	peiYangNum int32
	updateTime int64
	createTime int64
	deleteTime int64 
}

func NewPlayerWardrobePeiYangObject(pl player.Player) *PlayerWardrobePeiYangObject {
	pmo := &PlayerWardrobePeiYangObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerPeiYangObjectToEntity(o *PlayerWardrobePeiYangObject) (*wardrobeentity.PlayerWardrobePeiYangEntity, error) {
	e := &wardrobeentity.PlayerWardrobePeiYangEntity{
		Id:         o.id,
		PlayerId:   o.playerId,
		Type:       int32(o.typ),
		PeiYangNum: o.peiYangNum,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerWardrobePeiYangObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *PlayerWardrobePeiYangObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerWardrobePeiYangObject) GetType() int32 {
	return o.typ
}

func (o *PlayerWardrobePeiYangObject) GetPeiYangNum() int32 {
	return o.peiYangNum
}

func (o *PlayerWardrobePeiYangObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerPeiYangObjectToEntity(o)
	return e, err
}

func (o *PlayerWardrobePeiYangObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*wardrobeentity.PlayerWardrobePeiYangEntity)

	o.id = pse.Id
	o.playerId = pse.PlayerId
	o.typ = int32(pse.Type)
	o.peiYangNum = pse.PeiYangNum
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerWardrobePeiYangObject) SetModified() {
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
