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
type PlayerActivityObject struct {
	player       player.Player
	id           int64
	activityType activitytypes.ActivityType
	attendTimes  int32
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func (o *PlayerActivityObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerActivityObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerActivityObject) FromEntity(e storage.Entity) error {
	entity := e.(*entity.PlayerActivityEntity)

	o.id = entity.Id
	o.activityType = activitytypes.ActivityType(entity.ActivityType)
	o.attendTimes = entity.AttendTimes
	o.updateTime = entity.UpdateTime
	o.createTime = entity.CreateTime
	o.deleteTime = entity.DeleteTime
	return nil
}
func (o *PlayerActivityObject) ToEntity() (mailEntity storage.Entity) {

	mailEntity = &entity.PlayerActivityEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		ActivityType: int32(o.activityType),
		AttendTimes:  o.attendTimes,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return mailEntity
}

func (o *PlayerActivityObject) SetModified() {
	e := o.ToEntity()

	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("activity:转换应该成功"))
	}

	o.player.AddChangedObject(obj)

}

func (o *PlayerActivityObject) GetAttendTimes() int32 {
	return o.attendTimes
}

func CreatePlayerActivityObject(pl player.Player) *PlayerActivityObject {
	newObj := &PlayerActivityObject{
		player: pl,
	}
	return newObj
}
