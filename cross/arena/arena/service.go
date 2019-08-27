package arena

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/runner"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenascene "fgame/fgame/cross/arena/scene"
	areneascene "fgame/fgame/cross/arena/scene"
	"fgame/fgame/cross/player/player"
	arenatemplate "fgame/fgame/game/arena/template"
	arenatypes "fgame/fgame/game/arena/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
	skillcommon "fgame/fgame/game/skill/common"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"math"
	"math/rand"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type ArenaService interface {
	runner.Task
	Start()
	Match()
	Stop()
	//四神结束
	FourGodSceneEnd(fourGodType arenatypes.FourGodType)
	//开始匹配
	ArenaMatch(playerList []*MatchTeamMember) bool
	//停止匹配
	ArenaStopMatch(playerId int64) bool
	//匹配下一场
	ArenaNext(playerId int64) bool
	//竞技场结束
	ArenaEnd(s arenascene.ArenaSceneData, winnerId int64) *arenascene.ArenaTeam
	//进入四圣兽
	PlayerEnterFourGod(playerId int64, fourGodType arenatypes.FourGodType)
	//取消排队
	PlayerCancelFourGodQueue(playerId int64)
	//四神排队中
	TeamFourGodQueue(teamId int64)
	//四神游戏中
	TeamFourGod(teamId int64)
	//四神取消排队
	TeamCancelFourGodQueue(teamId int64)
	//获取队伍
	GetArenaTeamByPlayerId(playerId int64) *arenascene.ArenaTeam
	//获取竞技场景依赖玩家id
	GetArenaSceneByPlayerId(playerId int64) scene.Scene
	//获取四圣兽列表
	GetFourGodSceneList() []scene.Scene
	//获取四圣兽场景
	GetFourGodScene(arenatypes.FourGodType) scene.Scene
	//成员上线
	ArenaMemberOnline(pl scene.Player)
	//成员下线
	ArenaMemberOffline(pl scene.Player)
	//成员退出
	ArenaMemeberExit(pl scene.Player)
	//成员放弃
	ArenaMemeberGiveUp(pl scene.Player)
	//更新复活次数
	ArenaMemberUpdateReliveTime(pl scene.Player)
}

type arenaService struct {
	rwm sync.RWMutex
	//匹配玩家列表
	matchPlayerListMap map[int64][]*MatchTeamMember
	//人数分配
	matchPlayerNumMap map[int32]map[int64]int64

	//玩家所在的战队
	playerMap map[int64]*arenascene.ArenaTeam
	//所有战队
	allTeamMap map[int64]*arenascene.ArenaTeam
	//状态战队
	teamMapOfState map[arenascene.ArenaTeamState]map[int64]*arenascene.ArenaTeam
	//队伍和场景
	arenaSceneMap map[int64]scene.Scene
	//四圣兽场景
	fourGodSceneList []scene.Scene
	runner           heartbeat.HeartbeatTaskRunner
}

func (s *arenaService) init() (err error) {
	s.allTeamMap = make(map[int64]*arenascene.ArenaTeam)
	s.playerMap = make(map[int64]*arenascene.ArenaTeam)
	s.matchPlayerListMap = make(map[int64][]*MatchTeamMember)
	s.matchPlayerNumMap = make(map[int32]map[int64]int64)
	s.teamMapOfState = make(map[arenascene.ArenaTeamState]map[int64]*arenascene.ArenaTeam)
	s.arenaSceneMap = make(map[int64]scene.Scene)
	s.fourGodSceneList = make([]scene.Scene, 0, 4)

	s.runner = heartbeat.NewHeartbeatTaskRunner()
	s.runner.AddTask(CreateMatchTask(s))

	return nil
}

func (s *arenaService) Start() {
	s.checkArenaActivity()
}

func (s *arenaService) Match() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.match()
}

func (s *arenaService) Heartbeat() {
	s.check()
	s.runner.Heartbeat()
}

//TODO 修改加入定时器内
func (s *arenaService) check() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//TODO 定时将比赛完的加入匹配
	err := s.checkArenaActivity()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Warn("arena:竞技场检查活动,错误")
	}
	//检查结束的队伍
	s.checkTeamEnd()

}

func (s *arenaService) checkArenaActivity() (err error) {
	//活动场景还没结束
	if len(s.fourGodSceneList) > 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	arenaConstantTemp := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate()
	if !arenaConstantTemp.IsOnArenaTime(now) {
		return
	}
	// endTime := arenaConstantTemp.GetEndTime(now)
	// s.startArenaActivity(endTime)
	return nil
}

//开始竞技场活动
func (s *arenaService) startArenaActivity(endTime int64) {
	if len(s.fourGodSceneList) > 0 {
		return
	}

	for i := arenatypes.FourGodTypeQingLong; i <= arenatypes.FourGodTypeXuanWu; i++ {
		//获取地图
		mapTemplate := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().GetFourGodMapTemplate(i)
		fourGodScene := arenascene.CreateFourGodScene(int32(mapTemplate.TemplateId()), i, endTime)
		s.fourGodSceneList = append(s.fourGodSceneList, fourGodScene)
	}
}

//玩家开始匹配
func (s *arenaService) ArenaMatch(playerList []*MatchTeamMember) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if len(playerList) == 0 {
		log.WithFields(
			log.Fields{}).Warn("arena:成员是空的")
		return false
	}

	//判断成员是否在比赛中
	for _, mem := range playerList {
		t := s.getArenaTeamByPlayerId(mem.GetPlayerId())
		if t != nil {
			log.WithFields(
				log.Fields{
					"playerId": mem.GetPlayerId(),
				}).Warn("arena:成员已经在比赛中")
			return false
		}
	}

	//判断用户是否正在匹配中
	captainId := playerList[0].GetPlayerId()
	_, ok := s.matchPlayerListMap[captainId]
	if ok {
		log.WithFields(
			log.Fields{
				"playerId": captainId,
			}).Warn("arena:成员已经在匹配中")
		return false
	}

	s.matchPlayerListMap[captainId] = playerList
	numPlayers := int32(len(playerList))
	matchPlayerMap, ok := s.matchPlayerNumMap[numPlayers]
	if !ok {
		matchPlayerMap = make(map[int64]int64)
		s.matchPlayerNumMap[numPlayers] = matchPlayerMap
	}
	now := global.GetGame().GetTimeService().Now()
	matchPlayerMap[captainId] = now
	return true
}

//移除匹配队伍
func (s *arenaService) removeMatchPlayerList(captainId int64) {
	playerList, ok := s.matchPlayerListMap[captainId]
	if !ok {
		return
	}
	numPlayer := int32(len(playerList))
	delete(s.matchPlayerListMap, captainId)
	matchPlayerMap, ok := s.matchPlayerNumMap[numPlayer]
	if !ok {
		return
	}
	delete(matchPlayerMap, captainId)
}

