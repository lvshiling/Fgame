package player

import (
	"fgame/fgame/core/storage"
	itemskilltypes "fgame/fgame/game/itemskill/types"
	itemskillentity "fgame/fgame/game/itemskill/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//物品技能对象
type PlayerItemSkillObject struct {
	player     player.Player
	Id         int64
	Typ        itemskilltypes.ItemSkillType
	Level      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerItemSkillObject(pl player.Player) *PlayerItemSkillObject {
	pso := &PlayerItemSkillObject{
		player: pl,
	}
	return pso
}

func (pxfo *PlayerItemSkillObject) GetPlayerId() int64 {
	return pxfo.player.GetId()
}

func (pxfo *PlayerItemSkillObject) GetDBId() int64 {
	return pxfo.Id
}

func (pxfo *PlayerItemSkillObject) ToEntity() (e storage.Entity, err error) {
	e = &itemskillentity.PlayerItemSkillEntity{
		Id:         pxfo.Id,
		PlayerId:   pxfo.player.GetId(),
		Typ:        int32(pxfo.Typ),
		Level:      pxfo.Level,
		UpdateTime: pxfo.UpdateTime,
		CreateTime: pxfo.CreateTime,
		DeleteTime: pxfo.DeleteTime,
	}
	return e, err
}

func (pxfo *PlayerItemSkillObject) FromEntity(e storage.Entity) error {
	pxfe, _ := e.(*itemskillentity.PlayerItemSkillEntity)
	pxfo.Id = pxfe.Id
	pxfo.Typ = itemskilltypes.ItemSkillType(pxfe.Typ)
	pxfo.Level = pxfe.Level
	pxfo.UpdateTime = pxfe.UpdateTime
	pxfo.CreateTime = pxfe.CreateTime
	pxfo.DeleteTime = pxfe.DeleteTime
	return nil
}

func (pxfo *PlayerItemSkillObject) SetModified() {
	e, err := pxfo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("itemskill: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pxfo.player.AddChangedObject(obj)
	return
}
