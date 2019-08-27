package player

import (
	"fgame/fgame/core/storage"
	anqientity "fgame/fgame/game/anqi/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//暗器对象
type PlayerAnqiObject struct {
	player       player.Player
	Id           int64
	AdvanceId    int
	AnqiDanLevel int32
	AnqiDanNum   int32
	AnqiDanPro   int32
	TimesNum     int32
	Bless        int32
	BlessTime    int64
	Power        int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

func NewPlayerAnqiObject(pl player.Player) *PlayerAnqiObject {
	o := &PlayerAnqiObject{
		player: pl,
	}
	return o
}

func convertNewPlayerAnqiObjectToEntity(o *PlayerAnqiObject) (*anqientity.PlayerAnQiEntity, error) {

	e := &anqientity.PlayerAnQiEntity{
		Id:           o.Id,
		PlayerId:     o.player.GetId(),
		AdvancedId:   o.AdvanceId,
		AnqiDanLevel: o.AnqiDanLevel,
		AnqiDanNum:   o.AnqiDanNum,
		AnqiDanPro:   o.AnqiDanPro,
		TimesNum:     o.TimesNum,
		Bless:        o.Bless,
		BlessTime:    o.BlessTime,
		Power:        o.Power,
		UpdateTime:   o.UpdateTime,
		CreateTime:   o.CreateTime,
		DeleteTime:   o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerAnqiObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAnqiObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerAnqiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerAnqiObjectToEntity(o)
	return e, err
}

func (o *PlayerAnqiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*anqientity.PlayerAnQiEntity)

	o.Id = pse.Id
	o.AdvanceId = pse.AdvancedId
	o.AnqiDanLevel = pse.AnqiDanLevel
	o.AnqiDanNum = pse.AnqiDanNum
	o.AnqiDanPro = pse.AnqiDanPro
	o.TimesNum = pse.TimesNum
	o.Bless = pse.Bless
	o.BlessTime = pse.BlessTime
	o.Power = pse.Power
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerAnqiObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Anqi"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
