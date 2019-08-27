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

type PlayerAdditionSysLingZhuObject struct {
	player      player.Player
	id          int64
	sysType     additionsystypes.AdditionSysType
	lingZhuType additionsystypes.LingZhuType
	level       int32
	times       int32
	bless       int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func (o *PlayerAdditionSysLingZhuObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerAdditionSysLingZhuObject) GetTimes() int32 {
	return o.times
}

func (o *PlayerAdditionSysLingZhuObject) GetBless() int64 {
	return o.bless
}

// func (o *PlayerAdditionSysLingZhuObject) UpLevel(sucess bool, bless int64) {
// 	if sucess {
// 		o.level = o.level + 1
// 		o.times = 0
// 		o.bless = 0
// 	} else {
// 		o.times = o.times + 1
// 		o.bless += bless
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	o.updateTime = now
// 	o.SetModified()
// }

func (o *PlayerAdditionSysLingZhuObject) GetLingZhuType() additionsystypes.LingZhuType {
	return o.lingZhuType
}

func NewPlayerAdditionSysLingZhuObject(pl player.Player) *PlayerAdditionSysLingZhuObject {
	o := &PlayerAdditionSysLingZhuObject{
		player: pl,
	}
	return o
}

func createPlayerAdditionSysLingZhuObject(p player.Player, sysTyp additionsystypes.AdditionSysType, lingZhuType additionsystypes.LingZhuType) *PlayerAdditionSysLingZhuObject {
	now := global.GetGame().GetTimeService().Now()
	obj := NewPlayerAdditionSysLingZhuObject(p)
	obj.id, _ = idutil.GetId()
	obj.sysType = sysTyp
	obj.lingZhuType = lingZhuType
	obj.times = 0
	obj.bless = 0
	obj.level = 0
	obj.createTime = now
	return obj
}

func (o *PlayerAdditionSysLingZhuObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAdditionSysLingZhuObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerAdditionSysLingZhuObject) ToEntity() (e storage.Entity, err error) {
	e = &additionsysentity.PlayerAdditionSysLingZhuEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		SysType:    int32(o.sysType),
		LingZhuId:  int32(o.lingZhuType),
		Bless:      o.bless,
		Times:      o.times,
		Level:      o.level,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, err
}

func (o *PlayerAdditionSysLingZhuObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*additionsysentity.PlayerAdditionSysLingZhuEntity)

	o.id = pse.Id
	o.sysType = additionsystypes.AdditionSysType(pse.SysType)
	o.lingZhuType = additionsystypes.LingZhuType(pse.LingZhuId)
	o.times = pse.Times
	o.bless = pse.Bless
	o.level = pse.Level
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerAdditionSysLingZhuObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AdditionSysLingZhuObject"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
