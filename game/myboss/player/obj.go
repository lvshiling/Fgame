package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	mybossentity "fgame/fgame/game/myboss/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//个人BOSS对象
type PlayerMyBossObject struct {
	player     player.Player
	id         int64
	attendMap  map[int32]int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerMyBossObject(pl player.Player) *PlayerMyBossObject {
	o := &PlayerMyBossObject{
		player: pl,
	}
	return o
}

func convertNewPlayerMyBossObjectToEntity(o *PlayerMyBossObject) (*mybossentity.PlayerMyBossEntity, error) {
	attendMap, err := json.Marshal(o.attendMap)
	if err != nil {
		return nil, err
	}

	e := &mybossentity.PlayerMyBossEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		AttendMap:  string(attendMap),
		UpdateTime: o.updateTime,
		DeleteTime: o.deleteTime,
		CreateTime: o.createTime,
	}
	return e, nil
}

func (o *PlayerMyBossObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerMyBossObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerMyBossObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerMyBossObjectToEntity(o)
	return e, err
}

func (o *PlayerMyBossObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*mybossentity.PlayerMyBossEntity)

	dataMap := make(map[int32]int32)
	err := json.Unmarshal([]byte(pse.AttendMap), &dataMap)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.attendMap = dataMap
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerMyBossObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "MyBoss"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
