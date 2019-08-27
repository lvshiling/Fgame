package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenfaentity "fgame/fgame/game/shenfa/entity"
	shenfatypes "fgame/fgame/game/shenfa/types"
	"fmt"

	"github.com/pkg/errors"
)

//身法对象
type PlayerShenfaObject struct {
	player      player.Player
	Id          int64
	AdvanceId   int
	ShenfaId    int32
	UnrealLevel int32
	UnrealNum   int32
	UnrealPro   int32
	UnrealList  []int
	TimesNum    int32
	Bless       int32
	BlessTime   int64
	Hidden      int32
	Power       int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerShenfaObject(pl player.Player) *PlayerShenfaObject {
	o := &PlayerShenfaObject{
		player: pl,
	}
	return o
}

func convertNewPlayerShenfaObjectToEntity(o *PlayerShenfaObject) (*shenfaentity.PlayerShenfaEntity, error) {
	unrealInfoBytes, err := json.Marshal(o.UnrealList)
	if err != nil {
		return nil, err
	}
	e := &shenfaentity.PlayerShenfaEntity{
		Id:          o.Id,
		PlayerId:    o.player.GetId(),
		AdvancedId:  o.AdvanceId,
		ShenfaId:    o.ShenfaId,
		UnrealLevel: o.UnrealLevel,
		UnrealNum:   o.UnrealNum,
		UnrealPro:   o.UnrealPro,
		UnrealInfo:  string(unrealInfoBytes),
		TimesNum:    o.TimesNum,
		Bless:       o.Bless,
		BlessTime:   o.BlessTime,
		Hidden:      o.Hidden,
		Power:       o.Power,
		UpdateTime:  o.UpdateTime,
		CreateTime:  o.CreateTime,
		DeleteTime:  o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerShenfaObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerShenfaObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerShenfaObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerShenfaObjectToEntity(o)
	return e, err
}

func (o *PlayerShenfaObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*shenfaentity.PlayerShenfaEntity)

	var unrealList = make([]int, 0, 8)
	if err := json.Unmarshal([]byte(pse.UnrealInfo), &unrealList); err != nil {
		return err
	}
	o.Id = pse.Id
	o.AdvanceId = pse.AdvancedId
	o.ShenfaId = pse.ShenfaId
	o.UnrealLevel = pse.UnrealLevel
	o.UnrealNum = pse.UnrealNum
	o.UnrealPro = pse.UnrealPro
	o.UnrealList = unrealList
	o.TimesNum = pse.TimesNum
	o.Bless = pse.Bless
	o.BlessTime = pse.BlessTime
	o.Hidden = pse.Hidden
	o.Power = pse.Power
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerShenfaObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("ShenFa: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//身法非进阶对象
type PlayerShenfaOtherObject struct {
	player     player.Player
	Id         int64
	Typ        shenfatypes.ShenfaType
	ShenFaId   int32
	Level      int32
	UpNum      int32
	UpPro      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShenfaOtherObject(pl player.Player) *PlayerShenfaOtherObject {
	o := &PlayerShenfaOtherObject{
		player: pl,
	}
	return o
}

func convertShenfaOtherObjectToEntity(o *PlayerShenfaOtherObject) (*shenfaentity.PlayerShenfaOtherEntity, error) {

	e := &shenfaentity.PlayerShenfaOtherEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		Typ:        int32(o.Typ),
		ShenFaId:   o.ShenFaId,
		Level:      o.Level,
		UpNum:      o.UpNum,
		UpPro:      o.UpPro,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerShenfaOtherObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerShenfaOtherObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerShenfaOtherObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertShenfaOtherObjectToEntity(o)
	return e, err
}

func (o *PlayerShenfaOtherObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*shenfaentity.PlayerShenfaOtherEntity)

	o.Id = pse.Id
	o.Typ = shenfatypes.ShenfaType(pse.Typ)
	o.ShenFaId = pse.ShenFaId
	o.Level = pse.Level
	o.UpNum = pse.UpNum
	o.UpPro = pse.UpPro
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerShenfaOtherObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ShenFa"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
