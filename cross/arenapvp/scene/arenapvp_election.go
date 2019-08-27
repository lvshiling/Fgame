package scene

import (
	arenapvpeventtypes "fgame/fgame/cross/arenapvp/event/types"
	"fgame/fgame/cross/arenapvp/pbutil"
	activitytypes "fgame/fgame/game/activity/types"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	arenapvpscenetypes "fgame/fgame/game/arenapvp/types/scene"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"
	"math/rand"
)

//竞技场景数据
type ArenapvpSceneData interface {
	scene.SceneDelegate
	GetPvpTemp() *gametemplate.ArenapvpTemplate

	GetActivityEndTime() int64
	GetElectionIndex() int32
	IfLineup() bool
	GetLastLuckRewInfo() (int64, []scene.Player)
}

type aliveTimeData struct {
	enterTime    int64 //
	prePointTime int64 //上次积分奖励时间
}

type arenapvpSceneData struct {
	*scene.SceneDelegateBase
	activityEndTime int64                          //活动结束时间
	pvpTemp         *gametemplate.ArenapvpTemplate //
	sceneIndex      int32                          //会场号
	state           ArenaSceneState

	//海选
	pointTimeMap map[int64]*aliveTimeData
	winnerIdMap  map[int64]struct{}

	//上次添加假人时间
	lastAddRobotTime int64
	// 上次幸运奖时间
	lastLuckyRewTime    int64
	lastLuckyPlayerList []scene.Player
}

func (sd *arenapvpSceneData) OnSceneStart(s scene.Scene) {
	sd.SceneDelegateBase.OnSceneStart(s)

	//初始化
	sd.state = ArenaSceneStateInit
}

func (sd *arenapvpSceneData) OnSceneTick(s scene.Scene) {

	now := global.GetGame().GetTimeService().Now()
	switch sd.state {
	case ArenaSceneStateInit:
		{
			sd.state = ArenaSceneStateCompete
			break
		}
	case ArenaSceneStateCompete:
		{
			//时间到了
			electionEndTime := sd.pvpTemp.GetEndTime(now)
			if now >= electionEndTime {
				s.Finish(true)
				return
			}

			// 定时添加机器人
			sd.tickAddRobot()

			// 幸运奖
			sd.tickLuckyRew(s)

			//海选积分
			sd.tickElectionJiFen(s)
			break
		}
	case ArenaSceneStateEnd:
		{
			// 踢掉失败的人
			elapseTime := now - s.GetFinishTime()
			if elapseTime > int64(exitTime) {
				for _, spl := range s.GetAllPlayers() {
					_, ok := sd.winnerIdMap[spl.GetId()]
					if ok {
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

func (sd *arenapvpSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}

	p.StartArenapvpBattle()
	// 如果存在机器人，替换掉
	sd.checkPvpRobot()

	now := global.GetGame().GetTimeService().Now()
	aliveData := &aliveTimeData{
		enterTime:    now,
		prePointTime: 0,
	}
	sd.pointTimeMap[p.GetId()] = aliveData

	rankMap := sd.GetScene().GetAllRanks()
	scMsg := pbutil.BuildSCArenapvpElectionSceneData(p, rankMap)
	p.SendMsg(scMsg)
	return
}

func (sd *arenapvpSceneData) checkPvpRobot() {
	allPlayersNum := int32(len(sd.GetScene().GetAllPlayers()))
	if allPlayersNum >= sd.pvpTemp.GetPVPNum() {
		for _, robot := range sd.GetScene().GetAllRobots() {
			sd.GetScene().RemoveSceneObject(robot, true)
			break
		}
	}
}

func (sd *arenapvpSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}

	p.StopArenapvpBattle()
	delete(sd.pointTimeMap, p.GetId())
	s.RemovePlayer(arenapvpscenetypes.ArenapvpSceneRankTypePoint, p.GetId())

	if active {
		p.UpdateActivityRankValue(activitytypes.ActivityTypeArenapvp, arenapvpscenetypes.ArenapvpSceneRankTypePoint, 0)
	}
	return
}

func (sd *arenapvpSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}

	//海选淘汰
	remain := sd.pvpTemp.GetRemainReliveTimes(p.GetArenapvpReliveTimes())
	if p.IsDead() && remain <= 0 {
		if !p.IsRobot() {
			scMsg := pbutil.BuildSCArenapvpElectionFailedNotice()
			p.SendMsg(scMsg)
		}
		p.BackLastScene()
	}

	return
}

func (sd *arenapvpSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}
	sd.state = ArenaSceneStateEnd

	// 设置停战
	for _, p := range sd.GetScene().GetAllPlayers() {
		p.StopArenapvpBattle()
	}

	// 计算晋级玩家
	s.Sort()
	rankList := s.GetAllRankList(arenapvpscenetypes.ArenapvpSceneRankTypePoint)
	playerIdList := make([]int64, 0, len(rankList))
	for _, rankInfo := range rankList {
		playerIdList = append(playerIdList, rankInfo.GetPlayerId())
	}

	// 获胜玩家列表
	var winnerList []*arenapvpdata.PvpPlayerInfo
	for _, plId := range playerIdList {
		spl := s.GetPlayer(plId)
		//玩家不存在则顺延下去
		if spl == nil {
			continue
		}

		winCount := int32(len(winnerList))
		ranking := winCount + 1
		isWin := false
		if winCount < sd.GetPvpTemp().GetWinnerCount() {
			isWin = true
			winnerList = append(winnerList, arenapvpdata.ConvertToPvpPlayerInfo(spl))
			sd.winnerIdMap[plId] = struct{}{}
		}
		if !spl.IsRobot() {
			//发送奖励
			isMsg := pbutil.BuildISArenapvpResultElection(isWin, int32(ranking), int32(sd.GetPvpTemp().GetArenapvpType()))
			spl.SendMsg(isMsg)

			scMsg := pbutil.BuildSCArenapvpElectionEnd(isWin)
			spl.SendMsg(scMsg)
		}
	}

	// pvp结束
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpElectionSceneFinish, sd, winnerList)
	return
}

func (sd *arenapvpSceneData) OnSceneStop(s scene.Scene) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("arenapvp:pvp应该是同一个场景"))
	}
	//
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpElectionSceneStop, sd, nil)
	return
}

