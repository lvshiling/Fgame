package player

import (
	"fgame/fgame/core/storage"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenqientity "fgame/fgame/game/shenqi/entity"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/pkg/errors"
)

//玩家器灵数据
type PlayerShenQiQiLingObject struct {
	Player     player.Player
	Id         int64
	PlayerId   int64
	ShenQiType shenqitypes.ShenQiType
	QiLingType shenqitypes.QiLingType
	SlotId     shenqitypes.QiLingSubType
	Level      int32
	UpNum      int32
	UpPro      int32
	ItemId     int32
	BindType   itemtypes.ItemBindType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShenQiQiLingObject(pl player.Player) *PlayerShenQiQiLingObject {
	pio := &PlayerShenQiQiLingObject{
		Player:   pl,
		PlayerId: pl.GetId(),
	}
	return pio
}

func convertPlayerShenQiQiLingObjectToEntity(pio *PlayerShenQiQiLingObject) (*shenqientity.PlayerShenQiQiLingEntity, error) {
	e := &shenqientity.PlayerShenQiQiLingEntity{
		Id:         pio.Id,
		PlayerId:   pio.PlayerId,
		ShenQiType: int32(pio.ShenQiType),
		QiLingType: int32(pio.QiLingType),
		SlotId:     pio.SlotId.SubType(),
		Level:      pio.Level,
		UpNum:      pio.UpNum,
		UpPro:      pio.UpPro,
		ItemId:     pio.ItemId,
		BindType:   int32(pio.BindType),
		UpdateTime: pio.UpdateTime,
		CreateTime: pio.CreateTime,
		DeleteTime: pio.DeleteTime,
	}
	return e, nil
}

func (pio *PlayerShenQiQiLingObject) GetPlayerId() int64 {
	return pio.PlayerId
}

func (pio *PlayerShenQiQiLingObject) GetDBId() int64 {
	return pio.Id
}

func (pio *PlayerShenQiQiLingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerShenQiQiLingObjectToEntity(pio)
	return
}

func (pio *PlayerShenQiQiLingObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*shenqientity.PlayerShenQiQiLingEntity)
	pio.Id = pse.Id
	pio.PlayerId = pse.PlayerId
	pio.ShenQiType = shenqitypes.ShenQiType(pse.ShenQiType)
	pio.QiLingType = shenqitypes.QiLingType(pse.QiLingType)
	pio.SlotId = shenqitypes.CreateQiLingSubType(pio.QiLingType, pse.SlotId)
	pio.Level = pse.Level
	pio.UpNum = pse.UpNum
	pio.UpPro = pse.UpPro
	pio.ItemId = pse.ItemId
	pio.BindType = itemtypes.ItemBindType(pse.BindType)
	pio.UpdateTime = pse.UpdateTime
	pio.CreateTime = pse.CreateTime
	pio.DeleteTime = pse.DeleteTime
	return
}

func (pio *PlayerShenQiQiLingObject) SetModified() {
	e, err := pio.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ShenQiQiLing"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	pio.Player.AddChangedObject(obj)
	return
}

func (pio *PlayerShenQiQiLingObject) IsEmpty() bool {
	return pio.ItemId == 0
}

func (pio *PlayerShenQiQiLingObject) IsFull() bool {
	return pio.ItemId != 0
}

func createShenQiQiLingObject(p player.Player, typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, slotId shenqitypes.QiLingSubType, now int64) *PlayerShenQiQiLingObject {
	obj := NewPlayerShenQiQiLingObject(p)
	obj.Id, _ = idutil.GetId()
	obj.ShenQiType = typ
	obj.QiLingType = subType
	obj.SlotId = slotId
	obj.Level = 0
	obj.ItemId = 0
	obj.CreateTime = now
	return obj
}
