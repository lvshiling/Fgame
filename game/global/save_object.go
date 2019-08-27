package global

import (
	"context"
	"fgame/fgame/core/storage"
	"fgame/fgame/core/worker"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"

	log "github.com/Sirupsen/logrus"
)

type SaveObjectOperation struct {
	ctx context.Context
	e   storage.Entity
}

func (o *SaveObjectOperation) Context() context.Context {
	return o.ctx
}

func (o *SaveObjectOperation) Run() (result interface{}, err error) {
	if err = GetGame().GetDB().DB().Save(o.e).Error; err != nil {
		return
	}
	//TODO 写进缓存
	return
}

//操作回调
func (o *SaveObjectOperation) OnCallBack(result interface{}, err error) {
	//记录日志
	if err != nil {
		log.WithFields(
			log.Fields{
				"entity": o.e,
				"table":  o.e.TableName(),
				"err":    err,
			}).Error("global:保存数据,失败")
		eventData := exceptioneventtypes.CreateDBExceptionEventData(o.e.TableName(), o.e, err.Error())
		gameevent.Emit(exceptioneventtypes.ExceptionEventTypeDBException, nil, eventData)
	}

}

func CreateSaveObjectOperation(ctx context.Context, args ...interface{}) worker.Operation {
	arg := args[0]
	e := arg.(storage.Entity)
	o := &SaveObjectOperation{
		ctx: ctx,
		e:   e,
	}
	return o
}
