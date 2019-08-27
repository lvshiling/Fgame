package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	ringentity "fgame/fgame/game/ring/entity"
	ringtypes "fgame/fgame/game/ring/types"

	"github.com/pkg/errors"
)

type PlayerRingBaoKuObject struct {
	id                    int64
	player                player.Player
	typ                   ringtypes.BaoKuType
	luckyPoints           int32
	attendPoints          int32
	totalAttendTimes      int32
	lastSystemRefreshTime int64
	updateTime            int64
	createTime            int64
	deleteTime            int64
}

func NewRingBaoKuObject(pl player.Player) *PlayerRingBaoKuObject {
	o := &PlayerRingBaoKuObject{
		player: pl,
	}
	return o
}

func convertPlayerRingBaoKuObjectToEntity(o *PlayerRingBaoKuObject) (*ringentity.PlayerRingBaoKuEntity, error) {
	e := &ringentity.PlayerRingBaoKuEntity{
		Id:                    o.id,
		PlayerId:              o.player.GetId(),
		Typ:                   int32(o.typ),
		LuckyPoints:           o.luckyPoints,
		AttendPoints:          o.attendPoints,
		TotalAttendTimes:      o.totalAttendTimes,
		LastSystemRefreshTime: o.lastSystemRefreshTime,
		UpdateTime:            o.updateTime,
		CreateTime:            o.createTime,
		DeleteTime:            o.deleteTime,
	}
	return e, nil
}

func (o *PlayerRingBaoKuObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerRingBaoKuObject) GetLuckyPoints() int32 {
	return o.luckyPoints
}

func (o *PlayerRingBaoKuObject) GetAttendPoints() int32 {
	return o.attendPoints
}

func (o *PlayerRingBaoKuObject) GetType() ringtypes.BaoKuType {
	return o.typ
}

func (o *PlayerRingBaoKuObject) IfEnoughJiFen(num int32) bool {
	if num < 0 {
		return false
	}

	if o.attendPoints >= num {
		return true
	}

	return false
}

func (o *PlayerRingBaoKuObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerRingBaoKuObjectToEntity(o)
	return e, err
}

func (o *PlayerRingBaoKuObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*ringentity.PlayerRingBaoKuEntity)

	o.id = pse.Id
	o.typ = ringtypes.BaoKuType(pse.Typ)
	o.luckyPoints = pse.LuckyPoints
	o.attendPoints = pse.AttendPoints
	o.totalAttendTimes = pse.TotalAttendTimes
	o.lastSystemRefreshTime = pse.LastSystemRefreshTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerRingBaoKuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerRingBaoKu"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
