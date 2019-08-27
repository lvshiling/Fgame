package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shihunfanentity "fgame/fgame/game/shihunfan/entity"

	"github.com/pkg/errors"
)

//噬魂幡对象
type PlayerShiHunFanObject struct {
	player     player.Player
	Id         int64
	AdvanceId  int
	TimesNum   int32
	Bless      int32
	BlessTime  int64
	DanLevel   int32
	DanNum     int32
	DanPro     int32
	ChargeVal  int32
	Power      int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShiHunFanObject(pl player.Player) *PlayerShiHunFanObject {
	o := &PlayerShiHunFanObject{
		player: pl,
	}
	return o
}

func convertNewPlayerShiHunFanObjectToEntity(o *PlayerShiHunFanObject) (*shihunfanentity.PlayerShiHunFanEntity, error) {

	e := &shihunfanentity.PlayerShiHunFanEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		AdvancedId: o.AdvanceId,
		TimesNum:   o.TimesNum,
		Bless:      o.Bless,
		BlessTime:  o.BlessTime,
		DanLevel:   o.DanLevel,
		DanNum:     o.DanNum,
		DanPro:     o.DanPro,
		ChargeVal:  o.ChargeVal,
		Power:      o.Power,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerShiHunFanObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerShiHunFanObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerShiHunFanObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerShiHunFanObjectToEntity(o)
	return e, err
}

func (o *PlayerShiHunFanObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*shihunfanentity.PlayerShiHunFanEntity)

	o.Id = pse.Id
	o.AdvanceId = pse.AdvancedId

	o.TimesNum = pse.TimesNum
	o.Bless = pse.Bless
	o.BlessTime = pse.BlessTime
	o.DanLevel = pse.DanLevel
	o.DanNum = pse.DanNum
	o.DanPro = pse.DanPro
	o.ChargeVal = pse.ChargeVal
	o.Power = pse.Power
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerShiHunFanObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ShiHunFan"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
