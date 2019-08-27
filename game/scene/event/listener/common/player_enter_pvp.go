package common

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"

	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

func playerEnterPVP(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	if !p.IsMountHidden() {
		p.MountHidden(true)
	}

	//附加buff
	pvpBuffId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEnterPvpBuff)
	if pvpBuffId > 0 {
		scenelogic.AddBuff(p, pvpBuffId, p.GetId(), common.MAX_RATE)
	}

	playerEnterPVP := pbutil.BuildPlayerEnterPVP(p)
	p.SendMsg(playerEnterPVP)
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypePlayerEnterPVP, event.EventListenerFunc(playerEnterPVP))
}
