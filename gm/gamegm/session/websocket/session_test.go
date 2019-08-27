package websocket_test

import (
	"fmt"
	. "qipai/session/websocket"
	"testing"
	"time"
)

func TestSessionClose(t *testing.T) {
	s := NewWebsocketSession(nil, nil, nil)

	go func() {
		time.AfterFunc(time.Second, func() {
			s.Close()
		})
	}()
Loop:
	for {
		select {
		case <-s.ClosedChannel():

			break Loop
		default:
			fmt.Println("asdsd")
		}
	}

}
