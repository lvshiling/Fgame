package server

import (
	"errors"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	processor "fgame/fgame/cross/processor"
	gamecommon "fgame/fgame/game/common/common"
	playerlogic "fgame/fgame/game/player/logic"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"runtime/debug"

	log "github.com/Sirupsen/logrus"
)

var (
	ErrorMessageCanNotHandle = errors.New("error message can not handle")
)

func handleMessage(msg message.Message) (err error) {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			log.WithFields(
				log.Fields{
					"err":   terr,
					"stack": string(debug.Stack()),
				}).Error("message:消息处理panic")
			tterr, ok := terr.(error)
			if ok {
				err = tterr
				return
			}
			err = fmt.Errorf("message:处理消息 panic %#v", terr)
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
						s := gamesession.SessionInContext(tmsg.Context())
						if s != nil {
							playerlogic.SendSessionSystemMessage(s, codeErr.Code())
							err = nil
						}
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
						s := gamesession.SessionInContext(tmsg.Context())
						if s != nil {
							s.Close(true)
						}
					}
					break
				}
			}
		}
	}()
	switch tmsg := msg.(type) {
	case message.SessionMessage:
		err = handleSessionMessage(tmsg.Session(), tmsg.Message())
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
	err := processor.GetDispatch().Handle(s, msg)
	if err == nil {
		return nil
	}

	return nil
}
