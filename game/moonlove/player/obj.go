package player

import (
	"fgame/fgame/core/storage"
	moonloveentity "fgame/fgame/game/moonlove/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//月下情缘数据对象
type PlayerMoonloveObject struct {
	player          player.Player
	id              int64
	charmNum        int32
	generousNum     int32
	enterTime       int64
	preActivityTime int64
	updateTime      int64
	createTime      int64
	deleteTime      int64
}

func CreatePlayerMoonloveObject(pl player.Player) *PlayerMoonloveObject {
	pmobj := &PlayerMoonloveObject{
		player: pl,
	}

	return pmobj
}

//数据库id
func (o *PlayerMoonloveObject) GetDBId() int64 {
	return o.id
}

//对象转换为数据库实体
func (o *PlayerMoonloveObject) ToEntity() (e storage.Entity, err error) {
	e = &moonloveentity.PlayerMoonloveEntity{
		Id:              o.id,
		PlayerId:        o.player.GetId(),
		CharmNum:        o.charmNum,
		GenerousNum:     o.generousNum,
		PreActivityTime: o.preActivityTime,
		UpdateTime:      o.updateTime,
		CreateTime:      o.createTime,
		DeleteTime:      o.deleteTime,
	}
	return e, nil
}

//数据库实体转对象
func (o *PlayerMoonloveObject) FromEntity(e storage.Entity) (err error) {
	pmle, _ := e.(*moonloveentity.PlayerMoonloveEntity)
	o.id = pmle.Id
	o.charmNum = pmle.CharmNum
	o.generousNum = pmle.GenerousNum
	o.preActivityTime = pmle.PreActivityTime
	o.updateTime = pmle.UpdateTime
	o.createTime = pmle.CreateTime
	o.deleteTime = pmle.DeleteTime
	return nil
}

//提交修改
func (o *PlayerMoonloveObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Moonlove"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerMoonloveObject) GetPlayerId() int64 {
	return o.player.GetId()
}
