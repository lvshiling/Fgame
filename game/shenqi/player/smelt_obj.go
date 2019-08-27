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

//玩家淬炼数据
type PlayerShenQiSmeltObject struct {
	Player     player.Player
	Id         int64
	PlayerId   int64
	ShenQiType shenqitypes.ShenQiType
	SlotId     shenqitypes.SmeltType
	Level      int32
	UpNum      int32
	UpPro      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerShenQiSmeltObject(pl player.Player) *PlayerShenQiSmeltObject {
	pio := &PlayerShenQiSmeltObject{
		Player:   pl,
		PlayerId: pl.GetId(),
	}
	return pio
}

func convertPlayerShenQiSmeltObjectToEntity(pio *PlayerShenQiSmeltObject) (*shenqientity.PlayerShenQiSmeltEntity, error) {
	e := &shenqientity.PlayerShenQiSmeltEntity{
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

func (pio *PlayerShenQiSmeltObject) GetPlayerId() int64 {
	return pio.PlayerId
}

func (pio *PlayerShenQiSmeltObject) GetDBId() int64 {
	return pio.Id
}

func (pio *PlayerShenQiSmeltObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerShenQiSmeltObjectToEntity(pio)
	return
}

func (pio *PlayerShenQiSmeltObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*shenqientity.PlayerShenQiSmeltEntity)
	pio.Id = pse.Id
	pio.PlayerId = pse.PlayerId
	pio.ShenQiType = shenqitypes.ShenQiType(pse.ShenQiType)
	pio.SlotId = shenqitypes.SmeltType(pse.SlotId)
	pio.Level = pse.Level
	pio.UpNum = pse.UpNum
	pio.UpPro = pse.UpPro
	pio.UpdateTime = pse.UpdateTime
	pio.CreateTime = pse.CreateTime
	pio.DeleteTime = pse.DeleteTime
	return
}

func (pio *PlayerShenQiSmeltObject) SetModified() {
	e, err := pio.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "ShenQiSmelt"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	pio.Player.AddChangedObject(obj)
	return
}

func createShenQiSmeltObject(p player.Player, typ shenqitypes.ShenQiType, slotId shenqitypes.SmeltType, now int64) *PlayerShenQiSmeltObject {
	obj := NewPlayerShenQiSmeltObject(p)
	obj.Id, _ = idutil.GetId()
	obj.ShenQiType = typ
	obj.SlotId = slotId
	obj.Level = 0
	obj.CreateTime = now
	return obj
}
