package scene

import (
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenatemplate "fgame/fgame/game/arena/template"
	"fgame/fgame/game/common/common"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

const (
	countDownTime       = 10 * common.SECOND
	competeTime         = 30 * common.MINUTE
	failedCountDownTime = 12 * common.SECOND
)

type ArenaSceneState int32

const (
	ArenaSceneStateInit ArenaSceneState = iota
	ArenaSceneStateCompete
	ArenaSceneStateEnd
)

var (
	stateMap = map[ArenaSceneState]string{
		ArenaSceneStateInit:    "初始化",
		ArenaSceneStateCompete: "比赛中",
		ArenaSceneStateEnd:     "结束",
	}
)

func (s ArenaSceneState) String() string {
	return stateMap[s]
}

//竞技场景数据
type ArenaSceneData interface {
	scene.SceneDelegate
	GetTeam1() *ArenaTeam
	GetTeam2() *ArenaTeam
	GetLastTime() int64
	GetState() ArenaSceneState
	GetWinnerTeam() *ArenaTeam
	GetRandomTreasureId() int64
}

//3v3场景数据
func CreateArenaScene(mapTemplate *gametemplate.MapTemplate, team1 *ArenaTeam, team2 *ArenaTeam, endTime int64) (s scene.Scene) {
	asd := &arenaSceneData{}
	asd.SceneDelegateBase = scene.NewSceneDelegateBase()
	asd.team1 = team1
	asd.team2 = team2
	asd.state = ArenaSceneStateInit
	asd.countDownTime = int64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().BattleTime) * int64(common.SECOND)
	s = scene.CreateArenaScene(int32(mapTemplate.TemplateId()), endTime, asd)
	return s
}

type arenaSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//队伍1数据
	team1 *ArenaTeam
	//队伍2数据
	team2    *ArenaTeam
	lastTime int64
	state    ArenaSceneState
	//获胜队伍
	winnerTeam *ArenaTeam
	//随机奖励
	randomTreasureId int64
	//倒计时
	countDownTime int64
}

func (sd *arenaSceneData) GetTeam1() *ArenaTeam {
	return sd.team1
}

func (sd *arenaSceneData) GetTeam2() *ArenaTeam {
	return sd.team2
}
func (sd *arenaSceneData) GetLastTime() int64 {
	return sd.lastTime
}

func (sd *arenaSceneData) GetState() ArenaSceneState {
	return sd.state
}

func (sd *arenaSceneData) GetWinnerTeam() *ArenaTeam {
	return sd.winnerTeam
}

func (sd *arenaSceneData) GetRandomTreasureId() int64 {
	return sd.randomTreasureId
}

func (sd *arenaSceneData) OnSceneStart(s scene.Scene) {
	//初始化
	sd.state = ArenaSceneStateInit
	sd.s = s
	now := global.GetGame().GetTimeService().Now()
	sd.lastTime = now
}

func (sd *arenaSceneData) GetScene() scene.Scene {
	return sd.s
}

func (sd *arenaSceneData) OnSceneTick(s scene.Scene) {
	//TODO 限制连进来的时间
	elapseTime := global.GetGame().GetTimeService().Now() - sd.lastTime
	switch sd.state {
	case ArenaSceneStateInit:
		if elapseTime >= int64(sd.countDownTime) {
			//比赛开始
			sd.gameStart()
		}
		break
	case ArenaSceneStateCompete:
		//时间到了
		if elapseTime >= int64(competeTime) {
			sd.s.Finish(false)
		}
		break
	case ArenaSceneStateEnd:
		//TODO 踢人
		if elapseTime >= int64(failedCountDownTime) {
			//踢掉不是获胜的
			if sd.team1 != sd.winnerTeam {
				sd.kickTeam(sd.team1)
			}
			if sd.team2 != sd.winnerTeam {
				sd.kickTeam(sd.team2)
			}
		}

		break
	}
}

func (sd *arenaSceneData) kickTeam(teamObj *ArenaTeam) {

	for _, mem := range teamObj.GetTeam().GetMemberList() {
		memObj := sd.s.GetSceneObject(mem.GetPlayerId())
		if memObj == nil {
			continue
		}
		memPl, ok := memObj.(scene.Player)
		if !ok {
			continue
		}
		memPl.BackLastScene()
	}
}

