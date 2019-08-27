package player

import (
	"fgame/fgame/core/storage"
	equipbaokuentity "fgame/fgame/game/equipbaoku/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//玩家当日宝库商店购买道具对象
type PlayerEquipBaoKuShopObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	ShopId     int32
	DayCount   int32
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerEquipBaoKuShopObject(pl player.Player) *PlayerEquipBaoKuShopObject {
	pmo := &PlayerEquipBaoKuShopObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerEquipBaoKuShopObjectToEntity(pso *PlayerEquipBaoKuShopObject) (*equipbaokuentity.PlayerEquipBaoKuShopEntity, error) {
	e := &equipbaokuentity.PlayerEquipBaoKuShopEntity{
		Id:         pso.Id,
		PlayerId:   pso.PlayerId,
		ShopId:     pso.ShopId,
		DayCount:   pso.DayCount,
		LastTime:   pso.LastTime,
		UpdateTime: pso.UpdateTime,
		CreateTime: pso.CreateTime,
		DeleteTime: pso.DeleteTime,
	}
	return e, nil
}

func (pso *PlayerEquipBaoKuShopObject) GetPlayerId() int64 {
	return pso.PlayerId
}

func (pso *PlayerEquipBaoKuShopObject) GetDBId() int64 {
	return pso.Id
}

func (pso *PlayerEquipBaoKuShopObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerEquipBaoKuShopObjectToEntity(pso)
	return e, err
}

func (pso *PlayerEquipBaoKuShopObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*equipbaokuentity.PlayerEquipBaoKuShopEntity)

	pso.Id = pse.Id
	pso.PlayerId = pse.PlayerId
	pso.ShopId = pse.ShopId
	pso.DayCount = pse.DayCount
	pso.LastTime = pse.LastTime
	pso.UpdateTime = pse.UpdateTime
	pso.CreateTime = pse.CreateTime
	pso.DeleteTime = pse.DeleteTime
	return nil
}

func (pso *PlayerEquipBaoKuShopObject) SetModified() {
	e, err := pso.ToEntity()
	if err != nil {
		panic(fmt.Errorf("EquipBaoKuShop: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pso.player.AddChangedObject(obj)
	return
}
