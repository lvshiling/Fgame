package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tianmoentity "fgame/fgame/game/tianmo/entity"

	"github.com/pkg/errors"
)

//天魔对象
type PlayerTianMoObject struct {
	player         player.Player
	Id             int64
	AdvanceId      int32
	TianMoDanLevel int32
	TianMoDanNum   int32
	TianMoDanPro   int32
	TimesNum       int32
	Bless          int32
	BlessTime      int64
	Power          int64
	ChargeVal      int64
	UpdateTime     int64
	CreateTime     int64
	DeleteTime     int64
}

func NewPlayerTianMoObject(pl player.Player) *PlayerTianMoObject {
	o := &PlayerTianMoObject{
		player: pl,
	}
	return o
}

func convertNewPlayerTianMoObjectToEntity(o *PlayerTianMoObject) (*tianmoentity.PlayerTianMoEntity, error) {

	e := &tianmoentity.PlayerTianMoEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		AdvancedId: o.AdvanceId,
		DanLevel:   o.TianMoDanLevel,
		DanNum:     o.TianMoDanNum,
		DanPro:     o.TianMoDanPro,
		TimesNum:   o.TimesNum,
		Bless:      o.Bless,
		BlessTime:  o.BlessTime,
		Power:      o.Power,
		ChargeVal:  o.ChargeVal,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerTianMoObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerTianMoObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerTianMoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerTianMoObjectToEntity(o)
	return e, err
}

func (o *PlayerTianMoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*tianmoentity.PlayerTianMoEntity)

	o.Id = pse.Id
	o.AdvanceId = pse.AdvancedId
	o.TianMoDanLevel = pse.DanLevel
	o.TianMoDanNum = pse.DanNum
	o.TianMoDanPro = pse.DanPro
	o.TimesNum = pse.TimesNum
	o.Bless = pse.Bless
	o.BlessTime = pse.BlessTime
	o.Power = pse.Power
	o.ChargeVal = pse.ChargeVal
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerTianMoObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TianMo"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
