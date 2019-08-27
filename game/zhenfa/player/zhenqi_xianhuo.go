package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenfa/entity"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"fmt"
)

//玩家阵旗仙火对象
type PlayerZhenQiXianHuoObject struct {
	player     player.Player
	id         int64
	typ        zhenfatypes.ZhenFaType
	level      int32
	luckyStar  int32
	levelNum   int32
	levelPro   int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerZhenQiXianHuoObject(pl player.Player) *PlayerZhenQiXianHuoObject {
	pmo := &PlayerZhenQiXianHuoObject{
		player: pl,
	}
	return pmo
}

func convertZhenQiXianHuoObjectToEntity(psco *PlayerZhenQiXianHuoObject) (*entity.PlayerZhenQiXianHuoEntity, error) {

	e := &entity.PlayerZhenQiXianHuoEntity{
		Id:         psco.id,
		PlayerId:   psco.player.GetId(),
		Type:       int32(psco.typ),
		Level:      psco.level,
		LevelNum:   psco.levelNum,
		LevelPro:   psco.levelPro,
		LuckyStar:  psco.luckyStar,
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerZhenQiXianHuoObject) GetPlayerId() int64 {
	return psco.player.GetId()
}

func (psco *PlayerZhenQiXianHuoObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerZhenQiXianHuoObject) GetZhenFaType() zhenfatypes.ZhenFaType {
	return psco.typ
}

func (psco *PlayerZhenQiXianHuoObject) GetLevel() int32 {
	return psco.level
}

func (psco *PlayerZhenQiXianHuoObject) GetLuckyStar() int32 {
	return psco.luckyStar
}

func (psco *PlayerZhenQiXianHuoObject) GetLevelNum() int32 {
	return psco.levelNum
}

func (psco *PlayerZhenQiXianHuoObject) GetLevelPro() int32 {
	return psco.levelPro
}

func (psco *PlayerZhenQiXianHuoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertZhenQiXianHuoObjectToEntity(psco)
	return e, err
}

func (psco *PlayerZhenQiXianHuoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerZhenQiXianHuoEntity)

	psco.id = pse.Id
	psco.typ = zhenfatypes.ZhenFaType(pse.Type)
	psco.luckyStar = pse.LuckyStar
	psco.level = pse.Level
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerZhenQiXianHuoObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("zhenqi_xianhuo: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
