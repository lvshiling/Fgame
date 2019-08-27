package arenapvp

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/runner"
	"fgame/fgame/cross/arenapvp/dao"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	gamearenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/timeutils"
	"math"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type ArenapvpService interface {
	runner.Task
	Start()
	GetArenapvpSceneElection() scene.Scene                                          //获取海选场景
	CreateArenapvpSceneElection(endTime int64) scene.Scene                          //创建海选场景
	GetAllElectionSceneMap() map[int64]scene.Scene                                  //获取所有海选场景
	ArenapvpElectionFinish(s scene.Scene, winnerList []*arenapvpdata.PvpPlayerInfo) //海选结束

	CreateArenapvpSceneBattle(pvpTemp *gametemplate.ArenapvpTemplate, playerId, endTime int64) scene.Scene //创建下一场pvp对战场景
	ArenapvpBattleFinish(s scene.Scene)                                                                    //晋级pvp结束
	GetArenapvpBattleScene(playerId int64) scene.Scene                                                     //获取对战场景

	GetPvpResultList() []*arenapvpdata.PvpPlayerInfo             //所有pvp比赛结果map
	GetBaZhuList() []*ArenapvpBaZhuObject                        //霸主列表
	GetGuessDataList() []*arenapvpdata.GuessData                 //竞猜组
	GetPvpPlayerInfo(playerId int64) *arenapvpdata.PvpPlayerInfo //
	// GetLuckyPlayerList() []player.Player                         //幸运奖玩家列表

	//玩家退出
	PvpPlayerExit(playerId int64) (pvpPlayer *arenapvpdata.PvpPlayerInfo, flag bool)
	//玩家失败
	PvpPlayerFailed(playerId int64) (pvpPlayer *arenapvpdata.PvpPlayerInfo, flag bool)

	//定时器任务
	RefreshBattleResult() //刷新对战结果
}

type arenapvpService struct {
	rwm      sync.RWMutex
	hbRunner heartbeat.HeartbeatTaskRunner

	//历届霸主列表
	bazhuList []*ArenapvpBaZhuObject

	//海选pvp场景
	sceneMap map[int64]scene.Scene
	//晋级赛场景
	battleSceneMap map[int64]scene.Scene

	//32强玩家
	pvpPlayerMap map[int64]*arenapvpdata.PvpPlayerInfo
	//竞猜对象
	guessDataList []*arenapvpdata.GuessData

	pvpEndTime  int64
	arenaPvpEnd bool
}

func (s *arenapvpService) init() (err error) {
	s.sceneMap = make(map[int64]scene.Scene)
	s.battleSceneMap = make(map[int64]scene.Scene)
	s.pvpPlayerMap = make(map[int64]*arenapvpdata.PvpPlayerInfo)
	s.arenaPvpEnd = true

	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	s.hbRunner.AddTask(CreateArenapvpResultTask(s))

	err = s.loadBaZhuList()
	if err != nil {
		return
	}

	return nil
}

func (s *arenapvpService) loadBaZhuList() (err error) {

	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	//霸主列表
	ascEntityList, err := dao.GetArenapvpDao().GetArenapvpBaZhuList(platform, serverId)
	if err != nil {
		return
	}

	for _, entity := range ascEntityList {
		obj := NewArenapvpBaZhuObject()
		obj.FromEntity(entity)
		s.bazhuList = append(s.bazhuList, obj)
	}
	return
}

//获取上一届
func (s *arenapvpService) gePreRaceNum() int32 {
	if len(s.bazhuList) <= 0 {
		return 0
	}
	return s.bazhuList[len(s.bazhuList)-1].RaceNumber
}

func (s *arenapvpService) Start() {
}

func (s *arenapvpService) Heartbeat() {
	s.hbRunner.Heartbeat()
}

func (s *arenapvpService) GetPvpResultList() (pList []*arenapvpdata.PvpPlayerInfo) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	for _, p := range s.pvpPlayerMap {
		pList = append(pList, p)
	}
	return pList
}

func (s *arenapvpService) GetBaZhuList() []*ArenapvpBaZhuObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.bazhuList
}

func (s *arenapvpService) GetGuessDataList() []*arenapvpdata.GuessData {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.guessDataList
}

func (s *arenapvpService) GetPvpPlayerInfo(playerId int64) *arenapvpdata.PvpPlayerInfo {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.pvpPlayerMap[playerId]
}

func (s *arenapvpService) PvpPlayerExit(playerId int64) (pvpPlayer *arenapvpdata.PvpPlayerInfo, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	pvpPlayer, ok := s.pvpPlayerMap[playerId]
	if !ok {
		return nil, false
	}
	flag = pvpPlayer.Exit()
	if !flag {
		return pvpPlayer, false
	}
	return pvpPlayer, true
}

