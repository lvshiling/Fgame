package player

import (
	"fgame/fgame/core/storage"
	additionsysentity "fgame/fgame/game/additionsys/entity"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//附加系统升级对象
type PlayerAdditionSysLevelObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	SysType    additionsystypes.AdditionSysType
	Level      int32
	UpNum      int32
	UpPro      int32
	LingLevel  int32
	LingNum    int32
	LingPro    int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

//获取下一个升级数据
func (o *PlayerAdditionSysLevelObject) GetNextShengJiTemplate() additionsystemplate.SystemShengJiCommonTemplate {
	level := int32(1)
	if o.Level > 0 {
		level = o.Level + 1
	}
	nextShengJiTemplate := additionsystemplate.GetAdditionSysTemplateService().GetShengJiByArg(o.SysType, level)
	return nextShengJiTemplate
}

//获取化灵下一个升级数据
func (m *PlayerAdditionSysLevelObject) GetNextHuaLingTemplate() *gametemplate.SystemHuaLingTemplate {
	var nextHuaLingTemplate *gametemplate.SystemHuaLingTemplate
	if m.LingLevel == 0 {
		nextHuaLingTemplate, _ = additionsystemplate.GetAdditionSysTemplateService().GetHuaLingByArg(m.SysType, 1)
	} else {
		//判断系统是否可以升级
		huaLingTemplate, _ := additionsystemplate.GetAdditionSysTemplateService().GetHuaLingByArg(m.SysType, m.LingLevel)
		nextHuaLingTemplate = huaLingTemplate.GetNextTemplate()
	}
	return nextHuaLingTemplate
}

//系统化灵食用丹操作
func (o *PlayerAdditionSysLevelObject) HuaLingUpgrade(level int32) {
	if o.LingLevel == level || level <= 0 {
		return
	}
	template, _ := additionsystemplate.GetAdditionSysTemplateService().GetHuaLingByArg(o.SysType, level)
	if template == nil {
		return
	}
	o.LingLevel = level
	now := global.GetGame().GetTimeService().Now()
	o.UpdateTime = now
	o.SetModified()
	return
}

func NewPlayerAdditionSysLevelObject(pl player.Player) *PlayerAdditionSysLevelObject {
	o := &PlayerAdditionSysLevelObject{
		player:   pl,
		PlayerId: pl.GetId(),
	}
	return o
}

func createAdditionSysLevelObject(p player.Player, typ additionsystypes.AdditionSysType, now int64) *PlayerAdditionSysLevelObject {
	levelObject := NewPlayerAdditionSysLevelObject(p)
	levelObject.Id, _ = idutil.GetId()
	levelObject.SysType = typ
	levelObject.CreateTime = now
	return levelObject
}

func convertNewPlayerAdditionSysLevelObjectToEntity(o *PlayerAdditionSysLevelObject) (*additionsysentity.PlayerAdditionSysLevelEntity, error) {

	e := &additionsysentity.PlayerAdditionSysLevelEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		SysType:    int32(o.SysType),
		Level:      o.Level,
		UpNum:      o.UpNum,
		UpPro:      o.UpPro,
		LingLevel:  o.LingLevel,
		LingNum:    o.LingNum,
		LingPro:    o.LingPro,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerAdditionSysLevelObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerAdditionSysLevelObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerAdditionSysLevelObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerAdditionSysLevelObjectToEntity(o)
	return e, err
}

func (o *PlayerAdditionSysLevelObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*additionsysentity.PlayerAdditionSysLevelEntity)

	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.SysType = additionsystypes.AdditionSysType(pse.SysType)
	o.Level = pse.Level
	o.UpNum = pse.UpNum
	o.UpPro = pse.UpPro
	o.LingLevel = pse.LingLevel
	o.LingNum = pse.LingNum
	o.LingPro = pse.LingPro
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerAdditionSysLevelObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AdditionSysLevel"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
