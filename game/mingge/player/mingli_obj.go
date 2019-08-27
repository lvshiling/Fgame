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

type MingLiInfo struct {
	Slot           minggetypes.MingLiSlotType
	MingGeProperty minggetypes.MingGePropertyType
	Times          int32
}

func newMingLiInfo(slot minggetypes.MingLiSlotType, mingGeProperty minggetypes.MingGePropertyType, times int32) *MingLiInfo {
	d := &MingLiInfo{
		Slot:           slot,
		MingGeProperty: mingGeProperty,
		Times:          times,
	}
	return d
}

func (d *MingLiInfo) GetSlot() minggetypes.MingLiSlotType {
	return d.Slot
}

func (d *MingLiInfo) GetMingGeProperty() minggetypes.MingGePropertyType {
	return d.MingGeProperty
}

func (d *MingLiInfo) GetTimes() int32 {
	return d.Times
}

//玩家命理对象
type PlayerMingLiObject struct {
	player       player.Player
	id           int64
	mingGongType minggetypes.MingGongType
	subType      minggetypes.MingGongAllSubType
	mingLiMap    map[minggetypes.MingLiSlotType]*MingLiInfo
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerMingLiObject(pl player.Player) *PlayerMingLiObject {
	pmo := &PlayerMingLiObject{
		player: pl,
	}
	return pmo
}

func convertObjectToEntity(psco *PlayerMingLiObject) (*entity.PlayerMingGeMingLiEntity, error) {

	mingLiList, err := json.Marshal(psco.mingLiMap)
	if err != nil {
		return nil, err
	}

	e := &entity.PlayerMingGeMingLiEntity{
		Id:         psco.id,
		PlayerId:   psco.player.GetId(),
		Type:       int32(psco.mingGongType),
		SubType:    int32(psco.subType),
		MingLiList: string(mingLiList),
		UpdateTime: psco.updateTime,
		CreateTime: psco.createTime,
		DeleteTime: psco.deleteTime,
	}
	return e, nil
}

func (psco *PlayerMingLiObject) GetPlayerId() int64 {
	return psco.player.GetId()
}

func (psco *PlayerMingLiObject) GetDBId() int64 {
	return psco.id
}

func (psco *PlayerMingLiObject) GetMingGongType() minggetypes.MingGongType {
	return psco.mingGongType
}

func (psco *PlayerMingLiObject) GetSubType() minggetypes.MingGongAllSubType {
	return psco.subType
}

func (psco *PlayerMingLiObject) GetMingLiMap() map[minggetypes.MingLiSlotType]*MingLiInfo {
	return psco.mingLiMap
}

func (psco *PlayerMingLiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertObjectToEntity(psco)
	return e, err
}

func (psco *PlayerMingLiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerMingGeMingLiEntity)
	mingLiMap := make(map[minggetypes.MingLiSlotType]*MingLiInfo)

	if err := json.Unmarshal([]byte(pse.MingLiList), &mingLiMap); err != nil {
		return err
	}

	psco.id = pse.Id
	psco.mingGongType = minggetypes.MingGongType(pse.Type)
	psco.subType = minggetypes.MingGongAllSubType(pse.SubType)
	psco.mingLiMap = mingLiMap
	psco.updateTime = pse.UpdateTime
	psco.createTime = pse.CreateTime
	psco.deleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerMingLiObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("mingli: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}