//停止匹配
func (s *arenaService) ArenaStopMatch(playerId int64) bool {
	//获取队伍id
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.removeMatchPlayerList(playerId)
	return true
}

//玩家开始匹配
func (s *arenaService) ArenaNext(playerId int64) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	t := s.getArenaTeamByPlayerId(playerId)
	if t == nil {
		log.WithFields(
			log.Fields{}).Warn("arena:队伍不存在")
		return false
	}
	captain := t.GetTeam().GetCaptain()
	if captain == nil {
		panic(fmt.Errorf("arena:队长应该不为空"))
	}

	if captain.GetPlayerId() != playerId {
		log.WithFields(
			log.Fields{}).Warn("arena:不是队长")
		return false
	}

	//进入下一个
	flag := s.arenaNext(t)
	if !flag {
		return false
	}
	return true
}

//进入下一轮
func (s *arenaService) arenaNext(t *arenascene.ArenaTeam) bool {
	s.removeTeamByIdAndState(arenascene.ArenaTeamStateGameEnd, t.GetTeam().GetTeamId())
	//TODO 判断是否都是机器人
	if t.GetTeam().IfAllRobot() {
		gameevent.Emit(arenaeventtypes.EventTypeArenaRobotTeamEnd, t, nil)
		return false
	}
	// totalRound := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().TherionWinCount
	// if t.GetCurrent() >= totalRound {
	// 	flag := t.EnterFourGod()
	// 	if !flag {
	// 		log.WithFields(
	// 			log.Fields{}).Warn("arena:状态不对")
	// 		return false
	// 	}
	// 	s.cacheTeamAndState(t)
	// 	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGod, t, s.fourGodSceneList)
	// } else {
	flag := t.Match()
	if !flag {
		log.WithFields(
			log.Fields{}).Warn("arena:状态不对")
		return false
	}
	s.cacheTeamAndState(t)
	//发送事件
	gameevent.Emit(arenaeventtypes.EventTypeArenaNextMatch, t, nil)
	// }
	return true
}

var (
	firstNum           = int32(3)
	gameFirstMatchTime = 3 * int64(common.SECOND)
	gameEndWaitTime    = 10 * int64(common.SECOND)
	gameMatchTime      = 30 * int64(common.SECOND)
)

func (s *arenaService) checkTeamEnd() {
	now := global.GetGame().GetTimeService().Now()
	for _, t := range s.getTeamsByState(arenascene.ArenaTeamStateGameEnd) {
		elapse := now - t.GetLastTime()
		if elapse >= gameEndWaitTime {
			s.arenaNext(t)
		}
	}
}

func convertMemberFromRobotPlayer(pl scene.RobotPlayer) *MatchTeamMember {
	serverId := pl.GetServerId()
	playerId := pl.GetId()
	force := pl.GetForce()
	name := pl.GetOriginName()
	level := pl.GetLevel()
	role := pl.GetRole()
	sex := pl.GetSex()
	fashionId := pl.GetFashionId()
	skillList := make([]skillcommon.SkillObject, 0, len(pl.GetAllSkills()))
	for _, skill := range pl.GetAllSkills() {
		skillList = append(skillList, skill)
	}
	winCount := pl.GetArenaWinTime()
	memObj := CreateMatchTeamMemberObject(
		serverId,
		playerId,
		force,
		name,
		level,
		role,
		sex,
		fashionId,
		pl.GetAllSystemBattleProperties(),
		skillList,
		true,
		winCount,
	)
	return memObj
}

var (
	debug = false
)

