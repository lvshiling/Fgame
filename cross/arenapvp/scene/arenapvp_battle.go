package scene

import (
	coretypes "fgame/fgame/core/types"
	arenapvpeventtypes "fgame/fgame/cross/arenapvp/event/types"
	"fgame/fgame/cross/arenapvp/pbutil"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

type ArenaSceneState int32

const (
	ArenaSceneStateInit ArenaSceneState = iota
	ArenaSceneStateCompete
	ArenaSceneStateEnd
)

const (
	exitTime = 10 * common.SECOND
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
type ArenapvpBattleSceneData interface {
	scene.SceneDelegate
	GetPvpTemp() *gametemplate.ArenapvpTemplate
	GetActivityEndTime() int64
	GetBattlePlayer() (*arenapvpdata.PvpPlayerInfo, *arenapvpdata.PvpPlayerInfo)
	GetWinnerId() int64
	GetEnterPos(playerId int64) coretypes.Position
	Judge(flag bool)
	GetState() ArenaSceneState
}

type arenapvpBattleSceneData struct {
	*scene.SceneDelegateBase
	activityEndTime int64                          //活动结束时间
	pvpTemp         *gametemplate.ArenapvpTemplate //
	state           ArenaSceneState

	//晋级
	winnerId  int64
	battlePl1 *arenapvpdata.PvpPlayerInfo
	battlePl2 *arenapvpdata.PvpPlayerInfo
}

func (sd *arenapvpBattleSceneData) GetState() ArenaSceneState {
	return sd.state
}

func (sd *arenapvpBattleSceneData) OnSceneStart(s scene.Scene) {
	sd.SceneDelegateBase.OnSceneStart(s)

	//初始化
	sd.state = ArenaSceneStateInit
}

func (sd *arenapvpBattleSceneData) OnSceneTick(s scene.Scene) {

	//
	now := global.GetGame().GetTimeService().Now()
	switch sd.state {
	case ArenaSceneStateInit:
		{
			elapseTime := now - s.GetStartTime()
			if elapseTime >= sd.pvpTemp.ZhanDouTime {
				//比赛开始
				sd.gameStart()
			}
			break
		}
	case ArenaSceneStateCompete:
		{
			//时间到了
			pvpEndTime := sd.pvpTemp.GetEndTime(now)
			if now >= pvpEndTime {
				sd.GetScene().Finish(false)
			}
			break
		}
	case ArenaSceneStateEnd:
		{

			// 踢掉失败的人
			elapseTime := now - s.GetFinishTime()
			if elapseTime > int64(exitTime) {
				for _, spl := range s.GetAllPlayers() {
					if spl.GetId() == sd.winnerId {
						continue
					}
					spl.BackLastScene()
				}
			}
		}
		break
	}

	return
}

func (sd *arenapvpBattleSceneData) OnScenePlayerEnter(s scene.Scene, pl scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}

	if sd.state == ArenaSceneStateCompete {
		pl.StartArenapvpBattle()
	}

	player1, player2 := sd.GetBattlePlayer()
	scMsg := pbutil.BuildSCArenapvpSceneData(pl.GetArenapvpReliveTimes(), sd.GetScene(), int32(sd.GetPvpTemp().GetArenapvpType()), player1, player2)
	s.BroadcastMsg(scMsg)
	return
}

func (sd *arenapvpBattleSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}

	p.StopArenapvpBattle()

	if sd.state != ArenaSceneStateCompete {
		return
	}

	if active {
		sd.Judge(false)
		return
	}
	return
}

func (sd *arenapvpBattleSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}

	// if sd.state != ArenaSceneStateCompete {
	// 	return
	// }

	// sd.Judge(false)
	return
}

func (sd *arenapvpBattleSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}
	if !success {
		sd.Judge(true)
	}
	for _, p := range sd.GetScene().GetAllPlayers() {
		p.StopArenapvpBattle()
	}
	sd.state = ArenaSceneStateEnd
	// pvp结束
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpBattleSceneFinish, sd, nil)
	return
}

func (sd *arenapvpBattleSceneData) OnSceneStop(s scene.Scene) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}
	//
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpBattleSceneStop, sd, nil)
	return
}

func (sd *arenapvpBattleSceneData) GetPvpTemp() *gametemplate.ArenapvpTemplate {
	return sd.pvpTemp
}

func (sd *arenapvpBattleSceneData) GetActivityEndTime() int64 {
	return sd.activityEndTime
}

func (sd *arenapvpBattleSceneData) GetBattlePlayer() (*arenapvpdata.PvpPlayerInfo, *arenapvpdata.PvpPlayerInfo) {
	return sd.battlePl1, sd.battlePl2
}

func (sd *arenapvpBattleSceneData) GetWinnerId() int64 {
	return sd.winnerId
}

