package player

import (
	"fgame/fgame/core/storage"
	foundentity "fgame/fgame/game/found/entity"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//玩家资源找回对象
type PlayerFoundObject struct {
	player       player.Player
	id           int64
	resType      foundtypes.FoundResourceType
	playModeType foundtypes.PlayModeType
	joinTimes    int32
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func newPlayerFoundObject(pl player.Player) *PlayerFoundObject {
	o := &PlayerFoundObject{
		player: pl,
	}
	return o
}

func convertPlayerFoundObjectToEntity(o *PlayerFoundObject) (e *foundentity.PlayerFoundEntity, err error) {
	e = &foundentity.PlayerFoundEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		PlayModeType: int32(o.playModeType),
		ResType:      int32(o.resType),
		JoinTimes:    o.joinTimes,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *PlayerFoundObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFoundObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFoundObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFoundObjectToEntity(o)
	return e, err
}

func (o *PlayerFoundObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*foundentity.PlayerFoundEntity)

	o.id = te.Id
	o.playModeType = foundtypes.PlayModeType(te.PlayModeType)
	o.resType = foundtypes.FoundResourceType(te.ResType)
	o.joinTimes = te.JoinTimes
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerFoundObject) SetUpdateTime(time int64) {
	o.updateTime = time
}

func (o *PlayerFoundObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Found"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//
func (o *PlayerFoundObject) isFoundBackDay() bool {
	now := global.GetGame().GetTimeService().Now()
	diffDay, _ := timeutils.DiffDay(now, o.updateTime)
	if diffDay == 1 {
		isSameFive, _ := timeutils.IsSameFive(now, o.updateTime)
		if isSameFive {
			return false
		} else {
			return true
		}
	}
	return diffDay > 1
}

//资源找回结果数据
type PlayerFoundBackObject struct {
	id         int64
	player     player.Player
	resType    foundtypes.FoundResourceType
	resLevel   int32
	status     foundtypes.FoundBackStatus
	foundTimes int32
	group      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func (bo *PlayerFoundBackObject) GetResType() foundtypes.FoundResourceType {
	return bo.resType
}

func (bo *PlayerFoundBackObject) GetResLevel() int32 {
	return bo.resLevel
}

func (bo *PlayerFoundBackObject) GetFoundTimes() int32 {
	return bo.foundTimes
}

func (bo *PlayerFoundBackObject) GetFoundStatus() foundtypes.FoundBackStatus {
	return bo.status
}

func (bo *PlayerFoundBackObject) GetGroup() int32 {
	return bo.group
}

func newPlayerFoundBackObject(pl player.Player) *PlayerFoundBackObject {
	bo := &PlayerFoundBackObject{
		player: pl,
	}
	return bo
}

func convertPlayerFoundBackObjectToEntity(o *PlayerFoundBackObject) (e *foundentity.PlayerFoundBackEntity, err error) {
	e = &foundentity.PlayerFoundBackEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		ResType:    int32(o.resType),
		ResLevel:   o.resLevel,
		FoundTimes: o.foundTimes,
		Status:     int32(o.status),
		Group:      o.group,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (bo *PlayerFoundBackObject) GetPlayerId() int64 {
	return bo.player.GetId()
}

func (bo *PlayerFoundBackObject) GetPlayer() player.Player {
	return bo.player
}

func (bo *PlayerFoundBackObject) GetDBId() int64 {
	return bo.id
}

func (bo *PlayerFoundBackObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFoundBackObjectToEntity(bo)
	return e, err
}

func (bo *PlayerFoundBackObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*foundentity.PlayerFoundBackEntity)

	bo.id = te.Id
	bo.resType = foundtypes.FoundResourceType(te.ResType)
	bo.resLevel = te.ResLevel
	bo.foundTimes = te.FoundTimes
	bo.status = foundtypes.FoundBackStatus(te.Status)
	bo.group = te.Group
	bo.updateTime = te.UpdateTime
	bo.createTime = te.CreateTime
	bo.deleteTime = te.DeleteTime
	return nil
}

func (bo *PlayerFoundBackObject) SetModified() {
	e, err := bo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FoundBack"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	bo.player.AddChangedObject(obj)
	return
}
