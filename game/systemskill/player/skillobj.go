package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	systemskillentity "fgame/fgame/game/systemskill/entity"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	"fmt"
)

//系统技能对象
type PlayerSystemSkillObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Type       sysskilltypes.SystemSkillType
	SubType    sysskilltypes.SystemSkillSubType
	Level      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerSystemSkillObject(pl player.Player) *PlayerSystemSkillObject {
	pso := &PlayerSystemSkillObject{
		player: pl,
	}
	return pso
}

func (pxfo *PlayerSystemSkillObject) GetPlayerId() int64 {
	return pxfo.PlayerId
}

func (pxfo *PlayerSystemSkillObject) GetDBId() int64 {
	return pxfo.Id
}

func (pxfo *PlayerSystemSkillObject) ToEntity() (e storage.Entity, err error) {
	e = &systemskillentity.PlayerSystemSkillEntity{
		Id:         pxfo.Id,
		PlayerId:   pxfo.PlayerId,
		Type:       int32(pxfo.Type),
		SubType:    int32(pxfo.SubType),
		Level:      pxfo.Level,
		UpdateTime: pxfo.UpdateTime,
		CreateTime: pxfo.CreateTime,
		DeleteTime: pxfo.DeleteTime,
	}
	return e, err
}

func (pxfo *PlayerSystemSkillObject) FromEntity(e storage.Entity) error {
	pxfe, _ := e.(*systemskillentity.PlayerSystemSkillEntity)
	pxfo.Id = pxfe.Id
	pxfo.PlayerId = pxfe.PlayerId
	pxfo.Type = sysskilltypes.SystemSkillType(pxfe.Type)
	pxfo.SubType = sysskilltypes.SystemSkillSubType(pxfe.SubType)
	pxfo.Level = pxfe.Level
	pxfo.UpdateTime = pxfe.UpdateTime
	pxfo.CreateTime = pxfe.CreateTime
	pxfo.DeleteTime = pxfe.DeleteTime
	return nil
}

func (pxfo *PlayerSystemSkillObject) SetModified() {
	e, err := pxfo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("systemskill: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pxfo.player.AddChangedObject(obj)
	return
}

type SystemSkill struct {
	p           player.Player
	sysSkillMap map[sysskilltypes.SystemSkillSubType]*PlayerSystemSkillObject
}

func (s *SystemSkill) addSystemSkill(obj *PlayerSystemSkillObject) {
	subType := obj.SubType
	s.sysSkillMap[subType] = obj
}

func (s *SystemSkill) GetSystemSkillObjct(subType sysskilltypes.SystemSkillSubType) *PlayerSystemSkillObject {
	skillObj, exist := s.sysSkillMap[subType]
	if !exist {
		return nil
	}
	return skillObj
}

func (s *SystemSkill) GetSysSkillMap() map[sysskilltypes.SystemSkillSubType]*PlayerSystemSkillObject {
	return s.sysSkillMap
}

func createSystemSkill(pl player.Player, obj *PlayerSystemSkillObject) *SystemSkill {
	d := &SystemSkill{
		p:           pl,
		sysSkillMap: make(map[sysskilltypes.SystemSkillSubType]*PlayerSystemSkillObject),
	}
	d.addSystemSkill(obj)
	return d
}
