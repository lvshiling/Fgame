package player

import (
	"fgame/fgame/core/storage"
	fushientity "fgame/fgame/game/fushi/entity"
	fushitypes "fgame/fgame/game/fushi/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

type PlayerFuShiObject struct {
	player     player.Player
	id         int64
	typ        fushitypes.FuShiType
	fushiLevel int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerFuShiObject(pl player.Player) *PlayerFuShiObject {
	pfo := &PlayerFuShiObject{
		player: pl,
	}

	return pfo
}

func convertNewPlayerFuShiObjectToEntity(obj *PlayerFuShiObject) (*fushientity.PlayerFuShiEntity, error) {
	e := &fushientity.PlayerFuShiEntity{
		Id:         obj.id,
		PlayerId:   obj.player.GetId(),
		Typ:        int32(obj.typ),
		FushiLevel: obj.fushiLevel,
		UpdateTime: obj.updateTime,
		CreateTime: obj.createTime,
		DeleteTime: obj.deleteTime,
	}

	return e, nil
}

func (pfo *PlayerFuShiObject) GetType() fushitypes.FuShiType {
	return pfo.typ
}

func (pfo *PlayerFuShiObject) GetFushiLevel() int32 {
	return pfo.fushiLevel
}

func (pfo *PlayerFuShiObject) GetPlayerId() int64 {
	return pfo.player.GetId()
}

func (pfo *PlayerFuShiObject) GetDBId() int64 {
	return pfo.id
}

func (pfo *PlayerFuShiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFuShiObjectToEntity(pfo)
	return

}

func (pfo *PlayerFuShiObject) FromEntity(e storage.Entity) error {
	pfe, _ := e.(*fushientity.PlayerFuShiEntity)
	pfo.id = pfe.Id
	pfo.typ = fushitypes.FuShiType(pfe.Typ)
	pfo.fushiLevel = pfe.FushiLevel
	pfo.updateTime = pfe.UpdateTime
	pfo.createTime = pfe.CreateTime
	pfo.deleteTime = pfe.DeleteTime

	return nil
}

func (pfo *PlayerFuShiObject) SetModified() {
	e, err := pfo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FuShi"))
	}

	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pfo.player.AddChangedObject(obj)

}
