package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenqientity "fgame/fgame/game/shenqi/entity"
	"fmt"

	"github.com/pkg/errors"
)

//玩家神器数据
type PlayerShenQiObject struct {
	Player     player.Player
	Id         int64
	PlayerId   int64
	LingQiNum  int64
	Power      int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShenQiObject(pl player.Player) *PlayerShenQiObject {
	pio := &PlayerShenQiObject{
		Player:   pl,
		PlayerId: pl.GetId(),
	}
	return pio
}

func convertPlayerShenQiObjectToEntity(pio *PlayerShenQiObject) (*shenqientity.PlayerShenQiEntity, error) {
	e := &shenqientity.PlayerShenQiEntity{
		Id:         pio.Id,
		PlayerId:   pio.PlayerId,
		LingQiNum:  pio.LingQiNum,
		Power:      pio.Power,
		UpdateTime: pio.UpdateTime,
		CreateTime: pio.CreateTime,
		DeleteTime: pio.DeleteTime,
	}
	return e, nil
}

func (pio *PlayerShenQiObject) GetPlayerId() int64 {
	return pio.PlayerId
}

func (pio *PlayerShenQiObject) GetPower() int64 {
	return pio.Power
}

func (pio *PlayerShenQiObject) GetDBId() int64 {
	return pio.Id
}

func (pio *PlayerShenQiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerShenQiObjectToEntity(pio)
	return
}

func (pio *PlayerShenQiObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*shenqientity.PlayerShenQiEntity)
	pio.Id = pse.Id
	pio.PlayerId = pse.PlayerId
	pio.LingQiNum = pse.LingQiNum
	pio.Power = pse.Power
	pio.UpdateTime = pse.UpdateTime
	pio.CreateTime = pse.CreateTime
	pio.DeleteTime = pse.DeleteTime
	return
}

func (pio *PlayerShenQiObject) SetModified() {
	e, err := pio.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ShenQi"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	pio.Player.AddChangedObject(obj)
	return
}