//匹配
func (s *arenaService) match() {
	now := global.GetGame().GetTimeService().Now()
	arenaConstantTemp := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate()
	if !arenaConstantTemp.IsOnArenaTime(now) {
		log.Info("arena:活动未开始")
		//清空所有用户
		return
	}
	endTime := arenaConstantTemp.GetEndTime(now)

	//zrc: 临时解决
	s.checkRobotTeams()
	if debug {
		teamMap := s.getTeamsByState(arenascene.ArenaTeamStateMatch)
		if len(teamMap) > 1 {
			numOfTeam := len(teamMap)
			//随机匹配
			teamList := make([]*arenascene.ArenaTeam, 0, numOfTeam)
			for _, t := range teamMap {
				teamList = append(teamList, t)
			}
			//洗牌
			randIntList := rand.Perm(len(teamList))
			for i := 0; i < numOfTeam/2; i++ {
				team1 := teamList[randIntList[i*2]]
				team2 := teamList[randIntList[i*2+1]]
				s.matchTeam(team1, team2, endTime)
			}
		}
		matchPlayerListNum := len(s.matchPlayerListMap)
		maxMatchPlayerListNum := (matchPlayerListNum >> 1) << 1
		if maxMatchPlayerListNum > 0 {
			teamList := make([]*arenascene.ArenaTeam, 0, maxMatchPlayerListNum)
			for _, playerList := range s.matchPlayerListMap {
				t := s.createArenaTeam(playerList)
				teamList = append(teamList, t)
				if len(teamList) >= maxMatchPlayerListNum {
					break
				}
			}
			numOfTeam := len(teamList)
			randIntList := rand.Perm(len(teamList))
			for i := 0; i < numOfTeam/2; i++ {
				team1 := teamList[randIntList[i*2]]
				team2 := teamList[randIntList[i*2+1]]
				s.matchTeam(team1, team2, endTime)
			}
		}
		if len(teamMap) == 1 && len(s.matchPlayerListMap) == 1 {
			var team1 *arenascene.ArenaTeam
			var team2 *arenascene.ArenaTeam
			for _, tempTeam := range teamMap {
				team1 = tempTeam
				break
			}
			for _, playerList := range s.matchPlayerListMap {
				team2 = s.createArenaTeam(playerList)
				break
			}
			s.matchTeam(team1, team2, endTime)
		} else if len(teamMap) == 1 {
			var team1 *arenascene.ArenaTeam

			for _, tempTeam := range teamMap {
				team1 = tempTeam
				break
			}
			team2 := s.createRobotTeam(team1)
			if team2 == nil {
				log.Warn("arena:创建机器人战队,失败")
				return
			}
			s.matchTeam(team1, team2, endTime)
		}
		//匹配机器人
		threeMatchPlayerMap := s.matchPlayerNumMap[3]
		twoMatchPlayerMap := s.matchPlayerNumMap[2]
		oneMatchPlayerMap := s.matchPlayerNumMap[1]
		threeMatchPlayerNum := len(threeMatchPlayerMap)
		twoMatchPlayerNum := len(twoMatchPlayerMap)
		oneMatchPlayerNum := len(oneMatchPlayerMap)
		//3人
		if threeMatchPlayerNum > 0 {
			for tempthreeMatchPlayerId, lastTime := range threeMatchPlayerMap {
				elapse := now - lastTime
				if elapse >= gameMatchTime {
					threePlayerList := s.matchPlayerListMap[tempthreeMatchPlayerId]
					team1 := s.createArenaTeam(threePlayerList)
					team2 := s.createRobotTeamByMatchMemberList(threePlayerList)
					if team2 == nil {
						continue
					}
					s.matchTeam(team1, team2, endTime)
				}
			}
		}
		//2人
		if twoMatchPlayerNum > 0 {
			for tempTwoMatchPlayerId, lastTime := range twoMatchPlayerMap {
				elapse := now - lastTime
				if elapse >= gameMatchTime {
					twoPlayerList := s.matchPlayerListMap[tempTwoMatchPlayerId]
					team1 := s.createArenaTeam(twoPlayerList)
					team2 := s.createRobotTeamByMatchMemberList(twoPlayerList)
					if team2 == nil {
						continue
					}
					s.matchTeam(team1, team2, endTime)
				}
			}

		}
		if oneMatchPlayerNum > 0 {
			for tempOneMatchPlayerId, lastTime := range oneMatchPlayerMap {
				elapse := now - lastTime
				if elapse >= gameMatchTime {
					onePlayerList := s.matchPlayerListMap[tempOneMatchPlayerId]
					team1 := s.createArenaTeam(onePlayerList)
					team2 := s.createRobotTeamByMatchMemberList(onePlayerList)
					if team2 == nil {
						continue
					}
					s.matchTeam(team1, team2, endTime)
				}
			}
		}
		return
	}

	teamMap := s.getTeamsByState(arenascene.ArenaTeamStateMatch)
	if len(teamMap) > 1 {
		numOfTeam := len(teamMap)
		//随机匹配
		teamList := make([]*arenascene.ArenaTeam, 0, numOfTeam)
		for _, t := range teamMap {
			teamList = append(teamList, t)
		}
		//洗牌
		randIntList := rand.Perm(len(teamList))
		for i := 0; i < numOfTeam/2; i++ {
			team1 := teamList[randIntList[i*2]]
			team2 := teamList[randIntList[i*2+1]]
			s.matchTeam(team1, team2, endTime)
		}
	}

	//获取三人小队
	threeMatchPlayerMap := s.matchPlayerNumMap[3]
	threeMatchPlayerNum := len(threeMatchPlayerMap)
	//获取最大偶数
	maxTeam := (threeMatchPlayerNum >> 1) << 1
	if maxTeam != 0 {
		teamList := make([]*arenascene.ArenaTeam, 0, maxTeam)
		teamNum := 0
		for tempThreeMatchPlayerId, _ := range threeMatchPlayerMap {
			threePlayerList := s.matchPlayerListMap[tempThreeMatchPlayerId]
			t := s.createArenaTeam(threePlayerList)
			teamList = append(teamList, t)
			teamNum += 1
			threeMatchPlayerNum -= 1
			if teamNum >= maxTeam {
				break
			}
		}
		//洗牌
		randIntList := rand.Perm(len(teamList))
		for i := 0; i < teamNum/2; i++ {
			team1 := teamList[randIntList[i*2]]
			team2 := teamList[randIntList[i*2+1]]
			s.matchTeam(team1, team2, endTime)
		}
	}

	if len(teamMap) == 1 {
		var team1 *arenascene.ArenaTeam
		for _, tempTeam := range teamMap {
			team1 = tempTeam
			break
		}
		//获取三人队伍
		threeMatchPlayerMap := s.matchPlayerNumMap[3]
		if len(threeMatchPlayerMap) > 0 {
			var matchPlayerId int64
			for threeMatchPlayerId, _ := range threeMatchPlayerMap {
				matchPlayerId = threeMatchPlayerId
				break
			}
			playerList := s.matchPlayerListMap[matchPlayerId]
			if playerList == nil {
				panic(fmt.Errorf("arena:不可能获取不到队伍"))
			}
			team2 := s.createArenaTeam(playerList)
			s.matchTeam(team1, team2, endTime)
			goto AfterTeam
		}
		//获取双人小队
		twoMatchPlayerMap := s.matchPlayerNumMap[2]
		oneMatchPlayerMap := s.matchPlayerNumMap[1]
		if len(twoMatchPlayerMap) > 0 {
			var twoMatchPlayerId int64
			for tempTwoMatchPlayerId, _ := range twoMatchPlayerMap {
				twoMatchPlayerId = tempTwoMatchPlayerId
				break
			}

			twoPlayerList := s.matchPlayerListMap[twoMatchPlayerId]
			//匹配双人
			if len(oneMatchPlayerMap) > 0 {
				var oneMatchPlayerId int64
				for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
					oneMatchPlayerId = tempOneMatchPlayerId
					break
				}
				onePlayerList := s.matchPlayerListMap[oneMatchPlayerId]
				team2 := s.combinePlayerList(twoPlayerList, onePlayerList)
				s.matchTeam(team1, team2, endTime)
				goto AfterTeam
			}
		} else if len(oneMatchPlayerMap) >= 3 {
			memListOfList := make([][]*MatchTeamMember, 0, 3)
			for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
				onePlayerList := s.matchPlayerListMap[tempOneMatchPlayerId]
				memListOfList = append(memListOfList, onePlayerList)
				if len(memListOfList) >= 3 {
					break
				}
			}
			team2 := s.combinePlayerList(memListOfList...)
			s.matchTeam(team1, team2, endTime)
			goto AfterTeam
		}
		now := global.GetGame().GetTimeService().Now()
		elapse := now - team1.GetLastTime()
		maxWaitTime := gameFirstMatchTime
		if team1.GetMemberMaxWinCount() >= firstNum {
			maxWaitTime = gameMatchTime
		}
		//超过30秒
		if elapse >= maxWaitTime {
			team2 := s.createRobotTeam(team1)
			if team2 != nil {
				s.matchTeam(team1, team2, endTime)
			}

		}

	}
