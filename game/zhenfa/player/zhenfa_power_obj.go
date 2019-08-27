package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenfa/entity"
	"fmt"
)

//玩家阵法战力对象
type PlayerZhenFaPowerObject struct {
	player     player.Player
	id         int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerZhenFaPowerObject(pl player.Player) *PlayerZhenFaPowerObject {
	pmo := &PlayerZhenFaPowerObject{
		player: pl,
	}
	return pmo
}

func convertZhenFaPowerObjectToEntity(psco *PlayerZhenFaPowerObject) (*entity.PlayerZhenFaPowerEntity, error) {

	e := &entity.PlayerZhenFaPowerEntity{
		Id:         psco.id,
		PlayerId:   psco.player.GetId(),
		Power:      psco.power,
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerZhenFaPowerObject) GetPlayerId() int64 {
	return psco.player.GetId()
}

func (psco *PlayerZhenFaPowerObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerZhenFaPowerObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertZhenFaPowerObjectToEntity(psco)
	return e, err
}

func (psco *PlayerZhenFaPowerObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerZhenFaPowerEntity)

	psco.id = pse.Id
	psco.power = pse.Power
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerZhenFaPowerObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("ZhenFaPower: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
