package player

import (
	"fgame/fgame/core/storage"
	additionsysentity "fgame/fgame/game/additionsys/entity"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//附加系统其他数据对象
type PlayerAdditionSysTongLingObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	SysType     additionsystypes.AdditionSysType
	TongLingLev int32
	TongLingNum int32
	TongLingPro int32
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

//获取通灵下一个升级数据
func (o *PlayerAdditionSysTongLingObject) GetNextTongLingTemplate() additionsystemplate.SystemUpgradeCommonTemplate {
	level := int32(1)
	if o.TongLingLev > 0 {
		level = o.TongLingLev + 1
	}
	nextTemplate := additionsystemplate.GetAdditionSysTemplateService().GetTongLingByLevel(level)
	return nextTemplate
}

//系统通灵食用丹操作
// func (o *PlayerAdditionSysTongLingObject) TongLingUpgrade(pro, addTimes int32, sucess bool) {
// 	if pro < 0 {
// 		return
// 	}
// 	if sucess {
// 		template := o.GetNextTongLingTemplate()
// 		if template == nil {
// 			return
// 		}
// 		o.TongLingLev = template.GetLevel()
// 		o.TongLingNum = 0
// 		o.TongLingPro = 0
// 	} else {
// 		o.TongLingNum += addTimes
// 		o.TongLingPro += pro
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	o.UpdateTime = now
// 	o.SetModified()
// 	return
// }

//系统通灵食用丹操作
func (o *PlayerAdditionSysTongLingObject) TongLingUpgrade(level int32) {
	if o.TongLingLev == level || level <= 0 {
		return
	}
	template := additionsystemplate.GetAdditionSysTemplateService().GetTongLingByLevel(level)
	if template == nil {
		return
	}
	o.TongLingLev = level
	now := global.GetGame().GetTimeService().Now()
	o.UpdateTime = now
	o.SetModified()
	return
}

func NewPlayerAdditionSysTongLingObject(pl player.Player) *PlayerAdditionSysTongLingObject {
	o := &PlayerAdditionSysTongLingObject{
		player:   pl,
		PlayerId: pl.GetId(),
	}
	return o
}

func createAdditionSysTongLingObject(p player.Player, typ additionsystypes.AdditionSysType, now int64) *PlayerAdditionSysTongLingObject {
	obj := NewPlayerAdditionSysTongLingObject(p)
	obj.Id, _ = idutil.GetId()
	obj.SysType = typ
	obj.CreateTime = now
	return obj
}

func convertNewPlayerAdditionSysTongLingObjectToEntity(o *PlayerAdditionSysTongLingObject) (*additionsysentity.PlayerAdditionSysTongLingEntity, error) {

	e := &additionsysentity.PlayerAdditionSysTongLingEntity{
		Id:          o.Id,
		PlayerId:    o.player.GetId(),
		SysType:     int32(o.SysType),
		TongLingLev: o.TongLingLev,
		TongLingNum: o.TongLingNum,
		TongLingPro: o.TongLingPro,
		UpdateTime:  o.UpdateTime,
		CreateTime:  o.CreateTime,
		DeleteTime:  o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerAdditionSysTongLingObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAdditionSysTongLingObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerAdditionSysTongLingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerAdditionSysTongLingObjectToEntity(o)
	return e, err
}

func (o *PlayerAdditionSysTongLingObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*additionsysentity.PlayerAdditionSysTongLingEntity)

	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.SysType = additionsystypes.AdditionSysType(pse.SysType)
	o.TongLingLev = pse.TongLingLev
	o.TongLingNum = pse.TongLingNum
	o.TongLingPro = pse.TongLingPro
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerAdditionSysTongLingObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AdditionSysTongLing"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
