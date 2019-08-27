package player

import (
	"fgame/fgame/core/storage"
	densewatentity "fgame/fgame/game/densewat/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//金银密窟对象
type PlayerDenseWatObject struct {
	player     player.Player
	id         int64
	playerId   int64
	num        int32
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerDenseWatObject(pl player.Player) *PlayerDenseWatObject {
	pwo := &PlayerDenseWatObject{
		player: pl,
	}
	return pwo
}

func convertNewPlayerDenseWatObjectToEntity(pwo *PlayerDenseWatObject) (*densewatentity.PlayerDenseWatEntity, error) {

	e := &densewatentity.PlayerDenseWatEntity{
		Id:         pwo.id,
		PlayerId:   pwo.playerId,
		Num:        pwo.num,
		EndTime:    pwo.endTime,
		UpdateTime: pwo.updateTime,
		CreateTime: pwo.createTime,
		DeleteTime: pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerDenseWatObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerDenseWatObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerDenseWatObject) GetNum() int32 {
	return pwo.num
}

func (pwo *PlayerDenseWatObject) GetEndTime() int64 {
	return pwo.endTime
}

func (pwo *PlayerDenseWatObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerDenseWatObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerDenseWatObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*densewatentity.PlayerDenseWatEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.num = pse.Num
	pwo.endTime = pse.EndTime
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerDenseWatObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("densewat: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}
