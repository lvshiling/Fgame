package event

import (
	"fgame/fgame/core/event"
)

//添加事件快捷操作
func AddEventListener(et event.EventType, el event.EventListener) {
	syncEventEmitter.AddListener(et, el)
}

//发射事件快捷操作
func Emit(et event.EventType, target event.EventTarget, ed event.EventData) (err error) {
	// defer func() {
	// 	if terr := recover(); terr != nil {
	// 		debug.PrintStack()
	// 		log.WithFields(
	// 			log.Fields{
	// 				"err":   terr,
	// 				"stack": string(debug.Stack()),
	// 			}).Error("server:事件分发,错误")
	// 		tterr, ok := terr.(error)
	// 		if ok {
	// 			err = tterr
	// 			return
	// 		}
	// 	}
	// }()
	err = syncEventEmitter.Emit(et, target, ed)
	if err != nil {
		return err
	}
	return
}

var (
	syncEventEmitter event.EventEmitter
)

func init() {
	syncEventEmitter = event.NewEventEmitter()
}
