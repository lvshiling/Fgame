package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"

	skillentity "fgame/fgame/game/skill/entity"
)

//玩家职业技能数据
type PlayerSkillObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	SkillId    int32
	Level      int32
	TianFuMap  map[int32]int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerSkillObject(pl player.Player) *PlayerSkillObject {
	pso := &PlayerSkillObject{
		player: pl,
	}
	return pso
}

func convertPlayerSkillObjectToEntity(ppo *PlayerSkillObject) (pse *skillentity.PlayerSkillEntity, err error) {

	tianFuInfoBytes, err := json.Marshal(ppo.TianFuMap)
	if err != nil {
		return nil, err
	}

	e := &skillentity.PlayerSkillEntity{
		Id:         ppo.Id,
		PlayerId:   ppo.PlayerId,
		SkillId:    ppo.SkillId,
		Level:      ppo.Level,
		TianFuInfo: string(tianFuInfoBytes),
		UpdateTime: ppo.UpdateTime,
		CreateTime: ppo.CreateTime,
		DeleteTime: ppo.DeleteTime,
	}
	return e, nil
}

func (pso *PlayerSkillObject) GetPlayerId() int64 {
	return pso.PlayerId
}

func (pso *PlayerSkillObject) GetDBId() int64 {
	return pso.Id
}

func (pso *PlayerSkillObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerSkillObjectToEntity(pso)
	if err != nil {
		return
	}
	return
}

func (pso *PlayerSkillObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*skillentity.PlayerSkillEntity)
	tianFuMap := make(map[int32]int32)
	err = json.Unmarshal([]byte(pse.TianFuInfo), &tianFuMap)
	if err != nil {
		return
	}

	pso.Id = pse.Id
	pso.SkillId = pse.SkillId
	pso.Level = pse.Level
	pso.PlayerId = pse.PlayerId
	pso.TianFuMap = tianFuMap
	pso.UpdateTime = pse.UpdateTime
	pso.CreateTime = pse.CreateTime
	pso.DeleteTime = pse.DeleteTime
	return
}

func (pso *PlayerSkillObject) SetModified() {
	e, err := pso.ToEntity()
	if err != nil {
		panic(fmt.Errorf("Skill: err[%s]", err.Error()))
	}

	obj, _ := e.(types.PlayerDataEntity)
	if obj == nil {
		panic("never reach here")
	}
	pso.player.AddChangedObject(obj)
	return
}

func (pso *PlayerSkillObject) GetSkillId() int32 {
	return pso.SkillId
}

func (pso *PlayerSkillObject) GetLevel() int32 {
	return pso.Level
}

//玩家技能冷却时间数据
type PlayerSkillCdObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	SkillId    int32
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerSkillCdObject(pl player.Player) *PlayerSkillCdObject {
	pso := &PlayerSkillCdObject{
		player: pl,
	}
	return pso
}

func convertPlayerSkillCdObjectToEntity(ppo *PlayerSkillCdObject) (pse *skillentity.PlayerSkillCdEntity, err error) {
	e := &skillentity.PlayerSkillCdEntity{
		Id:         ppo.Id,
		PlayerId:   ppo.PlayerId,
		SkillId:    ppo.SkillId,
		LastTime:   ppo.LastTime,
		UpdateTime: ppo.UpdateTime,
		CreateTime: ppo.CreateTime,
		DeleteTime: ppo.DeleteTime,
	}
	return e, nil
}

func (pso *PlayerSkillCdObject) GetPlayerId() int64 {
	return pso.PlayerId
}

func (pso *PlayerSkillCdObject) GetDBId() int64 {
	return pso.Id
}

func (pso *PlayerSkillCdObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerSkillCdObjectToEntity(pso)
	if err != nil {
		return
	}
	return
}

func (pso *PlayerSkillCdObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*skillentity.PlayerSkillCdEntity)
	pso.Id = pse.Id
	pso.SkillId = pse.SkillId
	pso.PlayerId = pse.PlayerId
	pso.LastTime = pse.LastTime
	pso.UpdateTime = pse.UpdateTime
	pso.CreateTime = pse.CreateTime
	pso.DeleteTime = pse.DeleteTime
	return
}

func (pso *PlayerSkillCdObject) SetModified() {
	e, err := pso.ToEntity()
	if err != nil {
		panic(fmt.Errorf("Skill: err[%s]", err.Error()))
	}

	obj, _ := e.(types.PlayerDataEntity)
	if obj == nil {
		panic("never reach here")
	}
	pso.player.AddChangedObject(obj)
	return
}

func (pso *PlayerSkillCdObject) GetSkillId() int32 {
	return pso.SkillId
}

func (pso *PlayerSkillCdObject) GetLastTime() int64 {
	return pso.LastTime
}
