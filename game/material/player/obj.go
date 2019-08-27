package player

import (
	"fgame/fgame/core/storage"
	materialentity "fgame/fgame/game/material/entity"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//材料副本对象
type PlayerMaterialObject struct {
	player       player.Player
	id           int64
	materialType materialtypes.MaterialType
	useTimes     int32
	group        int32
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func CreateNewPlayerMaterialObject(pl player.Player) *PlayerMaterialObject {
	newObj := &PlayerMaterialObject{
		player: pl,
	}
	return newObj
}

//数据库id
func (o *PlayerMaterialObject) GetDBId() int64 {
	return o.id
}

//对象转换为数据库实体
func (o *PlayerMaterialObject) ToEntity() (e storage.Entity, err error) {
	e = &materialentity.PlayerMaterialEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		MaterialType: int32(o.materialType),
		UseTimes:     o.useTimes,
		Group:        o.group,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

//数据库实体转对象
func (o *PlayerMaterialObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*materialentity.PlayerMaterialEntity)
	o.id = pse.Id
	o.materialType = materialtypes.MaterialType(pse.MaterialType)
	o.useTimes = pse.UseTimes
	o.group = pse.Group
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

//提交修改
func (o *PlayerMaterialObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(fmt.Errorf("material: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//刷新 挑战次数
func (o *PlayerMaterialObject) refreshUseTimes(nowTime int64) error {
	isSame, err := timeutils.IsSameFive(o.updateTime, nowTime)
	if err != nil {
		return err
	}

	if !isSame {
		o.useTimes = 0
		o.updateTime = nowTime
		o.SetModified()
	}
	return nil
}

func (o *PlayerMaterialObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerMaterialObject) GetMaterialType() materialtypes.MaterialType {
	return o.materialType
}

func (o *PlayerMaterialObject) GetUseTimes() int32 {
	return o.useTimes
}

func (o *PlayerMaterialObject) GetGroup() int32 {
	return o.group
}
