package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/activity/entity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"

	"github.com/pkg/errors"
)

//活动采集数据
type PlayerActivityCollectObject struct {
	player       player.Player
	id           int64
	activityType activitytypes.ActivityType
	countMap     map[int32]int32
	endTime      int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func (o *PlayerActivityCollectObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerActivityCollectObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerActivityCollectObject) FromEntity(e storage.Entity) error {
	entity := e.(*entity.PlayerActivityCollectEntity)
	countMap := make(map[int32]int32)
	err := json.Unmarshal([]byte(entity.CountMap), &countMap)
	if err != nil {
		return err
	}

	o.id = entity.Id
	o.activityType = activitytypes.ActivityType(entity.ActivityType)
	o.countMap = countMap
	o.endTime = entity.EndTime
	o.updateTime = entity.UpdateTime
	o.createTime = entity.CreateTime
	o.deleteTime = entity.DeleteTime
	return nil
}

func (o *PlayerActivityCollectObject) ToEntity() (rankEntity storage.Entity, err error) {
	countMap, err := json.Marshal(o.countMap)
	if err != nil {
		return
	}
	rankEntity = &entity.PlayerActivityCollectEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		ActivityType: int32(o.activityType),
		CountMap:     string(countMap),
		EndTime:      o.endTime,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return
}

func (o *PlayerActivityCollectObject) SetModified() {

	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ActivityCollect"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("activity:转换应该成功"))
	}

	o.player.AddChangedObject(obj)

}

func CreatePlayerActivityCollectObject(pl player.Player) *PlayerActivityCollectObject {
	newObj := &PlayerActivityCollectObject{
		player: pl,
	}
	return newObj
}
