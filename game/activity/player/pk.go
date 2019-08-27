package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/activity/entity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//活动对象
type PlayerActivityPkObject struct {
	player         player.Player
	id             int64
	activityType   activitytypes.ActivityType
	killedNum      int32
	lastKilledTime int64
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func (o *PlayerActivityPkObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerActivityPkObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerActivityPkObject) FromEntity(e storage.Entity) error {
	entity := e.(*entity.PlayerActivityPkEntity)

	o.id = entity.Id
	o.activityType = activitytypes.ActivityType(entity.ActivityType)
	o.killedNum = entity.KilledNum
	o.lastKilledTime = entity.LastKilledTime

	o.updateTime = entity.UpdateTime
	o.createTime = entity.CreateTime
	o.deleteTime = entity.DeleteTime
	return nil
}
func (o *PlayerActivityPkObject) ToEntity() (pkEntity storage.Entity) {

	pkEntity = &entity.PlayerActivityPkEntity{
		Id:             o.id,
		PlayerId:       o.player.GetId(),
		ActivityType:   int32(o.activityType),
		KilledNum:      o.killedNum,
		LastKilledTime: o.lastKilledTime,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return pkEntity
}

func (o *PlayerActivityPkObject) SetModified() {
	e := o.ToEntity()

	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("activity:转换应该成功"))
	}

	o.player.AddChangedObject(obj)

}

func (o *PlayerActivityPkObject) GetKilledNum() int32 {
	return o.killedNum
}

func (o *PlayerActivityPkObject) GetLastKilledTime() int64 {
	return o.lastKilledTime
}
func CreatePlayerActivityPkObject(pl player.Player) *PlayerActivityPkObject {
	newObj := &PlayerActivityPkObject{
		player: pl,
	}
	return newObj
}
