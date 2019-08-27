package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xianzuncardentity "fgame/fgame/game/xianzuncard/entity"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"

	"github.com/pkg/errors"
)

type PlayerXianZunCardObject struct {
	id           int64
	player       player.Player
	cardType     xianzuncardtypes.XianZunCardType
	isActivite   int32
	isReceive    int32
	activiteTime int64
	receiveTime  int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewXianZunCardObject(pl player.Player) *PlayerXianZunCardObject {
	o := &PlayerXianZunCardObject{
		player: pl,
	}
	return o
}

func convertPlayerXianZunCardObjectToEntity(o *PlayerXianZunCardObject) (*xianzuncardentity.PlayerXianZunCardEntity, error) {
	e := &xianzuncardentity.PlayerXianZunCardEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		Typ:          int32(o.cardType),
		IsActivite:   o.isActivite,
		IsReceive:    o.isReceive,
		ActiviteTime: o.activiteTime,
		ReceiveTime:  o.receiveTime,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *PlayerXianZunCardObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerXianZunCardObject) IsActivite() bool {
	return o.isActivite != 0
}

func (o *PlayerXianZunCardObject) IsReceive() bool {
	return o.isReceive != 0
}

func (o *PlayerXianZunCardObject) GetActiviteTime() int64 {
	return o.activiteTime
}

func (o *PlayerXianZunCardObject) GetIsReceive() int32 {
	return o.isReceive
}

func (o *PlayerXianZunCardObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerXianZunCardObjectToEntity(o)
	return e, err
}

func (o *PlayerXianZunCardObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*xianzuncardentity.PlayerXianZunCardEntity)

	o.id = pse.Id
	o.cardType = xianzuncardtypes.XianZunCardType(pse.Typ)
	o.isActivite = pse.IsActivite
	o.isReceive = pse.IsReceive
	o.activiteTime = pse.ActiviteTime
	o.receiveTime = pse.ReceiveTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerXianZunCardObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerXianZunCard"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)

	return
}
