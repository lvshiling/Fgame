package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/mingge/entity"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//玩家命盘对象
type PlayerMingGePanObject struct {
	player      player.Player
	id          int64
	mingPanType minggetypes.MingGeType
	subType     minggetypes.MingGeAllSubType
	itemMap     map[minggetypes.MingGeSlotType]int32
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func NewPlayerMingGePanObject(pl player.Player) *PlayerMingGePanObject {
	pmo := &PlayerMingGePanObject{
		player: pl,
	}
	return pmo
}

func convertMingGePanObjectToEntity(psco *PlayerMingGePanObject) (*entity.PlayerMingGePanEntity, error) {
	itemList, err := json.Marshal(psco.itemMap)
	if err != nil {
		return nil, err
	}

	e := &entity.PlayerMingGePanEntity{
		Id:         psco.id,
		PlayerId:   psco.player.GetId(),
		Type:       int32(psco.mingPanType),
		SubType:    int32(psco.subType),
		ItemList:   string(itemList),
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerMingGePanObject) GetPlayerId() int64 {
	return psco.player.GetId()
}

func (psco *PlayerMingGePanObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerMingGePanObject) GetMingPanType() minggetypes.MingGeType {
	return psco.mingPanType
}

func (psco *PlayerMingGePanObject) GetSubType() minggetypes.MingGeAllSubType {
	return psco.subType
}

func (psco *PlayerMingGePanObject) GetMingPanItemMap() map[minggetypes.MingGeSlotType]int32 {
	return psco.itemMap
}

func (psco *PlayerMingGePanObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertMingGePanObjectToEntity(psco)
	return e, err
}

func (psco *PlayerMingGePanObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerMingGePanEntity)
	itemMap := make(map[minggetypes.MingGeSlotType]int32)

	if err := json.Unmarshal([]byte(pse.ItemList), &itemMap); err != nil {
		return err
	}

	psco.id = pse.Id
	psco.mingPanType = minggetypes.MingGeType(pse.Type)
	psco.subType = minggetypes.MingGeAllSubType(pse.SubType)
	psco.itemMap = itemMap
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerMingGePanObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("minggepan: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
