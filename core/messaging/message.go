package messaging

type Producer interface {
	Send(topic string, content []byte) error
	Close() error
}

type Consumer interface {
	Start(h Handler) error
	Stop() error
}

type Handler interface {
	Handle(msg []byte) error
}

type HandlerFunc func(msg []byte) error

func (hf HandlerFunc) Handle(msg []byte) error {
	return hf(msg)
}
