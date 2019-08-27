package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"

	crossentity "fgame/fgame/game/cross/entity"

	"github.com/pkg/errors"
)

//玩家跨服对象
type PlayerCrossObject struct {
	player     player.Player
	id         int64
	crossType  crosstypes.CrossType
	crossArgs  []string
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerCrossObject(pl player.Player) *PlayerCrossObject {
	o := &PlayerCrossObject{
		player: pl,
	}
	return o
}

func convertPlayerCrossObjectToEntity(o *PlayerCrossObject) (e *crossentity.PlayerCrossEntity, err error) {
	args, err := json.Marshal(o.crossArgs)
	if err != nil {
		return
	}
	e = &crossentity.PlayerCrossEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		CrossType:  int32(o.crossType),
		CrossArgs:  string(args),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerCrossObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerCrossObject) GetCrossType() crosstypes.CrossType {
	return o.crossType
}

func (o *PlayerCrossObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerCrossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerCrossObjectToEntity(o)
	return
}

func (o *PlayerCrossObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*crossentity.PlayerCrossEntity)
	o.id = pe.Id
	o.crossType = crosstypes.CrossType(pe.CrossType)
	o.crossArgs = make([]string, 0, 4)
	err = json.Unmarshal([]byte(pe.CrossArgs), &o.crossArgs)
	if err != nil {
		return
	}
	o.updateTime = pe.UpdateTime
	o.createTime = pe.CreateTime
	o.deleteTime = pe.DeleteTime
	return
}

func (o *PlayerCrossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Cross"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("cross:转换数据应该成功"))
	}

	o.player.AddChangedObject(obj)
	return
}
