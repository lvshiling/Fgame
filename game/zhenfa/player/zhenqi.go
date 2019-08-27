package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenfa/entity"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"fmt"
)

//玩家阵旗对象
type PlayerZhenQiObject struct {
	player     player.Player
	id         int64
	typ        zhenfatypes.ZhenFaType
	zhenQiPos  zhenfatypes.ZhenQiType
	number     int32
	numberNum  int32
	numberPro  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerZhenQiObject(pl player.Player) *PlayerZhenQiObject {
	pmo := &PlayerZhenQiObject{
		player: pl,
	}
	return pmo
}

func convertZhenQiObjectToEntity(psco *PlayerZhenQiObject) (*entity.PlayerZhenQiEntity, error) {

	e := &entity.PlayerZhenQiEntity{
		Id:         psco.id,
		PlayerId:   psco.player.GetId(),
		Type:       int32(psco.typ),
		ZhenQiPos:  int32(psco.zhenQiPos),
		Number:     psco.number,
		NumberNum:  psco.numberNum,
		NumberPro:  psco.numberPro,
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerZhenQiObject) GetPlayerId() int64 {
	return psco.player.GetId()
}

func (psco *PlayerZhenQiObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerZhenQiObject) GetZhenFaType() zhenfatypes.ZhenFaType {
	return psco.typ
}

func (psco *PlayerZhenQiObject) GetZhenQiPos() zhenfatypes.ZhenQiType {
	return psco.zhenQiPos
}

func (psco *PlayerZhenQiObject) GetNumber() int32 {
	return psco.number
}

func (psco *PlayerZhenQiObject) GetNumberNum() int32 {
	return psco.numberNum
}

func (psco *PlayerZhenQiObject) GetNumberPro() int32 {
	return psco.numberPro
}

func (psco *PlayerZhenQiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertZhenQiObjectToEntity(psco)
	return e, err
}

func (psco *PlayerZhenQiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerZhenQiEntity)

	psco.id = pse.Id
	psco.typ = zhenfatypes.ZhenFaType(pse.Type)
	psco.zhenQiPos = zhenfatypes.ZhenQiType(pse.ZhenQiPos)
	psco.number = pse.Number
	psco.numberNum = pse.NumberNum
	psco.numberPro = pse.NumberPro
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerZhenQiObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("zhenqi: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