func (s *arenapvpService) PvpPlayerFailed(playerId int64) (pvpPlayer *arenapvpdata.PvpPlayerInfo, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	pvpPlayer, ok := s.pvpPlayerMap[playerId]
	if !ok {
		return nil, false
	}
	flag = pvpPlayer.Failed()
	if !flag {
		return pvpPlayer, false
	}
	return pvpPlayer, true
}

func (s *arenapvpService) ArenapvpElectionFinish(sc scene.Scene, winnerList []*arenapvpdata.PvpPlayerInfo) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if sc == nil {
		return
	}

	sd, ok := sc.SceneDelegate().(arenapvpscene.ArenapvpSceneData)
	if !ok {
		return
	}

	//晋级玩家
	for _, pvpInfo := range winnerList {
		s.pvpPlayerMap[pvpInfo.PlayerId] = pvpInfo
	}

	delete(s.sceneMap, sc.Id())
	if len(s.sceneMap) == 0 {
		//海选结束匹配对手
		s.matchBattle(sd.GetPvpTemp())
	}
}

func (s *arenapvpService) ArenapvpBattleFinish(sc scene.Scene) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	sd, ok := sc.SceneDelegate().(arenapvpscene.ArenapvpBattleSceneData)
	if !ok {
		return
	}
	log.WithFields(
		log.Fields{
			"winnerId": sd.GetWinnerId(),
		}).Info("arenapvp:场景结束")
	if sd.GetWinnerId() == 0 {
		return
	}
	battlePl, ok := s.pvpPlayerMap[sd.GetWinnerId()]
	if !ok {
		log.WithFields(
			log.Fields{
				"winnerId": sd.GetWinnerId(),
				"玩家长度":     len(s.pvpPlayerMap),
			}).Warn("arenapvp:场景结束,获胜玩家不存在")
		return
	}
	battleData := battlePl.GetBattleData(sd.GetPvpTemp().GetArenapvpType())
	if battleData == nil {
		log.WithFields(
			log.Fields{
				"winnerId": sd.GetWinnerId(),
			}).Warn("arenapvp:场景结束,战斗数据不存在")
		return
	}

	battleData.WinnerId = sd.GetWinnerId()
	delete(s.battleSceneMap, battleData.BattleId1)
	delete(s.battleSceneMap, battleData.BattleId2)

	// 决赛
	if battleData.PvpType == gamearenapvptypes.ArenapvpTypeFinals {
		log.WithFields(
			log.Fields{
				"pvpType": battleData.PvpType.String(),
			}).Info("arenapvp:已经决赛")
		raceNum := s.gePreRaceNum() + 1
		platform := global.GetGame().GetPlatform()
		serverId := global.GetGame().GetServerIndex()
		baZhu := CreateArenapvpBaZhuObjectWithPvpPlayerInfo(battlePl, platform, serverId, raceNum)
		s.bazhuList = append(s.bazhuList, baZhu)
		s.arenaPvpEnd = true
		return
	}

	if len(s.battleSceneMap) == 0 {
		//晋级结束匹配对手
		s.matchBattle(sd.GetPvpTemp())
	}
}

func (s *arenapvpService) GetArenapvpBattleScene(playerId int64) scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	sc, ok := s.battleSceneMap[playerId]
	if !ok {
		return nil
	}
	return sc
}

func (s *arenapvpService) GetArenapvpSceneElection() scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.getEnterScene()
}

func (s *arenapvpService) GetAllElectionSceneMap() map[int64]scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.sceneMap
}

func (s *arenapvpService) CreateArenapvpSceneBattle(pvpTemp *gametemplate.ArenapvpTemplate, playerId, endTime int64) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	pvpPlayer1, ok := s.pvpPlayerMap[playerId]
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("arenapvp:玩家不存在")
		return nil
	}

	sc, ok := s.battleSceneMap[playerId]
	if !ok {
		battleData := pvpPlayer1.GetBattleData(pvpTemp.GetArenapvpType())
		if battleData == nil {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"pvptype":  pvpTemp.GetArenapvpType().String(),
				}).Warn("arenapvp:匹配数据不存在")
			return nil
		}
		pvpPlayer2 := s.pvpPlayerMap[battleData.GetBattleId(playerId)]
		asd := arenapvpscene.CreateArenapvpBattleSceneData(endTime, pvpTemp, pvpPlayer1, pvpPlayer2)
		now := global.GetGame().GetTimeService().Now()
		scEndTime := endTime
		if pvpTemp.GetNextTemp() != nil {
			scEndTime = pvpTemp.GetNextTemp().GetBeginTime(now)
		}
		sc = scene.CreateArenaScene(pvpTemp.MapId, scEndTime, asd)

		if battleData.BattleId1 != 0 {
			s.battleSceneMap[battleData.BattleId1] = sc
		}
		if battleData.BattleId2 != 0 {
			s.battleSceneMap[battleData.BattleId2] = sc
		}
	}

	return sc
}

