package operation

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/worker"

	log "github.com/Sirupsen/logrus"

	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"

	"fgame/fgame/game/session"
)

type LoadPlayerOperation struct {
	ctx      context.Context
	callback message.ScheduleMessageCallBack
}

func (o *LoadPlayerOperation) Context() context.Context {
	return o.ctx
}

func (o *LoadPlayerOperation) BindUUID() int64 {
	s := session.SessionInContext(o.ctx)
	return s.Player().GetId()
}

func (o *LoadPlayerOperation) Run() (result interface{}, err error) {
	s := session.SessionInContext(o.ctx)
	pl := s.Player().(player.Player)
	beforeTime := global.GetGame().GetTimeService().Now()
	err = pl.Load()
	//加载错误
	if err != nil {
		return
	}

	costTime := global.GetGame().GetTimeService().Now() - beforeTime
	//TODO 统计加载时间
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"costTime": costTime,
		}).Info("player:加载数据")

	result = true
	return
}

//操作回调
func (o *LoadPlayerOperation) OnCallBack(result interface{}, err error) {
	sm := message.NewScheduleMessage(o.callback, o.Context(), result, err)
	processor.GetMessageProcessor().ProcessInternal(sm)
}

func CreateLoadPlayerOperation(ctx context.Context, callback message.ScheduleMessageCallBack, args ...interface{}) worker.Operation {
	o := &LoadPlayerOperation{
		ctx:      ctx,
		callback: callback,
	}
	return o
}