func (sd *arenapvpSceneData) GetPvpTemp() *gametemplate.ArenapvpTemplate {
	return sd.pvpTemp
}

func (sd *arenapvpSceneData) GetActivityEndTime() int64 {
	return sd.activityEndTime
}

func (sd *arenapvpSceneData) GetElectionIndex() int32 {
	return sd.sceneIndex
}

func (sd *arenapvpSceneData) IfLineup() bool {
	allPlayersNum := int32(len(sd.GetScene().GetAllPlayers()))
	return allPlayersNum >= sd.pvpTemp.GetPVPNum()
}

func (sd *arenapvpSceneData) GetLastLuckRewInfo() (int64, []scene.Player) {
	return sd.lastLuckyRewTime, sd.lastLuckyPlayerList
}

//添加机器人
func (sd *arenapvpSceneData) tickAddRobot() {
	allPlayersNum := int32(len(sd.GetScene().GetAllPlayers()))
	if allPlayersNum >= sd.pvpTemp.GetPVPNum()-1 {
		return
	}

	pvpConstantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
	robotNum := int32(sd.GetScene().GetNumOfAllRobot())
	if robotNum >= pvpConstantTemp.RobotMax {
		return
	}

	realPlNum := allPlayersNum - robotNum
	if realPlNum == 0 || realPlNum >= pvpConstantTemp.ZhenshiPlayerRobot {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if now-sd.lastAddRobotTime < pvpConstantTemp.RobotAddTime {
		return
	}

	robot := sd.createArenapvpRobot()
	pos := sd.pvpTemp.GetMapTemp().RandomPosition()
	robot.SetEnterPos(pos)
	sd.GetScene().AddSceneObject(robot)

	sd.lastAddRobotTime = now
}

//幸运奖
func (sd *arenapvpSceneData) tickLuckyRew(s scene.Scene) {

	now := global.GetGame().GetTimeService().Now()
	pvpConstantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
	if sd.lastLuckyRewTime == 0 {
		diffFirst := now - s.GetStartTime()
		if diffFirst <= int64(pvpConstantTemp.XingYunFirstTime) {
			return
		}
	} else {
		diffRew := now - sd.lastLuckyRewTime
		if diffRew <= int64(pvpConstantTemp.XingYunTime) {
			return
		}
	}
	//记录时间
	sd.lastLuckyRewTime = now
	sd.lastLuckyPlayerList = []scene.Player{}

	allPlayerMap := s.GetAllPlayers()
	curPlNum := int32(len(allPlayerMap))
	if curPlNum == 0 {
		return
	}
	drewNum := pvpConstantTemp.XingYunPlayerCount
	realNum, robotNum := pvpConstantTemp.GetLuckyPlayerNumber(drewNum)

	// 假人
	if robotNum > 0 {
		count := int32(0)
		for _, spl := range allPlayerMap {
			if !spl.IsRobot() {
				continue
			}
			count += 1
			sd.lastLuckyPlayerList = append(sd.lastLuckyPlayerList, spl)
			if count >= robotNum {
				break
			}
		}
	}

	// 真实玩家
	if realNum > 0 {
		var realPlayerIdList []int64
		var realWeights []int64
		for plId, spl := range allPlayerMap {
			if spl.IsRobot() {
				continue
			}
			realPlayerIdList = append(realPlayerIdList, plId)
			realWeights = append(realWeights, 1)
		}

		realLen := int32(len(realPlayerIdList))
		if realNum > realLen {
			realNum = realLen
		}

		if realNum != 0 {
			randomRealIndexList := mathutils.RandomListFromWeights(realWeights, realNum)
			for _, randomIndex := range randomRealIndexList {
				randomPlId := realPlayerIdList[randomIndex]
				randomPl := allPlayerMap[randomPlId]
				isMsg := pbutil.BuildISAreanapvpElectionLuckyRew()
				randomPl.SendMsg(isMsg)

				sd.lastLuckyPlayerList = append(sd.lastLuckyPlayerList, randomPl)
			}
			return
		}
	}

}

// 存活加晋级积分
func (sd *arenapvpSceneData) tickElectionJiFen(s scene.Scene) {
	now := global.GetGame().GetTimeService().Now()
	pvpConstantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
	for _, spl := range s.GetAllPlayers() {
		timeData, ok := sd.pointTimeMap[spl.GetId()]
		if !ok {
			continue
		}

		isAddPoint := false
		if timeData.prePointTime == 0 {
			diffFirst := now - timeData.enterTime
			if diffFirst <= int64(pvpConstantTemp.CunhuoFristTiem) {
				continue
			}
			isAddPoint = true
		} else {

			diffRew := now - timeData.prePointTime
			if diffRew <= int64(pvpConstantTemp.CunhuoTime) {
				continue
			}
			isAddPoint = true
		}

		if isAddPoint {
			val := spl.GetActivityRankValue(activitytypes.ActivityTypeArenapvp, arenapvpscenetypes.ArenapvpSceneRankTypePoint)
			val += int64(pvpConstantTemp.CunhuoJifen)
			spl.UpdateActivityRankValue(activitytypes.ActivityTypeArenapvp, arenapvpscenetypes.ArenapvpSceneRankTypePoint, val)
			timeData.prePointTime = now
		}
	}
}

const (
	maxServer        = 10
	minPropertyRatio = 9500
	maxPropertyRatio = 11000
)

func (sd *arenapvpSceneData) createArenapvpRobot() scene.RobotPlayer {
	robotProperties := sd.countRobotProperties()
	robotForce := propertylogic.CulculateAllForce(robotProperties)

	//TODO:zrc 临时给1 后面加随机
	reliveTimes := sd.pvpTemp.RebornCountMax
	serverId := int32(rand.Intn(maxServer) + 1)
	platform := int32(1)
	return robot.GetRobotService().CreateArenapvpRobot(platform, serverId, robotProperties, reliveTimes, robotForce)
}

// 机器人属性
func (sd *arenapvpSceneData) countRobotProperties() map[propertytypes.BattlePropertyType]int64 {
	robotProperties := make(map[propertytypes.BattlePropertyType]int64)

	// 所有真实玩家属性
	allPlayerProperties := make(map[int32]int64)
	realNum := int32(0)
	for _, spl := range sd.GetScene().GetAllPlayers() {
		if spl.IsRobot() {
			continue
		}

		realNum += 1
		for typ, val := range spl.GetAllSystemBattleProperties() {
			allPlayerProperties[typ] += val
		}
	}

	pvpConstntTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
	randomRatio := pvpConstntTemp.GetRandomRobotRatio()
	for typ, val := range allPlayerProperties {
		avgVal := math.Ceil(float64(val) / float64(realNum))
		newVal := int64(math.Ceil(avgVal * randomRatio / float64(common.MAX_RATE)))
		robotProperties[propertytypes.BattlePropertyType(typ)] = newVal
	}

	return robotProperties
}

func CreateArenapvpSceneData(endTime int64, pvpTemp *gametemplate.ArenapvpTemplate, sceneIndex int32) ArenapvpSceneData {
	sd := &arenapvpSceneData{}
	sd.activityEndTime = endTime
	sd.pointTimeMap = make(map[int64]*aliveTimeData)
	sd.winnerIdMap = make(map[int64]struct{})
	sd.pvpTemp = pvpTemp
	sd.sceneIndex = sceneIndex
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
