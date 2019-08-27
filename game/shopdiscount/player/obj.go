package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shopdiscountentity "fgame/fgame/game/shopdiscount/entity"
	shopdiscounttypes "fgame/fgame/game/shopdiscount/types"
	"fmt"
)

//商城促销对象
type PlayerShopDiscountObject struct {
	player     player.Player
	Id         int64
	Typ        shopdiscounttypes.ShopDiscountType
	StartTime  int64
	EndTime    int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShopDiscountObject(pl player.Player) *PlayerShopDiscountObject {
	pso := &PlayerShopDiscountObject{
		player: pl,
	}
	return pso
}

func (pxfo *PlayerShopDiscountObject) GetPlayerId() int64 {
	return pxfo.player.GetId()
}

func (pxfo *PlayerShopDiscountObject) GetDBId() int64 {
	return pxfo.Id
}

func (pxfo *PlayerShopDiscountObject) ToEntity() (e storage.Entity, err error) {
	e = &shopdiscountentity.PlayerShopDiscountEntity{
		Id:         pxfo.Id,
		PlayerId:   pxfo.player.GetId(),
		Typ:        int32(pxfo.Typ),
		StartTime:  pxfo.StartTime,
		EndTime:    pxfo.EndTime,
		UpdateTime: pxfo.UpdateTime,
		CreateTime: pxfo.CreateTime,
		DeleteTime: pxfo.DeleteTime,
	}
	return e, err
}

func (pxfo *PlayerShopDiscountObject) FromEntity(e storage.Entity) error {
	pxfe, _ := e.(*shopdiscountentity.PlayerShopDiscountEntity)
	pxfo.Id = pxfe.Id
	pxfo.Typ = shopdiscounttypes.ShopDiscountType(pxfe.Typ)
	pxfo.StartTime = pxfe.StartTime
	pxfo.EndTime = pxfe.EndTime
	pxfo.UpdateTime = pxfe.UpdateTime
	pxfo.CreateTime = pxfe.CreateTime
	pxfo.DeleteTime = pxfe.DeleteTime
	return nil
}

func (pxfo *PlayerShopDiscountObject) SetModified() {
	e, err := pxfo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("shopdiscount: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pxfo.player.AddChangedObject(obj)
	return
}
