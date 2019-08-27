package trade

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	tradeentity "fgame/fgame/game/trade/entity"

	"github.com/pkg/errors"
)

type TradeRecycleObject struct {
	id                int64
	serverId          int32
	recycleGold       int64
	recycleTime       int64
	customRecycleGold int64
	updateTime        int64
	createTime        int64
	deleteTime        int64
}

func createTradeRecycleObject() *TradeRecycleObject {
	o := &TradeRecycleObject{}

	return o
}

func convertTradeRecyleToEntity(o *TradeRecycleObject) (*tradeentity.TradeRecycleEntity, error) {
	e := &tradeentity.TradeRecycleEntity{
		Id:                o.id,
		ServerId:          o.serverId,
		RecycleGold:       o.recycleGold,
		RecycleTime:       o.recycleTime,
		CustomRecycleGold: o.customRecycleGold,
		UpdateTime:        o.updateTime,
		CreateTime:        o.createTime,
		DeleteTime:        o.deleteTime,
	}
	return e, nil
}

func (o *TradeRecycleObject) GetId() int64 {
	return o.id
}

func (o *TradeRecycleObject) GetDBId() int64 {
	return o.id
}

func (o *TradeRecycleObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertTradeRecyleToEntity(o)
	return e, err
}

func (o *TradeRecycleObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*tradeentity.TradeRecycleEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.recycleGold = ae.RecycleGold
	o.recycleTime = ae.RecycleTime
	o.customRecycleGold = ae.CustomRecycleGold
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *TradeRecycleObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TradeRecycle"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
