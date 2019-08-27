package player

import (
	"fgame/fgame/core/storage"
	additionsysentity "fgame/fgame/game/additionsys/entity"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//附加系统升级对象
type PlayerAdditionSysAwakeObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	SysType    additionsystypes.AdditionSysType
	IsAwake    int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerAdditionSysAwakeObject(pl player.Player) *PlayerAdditionSysAwakeObject {
	o := &PlayerAdditionSysAwakeObject{
		player:   pl,
		PlayerId: pl.GetId(),
	}
	return o
}

func createAdditionSysAwakeObject(p player.Player, typ additionsystypes.AdditionSysType, now int64) *PlayerAdditionSysAwakeObject {
	levelObject := NewPlayerAdditionSysAwakeObject(p)
	levelObject.Id, _ = idutil.GetId()
	levelObject.SysType = typ
	levelObject.CreateTime = now
	return levelObject
}

func convertNewPlayerAdditionSysAwakeObjectToEntity(o *PlayerAdditionSysAwakeObject) (*additionsysentity.PlayerAdditionSysAwakeEntity, error) {

	e := &additionsysentity.PlayerAdditionSysAwakeEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		SysType:    int32(o.SysType),
		IsAwake:    o.IsAwake,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerAdditionSysAwakeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAdditionSysAwakeObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerAdditionSysAwakeObject) IsAlreadyAwake() bool {
	return o.IsAwake != 0
}

func (o *PlayerAdditionSysAwakeObject) GetAwakeLevel() int32 {
	return o.IsAwake
}

func (o *PlayerAdditionSysAwakeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerAdditionSysAwakeObjectToEntity(o)
	return e, err
}

func (o *PlayerAdditionSysAwakeObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*additionsysentity.PlayerAdditionSysAwakeEntity)

	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.SysType = additionsystypes.AdditionSysType(pse.SysType)
	o.IsAwake = pse.IsAwake
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerAdditionSysAwakeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AdditionSysAwake"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerAdditionSysAwakeObject) SysAwakeSuccess() {
	now := global.GetGame().GetTimeService().Now()
	o.IsAwake++
	o.UpdateTime = now
	o.SetModified()
	return
}
