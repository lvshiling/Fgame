package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/mingge/entity"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//玩家命盘祭炼对象
type PlayerMingGeRefinedObject struct {
	player     player.Player
	id         int64
	subType    minggetypes.MingGeAllSubType
	number     int32
	star       int32
	refinedNum int32
	refinedPro int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerMingGeRefinedObject(pl player.Player) *PlayerMingGeRefinedObject {
	pmo := &PlayerMingGeRefinedObject{
		player: pl,
	}
	return pmo
}

func convertMingGeRefinedObjectToEntity(psco *PlayerMingGeRefinedObject) (*entity.PlayerMingGeRefinedEntity, error) {
	e := &entity.PlayerMingGeRefinedEntity{
		Id:         psco.id,
		PlayerId:   psco.player.GetId(),
		SubType:    int32(psco.subType),
		Number:     psco.number,
		Star:       psco.star,
		RefinedNum: psco.refinedNum,
		RefinedPro: psco.refinedPro,
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerMingGeRefinedObject) GetPlayerId() int64 {
	return psco.player.GetId()
}

func (psco *PlayerMingGeRefinedObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerMingGeRefinedObject) GetSubType() minggetypes.MingGeAllSubType {
	return psco.subType
}

func (psco *PlayerMingGeRefinedObject) GetNumber() int32 {
	return psco.number
}

func (psco *PlayerMingGeRefinedObject) GetStar() int32 {
	return psco.star
}

func (psco *PlayerMingGeRefinedObject) GetRefinedNum() int32 {
	return psco.refinedNum
}

func (psco *PlayerMingGeRefinedObject) GetRefinedPro() int32 {
	return psco.refinedPro
}

func (psco *PlayerMingGeRefinedObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertMingGeRefinedObjectToEntity(psco)
	return e, err
}

func (psco *PlayerMingGeRefinedObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerMingGeRefinedEntity)

	psco.id = pse.Id
	psco.subType = minggetypes.MingGeAllSubType(pse.SubType)
	psco.number = pse.Number
	psco.star = pse.Star
	psco.refinedNum = pse.RefinedNum
	psco.refinedPro = pse.RefinedPro
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerMingGeRefinedObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("minggepan: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
