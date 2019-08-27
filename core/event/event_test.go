package event_test

import (
	. "fgame/fgame/core/event"
	"testing"
)

type ctxEventType string
type eEventType string

const (
	ctxKey ctxEventType = "aa"
)

const (
	eKey eEventType = "aa"
)

func TestEventTypeEqual(t *testing.T) {
	var keyMap map[EventType]string = make(map[EventType]string)
	keyMap[ctxKey] = "hello"
	if keyMap[eKey] == keyMap[ctxKey] {
		t.Fatalf("event type[%#v,%T],[%#v,%T]should different", eKey, eKey, ctxKey, ctxKey)
	}
}
