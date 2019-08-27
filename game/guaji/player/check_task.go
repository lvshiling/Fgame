package player

import (
	gameevent "fgame/fgame/game/event"
	guajieventtypes "fgame/fgame/game/guaji/event/types"
	"fgame/fgame/game/guaji/guaji"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	checkTaskTime = time.Second * 3
)

type checkTask struct {
	p         player.Player
	lastIndex int32
}

func (t *checkTask) Run() {
	gameevent.Emit(guajieventtypes.GuaJiEventTypeGuaJiCheck, t.p, nil)
	//判断是否正在挂机
	m := t.p.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*PlayerGuaJiManager)
	//执行挂机中
	if t.p.IsGuaJi() {
		return
	}
	//检查是否正在挂机
	currentGuaJiData, index := m.GetCurrentGuaJiType()
	if index >= 0 {
		t.lastIndex = index
		m.StopCurrentGuaJi()
		log.WithFields(
			log.Fields{
				"playerId":         t.p.GetId(),
				"currentGuaJiType": currentGuaJiData.GetType().String(),
			}).Info("guaji:退出挂机")
		return
	}
	index = t.lastIndex
	//尝试下一个挂机
	guaJiList := m.GetGuaJiTypeList()
	if index >= int32(len(guaJiList)) {
		index = -1
	}

	for i := index + 1; i < int32(len(guaJiList)); i++ {
		guaJiData := guaJiList[i]
		guaJiType := guaJiData.GetType()
		//检查是否可以停止挂机
		guaJiEnterCheckHandler := guaji.GetGuaJiEnterCheckHandler(guaJiType)
		if guaJiEnterCheckHandler != nil {
			//检查是否可以进入挂机
			flag := guaJiEnterCheckHandler.GuaJiEnterCheck(t.p)
			if !flag {
				log.WithFields(
					log.Fields{
						"playerId":  t.p.GetId(),
						"guaJiType": guaJiData.GetType().String(),
					}).Warn("guaji:进入挂机检查失败,稍后重试")
				continue
			}
		}
		m.StartGuaJi(i)
		log.WithFields(
			log.Fields{
				"playerId":  t.p.GetId(),
				"guaJiType": guaJiData.GetType().String(),
			}).Info("guaji:进入挂机")
		return
	}

	for i := int32(0); i <= index; i++ {
		guaJiData := guaJiList[i]
		guaJiType := guaJiData.GetType()
		//检查是否可以停止挂机
		guaJiEnterCheckHandler := guaji.GetGuaJiEnterCheckHandler(guaJiType)
		if guaJiEnterCheckHandler != nil {
			//检查是否可以进入挂机
			flag := guaJiEnterCheckHandler.GuaJiEnterCheck(t.p)
			if !flag {
				log.WithFields(
					log.Fields{
						"playerId":  t.p.GetId(),
						"guaJiType": guaJiType.String(),
					}).Warn("guaji:进入挂机检查失败,稍后重试")
				continue
			}
		}
		m.StartGuaJi(i)
		log.WithFields(
			log.Fields{
				"playerId":  t.p.GetId(),
				"guaJiType": guaJiType.String(),
			}).Info("guaji:进入挂机")
		return
	}
	return
}

func (t *checkTask) ElapseTime() time.Duration {
	return checkTaskTime
}

func CreateCheckTask(p player.Player) *checkTask {
	t := &checkTask{
		p:         p,
		lastIndex: -1,
	}
	return t
}
