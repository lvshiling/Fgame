package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xuedunentity "fgame/fgame/game/xuedun/entity"

	"github.com/pkg/errors"
)

//血盾对象
type PlayerXueDunObject struct {
	player     player.Player
	id         int64
	playerId   int64
	blood      int64
	number     int32 //阶别
	star       int32
	starNum    int32
	starPro    int32
	culLevel   int32
	culNum     int32
	culPro     int32
	isActive   int32
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerXueDunObject(pl player.Player) *PlayerXueDunObject {
	pmo := &PlayerXueDunObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerXueDunObjectToEntity(pmo *PlayerXueDunObject) (*xuedunentity.PlayerXueDunEntity, error) {

	e := &xuedunentity.PlayerXueDunEntity{
		Id:         pmo.id,
		PlayerId:   pmo.playerId,
		Blood:      pmo.blood,
		Number:     pmo.number,
		Star:       pmo.star,
		StarNum:    pmo.starNum,
		StarPro:    pmo.starPro,
		CulLevel:   pmo.culLevel,
		CulNum:     pmo.culNum,
		CulPro:     pmo.culPro,
		IsActive:   pmo.isActive,
		Power:      pmo.power,
		UpdateTime: pmo.updateTime,
		CreateTime: pmo.createTime,
		DeleteTime: pmo.deleteTime,
	}
	return e, nil
}

func (pmo *PlayerXueDunObject) GetPlayerId() int64 {
	return pmo.playerId
}

func (pmo *PlayerXueDunObject) GetDBId() int64 {
	return pmo.id
}

func (pmo *PlayerXueDunObject) GetBlood() int64 {
	return pmo.blood
}

func (pmo *PlayerXueDunObject) GetNumber() int32 {
	return pmo.number
}

func (pmo *PlayerXueDunObject) GetStar() int32 {
	return pmo.star
}

func (pmo *PlayerXueDunObject) GetStarNum() int32 {
	return pmo.starNum
}

func (pmo *PlayerXueDunObject) GetStarPro() int32 {
	return pmo.starPro
}

func (pmo *PlayerXueDunObject) GetCulLevel() int32 {
	return pmo.culLevel
}

func (pmo *PlayerXueDunObject) GetCulNum() int32 {
	return pmo.culNum
}

func (pmo *PlayerXueDunObject) GetCulPro() int32 {
	return pmo.culPro
}

func (pmo *PlayerXueDunObject) GetIsActive() bool {
	return pmo.isActive == 1
}

func (pmo *PlayerXueDunObject) GetPower() int64 {
	return pmo.power
}

func (pmo *PlayerXueDunObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerXueDunObjectToEntity(pmo)
	return e, err
}

func (pmo *PlayerXueDunObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*xuedunentity.PlayerXueDunEntity)

	pmo.id = pse.Id
	pmo.playerId = pse.PlayerId
	pmo.number = pse.Number
	pmo.blood = pse.Blood
	pmo.star = pse.Star
	pmo.starNum = pse.StarNum
	pmo.starPro = pse.StarPro
	pmo.culLevel = pse.CulLevel
	pmo.culNum = pse.CulNum
	pmo.culPro = pse.CulPro
	pmo.isActive = pse.IsActive
	pmo.power = pse.Power
	pmo.updateTime = pse.UpdateTime
	pmo.createTime = pse.CreateTime
	pmo.deleteTime = pse.DeleteTime
	return nil
}

func (pmo *PlayerXueDunObject) SetModified() {
	e, err := pmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "xuedun"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmo.player.AddChangedObject(obj)
	return
}
