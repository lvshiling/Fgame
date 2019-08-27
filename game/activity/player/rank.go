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

//活动排行数据
type PlayerActivityRankObject struct {
	player       player.Player
	id           int64
	activityType activitytypes.ActivityType
	rankMap      map[int32]int64
	endTime      int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func (o *PlayerActivityRankObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerActivityRankObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerActivityRankObject) FromEntity(e storage.Entity) error {
	entity := e.(*entity.PlayerActivityRankEntity)

	o.id = entity.Id
	o.activityType = activitytypes.ActivityType(entity.ActivityType)
	rankMap := make(map[int32]int64)
	err := json.Unmarshal([]byte(entity.RankMap), &rankMap)
	if err != nil {
		return err
	}
	o.rankMap = rankMap
	o.endTime = entity.EndTime
	o.updateTime = entity.UpdateTime
	o.createTime = entity.CreateTime
	o.deleteTime = entity.DeleteTime
	return nil
}

func (o *PlayerActivityRankObject) ToEntity() (rankEntity storage.Entity, err error) {
	rankMap, err := json.Marshal(o.rankMap)
	if err != nil {
		return
	}
	rankEntity = &entity.PlayerActivityRankEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		ActivityType: int32(o.activityType),
		RankMap:      string(rankMap),
		EndTime:      o.endTime,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return
}

func (o *PlayerActivityRankObject) SetModified() {

	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ActivityRank"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("activity:转换应该成功"))
	}

	o.player.AddChangedObject(obj)

}

func CreatePlayerActivityRankObject(pl player.Player) *PlayerActivityRankObject {
	newObj := &PlayerActivityRankObject{
		player: pl,
	}
	return newObj
}
