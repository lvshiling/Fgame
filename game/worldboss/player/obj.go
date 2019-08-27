package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	worldbossentity "fgame/fgame/game/worldboss/entity"
	worldbosstypes "fgame/fgame/game/worldboss/types"

	"github.com/pkg/errors"
)

//周卡
type PlayerBossReliveObject struct {
	player     player.Player
	id         int64
	bossType   worldbosstypes.BossType
	reliveTime int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerBossReliveObject(pl player.Player) *PlayerBossReliveObject {
	o := &PlayerBossReliveObject{
		player: pl,
	}
	return o
}

func convertPlayerBossReliveObjectToEntity(o *PlayerBossReliveObject) (e *worldbossentity.PlayerBossReliveEntity, err error) {
	e = &worldbossentity.PlayerBossReliveEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		BossType:   int32(o.bossType),
		ReliveTime: o.reliveTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}

	return e, nil
}

func (o *PlayerBossReliveObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerBossReliveObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerBossReliveObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerBossReliveObjectToEntity(o)
	return e, err
}

func (o *PlayerBossReliveObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*worldbossentity.PlayerBossReliveEntity)

	o.id = te.Id
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	o.bossType = worldbosstypes.BossType(te.BossType)
	o.reliveTime = te.ReliveTime
	return nil
}

func (o *PlayerBossReliveObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "player_boss_relive"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
