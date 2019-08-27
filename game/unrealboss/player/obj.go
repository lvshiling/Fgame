package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	unrealbossentity "fgame/fgame/game/unrealboss/entity"

	"github.com/pkg/errors"
)

//幻境BOSS对象
type PlayerUnrealBossObject struct {
	player        player.Player
	id            int64
	pilaoNum      int32
	buyPiLaoNum   int32
	buyPiLaoTimes int32
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewPlayerUnrealBossObject(pl player.Player) *PlayerUnrealBossObject {
	o := &PlayerUnrealBossObject{
		player: pl,
	}
	return o
}

func convertNewPlayerUnrealBossObjectToEntity(o *PlayerUnrealBossObject) (*unrealbossentity.PlayerUnrealBossEntity, error) {
	e := &unrealbossentity.PlayerUnrealBossEntity{
		Id:            o.id,
		PlayerId:      o.player.GetId(),
		PiLaoNum:      o.pilaoNum,
		BuyPiLaoNum:   o.buyPiLaoNum,
		BuyPiLaoTimes: o.buyPiLaoTimes,
		UpdateTime:    o.updateTime,
		DeleteTime:    o.deleteTime,
		CreateTime:    o.createTime,
	}
	return e, nil
}

func (o *PlayerUnrealBossObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerUnrealBossObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerUnrealBossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerUnrealBossObjectToEntity(o)
	return e, err
}

func (o *PlayerUnrealBossObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*unrealbossentity.PlayerUnrealBossEntity)

	o.id = pse.Id
	o.pilaoNum = pse.PiLaoNum
	o.buyPiLaoNum = pse.BuyPiLaoNum
	o.buyPiLaoTimes = pse.BuyPiLaoTimes
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerUnrealBossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "UnrealBoss"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
