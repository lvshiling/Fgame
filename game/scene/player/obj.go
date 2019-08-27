package player

import (
	"fgame/fgame/core/storage"
	coretypes "fgame/fgame/core/types"
	"fmt"

	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	gameentity "fgame/fgame/game/scene/entity"
)

//玩家场景数据
type PlayerSceneObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	MapId       int32
	SceneId     int64
	PosX        float64
	PosY        float64
	PosZ        float64
	LastSceneId int64
	LastMapId   int32
	LastPosX    float64
	LastPosY    float64
	LastPosZ    float64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerSceneObject(pl player.Player) *PlayerSceneObject {
	pso := &PlayerSceneObject{
		player: pl,
	}
	return pso
}

func convertPlayerSceneObjectToEntity(pso *PlayerSceneObject) *gameentity.PlayerSceneEntity {
	e := &gameentity.PlayerSceneEntity{
		Id:          pso.Id,
		PlayerId:    pso.PlayerId,
		MapId:       pso.MapId,
		SceneId:     pso.SceneId,
		PosX:        pso.PosX,
		PosY:        pso.PosY,
		PosZ:        pso.PosZ,
		LastSceneId: pso.LastSceneId,
		LastMapId:   pso.LastMapId,
		LastPosX:    pso.LastPosX,
		LastPosY:    pso.LastPosY,
		LastPosZ:    pso.LastPosZ,
		UpdateTime:  pso.UpdateTime,
		CreateTime:  pso.CreateTime,
		DeleteTime:  pso.DeleteTime,
	}
	return e
}

func (psd *PlayerSceneObject) GetPlayerId() int64 {
	return psd.PlayerId
}

func (psd *PlayerSceneObject) GetSceneId() int64 {
	return psd.SceneId
}

func (psd *PlayerSceneObject) GetMapId() int32 {
	return psd.MapId
}

func (psd *PlayerSceneObject) GetPosX() float64 {
	return psd.PosX
}

func (psd *PlayerSceneObject) GetPosY() float64 {
	return psd.PosY
}
func (psd *PlayerSceneObject) GetPosZ() float64 {
	return psd.PosZ
}
func (psd *PlayerSceneObject) GetLastMapId() int32 {
	return psd.LastMapId
}

func (psd *PlayerSceneObject) GetLastSceneId() int64 {
	return psd.LastSceneId
}

func (psd *PlayerSceneObject) GetLastPosX() float64 {
	return psd.LastPosX
}

func (psd *PlayerSceneObject) GetLastPosY() float64 {
	return psd.LastPosY
}
func (psd *PlayerSceneObject) GetLastPosZ() float64 {
	return psd.LastPosZ
}

func (psd *PlayerSceneObject) GetDBId() int64 {
	return psd.Id
}

func (psd *PlayerSceneObject) ToEntity() (e storage.Entity, err error) {
	e = convertPlayerSceneObjectToEntity(psd)
	return
}

func (psd *PlayerSceneObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*gameentity.PlayerSceneEntity)
	psd.Id = pse.Id
	psd.PlayerId = pse.PlayerId
	psd.MapId = pse.MapId
	psd.SceneId = pse.SceneId
	psd.PosX = pse.PosX
	psd.PosY = pse.PosY
	psd.PosZ = pse.PosZ
	psd.LastSceneId = pse.LastSceneId
	psd.LastMapId = pse.LastMapId
	psd.LastPosX = pse.LastPosX
	psd.LastPosY = pse.LastPosY
	psd.LastPosZ = pse.LastPosZ
	psd.UpdateTime = pse.UpdateTime
	psd.CreateTime = pse.CreateTime
	psd.DeleteTime = pse.DeleteTime
	return
}

func (psd *PlayerSceneObject) SetModified() {
	e, err := psd.ToEntity()
	if err != nil {
		panic(fmt.Errorf("scene: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psd.player.AddChangedObject(obj)
	return
}

func (psd *PlayerSceneObject) GetPos() coretypes.Position {
	pos := coretypes.Position{
		X: psd.PosX,
		Y: psd.PosY,
		Z: psd.PosZ,
	}
	return pos
}

func (psd *PlayerSceneObject) GetLastPos() coretypes.Position {
	pos := coretypes.Position{
		X: psd.LastPosX,
		Y: psd.LastPosY,
		Z: psd.LastPosZ,
	}
	return pos
}
