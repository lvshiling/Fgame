package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/supremetitle/entity"

	"github.com/pkg/errors"
)

//至尊称号对象
type PlayerSupremeTitleObject struct {
	player     player.Player
	id         int64
	titleId    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerSupremeTitleObject(pl player.Player) *PlayerSupremeTitleObject {
	pmo := &PlayerSupremeTitleObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerSupremeTitleObjectToEntity(pqo *PlayerSupremeTitleObject) (e *entity.PlayerSupremeTitleEntity, err error) {
	e = &entity.PlayerSupremeTitleEntity{
		Id:         pqo.id,
		TitleId:    pqo.titleId,
		PlayerId:   pqo.player.GetId(),
		UpdateTime: pqo.updateTime,
		CreateTime: pqo.createTime,
		DeleteTime: pqo.deleteTime,
	}
	return
}

func (pqo *PlayerSupremeTitleObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerSupremeTitleObject) GetDBId() int64 {
	return pqo.id
}

func (pqo *PlayerSupremeTitleObject) GetTitleId() int32 {
	return pqo.titleId
}

func (pqo *PlayerSupremeTitleObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerSupremeTitleObjectToEntity(pqo)
	return
}

func (pqo *PlayerSupremeTitleObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*entity.PlayerSupremeTitleEntity)

	pqo.id = pqe.Id
	pqo.titleId = pqe.TitleId
	pqo.updateTime = pqe.UpdateTime
	pqo.createTime = pqe.CreateTime
	pqo.deleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerSupremeTitleObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "supreme_title"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}
	pqo.player.AddChangedObject(obj)
	return
}
