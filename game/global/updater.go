package global

import (
	"context"
	"fgame/fgame/core/storage"
	"sync"
)

//TODO goroutine安全
type GlobalUpdater struct {
	changedObjects *sync.Map
}

func (pu *GlobalUpdater) AddChangedObject(obj storage.Entity) {
	op := CreateSaveObjectOperation(context.Background(), obj)
	GetGame().GetOperationService().PostOperation(op)
	// pu.changedObjects.Store(obj.GetId(), obj)

}

func (pu *GlobalUpdater) Update() (err error) {

	// pu.changedObjects.Range(func(key, val interface{}) bool {
	// 	op := CreateSaveObjectOperation(context.Background(), val)
	// 	GetGame().GetOperationService().PostOperation(op)
	// 	return true
	// })
	// pu.changedObjects = &sync.Map{}
	return nil
}

func NewGlobalUpdater() *GlobalUpdater {
	pu := &GlobalUpdater{}
	pu.changedObjects = &sync.Map{}
	return pu
}
