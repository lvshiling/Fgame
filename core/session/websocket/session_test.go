package websocket_test

import (
	. "fgame/fgame/core/session/websocket"
	"testing"
)

func TestSessionNewPanic(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("nil conn should panic")
		}
	}()
	NewWebsocketSession(nil, nil, nil)
}

//TODO 测试真正的conn