func (s *arenapvpService) CreateArenapvpSceneElection(endTime int64) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if !s.arenaPvpEnd {
		// TODO :xzk26 会出现这个情况
		log.Infoln("arenapvp:比武大会未结束")
		return nil
	}

	if len(s.sceneMap) > 0 {
		return s.getEnterScene()
	}

	if s.pvpEndTime != endTime {
		s.arenaPvpEnd = false
		s.pvpEndTime = endTime
		s.pvpPlayerMap = make(map[int64]*arenapvpdata.PvpPlayerInfo)
		s.guessDataList = []*arenapvpdata.GuessData{}
	}

	pvpTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTemplate(gamearenapvptypes.ArenapvpTypeElection)
	mapId := pvpTemp.MapId
	now := global.GetGame().GetTimeService().Now()
	scEndTime := pvpTemp.GetNextTemp().GetBeginTime(now)
	for i := int32(0); i < pvpTemp.PVPCount; i++ {
		asd := arenapvpscene.CreateArenapvpSceneData(endTime, pvpTemp, i)
		sc := scene.CreateArenaScene(mapId, scEndTime, asd)
		s.sceneMap[sc.Id()] = sc
	}

	return s.getEnterScene()
}

func (s *arenapvpService) RefreshBattleResult() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameDay(now, s.pvpEndTime)
	if !isSame {
		s.pvpPlayerMap = make(map[int64]*arenapvpdata.PvpPlayerInfo)
		s.guessDataList = []*arenapvpdata.GuessData{}
	}
}

//获取要进入的会场
func (s *arenapvpService) getEnterScene() scene.Scene {
	sCount := int32(len(s.sceneMap))
	if sCount == 0 {
		return nil
	}

	//人数最少的场景
	minPlCount := int32(0)
	var enterS scene.Scene
	for _, sc := range s.sceneMap {
		plCount := int32(len(sc.GetAllPlayers()))
		if enterS == nil {
			enterS = sc
			minPlCount = plCount
			continue
		}

		if minPlCount <= plCount {
			continue
		}

		enterS = sc
		minPlCount = plCount
	}

	return enterS
}

