package listener

import (
	"fgame/fgame/core/event"
	crosspbutil "fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
)

//玩家系统属性变更
func playerSystemPropertyChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	//获取变更的属性
	//获取属性管理器
	propertyManager := p.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	changedProperties := propertyManager.GetChangedBattlePropertiesAndReset()
	//更新系统属性
	p.UpdateSystemBattleProperty(changedProperties)
	power := propertyManager.GetForce()
	if p.IsCross() {
		//推送到跨服
		battleChanged := p.GetSystemBattlePropertyChangedTypesAndReset()
		siPlayerSystemBattlePropertyChanged := crosspbutil.BuildSIPlayerSystemBattlePropertyChanged(battleChanged, power)

		p.SendCrossMsg(siPlayerSystemBattlePropertyChanged)
	} else {
		//重新计算战斗属性
		p.UpdateForce(power)
		p.Calculate()
	}
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerSystemPropertyChanged, event.EventListenerFunc(playerSystemPropertyChanged))
}
