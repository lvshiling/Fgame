package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/supremetitle/entity"

	"github.com/pkg/errors"
)

//至尊称号对象
type PlayerWearSupremeTitleObject struct {
	player     player.Player
	id         int64
	titleWear  int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerWearSupremeTitleObject(pl player.Player) *PlayerWearSupremeTitleObject {
	pmo := &PlayerWearSupremeTitleObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerWearSupremeTitleObjectToEntity(pqo *PlayerWearSupremeTitleObject) (e *entity.PlayerWearSupremeTitleEntity, err error) {
	e = &entity.PlayerWearSupremeTitleEntity{
		Id:         pqo.id,
		TitleWear:  pqo.titleWear,
		PlayerId:   pqo.player.GetId(),
		UpdateTime: pqo.updateTime,
		CreateTime: pqo.createTime,
		DeleteTime: pqo.deleteTime,
	}
	return
}

func (pqo *PlayerWearSupremeTitleObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerWearSupremeTitleObject) GetDBId() int64 {
	return pqo.id
}

func (pqo *PlayerWearSupremeTitleObject) GetTitleWear() int32 {
	return pqo.titleWear
}

func (pqo *PlayerWearSupremeTitleObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerWearSupremeTitleObjectToEntity(pqo)
	return
}

func (pqo *PlayerWearSupremeTitleObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*entity.PlayerWearSupremeTitleEntity)

	pqo.id = pqe.Id
	pqo.titleWear = pqe.TitleWear
	pqo.updateTime = pqe.UpdateTime
	pqo.createTime = pqe.CreateTime
	pqo.deleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerWearSupremeTitleObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "supreme_title_wear"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}
	pqo.player.AddChangedObject(obj)
	return
}
