package player

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/pkg/timeutils"
	"time"
)

const (
	questTaskTime = 3 * time.Second
)

type QuestTask struct {
	p player.Player
}

func (wt *QuestTask) Run() {
	manager := wt.p.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*PlayerQuestDataManager)
	manager.QuestReset()

	//开服七天内判断是否过12点
	kaiFuMuBiaoMap := manager.GetKaiFuMuBiaoMap()
	if len(kaiFuMuBiaoMap) >= questtypes.KaiFuMuBiaoDayMax {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	diff, _ := timeutils.DiffDay(now, openTime)
	// //开服区间外&开服区间玩家登录过
	// if diff+1 > questtypes.KaiFuMuBiaoDayMax && len(kaiFuMuBiaoMap) == 0 {
	// 	return
	// }
	maxOpenDay := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeKaiFuMuBiaoMaxOpenDay)
	if maxOpenDay != 0 && diff+1 > maxOpenDay && len(kaiFuMuBiaoMap) == 0 {
		return
	}
	crossDayTime := manager.GetQuestCrossDayTime()
	crossDayDiff, _ := timeutils.DiffDay(now, crossDayTime)
	if crossDayDiff >= 1 {
		manager.setCrossDayTime(now)
		manager.refreshKaiFuMuBiaoQuest(diff)
		gameevent.Emit(questeventtypes.EventTypeQuestKaiFuMuBiaoCrossDay, wt.p, nil)
	}
}

func (wt *QuestTask) ElapseTime() time.Duration {
	return questTaskTime
}

func CreateQuestTask(p player.Player) *QuestTask {
	qTask := &QuestTask{
		p: p,
	}
	return qTask
}
