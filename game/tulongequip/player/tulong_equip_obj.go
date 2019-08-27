package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tulongequipentity "fgame/fgame/game/tulongequip/entity"
	"fmt"

	"github.com/pkg/errors"
)

//玩家套装技能数据
type PlayerTuLongEquipObject struct {
	player     player.Player
	id         int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerTuLongEquipObject(pl player.Player) *PlayerTuLongEquipObject {
	o := &PlayerTuLongEquipObject{
		player: pl,
	}
	return o
}

func convertPlayerTuLongEquipObjectToEntity(o *PlayerTuLongEquipObject) (*tulongequipentity.PlayerTuLongEquipEntity, error) {

	e := &tulongequipentity.PlayerTuLongEquipEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Power:      o.power,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerTuLongEquipObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerTuLongEquipObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTuLongEquipObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerTuLongEquipObjectToEntity(o)
	return
}

func (o *PlayerTuLongEquipObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*tulongequipentity.PlayerTuLongEquipEntity)

	o.id = pse.Id
	o.power = pse.Power
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return
}

func (o *PlayerTuLongEquipObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TuLongEquip"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}
