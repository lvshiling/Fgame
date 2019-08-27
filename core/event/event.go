package event

//事件类型
type EventType interface{}

//事件目标
type EventTarget interface{}

//事件数据
type EventData interface{}

//事件
// type Event struct {
// 	typ  EventType
// 	data EventData
// }

// func (e *Event) EventType() EventType {
// 	return e.typ
// }

// func (e *Event) EventData() EventData {
// 	return e.data
// }

//TODO 复用
// func createEvent(t EventType, d EventData) *Event {
// 	e := &Event{}
// 	e.typ = t
// 	e.data = d
// 	return e
// }

//事件监听
type EventListener interface {
	Handle(EventTarget, EventData) (err error)
}

type EventListenerFunc func(EventTarget, EventData) (err error)

func (elf EventListenerFunc) Handle(t EventTarget, d EventData) (err error) {
	return elf(t, d)
}

type EventEmitter interface {
	Emit(e EventType, target EventTarget, d EventData) (err error)
	AddListener(e EventType, el EventListener)
}

type eventEmitter struct {
	listenerMap map[EventType][]EventListener
}

func (em *eventEmitter) Emit(t EventType, target EventTarget, d EventData) (err error) {
	hs, exist := em.listenerMap[t]
	if !exist {
		return
	}
	// e := createEvent(t, d)
	for _, h := range hs {
		err = h.Handle(target, d)
		if err != nil {
			return
		}
	}
	return nil
}

func (em *eventEmitter) AddListener(e EventType, el EventListener) {
	hs, _ := em.listenerMap[e]
	hs = append(hs, el)
	em.listenerMap[e] = hs
}

func NewEventEmitter() EventEmitter {
	em := &eventEmitter{}
	em.listenerMap = make(map[EventType][]EventListener)
	return em
}