AfterTeam:

	twoMatchPlayerMap := s.matchPlayerNumMap[2]
	oneMatchPlayerMap := s.matchPlayerNumMap[1]

	twoMatchPlayerNum := len(twoMatchPlayerMap)
	oneMatchPlayerNum := len(oneMatchPlayerMap)
	minNum := oneMatchPlayerNum
	if twoMatchPlayerNum < oneMatchPlayerNum {
		minNum = twoMatchPlayerNum
	}
	//获取最大偶数
	maxTeam = (minNum >> 1) << 1
	if maxTeam != 0 {
		teamList := make([]*arenascene.ArenaTeam, 0, maxTeam)
		teamNum := 0
		for tempTwoMatchPlayerId, _ := range twoMatchPlayerMap {
			twoPlayerList := s.matchPlayerListMap[tempTwoMatchPlayerId]
			for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
				onePlayerList := s.matchPlayerListMap[tempOneMatchPlayerId]
				t := s.combinePlayerList(twoPlayerList, onePlayerList)
				teamList = append(teamList, t)
				teamNum += 1
				oneMatchPlayerNum -= 1
				twoMatchPlayerNum -= 1
				break
			}
			if teamNum >= maxTeam {
				break
			}
		}
		//洗牌
		randIntList := rand.Perm(len(teamList))
		for i := 0; i < teamNum/2; i++ {
			team1 := teamList[randIntList[i*2]]
			team2 := teamList[randIntList[i*2+1]]
			s.matchTeam(team1, team2, endTime)
		}
	}

	//组合1个的
	remainOneTeam := oneMatchPlayerNum / 3
	if remainOneTeam > 0 {
		maxTeam := (remainOneTeam >> 1) << 1
		if maxTeam == 0 {
			goto AfterRemain
		}
		teamList := make([]*arenascene.ArenaTeam, 0, maxTeam)

		oneMatchPlayerList := make([]int64, 0, oneMatchPlayerNum)
		for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
			oneMatchPlayerList = append(oneMatchPlayerList, tempOneMatchPlayerId)
		}
		for i := 0; i < maxTeam; i++ {
			firstPlayerList := s.matchPlayerListMap[oneMatchPlayerList[i*3]]
			secondPlayerList := s.matchPlayerListMap[oneMatchPlayerList[i*3+1]]
			thirdPlayerList := s.matchPlayerListMap[oneMatchPlayerList[i*3+2]]
			t := s.combinePlayerList(firstPlayerList, secondPlayerList, thirdPlayerList)
			teamList = append(teamList, t)
		}
		teamNum := len(teamList)
		randIntList := rand.Perm(teamNum)
		for i := 0; i < teamNum/2; i++ {
			team1 := teamList[randIntList[i*2]]
			team2 := teamList[randIntList[i*2+1]]
			s.matchTeam(team1, team2, endTime)
		}
	}
AfterRemain:
	//匹配剩余的
	threeMatchPlayerMap = s.matchPlayerNumMap[3]
	twoMatchPlayerMap = s.matchPlayerNumMap[2]
	oneMatchPlayerMap = s.matchPlayerNumMap[1]
	threeMatchPlayerNum = len(threeMatchPlayerMap)
	twoMatchPlayerNum = len(twoMatchPlayerMap)
	oneMatchPlayerNum = len(oneMatchPlayerMap)
	if threeMatchPlayerNum == 1 {
		if twoMatchPlayerNum >= 1 && oneMatchPlayerNum >= 1 {
			var threePlayerList []*MatchTeamMember
			for tempthreeMatchPlayerId, _ := range threeMatchPlayerMap {
				threePlayerList = s.matchPlayerListMap[tempthreeMatchPlayerId]
				break
			}
			team1 := s.createArenaTeam(threePlayerList)

			var twoPlayerList []*MatchTeamMember
			for tempTwoMatchPlayerId, _ := range twoMatchPlayerMap {
				twoPlayerList = s.matchPlayerListMap[tempTwoMatchPlayerId]
				break
			}
			var onePlayerList []*MatchTeamMember
			for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
				onePlayerList = s.matchPlayerListMap[tempOneMatchPlayerId]
				break
			}
			team2 := s.combinePlayerList(twoPlayerList, onePlayerList)
			s.matchTeam(team1, team2, endTime)
		} else if oneMatchPlayerNum >= 3 {
			var threePlayerList []*MatchTeamMember
			for tempthreeMatchPlayerId, _ := range threeMatchPlayerMap {
				threePlayerList = s.matchPlayerListMap[tempthreeMatchPlayerId]
				break
			}
			team1 := s.createArenaTeam(threePlayerList)
			var onePlayerListList [][]*MatchTeamMember

			for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
				tempOnePlayerList := s.matchPlayerListMap[tempOneMatchPlayerId]
				onePlayerListList = append(onePlayerListList, tempOnePlayerList)
				if len(onePlayerListList) >= 3 {
					break
				}
			}
			team2 := s.combinePlayerList(onePlayerListList...)
			s.matchTeam(team1, team2, endTime)
		}
	} else {
		//组合2和1 3个1
		if twoMatchPlayerNum >= 1 && oneMatchPlayerNum >= 4 {
			memListOfList := make([][]*MatchTeamMember, 0, 3)
			for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
				onePlayerList := s.matchPlayerListMap[tempOneMatchPlayerId]
				memListOfList = append(memListOfList, onePlayerList)
				if len(memListOfList) >= 3 {
					break
				}
			}
			team1 := s.combinePlayerList(memListOfList...)
			var twoMatchPlayerId int64
			for tempTwoMatchPlayerId, _ := range twoMatchPlayerMap {
				twoMatchPlayerId = tempTwoMatchPlayerId
				break
			}

			twoPlayerList := s.matchPlayerListMap[twoMatchPlayerId]
			var oneMatchPlayerId int64
			for tempOneMatchPlayerId, _ := range oneMatchPlayerMap {
				oneMatchPlayerId = tempOneMatchPlayerId
				break
			}
			onePlayerList := s.matchPlayerListMap[oneMatchPlayerId]
			team2 := s.combinePlayerList(twoPlayerList, onePlayerList)
			s.matchTeam(team1, team2, endTime)
		}
	}

	//匹配机器人
	threeMatchPlayerMap = s.matchPlayerNumMap[3]
	twoMatchPlayerMap = s.matchPlayerNumMap[2]
	oneMatchPlayerMap = s.matchPlayerNumMap[1]
	threeMatchPlayerNum = len(threeMatchPlayerMap)
	twoMatchPlayerNum = len(twoMatchPlayerMap)
	oneMatchPlayerNum = len(oneMatchPlayerMap)
	//3人
	if threeMatchPlayerNum > 0 {
		for tempthreeMatchPlayerId, lastTime := range threeMatchPlayerMap {
			elapse := now - lastTime
			if elapse >= gameFirstMatchTime {
				threePlayerList := s.matchPlayerListMap[tempthreeMatchPlayerId]
				team1 := s.createArenaTeam(threePlayerList)
				team2 := s.createRobotTeamByMatchMemberList(threePlayerList)
				s.matchTeam(team1, team2, endTime)
			}
		}
	}
	//2人
	if twoMatchPlayerNum > 0 {
		for tempTwoMatchPlayerId, lastTime := range twoMatchPlayerMap {
			elapse := now - lastTime
			if elapse >= gameFirstMatchTime {
				twoPlayerList := s.matchPlayerListMap[tempTwoMatchPlayerId]
				team1 := s.completmentRobotTeamByMatchMemberList(twoPlayerList)
				team2 := s.createRobotTeamByMatchMemberList(twoPlayerList)
				s.matchTeam(team1, team2, endTime)
			}
		}

	}
	if oneMatchPlayerNum > 0 {
		for tempOneMatchPlayerId, lastTime := range oneMatchPlayerMap {
			elapse := now - lastTime
			if elapse >= gameFirstMatchTime {
				onePlayerList := s.matchPlayerListMap[tempOneMatchPlayerId]
				team1 := s.completmentRobotTeamByMatchMemberList(onePlayerList)
				team2 := s.createRobotTeamByMatchMemberList(onePlayerList)
				s.matchTeam(team1, team2, endTime)
			}
		}
	}
}

