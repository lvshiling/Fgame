package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenfa/entity"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"fmt"
)

//玩家阵法对象
type PlayerZhenFaObject struct {
	player     player.Player
	id         int64
	typ        zhenfatypes.ZhenFaType
	level      int32
	levelNum   int32
	levelPro   int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerZhenFaObject(pl player.Player) *PlayerZhenFaObject {
	pmo := &PlayerZhenFaObject{
		player: pl,
	}
	return pmo
}

func convertZhenFaObjectToEntity(psco *PlayerZhenFaObject) (*entity.PlayerZhenFaEntity, error) {

	e := &entity.PlayerZhenFaEntity{
		Id:         psco.id,
		PlayerId:   psco.player.GetId(),
		Type:       int32(psco.typ),
		Level:      psco.level,
		LevelNum:   psco.levelNum,
		LevelPro:   psco.levelPro,
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerZhenFaObject) GetPlayerId() int64 {
	return psco.player.GetId()
}

func (psco *PlayerZhenFaObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerZhenFaObject) GetZhenFaType() zhenfatypes.ZhenFaType {
	return psco.typ
}

func (psco *PlayerZhenFaObject) GetLevel() int32 {
	return psco.level
}

func (psco *PlayerZhenFaObject) GetLevelNum() int32 {
	return psco.levelNum
}

func (psco *PlayerZhenFaObject) GetLevelPro() int32 {
	return psco.levelPro
}

func (psco *PlayerZhenFaObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertZhenFaObjectToEntity(psco)
	return e, err
}

func (psco *PlayerZhenFaObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerZhenFaEntity)

	psco.id = pse.Id
	psco.typ = zhenfatypes.ZhenFaType(pse.Type)
	psco.level = pse.Level
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerZhenFaObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("zhenfa: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
