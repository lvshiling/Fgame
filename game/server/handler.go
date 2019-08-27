package server

import (
	"errors"
	"fgame/fgame/common/codec"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	gamecommon "fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"runtime/debug"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

var (
	ErrorMessageCanNotHandle = errors.New("error message can not handle")
)

func handleMessage(msg message.Message) (err error) {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"err":   terr,
					"stack": exceptionContent,
				}).Error("message:消息处理panic")
			tterr, ok := terr.(error)
			if ok {
				err = tterr
			} else {
				err = fmt.Errorf("message:处理消息 panic %#v", terr)
			}
			//发送异常日志
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
		if err != nil {
			codeErr, ok := err.(gamecommon.Error)
			if ok {
				switch tmsg := msg.(type) {
				case message.SessionMessage:
					s := gamesession.SessionInContext(tmsg.Session().Context())
					playerlogic.SendSessionSystemMessage(s, codeErr.Code())
					err = nil
					break
				case message.ScheduleMessage:
					{
						p := scene.PlayerInContext(tmsg.Context())
						if p != nil {
							pl, ok := p.(scene.Player)
							if !ok {
								return
							}
							playerlogic.SendSystemMessage(pl, codeErr.Code())
						}
						err = nil
					}
					break
				}
			} else {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Error("handle:消息处理错误")
				switch tmsg := msg.(type) {
				case message.SessionMessage:
					tmsg.Session().Close()
					break
				case message.ScheduleMessage:
					{
						p := scene.PlayerInContext(tmsg.Context())
						if p != nil {
							pl, ok := p.(scene.Player)
							if !ok {
								return
							}
							pl.Close(nil)
						}
					}
					break
				}
			}
		}
	}()
	switch tmsg := msg.(type) {
	case message.SessionMessage:
		if tmsg.IsCross() {
			err = handleCrossMessage(tmsg.Session(), tmsg.Message())
		} else {
			err = handleSessionMessage(tmsg.Session(), tmsg.Message())
		}
		return err
	case message.ScheduleMessage:
		err = tmsg.Run()
		//TODO 断线处理
		return err
	default:
		err = ErrorMessageCanNotHandle
		return
	}
	return
}

func handleSessionMessage(s session.Session, msg interface{}) error {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	//认证过
	if pl != nil {
		m, ok := msg.(*codec.Message)
		if !ok {
			return nil
		}

		//加载过的
		tpl, ok := pl.(player.Player)
		if ok {
			//跨服中
			if tpl.IsCross() {
				//判断是否是跨服消息
				if processor.IsProxy(m.MessageType) {
					tmsg := m.Body.(proto.Message)
					tpl.SendCrossMsg(tmsg)
					return nil
				}

			}
		}
	}
	err := processor.GetDispatch().Handle(s, msg)
	if err == nil {
		return nil
	}

	return err
}

func handleCrossMessage(s session.Session, msg interface{}) error {
	gcs := gamesession.SessionInContext(s.Context())
	m, ok := msg.(*codec.Message)
	if !ok {
		return nil
	}
	pl := gcs.Player()
	//加载过的
	tpl := pl.(player.Player)

	//转发消息
	if !codec.IsCrossMsg(m.MessageType) {
		tmsg := m.Body.(proto.Message)
		tpl.SendMsg(tmsg)
		return nil
	}
	err := processor.GetCrossDispatch().Handle(s, msg)
	if err == nil {
		return nil
	}

	return err
}
