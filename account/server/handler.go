package server

import (
	"errors"
	"fgame/fgame/account/processor"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
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
				}).Error("server:消息处理错误")

			tterr, ok := terr.(error)
			if ok {
				err = tterr
				return
			}
			err = fmt.Errorf("server:处理消息 panic %#v", terr)
		}
		if err != nil {
			log.WithFields(
				log.Fields{
					"err": err,
				}).Error("server:消息处理错误")

			switch tmsg := msg.(type) {
			case message.SessionMessage:
				tmsg.Session().Close()
				break
			}
		}
	}()
	switch tmsg := msg.(type) {
	case message.SessionMessage:
		err = handleSessionMessage(tmsg.Session(), tmsg.Message())
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
	return err
}