func (sd *arenaSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	now := global.GetGame().GetTimeService().Now()
	sd.lastTime = now
	sd.state = ArenaSceneStateEnd

	randomTreasureId := int64(0)
	// winTeam := sd.winnerTeam
	// if winTeam != nil {
	// level := winTeam.GetCurrent() + 1
	// arenaTemplate := arenatemplate.GetArenaTemplateService().GetArenaTemplate(level)
	// if arenaTemplate != nil {
	// 	randomIndex := -1
	// 	if len(arenaTemplate.GetExtraItemMap()) > 0 {
	// 		randomIndex = mathutils.RandomRange(0, len(winTeam.GetTeam().GetMemberList()))
	// 		randomTreasureId = winTeam.GetTeam().GetMemberList()[randomIndex].GetPlayerId()
	// 	}
	// }
	// }
	sd.randomTreasureId = randomTreasureId
	gameevent.Emit(arenaeventtypes.EventTypeArenaSceneEnd, sd, nil)
	//检查关闭
	sd.checkStop()
}

func (sd *arenaSceneData) OnSceneStop(s scene.Scene) {

}

func (sd *arenaSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
}
func (sd *arenaSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {
}
func (sd *arenaSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
}

func (sd *arenaSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
}

func (sd *arenaSceneData) OnSceneBiologyAllDead(s scene.Scene) {
}

func (sd *arenaSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

func (sd *arenaSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.state != ArenaSceneStateCompete {
		return
	}
	sd.judge()
}

func (sd *arenaSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	p.StopArenaBattle()

	gameevent.Emit(arenaeventtypes.EventTypeArenaScenePlayerExit, p, active)

	switch sd.state {
	case ArenaSceneStateInit:
		return
	case ArenaSceneStateCompete:
		sd.judge()
		return
	case ArenaSceneStateEnd:
		sd.checkStop()
		return
	}
}

func (sd *arenaSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

func (sd *arenaSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.state == ArenaSceneStateCompete {
		p.StartArenaBattle()
	}
	//切换pk模式
	p.SwitchPkState(pktypes.PkStateGroup, pktypes.PkCommonCampDefault)
	gameevent.Emit(arenaeventtypes.EventTypeArenaScenePlayerEnter, sd, p)
}

func (sd *arenaSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {

}
func (sd *arenaSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {

}

func (sd *arenaSceneData) OnSceneRefreshGroup(s scene.Scene, group int32) {

}

//游戏开始
func (sd *arenaSceneData) gameStart() {
	if sd.state != ArenaSceneStateInit {
		panic(fmt.Errorf("arena:游戏开始,状态应该为[%d]", sd.state))
	}
	sd.state = ArenaSceneStateCompete
	now := global.GetGame().GetTimeService().Now()
	sd.lastTime = now
	gameevent.Emit(arenaeventtypes.EventTypeArenaSceneStart, sd, nil)
	for _, p := range sd.s.GetAllPlayers() {
		p.StartArenaBattle()
	}
	//发送游戏开始
	sd.judge()
}

//裁定结果
func (sd *arenaSceneData) judge() {
	if sd.state != ArenaSceneStateCompete {
		return
	}
	team1Failed := sd.ifTeamFailed(sd.team1)
	team2Failed := sd.ifTeamFailed(sd.team2)
	if !team1Failed && !team2Failed {
		return
	}
	if team1Failed && team2Failed {
		sd.s.Finish(true)
		return
	}
	if team1Failed {
		sd.winnerTeam = sd.team2
		sd.s.Finish(true)
		return
	}
	sd.winnerTeam = sd.team1
	sd.s.Finish(true)
}

//检查关闭
func (sd *arenaSceneData) checkStop() {
	if len(sd.s.GetAllPlayers()) == 0 {
		sd.s.Stop(true, false)
	}
}

func (sd *arenaSceneData) ifTeamFailed(teamData *ArenaTeam) bool {
	for _, mem := range teamData.GetTeam().GetMemberList() {
		if mem.GetStatus() != MemberStatusOnline {
			continue
		}
		so := sd.s.GetSceneObject(mem.GetPlayerId())
		if so == nil {
			continue
		}
		memPl, ok := so.(scene.Player)
		if !ok {
			continue
		}
		if memPl.IsDead() && mem.GetRemainReliveTime() <= 0 {
			continue
		}
		return false
	}
	return true
}