//匹配对手
func (s *arenapvpService) matchBattle(pvpTemp *gametemplate.ArenapvpTemplate) {
	log.WithFields(
		log.Fields{
			"pvpType": pvpTemp.GetArenapvpType().String(),
		}).Info("arenapvp:场景结束,匹配对手")
	nextPvpTemp := pvpTemp.GetNextTemp()
	curPvpType := pvpTemp.GetArenapvpType()
	nextPvpType := nextPvpTemp.GetArenapvpType()

	// 决赛
	if curPvpType == gamearenapvptypes.ArenapvpTypeFinals {
		log.WithFields(
			log.Fields{
				"pvpType": pvpTemp.GetArenapvpType().String(),
			}).Info("arenapvp:场景结束,决赛了")
		return
	}

	// 当前获胜选手
	var battlePlList []*arenapvpdata.PvpPlayerInfo
	switch curPvpType {
	case gamearenapvptypes.ArenapvpTypeElection:
		{
			var realPlList []*arenapvpdata.PvpPlayerInfo
			var robotPlList []*arenapvpdata.PvpPlayerInfo
			for _, battlePl := range s.pvpPlayerMap {
				if battlePl.IsRobot {
					robotPlList = append(robotPlList, battlePl)
				} else {
					realPlList = append(realPlList, battlePl)
				}
			}
			battlePlList = append(battlePlList, realPlList...)
			battlePlList = append(battlePlList, robotPlList...)
		}
	case gamearenapvptypes.ArenapvpTypeTop32,
		gamearenapvptypes.ArenapvpTypeTop16,
		gamearenapvptypes.ArenapvpTypeTop8,
		gamearenapvptypes.ArenapvpTypeTop4:
		{
			var realPlList []*arenapvpdata.PvpPlayerInfo
			var robotPlList []*arenapvpdata.PvpPlayerInfo
			for _, battlePl := range s.pvpPlayerMap {
				battleData := battlePl.GetBattleData(curPvpType)
				if battleData == nil {
					continue
				}

				//获胜选手
				if battleData.WinnerId != battlePl.PlayerId {
					continue
				}

				if battlePl.IsRobot {
					robotPlList = append(robotPlList, battlePl)
				} else {
					realPlList = append(realPlList, battlePl)
				}
			}

			battlePlList = append(battlePlList, realPlList...)
			battlePlList = append(battlePlList, robotPlList...)
		}
	}

	count := len(battlePlList)
	middle := count / 2

	battleGroupMap := make(map[int][]*arenapvpdata.PvpPlayerInfo)
	// 下一场选对手
	switch nextPvpType {
	case gamearenapvptypes.ArenapvpTypeTop32,
		gamearenapvptypes.ArenapvpTypeTop16,
		gamearenapvptypes.ArenapvpTypeTop8:
		{
			// 优先匹配真假人
			pvpConstantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()

			// 组成一组
			for index := 0; index < middle; index++ {
				// 优先匹配真人：相邻玩家
				battleIndex1 := index * 2
				battleIndex2 := index*2 + 1
				if pvpConstantTemp.IsJiaRen() {
					// 优先匹配假人：首尾玩家
					battleIndex1 = index
					battleIndex2 = count - 1 - index
				}

				battlePl1 := battlePlList[battleIndex1]
				battlePl2 := battlePlList[battleIndex2]

				battleData := arenapvpdata.CreateBattleResultData(nextPvpType, battlePl1.PlayerId, battlePl2.PlayerId, int32(index))
				battlePl1.BattleDataList = append(battlePl1.BattleDataList, battleData)
				battlePl2.BattleDataList = append(battlePl2.BattleDataList, battleData)

				battleGroupMap[index] = append(battleGroupMap[index], battlePl1)
				battleGroupMap[index] = append(battleGroupMap[index], battlePl2)
			}

			if count%2 != 0 {
				// 优先匹配真人：相邻玩家
				singleIndex := count - 1
				if pvpConstantTemp.IsJiaRen() {
					// 优先匹配假人：首尾玩家
					singleIndex = middle
				}

				singlePl := battlePlList[singleIndex]
				battleData := arenapvpdata.CreateBattleResultData(nextPvpType, singlePl.PlayerId, 0, int32(middle))
				singlePl.BattleDataList = append(singlePl.BattleDataList, battleData)

				battleGroupMap[middle] = append(battleGroupMap[middle], singlePl)
			}
		}
	case gamearenapvptypes.ArenapvpTypeTop4,
		gamearenapvptypes.ArenapvpTypeFinals:
		{

			preIndexPlMap := make(map[int32]*arenapvpdata.PvpPlayerInfo)
			for _, battlePl := range battlePlList {
				battleData := battlePl.GetBattleData(curPvpType)
				preIndexPlMap[battleData.Index] = battlePl
			}

			for index := 0; index < middle; index++ {
				//上一场相邻两组获胜选手自动成为下一组对战
				preIndex1 := int32(index * 2)
				preIndex2 := int32(index*2 + 1)

				battlePl1 := preIndexPlMap[preIndex1]
				battlePl2 := preIndexPlMap[preIndex2]

				battleData := arenapvpdata.CreateBattleResultData(nextPvpType, battlePl1.PlayerId, battlePl2.PlayerId, int32(index))
				battlePl1.BattleDataList = append(battlePl1.BattleDataList, battleData)
				battlePl2.BattleDataList = append(battlePl2.BattleDataList, battleData)

				battleGroupMap[index] = append(battleGroupMap[index], battlePl1)
				battleGroupMap[index] = append(battleGroupMap[index], battlePl2)
			}

			if count%2 != 0 {
				singlePl := battlePlList[middle]
				battleData := arenapvpdata.CreateBattleResultData(nextPvpType, singlePl.PlayerId, 0, int32(middle))
				singlePl.BattleDataList = append(singlePl.BattleDataList, battleData)

				battleGroupMap[middle] = append(battleGroupMap[middle], singlePl)
			}
		}
	}

	//战力相近
	guessIndex := -1
	preDiff := float64(0)
	for index, battleList := range battleGroupMap {
		if len(battleList) != 2 {
			continue
		}

		diff := math.Abs(float64(battleList[0].Force - battleList[1].Force))
		if index == 0 {
			preDiff = diff
			guessIndex = index
			continue
		}

		if diff >= preDiff {
			continue
		}
		preDiff = diff
		guessIndex = index
	}

	log.WithFields(
		log.Fields{
			"guessIndex":     guessIndex,
			"battleGroupLen": len(battleGroupMap),
		}).Infoln("arenapvp:匹配完成,选择竞猜数据")

	if guessIndex != -1 {
		guessList := battleGroupMap[guessIndex]
		raceNum := s.gePreRaceNum() + 1
		s.guessDataList = append(s.guessDataList, arenapvpdata.CreateGuessData(nextPvpType, raceNum, guessList))
	}
	return
}

var (
	once sync.Once
	as   *arenapvpService
)

func Init() (err error) {
	once.Do(func() {
		as = &arenapvpService{}
		err = as.init()
	})
	return err
}

func GetArenapvpService() ArenapvpService {
	return as
}
