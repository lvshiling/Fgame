package listener

import (
	"fgame/fgame/core/event"
	crosspbutil "fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
)

//灵童系统属性变更
func playerLingTongSystemPropertyChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTong := p.GetLingTong()
	if lingTong == nil {
		return
	}
	//获取变更的属性
	//获取属性管理器
	lingTongManager := p.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	changedProperties := lingTongManager.GetChangedBattlePropertiesAndReset()
	lingTong.UpdateSystemBattleProperty(changedProperties)

	if p.IsCross() {
		changedProperties = lingTong.GetSystemBattlePropertyChangedTypesAndReset()
		siPlayerSystemBattlePropertyChanged := crosspbutil.BuildLingTongSystemBattlePropertyChanged(changedProperties)
		p.SendCrossMsg(siPlayerSystemBattlePropertyChanged)
		return
	}

	//不在同一个场景,更新属性会引起多线程
	if !scenelogic.CheckIfLingTongAndPlayerSameScene(lingTong) {
		return
	}
	changedProperties = lingTong.GetSystemBattlePropertyChangedTypesAndReset()
	lingTong.Calculate()
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongSystemPropertyChanged, event.EventListenerFunc(playerLingTongSystemPropertyChanged))
}
