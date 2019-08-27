package message

type Handler interface {
	HandleMessage(msg Message) (err error)
}

type HandlerFunc func(msg Message) (err error)

func (hf HandlerFunc) HandleMessage(msg Message) (err error) {
	return hf(msg)
}
