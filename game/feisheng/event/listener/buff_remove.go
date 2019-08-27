package listener

import (
	"fgame/fgame/core/event"
	buffcommon "fgame/fgame/game/buff/common"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	feishenglogic "fgame/fgame/game/feisheng/logic"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//buff移除
func buffRemove(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)
	buffObject := data.(buffcommon.BuffObject)

	pl, ok := bo.(player.Player)
	if !ok {
		return
	}

	buffId := buffObject.GetBuffId()
	feiShengSuccessBuffId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFeiShengSuccessBuffId)
	feiShengFaildBuffId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFeiShengFaildBuffId)
	if buffId != feiShengSuccessBuffId && buffId != feiShengFaildBuffId {
		return
	}

	//渡劫成功
	result := false
	if buffId == feiShengSuccessBuffId {
		feishenglogic.FeiShengPropertyChanged(pl)
		result = true
	}

	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	scMsg := pbutil.BuildSCFeiShengDuJieNotice(result, feiManager.GetFeiShengLevel())
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffRemove, event.EventListenerFunc(buffRemove))
}
