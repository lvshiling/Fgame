package listener

import (
	"fgame/fgame/core/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/scene/scene"

	logmodel "fgame/fgame/logserver/model"
)

//玩家附加系统觉醒日志
func sceneInfoLog(target event.EventTarget, data event.EventData) (err error) {

	s, ok := target.(scene.Scene)
	if !ok {
		return
	}

	systemScene := &logmodel.SystemScene{}
	systemScene.SystemLogMsg = gamelog.SystemLogMsg()
	systemScene.Content = s.String()

	log.GetLogService().SendLog(systemScene)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeSceneInfoLog, event.EventListenerFunc(sceneInfoLog))
}