const (
	maxServer = 10
)

//补充机器人
func (s *arenaService) completmentRobotTeamByMatchMemberList(memList []*MatchTeamMember) *arenascene.ArenaTeam {
	robotMemList := make([]*MatchTeamMember, 0, 3)
	allProperties := make(map[int32]int64)
	num := 0

	for _, mem := range memList {
		num += 1
		for typ, val := range mem.GetBattleProperties() {
			allProperties[typ] += val
		}
	}
	if num <= 0 {
		return nil
	}
	//复活次数
	threeRobotTemplate := arenatemplate.GetArenaTemplateService().GetThreeRobotTemplate(1)
	reliveTime := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RandomRevive()
	if threeRobotTemplate != nil {
		reliveTime = threeRobotTemplate.RandomReborn()
	}

	needComplement := 3 - len(memList)

	//随机属性
	avgProperties := make(map[int32]int64)
	for typ, val := range allProperties {
		avgProperties[typ] = int64(math.Ceil(float64(val) / float64(num)))
	}
	randonServer := rand.Intn(maxServer) + 1
	serverId := int32(randonServer)
	for i := 0; i < needComplement; i++ {
		robotProperties := make(map[propertytypes.BattlePropertyType]int64)
		percent := float64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RandomPropertyPercent()) / float64(common.MAX_RATE)
		for typ, val := range avgProperties {
			if typ != int32(propertytypes.BattlePropertyTypeMaxHP) && typ != int32(propertytypes.BattlePropertyTypeAttack) && typ != int32(propertytypes.BattlePropertyTypeDefend) {
				robotProperties[propertytypes.BattlePropertyType(typ)] = val
			} else {
				robotProperties[propertytypes.BattlePropertyType(typ)] = int64(math.Ceil(float64(val) * percent))
			}
		}

		force := propertylogic.CulculateAllForce(robotProperties)

		robotPlayer := robot.GetRobotService().CreateArenaRobot(serverId, robotProperties, reliveTime, force)
		robotMemList = append(robotMemList, convertMemberFromRobotPlayer(robotPlayer))
	}
	team1 := s.combinePlayerList(memList, robotMemList)
	return team1
}

//创建机器人队伍
func (s *arenaService) createRobotTeamByMatchMemberList(memList []*MatchTeamMember) *arenascene.ArenaTeam {
	robotMemList := make([]*MatchTeamMember, 0, 3)

	allProperties := make(map[int32]int64)
	num := 0

	for _, mem := range memList {
		num += 1
		for typ, val := range mem.GetBattleProperties() {
			allProperties[typ] += val
		}
	}
	if num <= 0 {
		return nil
	}

	//复活次数
	threeRobotTemplate := arenatemplate.GetArenaTemplateService().GetThreeRobotTemplate(1)
	reliveTime := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RandomRevive()
	if threeRobotTemplate != nil {
		reliveTime = threeRobotTemplate.RandomReborn()
	}

	//随机属性
	avgProperties := make(map[int32]int64)
	for typ, val := range allProperties {
		avgProperties[typ] = int64(math.Ceil(float64(val) / float64(num)))
	}
	randonServer := rand.Intn(maxServer) + 1
	serverId := int32(randonServer)
	for i := 0; i < 3; i++ {
		robotProperties := make(map[propertytypes.BattlePropertyType]int64)
		percent := float64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RandomPropertyPercent()) / float64(common.MAX_RATE)
		for typ, val := range avgProperties {
			if typ != int32(propertytypes.BattlePropertyTypeMaxHP) && typ != int32(propertytypes.BattlePropertyTypeAttack) && typ != int32(propertytypes.BattlePropertyTypeDefend) {
				robotProperties[propertytypes.BattlePropertyType(typ)] = val
			} else {
				robotProperties[propertytypes.BattlePropertyType(typ)] = int64(math.Ceil(float64(val) * percent))
			}
		}

		force := propertylogic.CulculateAllForce(robotProperties)

		robotPlayer := robot.GetRobotService().CreateArenaRobot(serverId, robotProperties, reliveTime, force)
		robotMemList = append(robotMemList, convertMemberFromRobotPlayer(robotPlayer))
	}
	team1 := s.createArenaTeam(robotMemList)
	return team1
}

//创建机器人队伍
func (s *arenaService) createRobotTeam(team *arenascene.ArenaTeam) *arenascene.ArenaTeam {
	robotMemList := make([]*MatchTeamMember, 0, 3)

	allProperties := make(map[int32]int64)
	num := 0

	for _, mem := range team.GetTeam().GetMemberList() {
		copyPlayer := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if copyPlayer == nil {
			continue
		}
		if copyPlayer.GetScene() == nil {
			continue
		}
		num += 1
		for typ, val := range copyPlayer.GetAllSystemBattleProperties() {
			allProperties[typ] += val
		}
	}
	if num <= 0 {
		return nil
	}

	//随机属性
	avgProperties := make(map[int32]int64)
	for typ, val := range allProperties {
		avgProperties[typ] = int64(math.Ceil(float64(val) / float64(num)))
	}
	randonServer := rand.Intn(maxServer) + 1
	serverId := int32(randonServer)

	robotNum := int32(3)
	//随机数量
	threeRobotTemplate := arenatemplate.GetArenaTemplateService().GetThreeRobotTemplate(team.GetMemberMaxWinCount() + 1)
	reliveTime := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RandomRevive()
	if threeRobotTemplate != nil {
		//随机机器人数量
		randomRobotNum := threeRobotTemplate.RandomRobot()
		if randomRobotNum < robotNum {
			robotNum = randomRobotNum
		}
		reliveTime = threeRobotTemplate.RandomReborn()
	}
	for i := 0; i < int(robotNum); i++ {
		robotProperties := make(map[propertytypes.BattlePropertyType]int64)
		percent := float64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RandomPropertyPercent()) / float64(common.MAX_RATE)
		for typ, val := range avgProperties {
			if typ != int32(propertytypes.BattlePropertyTypeMaxHP) && typ != int32(propertytypes.BattlePropertyTypeAttack) && typ != int32(propertytypes.BattlePropertyTypeDefend) {
				robotProperties[propertytypes.BattlePropertyType(typ)] = val
			} else {
				robotProperties[propertytypes.BattlePropertyType(typ)] = int64(math.Ceil(float64(val) * percent))
			}
		}

		force := propertylogic.CulculateAllForce(robotProperties)

		robotPlayer := robot.GetRobotService().CreateArenaRobot(serverId, robotProperties, reliveTime, force)
		robotMemList = append(robotMemList, convertMemberFromRobotPlayer(robotPlayer))
	}
	team2 := s.createArenaTeam(robotMemList)
	return team2
}

