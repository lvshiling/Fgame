package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	pktypes "fgame/fgame/game/pk/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//玩家被杀
func playerKilled(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	killedPlayer, ok := data.(scene.Player)
	if !ok {
		return
	}

	scPlayerKill := pbutil.BuildSCPlayerKill(killedPlayer)
	pl.SendMsg(scPlayerKill)

	s := pl.GetScene()
	//和平模式被击杀
	if killedPlayer.GetPkState() == pktypes.PkStatePeach {
		//添加buff
		pkBuffId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypePKProtectBuff)
		scenelogic.AddBuff(killedPlayer, pkBuffId, 0, common.MAX_RATE)
	}

	if !s.MapTemplate().IsWorld() {
		return
	}

	//陈楚楠说只有世界地图加红名值

	killedPlayerRedState := killedPlayer.GetPkRedState()

	switch killedPlayerRedState {
	case pktypes.PkRedStateInit:
		pl.Kill(true)
	default:
		pl.Kill(false)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypePlayerKilled, event.EventListenerFunc(playerKilled))
}
