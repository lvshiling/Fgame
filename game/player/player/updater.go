package player

import (
	"context"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player/operation"
	"fgame/fgame/game/player/types"
	"sync"
)

//TODO goroutine安全
type PlayerUpdater struct {
	changedObjects *sync.Map
	//TODO 转换为sync.Map
	// changedObjects map[int64]types.PlayerDataEntity
}

func (pu *PlayerUpdater) AddChangedObject(obj types.PlayerDataEntity) {
	pu.changedObjects.Store(obj.GetId(), obj)
	// pu.changedObjects[obj.GetId()] = obj
}

func (pu *PlayerUpdater) Update() (err error) {
	// for _, obj := range pu.changedObjects. {
	// 	//TODO 修改context
	// 	op := operation.CreateSavePlayerObjectOperation(context.Background(), obj)
	// 	global.GetGame().GetOperationService().PostOperation(op)

	// }
	// pu.changedObjects = make(map[int64]types.PlayerDataEntity)
	// return nil

	pu.changedObjects.Range(func(key, val interface{}) bool {
		op := operation.CreateSavePlayerObjectOperation(context.Background(), val)
		global.GetGame().GetOperationService().PostOperation(op)
		return true
	})
	pu.changedObjects = &sync.Map{}
	return nil
}

func NewPlayerUpdater() *PlayerUpdater {
	pu := &PlayerUpdater{}
	pu.changedObjects = &sync.Map{}
	return pu
}