//组成战队
func (s *arenaService) createArenaTeam(aMemList []*MatchTeamMember) *arenascene.ArenaTeam {
	teamId, _ := idutil.GetId()
	memList := convertToTeamMemberObjectList(aMemList)

	t := arenascene.CreateArenaTeamWithMembers(teamId, memList)
	for _, mem := range memList {
		s.playerMap[mem.GetPlayerId()] = t
	}
	//移除匹配队伍
	s.removeMatchPlayerList(memList[0].GetPlayerId())

	s.allTeamMap[teamId] = t
	eventData := CreateTeamCreateEventData([][]*arenascene.TeamMemberObject{memList})
	gameevent.Emit(arenaeventtypes.EventTypeArenaTeamCreate, t, eventData)
	return t
}

func (s *arenaService) combinePlayerList(aMemListOfList ...[]*MatchTeamMember) *arenascene.ArenaTeam {
	totalNum := 0
	memListOfList := make([][]*arenascene.TeamMemberObject, 0, len(aMemListOfList))

	for _, aMemList := range aMemListOfList {
		totalNum += len(aMemList)
		memListOfList = append(memListOfList, convertToTeamMemberObjectList(aMemList))
	}
	if totalNum != 3 {
		panic(fmt.Errorf("arena:组队成员不等于3"))
	}

	teamId, _ := idutil.GetId()
	t := arenascene.CreateArenaTeamWithMemberLists(teamId, memListOfList...)
	for _, memList := range memListOfList {
		for _, mem := range memList {
			s.playerMap[mem.GetPlayerId()] = t
		}
	}
	for _, memList := range memListOfList {
		s.removeMatchPlayerList(memList[0].GetPlayerId())
	}

	s.allTeamMap[teamId] = t
	eventData := CreateTeamCreateEventData(memListOfList)
	gameevent.Emit(arenaeventtypes.EventTypeArenaTeamCreate, t, eventData)
	return t
}

func (s *arenaService) matchTeam(team1 *arenascene.ArenaTeam, team2 *arenascene.ArenaTeam, endTime int64) {
	//创建比赛场景
	team1Id := team1.GetTeam().GetTeamId()
	team1OldState := team1.GetState()
	team2Id := team2.GetTeam().GetTeamId()
	team2OldState := team2.GetState()
	s.removeTeamByIdAndState(team1OldState, team1Id)
	s.removeTeamByIdAndState(team2OldState, team2Id)
	flag := team1.Game()
	if !flag {
		panic("arena:队伍加入比赛应该ok")
	}
	flag = team2.Game()
	if !flag {
		panic("arena:队伍加入比赛应该ok")
	}
	s.cacheTeamAndState(team1)
	s.cacheTeamAndState(team2)
	arenaMapTemplate := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().GetArenaMapTemplate()
	//创建场景
	arenaScene := arenascene.CreateArenaScene(arenaMapTemplate, team1, team2, endTime)

	//移除队伍场景
	delete(s.arenaSceneMap, team1Id)
	delete(s.arenaSceneMap, team2Id)

	//缓存队伍场景
	s.arenaSceneMap[team1Id] = arenaScene
	s.arenaSceneMap[team2Id] = arenaScene

	eventData := CreateMatchEventData(team1, team2)
	//发送事件
	gameevent.Emit(arenaeventtypes.EventTypeArenaMatched, arenaScene, eventData)
}

//获取队伍
func (s *arenaService) getTeamsByState(state arenascene.ArenaTeamState) map[int64]*arenascene.ArenaTeam {
	teamMap, ok := s.teamMapOfState[state]
	if !ok {
		return nil
	}
	return teamMap
}

func (s *arenaService) getArenaTeam(teamId int64) *arenascene.ArenaTeam {
	t, ok := s.allTeamMap[teamId]
	if !ok {
		return nil
	}
	return t
}

//获取竞技场数据
func (s *arenaService) GetArenaTeamByPlayerId(playerId int64) *arenascene.ArenaTeam {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.getArenaTeamByPlayerId(playerId)
}

func (s *arenaService) getArenaTeamByPlayerId(playerId int64) *arenascene.ArenaTeam {
	t, ok := s.playerMap[playerId]
	if !ok {
		return nil
	}
	return t
}

//获取当前场景
func (s *arenaService) GetArenaSceneByPlayerId(playerId int64) scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	t := s.getArenaTeamByPlayerId(playerId)
	if t == nil {
		return nil
	}
	switch t.GetState() {
	case arenascene.ArenaTeamStateInit:
		return nil
	case arenascene.ArenaTeamStateMatch,
		arenascene.ArenaTeamStateGame,
		arenascene.ArenaTeamStateGameEnd,
		arenascene.ArenaTeamStateFourGodInit,
		arenascene.ArenaTeamStateFourGodEnter,
		arenascene.ArenaTeamStateFourGodQueue:
		return s.getArenaScene(t.GetTeam().GetTeamId())
	case arenascene.ArenaTeamStateFourGod:
		ts, _ := s.getFourGodScene(t.GetFourGodType())
		return ts
	}
	return nil
}

func (s *arenaService) getArenaScene(teamId int64) scene.Scene {
	tempS, ok := s.arenaSceneMap[teamId]
	if !ok {
		return nil
	}
	return tempS
}

//进入四圣兽场景
func (s *arenaService) PlayerEnterFourGod(playerId int64, fourGodType arenatypes.FourGodType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	teamObject := s.getArenaTeamByPlayerId(playerId)
	if teamObject == nil {
		//TODO 提示队伍不在竞技场
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"state":    teamObject.GetState().String(),
			}).Warn("arena:队伍不存在")
		return
	}

	fourGodScene, _ := s.getFourGodScene(fourGodType)
	if fourGodScene == nil {
		//TODO 活动未开始
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("arena:场景不存在")
		return
	}
	oldState := teamObject.GetState()
	//判断队伍状态
	if !teamObject.EnterFourGodGame(fourGodType) {
		//TODO 提示
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"state":    teamObject.GetState().String(),
			}).Warn("arena:进入四神错误")
		return
	}

	//移除缓存
	s.removeTeamByIdAndState(oldState, teamObject.GetTeam().GetTeamId())
	s.cacheTeamAndState(teamObject)
	teamEnterFourGodScene(fourGodScene, teamObject)
}

