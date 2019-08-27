package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xiantaoentity "fgame/fgame/game/xiantao/entity"

	"github.com/pkg/errors"
)

//玩家仙桃大会对象
type PlayerXianTaoObject struct {
	player           player.Player
	Id               int64
	PlayerId         int64
	JuniorPeachCount int32
	HighPeachCount   int32
	RobCount         int32
	BeRobCount       int32
	EndTime          int64
	UpdateTime       int64
	CreateTime       int64
	DeleteTime       int64
}

func NewPlayerXianTaoObject(pl player.Player) *PlayerXianTaoObject {
	pdo := &PlayerXianTaoObject{
		player: pl,
	}
	return pdo
}

func (pdo *PlayerXianTaoObject) GetPlayerId() int64 {
	return pdo.PlayerId
}

func (pdo *PlayerXianTaoObject) GetDBId() int64 {
	return pdo.Id
}

func convertObjectToEntity(pdo *PlayerXianTaoObject) (*xiantaoentity.PlayerXianTaoEntity, error) {
	e := &xiantaoentity.PlayerXianTaoEntity{
		Id:               pdo.Id,
		PlayerId:         pdo.PlayerId,
		JuniorPeachCount: pdo.JuniorPeachCount,
		HighPeachCount:   pdo.HighPeachCount,
		RobCount:         pdo.RobCount,
		BeRobCount:       pdo.BeRobCount,
		EndTime:          pdo.EndTime,
		UpdateTime:       pdo.UpdateTime,
		CreateTime:       pdo.CreateTime,
		DeleteTime:       pdo.DeleteTime,
	}
	return e, nil
}

func (pdo *PlayerXianTaoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertObjectToEntity(pdo)
	return
}

func (pdo *PlayerXianTaoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*xiantaoentity.PlayerXianTaoEntity)

	pdo.Id = pse.Id
	pdo.PlayerId = pse.PlayerId
	pdo.JuniorPeachCount = pse.JuniorPeachCount
	pdo.HighPeachCount = pse.HighPeachCount
	pdo.RobCount = pse.RobCount
	pdo.BeRobCount = pse.BeRobCount
	pdo.EndTime = pse.EndTime
	pdo.UpdateTime = pse.UpdateTime
	pdo.CreateTime = pse.CreateTime
	pdo.DeleteTime = pse.DeleteTime
	return nil
}

func (pdo *PlayerXianTaoObject) SetModified() {
	e, err := pdo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "XianTao"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pdo.player.AddChangedObject(obj)
	return
}
