package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	songbutingentity "fgame/fgame/game/songbuting/entity"
	"fmt"
)

//玩家元宝送不停
type PlayerSongBuTingObject struct {
	player     player.Player
	id         int64
	playerId   int64
	isReceive  int32
	times      int32
	lastTime   int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerSongBuTingObject(pl player.Player) *PlayerSongBuTingObject {
	pmo := &PlayerSongBuTingObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerObjectToEntity(pso *PlayerSongBuTingObject) (*songbutingentity.PlayerSongBuTingEntity, error) {
	e := &songbutingentity.PlayerSongBuTingEntity{
		Id:         pso.id,
		PlayerId:   pso.playerId,
		IsReceive:  pso.isReceive,
		Times:      pso.times,
		LastTime:   pso.lastTime,
		UpdateTime: pso.updateTime,
		CreateTime: pso.createTime,
		DeleteTime: pso.deleteTime,
	}
	return e, nil
}

func (pso *PlayerSongBuTingObject) GetPlayerId() int64 {
	return pso.playerId
}

func (pso *PlayerSongBuTingObject) GetDBId() int64 {
	return pso.id
}

func (pso *PlayerSongBuTingObject) GetIsReceive() bool {
	return pso.isReceive == 1
}

func (pso *PlayerSongBuTingObject) GetTimes() int32 {
	return pso.times
}

func (pso *PlayerSongBuTingObject) GetLastTime() int64 {
	return pso.lastTime
}

func (pso *PlayerSongBuTingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerObjectToEntity(pso)
	return e, err
}

func (pso *PlayerSongBuTingObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*songbutingentity.PlayerSongBuTingEntity)

	pso.id = pse.Id
	pso.playerId = pse.PlayerId
	pso.times = pse.Times
	pso.isReceive = pse.IsReceive
	pso.lastTime = pse.LastTime
	pso.updateTime = pse.UpdateTime
	pso.createTime = pse.CreateTime
	pso.deleteTime = pse.DeleteTime
	return nil
}

func (pso *PlayerSongBuTingObject) SetModified() {
	e, err := pso.ToEntity()
	if err != nil {
		panic(fmt.Errorf("songbuting: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pso.player.AddChangedObject(obj)
	return
}
