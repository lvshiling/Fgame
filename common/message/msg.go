package message

import (
	"context"
	"fgame/fgame/core/session"
)

//TODO 优化
type Message interface{}

type SessionMessage interface {
	Session() session.Session
	Message() interface{}
	IsCross() bool
}

type sessionMessage struct {
	session session.Session
	msg     interface{}
	cross   bool
}

func (sm *sessionMessage) Session() session.Session {
	return sm.session
}
func (sm *sessionMessage) Message() interface{} {
	return sm.msg
}
func (sm *sessionMessage) IsCross() bool {
	return sm.cross
}

func NewSessionMessage(s session.Session, msg interface{}) SessionMessage {
	sm := &sessionMessage{}
	sm.session = s
	sm.msg = msg
	sm.cross = false
	return sm
}

func NewCrossSessionMessage(s session.Session, msg interface{}) SessionMessage {
	sm := &sessionMessage{}
	sm.session = s
	sm.msg = msg
	sm.cross = true
	return sm
}

type ScheduleMessageCallBack func(ctx context.Context, result interface{}, err error) error

type ScheduleMessage interface {
	Context() context.Context
	Run() error
}

type scheduleMessage struct {
	ctx      context.Context
	result   interface{}
	err      error
	callBack ScheduleMessageCallBack
}

func (sm *scheduleMessage) Context() context.Context {
	return sm.ctx
}

func (sm *scheduleMessage) Run() error {
	return sm.callBack(sm.ctx, sm.result, sm.err)
}

func NewScheduleMessage(callBack ScheduleMessageCallBack, ctx context.Context, result interface{}, err error) ScheduleMessage {
	sm := &scheduleMessage{}
	sm.callBack = callBack
	sm.ctx = ctx
	sm.result = result
	sm.err = err
	return sm
}
