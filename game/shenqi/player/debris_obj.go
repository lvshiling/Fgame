package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	shenqientity "fgame/fgame/game/shenqi/entity"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/pkg/errors"
)

//玩家碎片数据
type PlayerShenQiDebrisObject struct {
	Player     player.Player
	Id         int64
	PlayerId   int64
	ShenQiType shenqitypes.ShenQiType
	SlotId     shenqitypes.DebrisType
	Level      int32
	UpNum      int32
	UpPro      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShenQiDebrisObject(pl player.Player) *PlayerShenQiDebrisObject {
	pio := &PlayerShenQiDebrisObject{
		Player:   pl,
		PlayerId: pl.GetId(),
	}
	return pio
}

func convertPlayerShenQiDebrisObjectToEntity(pio *PlayerShenQiDebrisObject) (*shenqientity.PlayerShenQiDebrisEntity, error) {
	e := &shenqientity.PlayerShenQiDebrisEntity{
		Id:         pio.Id,
		PlayerId:   pio.PlayerId,
		ShenQiType: int32(pio.ShenQiType),
		SlotId:     int32(pio.SlotId),
		Level:      pio.Level,
		UpNum:      pio.UpNum,
		UpPro:      pio.UpPro,
		UpdateTime: pio.UpdateTime,
		CreateTime: pio.CreateTime,
		DeleteTime: pio.DeleteTime,
	}
	return e, nil
}

func (pio *PlayerShenQiDebrisObject) GetPlayerId() int64 {
	return pio.PlayerId
}

func (pio *PlayerShenQiDebrisObject) GetDBId() int64 {
	return pio.Id
}

func (pio *PlayerShenQiDebrisObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerShenQiDebrisObjectToEntity(pio)
	return
}

func (pio *PlayerShenQiDebrisObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*shenqientity.PlayerShenQiDebrisEntity)
	pio.Id = pse.Id
	pio.PlayerId = pse.PlayerId
	pio.ShenQiType = shenqitypes.ShenQiType(pse.ShenQiType)
	pio.SlotId = shenqitypes.DebrisType(pse.SlotId)
	pio.Level = pse.Level
	pio.UpNum = pse.UpNum
	pio.UpPro = pse.UpPro
	pio.UpdateTime = pse.UpdateTime
	pio.CreateTime = pse.CreateTime
	pio.DeleteTime = pse.DeleteTime
	return
}

func (pio *PlayerShenQiDebrisObject) SetModified() {
	e, err := pio.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ShenQiDebris"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	pio.Player.AddChangedObject(obj)
	return
}

func createShenQiDebrisObject(p player.Player, typ shenqitypes.ShenQiType, slotId shenqitypes.DebrisType, now int64) *PlayerShenQiDebrisObject {
	obj := NewPlayerShenQiDebrisObject(p)
	obj.Id, _ = idutil.GetId()
	obj.ShenQiType = typ
	obj.SlotId = slotId
	obj.Level = 0
	obj.CreateTime = now
	return obj
}
