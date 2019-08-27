package listener

import (
	"fgame/fgame/core/event"
	crosslogic "fgame/fgame/game/cross/logic"
	gameevent "fgame/fgame/game/event"
	guajieventtypes "fgame/fgame/game/guaji/event/types"
	"fgame/fgame/game/guaji/pbutil"
	playerguaji "fgame/fgame/game/guaji/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

//玩家开始挂机
func guaJiStart(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	m := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	guaJiData, _ := m.GetCurrentGuaJiType()
	guaJiType := guaJiData.GetType()
	guaJiAiType, flag := guaJiType.GuaJiAIType()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"guaJiType": guaJiType.String(),
			}).Warn("guaji:开始挂机,不存在ai")
		return
	}
	if !pl.StartGuaJi(guaJiAiType) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"guaJiType": guaJiAiType.String(),
			}).Warn("guaji:开始挂机,失败")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"guaJiType": guaJiAiType.String(),
		}).Info("guaji:开始挂机")

	scCurrentGuaJi := pbutil.BuildSCCurrentGuaJi(int32(guaJiType))
	pl.SendMsg(scCurrentGuaJi)

	//跨服中退出跨服
	if pl.IsCross() {
		crosslogic.PlayerExitCross(pl)
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().IsWorld() {
		return
	}
	//返回世界地图
	pl.BackLastScene()
	return
}

func init() {
	gameevent.AddEventListener(guajieventtypes.GuaJiEventTypeGuaJiStart, event.EventListenerFunc(guaJiStart))
}
