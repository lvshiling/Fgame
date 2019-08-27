package dispatch

import "fgame/fgame/core/session"

type middlewareHandler struct {
	handlers []Handler
}

func (mh *middlewareHandler) Handle(s session.Session, msg interface{}) (err error) {
	for _, h := range mh.handlers {
		err = h.Handle(s, msg)
		if err != nil {
			return
		}
	}
	return nil
}

func NewMiddlewareHandler(handlers ...Handler) Handler {
	return &middlewareHandler{handlers: handlers}
}
