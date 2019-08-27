package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	vipentity "fgame/fgame/game/vip/entity"

	"github.com/pkg/errors"
)

//VIP对象
type PlayerVipObject struct {
	player       player.Player
	id           int64
	vipLevel     int32
	vipStar      int32
	consumeLevel int32
	chargeNum    int64
	freeGiftMap  map[int32]int32
	discountMap  map[int32]int32
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerVipObject(pl player.Player) *PlayerVipObject {
	o := &PlayerVipObject{
		player: pl,
	}
	return o
}

func convertNewPlayerVipObjectToEntity(o *PlayerVipObject) (*vipentity.PlayerVipEntity, error) {
	discountData, err := json.Marshal(o.discountMap)
	if err != nil {
		return nil, err
	}
	freeGiftData, err := json.Marshal(o.freeGiftMap)
	if err != nil {
		return nil, err
	}

	e := &vipentity.PlayerVipEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		VipLevel:     o.vipLevel,
		VipStar:      o.vipStar,
		ConsumeLevel: o.consumeLevel,
		ChargeNum:    o.chargeNum,
		FreeGiftMap:  string(freeGiftData),
		DiscountMap:  string(discountData),
		UpdateTime:   o.updateTime,
		DeleteTime:   o.deleteTime,
		CreateTime:   o.createTime,
	}
	return e, nil
}

func (o *PlayerVipObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerVipObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerVipObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerVipObjectToEntity(o)
	return e, err
}

func (o *PlayerVipObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*vipentity.PlayerVipEntity)

	discountMap := make(map[int32]int32)
	err := json.Unmarshal([]byte(pse.DiscountMap), &discountMap)
	if err != nil {
		return err
	}

	freeGiftMap := make(map[int32]int32)
	err = json.Unmarshal([]byte(pse.FreeGiftMap), &freeGiftMap)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.vipLevel = pse.VipLevel
	o.vipStar = pse.VipStar
	o.consumeLevel = pse.ConsumeLevel
	o.chargeNum = pse.ChargeNum
	o.freeGiftMap = freeGiftMap
	o.discountMap = discountMap
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerVipObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "VIP"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
