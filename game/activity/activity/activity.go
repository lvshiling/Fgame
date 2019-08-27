package activity

import (
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

//活动参加处理器
type ActivityAttendHandler interface {
	Attend(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error)
}

type ActivityAttendHandlerFunc func(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error)

func (f ActivityAttendHandlerFunc) Attend(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	return f(pl, activityTemplate, args...)
}

var (
	activityHandlerMap = map[activitytypes.ActivityType]ActivityAttendHandler{}
)

func RegisterActivityHandler(activityType activitytypes.ActivityType, h ActivityAttendHandler) {
	_, ok := activityHandlerMap[activityType]
	if ok {
		panic(fmt.Errorf("activity:repeat register %s activity", activityType.String()))
	}
	activityHandlerMap[activityType] = h
}

func GetActivityHandler(activityType activitytypes.ActivityType) ActivityAttendHandler {
	h, ok := activityHandlerMap[activityType]
	if !ok {
		return nil
	}
	return h
}