//队伍四圣兽场景
func teamEnterFourGodScene(s scene.Scene, teamObject *arenascene.ArenaTeam) {
	ctx := scene.WithScene(context.Background(), s)
	s.Post(message.NewScheduleMessage(onTeamEnterFourGodScene, ctx, teamObject, nil))
}

//进入四圣兽
func onTeamEnterFourGodScene(ctx context.Context, result interface{}, err error) error {
	teamObject := result.(*arenascene.ArenaTeam)
	fourGodScene := scene.SceneInContext(ctx)
	sd := fourGodScene.SceneDelegate()
	fourGodSceneData := sd.(arenascene.FourGodSceneData)
	fourGodSceneData.TeamJoin(teamObject)
	return nil
}

//取消四圣兽排队
func (s *arenaService) PlayerCancelFourGodQueue(playerId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	teamObject := s.getArenaTeamByPlayerId(playerId)
	if teamObject == nil {
		//TODO 提示队伍不存在
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("arena:队伍不存在")
		return
	}
	//不在排队中
	if teamObject.GetState() != arenascene.ArenaTeamStateFourGodQueue {
		return
	}
	//判断是否在游戏中
	fourGodType := teamObject.GetFourGodType()

	fourGodScene, _ := s.getFourGodScene(fourGodType)
	if fourGodScene == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("arena:场景不存在")
		return
	}
	teamId := teamObject.GetTeam().GetTeamId()
	teamCancelFourGodSceneQueue(fourGodScene, teamId)
}

//四神排队中
func (s *arenaService) TeamFourGodQueue(teamId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	teamObject := s.getArenaTeam(teamId)
	if teamObject == nil {
		return
	}
	oldState := teamObject.GetState()
	flag := teamObject.FourGodGameQueue()
	if !flag {
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Warn("arena:进入排队失败")
		return
	}
	s.removeTeamByIdAndState(oldState, teamId)
	s.cacheTeamAndState(teamObject)
}

func (s *arenaService) TeamFourGod(teamId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	teamObject := s.getArenaTeam(teamId)
	if teamObject == nil {
		return
	}
	oldState := teamObject.GetState()
	flag := teamObject.FourGodGame()
	if !flag {
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Warn("arena:进入四神失败")
		return
	}
	s.removeTeamByIdAndState(oldState, teamId)
	s.cacheTeamAndState(teamObject)
}

func (s *arenaService) TeamCancelFourGodQueue(teamId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	teamObject := s.getArenaTeam(teamId)
	if teamObject == nil {
		return
	}
	oldState := teamObject.GetState()
	flag := teamObject.CancelFourGodGameQueue()
	if !flag {
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Warn("arena:取消四神排队失败")
		return
	}
	s.removeTeamByIdAndState(oldState, teamId)
	s.cacheTeamAndState(teamObject)
}

//队伍四圣兽取消排队
func teamCancelFourGodSceneQueue(s scene.Scene, teamId int64) {
	ctx := scene.WithScene(context.Background(), s)
	s.Post(message.NewScheduleMessage(onTeamCancelFourGodSceneQueue, ctx, teamId, nil))
}

//退出排队
func onTeamCancelFourGodSceneQueue(ctx context.Context, result interface{}, err error) error {
	teamId := result.(int64)
	fourGodScene := scene.SceneInContext(ctx)
	fourGodSceneData := fourGodScene.SceneDelegate().(arenascene.FourGodSceneData)
	fourGodSceneData.TeamLeaveQueue(teamId)
	return nil
}

func (s *arenaService) getFourGodScene(fourGodType arenatypes.FourGodType) (scene.Scene, int32) {
	for i, fourGodScene := range s.fourGodSceneList {
		sd := fourGodScene.SceneDelegate().(arenascene.FourGodSceneData)
		if sd.GetFourGodType() == fourGodType {
			return fourGodScene, int32(i)
		}
	}

	return nil, -1
}

//添加战队
func (s *arenaService) addTeam(t *arenascene.ArenaTeam) {
	s.allTeamMap[t.GetTeam().GetTeamId()] = t
	s.cacheTeamAndState(t)
}

func (s *arenaService) cacheTeamAndState(t *arenascene.ArenaTeam) {
	teamMap, ok := s.teamMapOfState[t.GetState()]
	if !ok {
		teamMap = make(map[int64]*arenascene.ArenaTeam)
		s.teamMapOfState[t.GetState()] = teamMap
	}
	teamMap[t.GetTeam().GetTeamId()] = t
}

//移除队伍
func (s *arenaService) removeTeamByIdAndState(state arenascene.ArenaTeamState, teamId int64) {
	teamMap, ok := s.teamMapOfState[state]
	if !ok {
		return
	}
	delete(teamMap, teamId)
}

//移除战队
func (s *arenaService) removeTeam(teamId int64) {
	te := s.getArenaTeam(teamId)
	if te != nil {
		for _, mem := range te.GetTeam().GetMemberList() {
			if mem.GetStatus() == areneascene.MemberStatusGoAway {
				continue
			}
			if mem.GetStatus() == areneascene.MemberStatusFailed {
				continue
			}
			s.removePlayer(mem.GetPlayerId())
		}
		s.removeTeamByIdAndState(te.GetState(), teamId)
		delete(s.allTeamMap, teamId)
	}
}

//移除玩家
func (s *arenaService) removePlayer(playerId int64) {
	delete(s.playerMap, playerId)
}

//四神结束了
func (s *arenaService) FourGodSceneEnd(fourGodType arenatypes.FourGodType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	_, pos := s.getFourGodScene(fourGodType)
	if pos == -1 {
		return
	}
	s.fourGodSceneList = append(s.fourGodSceneList[:pos], s.fourGodSceneList[pos+1:]...)
}