func (sd *arenapvpBattleSceneData) GetEnterPos(playerId int64) coretypes.Position {
	var bornPos coretypes.Position
	if sd.battlePl1 != nil && sd.battlePl1.PlayerId == playerId {
		return sd.GetPvpTemp().GetPos1()
	}

	if sd.battlePl2 != nil && sd.battlePl2.PlayerId == playerId {
		return sd.GetPvpTemp().GetPos2()
	}

	return bornPos
}

func CreateArenapvpBattleSceneData(endTime int64, pvpTemp *gametemplate.ArenapvpTemplate, battlePl1, battlePl2 *arenapvpdata.PvpPlayerInfo) ArenapvpBattleSceneData {
	sd := &arenapvpBattleSceneData{}
	sd.activityEndTime = endTime
	sd.pvpTemp = pvpTemp
	sd.state = ArenaSceneStateInit
	sd.battlePl1 = battlePl1
	sd.battlePl2 = battlePl2
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}

//游戏开始
func (sd *arenapvpBattleSceneData) gameStart() {
	if sd.state != ArenaSceneStateInit {
		panic(fmt.Errorf("arena:游戏开始,状态应该为[%d]", sd.state))
	}

	sd.state = ArenaSceneStateCompete

	scMsg := pbutil.BuildSCArenapvpBattleStart()
	sd.GetScene().BroadcastMsg(scMsg)

	for _, p := range sd.GetScene().GetAllPlayers() {
		p.StartArenapvpBattle()
	}

	//判断结果
	sd.Judge(false)
}

//裁定结果
// NOTICE:judge里面的finish都要置true
func (sd *arenapvpBattleSceneData) Judge(isEnd bool) {
	if sd.state != ArenaSceneStateCompete {
		return
	}

	s := sd.GetScene()
	// 1v1对局
	if sd.battlePl1 != nil && sd.battlePl2 != nil {
		// 若两方同时掉线，则在比赛开始时，随机其中一个玩家进入下一轮
		plNum := s.GetAllPlayers()
		if len(plNum) == 0 {
			isWin := mathutils.RandomHit(common.MAX_RATE, common.MAX_RATE/2)
			if isWin {
				sd.winnerId = sd.battlePl1.PlayerId
			} else {
				sd.winnerId = sd.battlePl2.PlayerId
			}

			s.Finish(true)
			return
		}

		pl1 := s.GetPlayer(sd.battlePl1.PlayerId)
		pl2 := s.GetPlayer(sd.battlePl2.PlayerId)
		//当时间结束时
		if isEnd {
			//只有一个玩家
			if pl1 == nil {
				sd.winnerId = pl2.GetId()
				s.Finish(true)
				return
			}

			if pl2 == nil {
				sd.winnerId = pl1.GetId()
				s.Finish(true)
				return
			}

			//命数多的获胜
			if pl1.GetArenapvpReliveTimes() == pl2.GetArenapvpReliveTimes() {
				goto HP
			}
			if pl1.GetArenapvpReliveTimes() < pl2.GetArenapvpReliveTimes() {
				sd.winnerId = pl1.GetId()
				s.Finish(true)
				return
			} else {
				sd.winnerId = pl2.GetId()
				s.Finish(true)
				return
			}

		HP:
			if pl1.GetHP() == pl2.GetHP() {
				goto FORCE
			}
			//血量多的获胜
			if pl1.GetHP() < pl2.GetHP() {
				sd.winnerId = pl2.GetId()
				s.Finish(true)
				return
			} else {
				sd.winnerId = pl1.GetId()
				s.Finish(true)
				return
			}

		FORCE:
			// 战力较高的玩家进入下一轮
			if sd.battlePl1.Force < sd.battlePl2.Force {
				sd.winnerId = pl2.GetId()
				s.Finish(true)
				return
			} else {
				sd.winnerId = pl1.GetId()
				s.Finish(true)
				return
			}
		} else {
			//复活次数
			if pl1 != nil {
				remain1 := sd.pvpTemp.GetRemainReliveTimes(pl1.GetArenapvpReliveTimes())
				if pl1.IsDead() && remain1 <= 0 {
					sd.winnerId = sd.battlePl2.PlayerId
					s.Finish(true)
					return
				}
			}

			if pl2 != nil {
				remain2 := sd.pvpTemp.GetRemainReliveTimes(pl2.GetArenapvpReliveTimes())
				if pl2.IsDead() && remain2 <= 0 {
					sd.winnerId = sd.battlePl1.PlayerId
					s.Finish(true)
					return
				}
			}
			return
		}
	}

	// 轮空局（没有对手）
	if sd.battlePl1 == nil {
		sd.winnerId = sd.battlePl2.PlayerId
		s.Finish(true)
		return
	}
	if sd.battlePl2 == nil {
		sd.winnerId = sd.battlePl1.PlayerId
		s.Finish(true)
		return
	}
}
