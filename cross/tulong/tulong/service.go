package tulong

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/cross/tulong/dao"
	tulongscene "fgame/fgame/cross/tulong/scene"
	tulongcrosstypes "fgame/fgame/cross/tulong/types"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	tulongtemplate "fgame/fgame/game/tulong/template"
	tulongtypes "fgame/fgame/game/tulong/types"
	"sort"
	"sync"
)

type TuLongService interface {
	Heartbeat()
	CheckTuLongActivityTask() error

	//创建屠龙场景
	CreateTuLongScene(mapId int32, endTime int64) scene.Scene
	//获取屠龙场景
	GetTuLongScene() scene.Scene
	//屠龙活动结束
	TuLongSceneFinish()

	//获取排行榜
	GetRankList() (dataList []*TuLongRankObject)
	//获取出生标识
	GetPlayerBornBiaoShi(allianceId int64) (biaoShi int32, flag bool)
	//击杀boss
	KillBoss(serverId int32, allianceId int64, allianceName string)
}

type tuLongService struct {
	rwm         sync.RWMutex
	hbRunner    heartbeat.HeartbeatTaskRunner
	tuLongScene scene.Scene

	//仙盟出生标识信息
	allianceMap map[int64]int32
	//大boss出生标识
	bigBossBiaoShi int32
	//屠龙排行榜数据
	tuLongRankList []*TuLongRankObject
}

func (s *tuLongService) init() (err error) {
	s.allianceMap = make(map[int64]int32)
	s.tuLongRankList = make([]*TuLongRankObject, 0, tulongcrosstypes.TuLongRankSize)

	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	s.hbRunner.AddTask(CreateTuLongStartTask(s))

	err = s.loadTuLongRankData()
	if err != nil {
		return
	}

	return
}

func (s *tuLongService) loadTuLongRankData() (err error) {
	platform := global.GetGame().GetPlatform()
	areaId := global.GetGame().GetServerIndex()

	//排行榜列表
	tuLongRankList, err := dao.GetTuLongDao().GetTuLongRankList(platform, areaId)
	if err != nil {
		return
	}
	for _, tuLongRank := range tuLongRankList {
		tlro := NewTuLongRankObject()
		tlro.FromEntity(tuLongRank)
		s.tuLongRankList = append(s.tuLongRankList, tlro)
	}

	return
}

func (s *tuLongService) GetRankList() (dataList []*TuLongRankObject) {
	len := len(s.tuLongRankList)
	if len == 0 {
		return
	}

	addLen := int32(tulongcrosstypes.TuLongRankSize)
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return s.tuLongRankList[0:addLen]
}

func (s *tuLongService) getRankData(allianceId int64) *TuLongRankObject {
	for _, tuLongRank := range s.tuLongRankList {
		if tuLongRank.AllianceId != allianceId {
			continue
		}
		return tuLongRank
	}

	return nil
}

func (s *tuLongService) KillBoss(serverId int32, allianceId int64, allianceName string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	rankObj := s.getRankData(allianceId)
	if rankObj == nil {
		rankObj = initNewTuLongRankObject(serverId, allianceId, allianceName)
		s.tuLongRankList = append(s.tuLongRankList, rankObj)
	} else {
		rankObj.KillNum += 1
		rankObj.LastTime = now
		rankObj.UpdateTime = now
		rankObj.SetModified()
	}

	sort.Sort(sort.Reverse(TuLongRankObjectList(s.tuLongRankList)))
}

func (s *tuLongService) CreateTuLongScene(mapId int32, endTime int64) (data scene.Scene) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	return s.createTuLongScene(mapId, endTime)
}

func (s *tuLongService) createTuLongScene(mapId int32, endTime int64) scene.Scene {
	if s.tuLongScene != nil {
		return s.tuLongScene
	}

	bigEggBornBiaoShi, flag := tulongtemplate.GetTuLongTemplateService().GetTuLongPosBiaoShi(tulongtypes.TuLongPosTypeBoss, nil)
	if !flag {
		return nil
	}
	s.bigBossBiaoShi = bigEggBornBiaoShi

	tuLongSceneData := tulongscene.CreateTuLongSceneData(bigEggBornBiaoShi)
	sc := tulongscene.CreateTuLongScene(mapId, endTime, tuLongSceneData)
	if sc != nil {
		s.tuLongScene = sc
		s.clearTuLongRank()
	}

	return sc
}

func (s *tuLongService) clearTuLongRank() {
	now := global.GetGame().GetTimeService().Now()
	for _, rankObj := range s.tuLongRankList {
		rankObj.DeleteTime = now
		rankObj.SetModified()
	}
	s.tuLongRankList = make([]*TuLongRankObject, 0, tulongcrosstypes.TuLongRankSize)
}

func (s *tuLongService) GetTuLongScene() scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.tuLongScene
}

func (s *tuLongService) TuLongSceneFinish() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.tuLongScene = nil
	s.allianceMap = make(map[int64]int32)
	s.bigBossBiaoShi = 0
}

func (s *tuLongService) GetPlayerBornBiaoShi(allianceId int64) (biaoShi int32, flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if s.tuLongScene == nil {
		return
	}

	biaoShi, exist := s.allianceMap[allianceId]
	if exist {
		flag = exist
		return
	}

	biaoShi, ok := s.randomPlayerBornBiaoShi()
	if !ok {
		return
	}

	s.allianceMap[allianceId] = biaoShi
	flag = true
	return
}

func (s *tuLongService) randomPlayerBornBiaoShi() (biaoShi int32, flag bool) {
	if s.bigBossBiaoShi == 0 {
		return
	}
	var biaoShiList []int32
	biaoShiList = append(biaoShiList, s.bigBossBiaoShi)

	biaoShi, flag = tulongtemplate.GetTuLongTemplateService().GetTuLongPosBiaoShi(tulongtypes.TuLongPosTypePlayer, biaoShiList)
	if !flag {
		panic("tulong: 获取出生标识应该是ok的")
	}

	flag = true
	return
}

func (s *tuLongService) CheckTuLongActivityTask() (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if s.tuLongScene != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeCoressTuLong)
	activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
	if err != nil {
		return
	}
	if activityTimeTemplate == nil {
		return
	}
	endTime, err := activityTimeTemplate.GetEndTime(now)
	if err != nil {
		return
	}

	s.createTuLongScene(activityTemplate.Mapid, endTime)
	return nil
}

func (s *tuLongService) Heartbeat() {
	s.hbRunner.Heartbeat()
}

var (
	once sync.Once
	cs   *tuLongService
)

func Init() (err error) {
	once.Do(func() {
		cs = &tuLongService{}
		err = cs.init()
	})
	return err
}

func GetTuLongService() TuLongService {
	return cs
}
