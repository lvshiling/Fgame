package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/charge/charge"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	logmodel "fgame/fgame/logserver/model"
)

// 新首充活动时间改变日志
func newFirstChargeTimeChangeLog(target event.EventTarget, data event.EventData) (err error) {

	obj, ok := target.(*charge.NewFirstChargeObject)
	if !ok {
		return
	}
	duration, ok := data.(int32)
	if !ok {
		return
	}

	chargeMerge := &logmodel.NewFirstChargeTime{}
	chargeMerge.SystemLogMsg = gamelog.SystemLogMsg()
	chargeMerge.StartTime = obj.GetStartTime()
	chargeMerge.Duration = duration
	log.GetLogService().SendLog(chargeMerge)
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeNewFirstChargeTimeChangeLog, event.EventListenerFunc(newFirstChargeTimeChangeLog))
}
