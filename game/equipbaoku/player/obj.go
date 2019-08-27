package player

import (
	"fgame/fgame/core/storage"
	equipbaokuentity "fgame/fgame/game/equipbaoku/entity"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//装备宝库对象
type PlayerEquipBaoKuObject struct {
	player                player.Player
	id                    int64
	typ                   equipbaokutypes.BaoKuType
	luckyPoints           int32
	attendPoints          int32
	totalAttendTimes      int32
	lastSystemRefreshTime int64
	updateTime            int64
	createTime            int64
	deleteTime            int64
}

func NewPlayerEquipBaoKuObject(pl player.Player) *PlayerEquipBaoKuObject {
	o := &PlayerEquipBaoKuObject{
		player: pl,
	}
	return o
}

func convertNewPlayerEquipBaoKuObjectToEntity(o *PlayerEquipBaoKuObject) (*equipbaokuentity.PlayerEquipBaoKuEntity, error) {
	e := &equipbaokuentity.PlayerEquipBaoKuEntity{
		Id:                    o.id,
		PlayerId:              o.GetPlayerId(),
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

func (o *PlayerEquipBaoKuObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerEquipBaoKuObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerEquipBaoKuObject) GetBaoKuType() equipbaokutypes.BaoKuType {
	return o.typ
}

func (o *PlayerEquipBaoKuObject) GetLuckyPoints() int32 {
	return o.luckyPoints
}

func (o *PlayerEquipBaoKuObject) GetAttendPoints() int32 {
	return o.attendPoints
}

func (o *PlayerEquipBaoKuObject) GetToatalAttendTimes() int32 {
	return o.totalAttendTimes
}

func (o *PlayerEquipBaoKuObject) GetLastSystemRefreshTime() int64 {
	return o.lastSystemRefreshTime
}

func (o *PlayerEquipBaoKuObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerEquipBaoKuObjectToEntity(o)
	return e, err
}

func (o *PlayerEquipBaoKuObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*equipbaokuentity.PlayerEquipBaoKuEntity)

	o.id = pse.Id
	o.typ = equipbaokutypes.BaoKuType(pse.Typ)
	o.luckyPoints = pse.LuckyPoints
	o.attendPoints = pse.AttendPoints
	o.totalAttendTimes = pse.TotalAttendTimes
	o.lastSystemRefreshTime = pse.LastSystemRefreshTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerEquipBaoKuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "EquipBaoKu"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
