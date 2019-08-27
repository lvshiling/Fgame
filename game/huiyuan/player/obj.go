package player

import (
	"fgame/fgame/core/storage"
	huiyuanentity "fgame/fgame/game/huiyuan/entity"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//会员
type PlayerHuiYuanObject struct {
	player                 player.Player
	id                     int64
	level                  int32
	huiyuanType            huiyuantypes.HuiYuanType
	lastReceiveTime        int64
	lastInterimReceiveTime int64
	plusBuyTime            int64
	interimBuyTime         int64
	expireTime             int64
	updateTime             int64
	createTime             int64
	deleteTime             int64
}

func newPlayerHuiYuanObject(pl player.Player) *PlayerHuiYuanObject {
	o := &PlayerHuiYuanObject{
		player: pl,
	}
	return o
}

func convertPlayerHuiYuanObjectToEntity(o *PlayerHuiYuanObject) (e *huiyuanentity.PlayerHuiYuanEntity, err error) {

	e = &huiyuanentity.PlayerHuiYuanEntity{
		Id:                     o.id,
		PlayerId:               o.player.GetId(),
		Level:                  o.level,
		Type:                   int32(o.huiyuanType),
		LastReceiveTime:        o.lastReceiveTime,
		LastInterimReceiveTime: o.lastInterimReceiveTime,
		PlusBuyTime:            o.plusBuyTime,
		InterimBuyTime:         o.interimBuyTime,
		ExpireTime:             o.expireTime,
		UpdateTime:             o.updateTime,
		CreateTime:             o.createTime,
		DeleteTime:             o.deleteTime,
	}
	return e, nil
}

func (o *PlayerHuiYuanObject) GetExpireTime() int64 {
	return o.expireTime
}

func (o *PlayerHuiYuanObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerHuiYuanObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerHuiYuanObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerHuiYuanObjectToEntity(o)
	return e, err
}

func (o *PlayerHuiYuanObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*huiyuanentity.PlayerHuiYuanEntity)

	o.id = te.Id
	o.level = te.Level
	o.huiyuanType = huiyuantypes.HuiYuanType(te.Type)
	o.lastReceiveTime = te.LastReceiveTime
	o.lastInterimReceiveTime = te.LastInterimReceiveTime
	o.plusBuyTime = te.PlusBuyTime
	o.interimBuyTime = te.InterimBuyTime
	o.expireTime = te.ExpireTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerHuiYuanObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "HuiYuan"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
