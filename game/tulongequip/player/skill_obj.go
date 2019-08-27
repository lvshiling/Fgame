package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	tulongequipentity "fgame/fgame/game/tulongequip/entity"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fmt"

	"github.com/pkg/errors"
)

//玩家套装技能数据
type PlayerTuLongSuitSkillObject struct {
	player     player.Player
	id         int64
	playerId   int64
	suitType   tulongequiptypes.TuLongSuitType
	level      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerTuLongSuitSkillObject(pl player.Player) *PlayerTuLongSuitSkillObject {
	o := &PlayerTuLongSuitSkillObject{
		player:   pl,
		playerId: pl.GetId(),
	}
	return o
}

func convertPlayerTuLongSuitSkillObjectToEntity(o *PlayerTuLongSuitSkillObject) (*tulongequipentity.PlayerTuLongSuitSkillEntity, error) {

	e := &tulongequipentity.PlayerTuLongSuitSkillEntity{
		Id:         o.id,
		PlayerId:   o.playerId,
		SuitType:   int32(o.suitType),
		Level:      o.level,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerTuLongSuitSkillObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *PlayerTuLongSuitSkillObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerTuLongSuitSkillObject) GetSuitType() tulongequiptypes.TuLongSuitType {
	return o.suitType
}

func (o *PlayerTuLongSuitSkillObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerTuLongSuitSkillObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerTuLongSuitSkillObjectToEntity(o)
	return
}

func (o *PlayerTuLongSuitSkillObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*tulongequipentity.PlayerTuLongSuitSkillEntity)

	o.id = pse.Id
	o.playerId = pse.PlayerId
	o.suitType = tulongequiptypes.TuLongSuitType(pse.SuitType)
	o.level = pse.Level
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return
}

func (o *PlayerTuLongSuitSkillObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TuLongSuitSkill"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}
