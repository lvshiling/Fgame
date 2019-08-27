package operation

import (
	"context"
	"fgame/fgame/core/worker"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type tabler interface {
	TableName() string
}

type SavePlayerObjectOperation struct {
	ctx context.Context
	e   types.PlayerDataEntity
}

func (o *SavePlayerObjectOperation) Context() context.Context {
	return o.ctx
}

func (o *SavePlayerObjectOperation) BindUUID() int64 {
	return o.e.GetPlayerId()
}

func (o *SavePlayerObjectOperation) Run() (result interface{}, err error) {
	if err = global.GetGame().GetDB().DB().Save(o.e).Error; err != nil {
		te := o.e.(tabler)
		err = fmt.Errorf("%s error %s", te.TableName(), err.Error())
		return
	}
	return
}

//操作回调
func (o *SavePlayerObjectOperation) OnCallBack(result interface{}, err error) {
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerEntity": o.e,
				"err":          err,
			}).Error("player:保存玩家数据,失败")
		eventData := exceptioneventtypes.CreateDBExceptionEventData(o.e.TableName(), o.e, err.Error())
		gameevent.Emit(exceptioneventtypes.ExceptionEventTypeDBException, nil, eventData)
	}

}

func CreateSavePlayerObjectOperation(ctx context.Context, args ...interface{}) worker.Operation {
	arg := args[0]
	e := arg.(types.PlayerDataEntity)
	o := &SavePlayerObjectOperation{
		ctx: ctx,
		e:   e,
	}
	return o
}
