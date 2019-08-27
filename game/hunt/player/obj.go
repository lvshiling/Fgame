package player

import (
	"fgame/fgame/core/storage"
	huntentity "fgame/fgame/game/hunt/entity"
	hunttypes "fgame/fgame/game/hunt/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"

	"github.com/pkg/errors"
)

//玩家寻宝数据
type PlayerHuntObject struct {
	player         player.Player
	id             int64
	huntType       hunttypes.HuntType
	freeHuntCount  int32
	totalHuntCount int32
	lastHuntTime   int64
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewPlayerHuntObject(pl player.Player) *PlayerHuntObject {
	o := &PlayerHuntObject{
		player: pl,
	}
	return o
}

func convertPlayerHuntObjectToEntity(o *PlayerHuntObject) (*huntentity.PlayerHuntEntity, error) {

	e := &huntentity.PlayerHuntEntity{
		Id:             o.id,
		PlayerId:       o.player.GetId(),
		HuntType:       int32(o.huntType),
		FreeHuntCount:  o.freeHuntCount,
		TotalHuntCount: o.totalHuntCount,
		LastHuntTime:   o.lastHuntTime,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *PlayerHuntObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerHuntObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerHuntObject) GetLastHuntTime() int64 {
	return o.lastHuntTime
}

func (o *PlayerHuntObject) GetFreeHuntCount() int32 {
	return o.freeHuntCount
}

func (o *PlayerHuntObject) GetTotalHuntCount() int32 {
	return o.totalHuntCount
}

func (o *PlayerHuntObject) GetHuntType() hunttypes.HuntType {
	return o.huntType
}

func (o *PlayerHuntObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerHuntObjectToEntity(o)
	return
}

func (o *PlayerHuntObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*huntentity.PlayerHuntEntity)

	o.id = pse.Id
	o.huntType = hunttypes.HuntType(pse.HuntType)
	o.freeHuntCount = pse.FreeHuntCount
	o.totalHuntCount = pse.TotalHuntCount
	o.lastHuntTime = pse.LastHuntTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return
}

func (o *PlayerHuntObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Hunt"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}
