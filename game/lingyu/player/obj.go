package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	lingyuentity "fgame/fgame/game/lingyu/entity"
	lingyutypes "fgame/fgame/game/lingyu/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//领域对象
type PlayerLingyuObject struct {
	player      player.Player
	Id          int64
	AdvanceId   int
	LingyuId    int32
	UnrealLevel int32
	UnrealNum   int32
	UnrealPro   int32
	UnrealList  []int
	TimesNum    int32
	Bless       int32
	BlessTime   int64
	Hidden      int32
	Power       int64
	ChargeVal   int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerLingyuObject(pl player.Player) *PlayerLingyuObject {
	o := &PlayerLingyuObject{
		player: pl,
	}
	return o
}

func convertNewPlayerLingyuObjectToEntity(o *PlayerLingyuObject) (*lingyuentity.PlayerLingyuEntity, error) {
	unrealInfoBytes, err := json.Marshal(o.UnrealList)
	if err != nil {
		return nil, err
	}
	e := &lingyuentity.PlayerLingyuEntity{
		Id:          o.Id,
		PlayerId:    o.player.GetId(),
		AdvancedId:  o.AdvanceId,
		LingyuId:    o.LingyuId,
		UnrealLevel: o.UnrealLevel,
		UnrealNum:   o.UnrealNum,
		UnrealPro:   o.UnrealPro,
		UnrealInfo:  string(unrealInfoBytes),
		TimesNum:    o.TimesNum,
		Bless:       o.Bless,
		BlessTime:   o.BlessTime,
		Hidden:      o.Hidden,
		Power:       o.Power,
		ChargeVal:   o.ChargeVal,
		UpdateTime:  o.UpdateTime,
		CreateTime:  o.CreateTime,
		DeleteTime:  o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerLingyuObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerLingyuObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerLingyuObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerLingyuObjectToEntity(o)
	return e, err
}

func (o *PlayerLingyuObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*lingyuentity.PlayerLingyuEntity)

	var unrealList = make([]int, 0, 8)
	if err := json.Unmarshal([]byte(pse.UnrealInfo), &unrealList); err != nil {
		return err
	}
	o.Id = pse.Id
	o.AdvanceId = pse.AdvancedId
	o.LingyuId = pse.LingyuId
	o.UnrealLevel = pse.UnrealLevel
	o.UnrealNum = pse.UnrealNum
	o.UnrealPro = pse.UnrealPro
	o.UnrealList = unrealList
	o.TimesNum = pse.TimesNum
	o.Bless = pse.Bless
	o.BlessTime = pse.BlessTime
	o.Hidden = pse.Hidden
	o.Power = pse.Power
	o.ChargeVal = pse.ChargeVal
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerLingyuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Lingyu"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//领域非进阶对象
type PlayerLingyuOtherObject struct {
	player     player.Player
	Id         int64
	Typ        lingyutypes.LingyuType
	LingYuId   int32
	Level      int32
	UpNum      int32
	UpPro      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerLingyuOtherObject(pl player.Player) *PlayerLingyuOtherObject {
	o := &PlayerLingyuOtherObject{
		player: pl,
	}
	return o
}

func convertLingyuOtherObjectToEntity(o *PlayerLingyuOtherObject) (*lingyuentity.PlayerLingyuOtherEntity, error) {

	e := &lingyuentity.PlayerLingyuOtherEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		Typ:        int32(o.Typ),
		LingYuId:   o.LingYuId,
		Level:      o.Level,
		UpNum:      o.UpNum,
		UpPro:      o.UpPro,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerLingyuOtherObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerLingyuOtherObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerLingyuOtherObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingyuOtherObjectToEntity(o)
	return e, err
}

func (o *PlayerLingyuOtherObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*lingyuentity.PlayerLingyuOtherEntity)

	o.Id = pse.Id
	o.Typ = lingyutypes.LingyuType(pse.Typ)
	o.LingYuId = pse.LingYuId
	o.Level = pse.Level
	o.UpNum = pse.UpNum
	o.UpPro = pse.UpPro
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerLingyuOtherObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "LingyuOther"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