//竞技场结束
func (s *arenaService) ArenaEnd(sd arenascene.ArenaSceneData, winnerId int64) *arenascene.ArenaTeam {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//TODO 可能队伍已经解散

	team1Id := sd.GetTeam1().GetTeam().GetTeamId()
	team2Id := sd.GetTeam2().GetTeam().GetTeamId()
	team1 := s.getArenaTeam(team1Id)
	team2 := s.getArenaTeam(team2Id)

	if team1 != nil {
		s.removeTeamByIdAndState(arenascene.ArenaTeamStateGame, team1Id)
		flag := team1.GameEnd()
		if !flag {
			panic(fmt.Errorf("arena:进入结束应该成功"))
		}
	}

	if team2 != nil {
		s.removeTeamByIdAndState(arenascene.ArenaTeamStateGame, team2Id)
		flag := team2.GameEnd()
		if !flag {
			panic(fmt.Errorf("arena:进入结束应该成功"))
		}
	}

	//没有获胜的
	if winnerId == 0 {
		//删除队伍场景
		delete(s.arenaSceneMap, team1Id)
		delete(s.arenaSceneMap, team2Id)
		s.removeTeam(team1Id)
		s.removeTeam(team2Id)
		return nil
	}

	if winnerId == team1Id {
		delete(s.arenaSceneMap, team2Id)
		team1.Win()
		s.removeTeam(team2Id)
		s.cacheTeamAndState(team1)
		return team1
	} else {
		delete(s.arenaSceneMap, team1Id)
		team2.Win()
		s.removeTeam(team1Id)
		s.cacheTeamAndState(team2)
		return team2
	}
}

func (s *arenaService) GetFourGodSceneList() []scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.fourGodSceneList
}

func (s *arenaService) GetFourGodScene(fourGodType arenatypes.FourGodType) scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	ts, _ := s.getFourGodScene(fourGodType)
	return ts
}

//成员上线
func (s *arenaService) ArenaMemberOnline(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	t := s.getArenaTeamByPlayerId(playerId)
	if t == nil {
		return
	}
	mem, _ := t.GetTeam().GetMember(playerId)
	if mem == nil {
		return
	}
	if mem.GetStatus() != arenascene.MemberStatusOffline {
		return
	}
	mem.SetStatus(arenascene.MemberStatusOnline)
	mem.SetReliveTime(pl.GetArenaReliveTime())
	gameevent.Emit(arenaeventtypes.EventTypeArenaTeamMemberOnline, pl, t)
}

//成员更新复活次数
func (s *arenaService) ArenaMemberUpdateReliveTime(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	t := s.getArenaTeamByPlayerId(playerId)
	if t == nil {
		return
	}
	mem, _ := t.GetTeam().GetMember(playerId)
	if mem == nil {
		return
	}
	mem.SetReliveTime(pl.GetArenaReliveTime())

}

//成员下线
func (s *arenaService) ArenaMemberOffline(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	_, ok := s.matchPlayerListMap[playerId]
	if ok {
		//退出匹配
		s.removeMatchPlayerList(playerId)
		return
	}

	t := s.getArenaTeamByPlayerId(playerId)
	if t == nil {
		return
	}

	mem, _ := t.GetTeam().GetMember(playerId)
	if mem == nil {
		return
	}

	if mem.GetStatus() != arenascene.MemberStatusOnline {
		return
	}
	mem.SetStatus(arenascene.MemberStatusOffline)
	gameevent.Emit(arenaeventtypes.EventTypeArenaTeamMemberOffline, pl, t)
	s.checkArenaTeam(t)
}

//成员退出
func (s *arenaService) ArenaMemeberExit(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	t := s.getArenaTeamByPlayerId(playerId)
	if t == nil {
		return
	}
	mem, _ := t.GetTeam().GetMember(playerId)
	if mem == nil {
		return
	}
	if mem.GetStatus() == arenascene.MemberStatusFailed {
		return
	}

	//移除玩家
	s.removePlayer(mem.GetPlayerId())
	mem.SetStatus(arenascene.MemberStatusGoAway)
	gameevent.Emit(arenaeventtypes.EventTypeArenaTeamMemberExit, pl, t)
	s.checkArenaTeam(t)
}

//成员放弃
func (s *arenaService) ArenaMemeberGiveUp(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	t := s.getArenaTeamByPlayerId(playerId)
	if t == nil {
		return
	}
	mem, _ := t.GetTeam().GetMember(playerId)
	if mem == nil {
		return
	}

	if mem.GetStatus() == arenascene.MemberStatusGoAway {
		return
	}

	//移除玩家
	s.removePlayer(mem.GetPlayerId())
	mem.SetStatus(arenascene.MemberStatusFailed)
	gameevent.Emit(arenaeventtypes.EventTypeArenaTeamMemberGiveUp, pl, t)
	s.checkArenaTeam(t)
}

func (s *arenaService) checkRobotTeams() {
	teamMap := s.getTeamsByState(arenascene.ArenaTeamStateMatch)
	for _, t := range teamMap {
		if t.GetTeam().IfAllRobot() {
			gameevent.Emit(arenaeventtypes.EventTypeArenaRobotTeamEnd, t, nil)
		}
	}
}

func (s *arenaService) checkArenaTeam(t *arenascene.ArenaTeam) {
	switch t.GetState() {
	case arenascene.ArenaTeamStateMatch,
		arenascene.ArenaTeamStateGameEnd,
		arenascene.ArenaTeamStateFourGodInit,
		arenascene.ArenaTeamStateFourGodEnter,
		arenascene.ArenaTeamStateFourGodQueue:
		t.KickFailed()
		break
	}
	teamId := t.GetTeam().GetTeamId()
	if !t.GetTeam().IfAllLeave() {
		return
	}
	//全部下线了
	switch t.GetState() {
	case arenascene.ArenaTeamStateInit:
		panic(fmt.Errorf("arena:队伍不应该是初始化"))
	case arenascene.ArenaTeamStateMatch,
		arenascene.ArenaTeamStateGame,
		arenascene.ArenaTeamStateGameEnd,
		arenascene.ArenaTeamStateFourGodInit,
		arenascene.ArenaTeamStateFourGodEnter:
		delete(s.arenaSceneMap, teamId)
		s.removeTeam(teamId)
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Info("arena:成员全部离线")
		break
	case arenascene.ArenaTeamStateFourGodQueue:
		delete(s.arenaSceneMap, teamId)
		s.removeTeam(teamId)
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Info("arena:成员全部离线")
		fourGodScene, _ := s.getFourGodScene(t.GetFourGodType())
		if fourGodScene == nil {
			log.WithFields(
				log.Fields{
					"teamId":      teamId,
					"fourGodType": t.GetFourGodType(),
				}).Info("arena:四圣兽场景不存在")
			return
		}
		teamCancelFourGodSceneQueue(fourGodScene, teamId)

		break
	case arenascene.ArenaTeamStateFourGod:
		s.removeTeam(teamId)
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Info("arena:成员全部离线")
		break
	}
}

func (s *arenaService) Stop() {
	return
}

var (
	once sync.Once
	as   *arenaService
)

func Init() (err error) {
	once.Do(func() {
		as = &arenaService{}
		err = as.init()
	})
	return err
}

func GetArenaService() ArenaService {
	return as
}
