package baby

import (
	"fgame/fgame/game/baby/dao"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	babytypes "fgame/fgame/game/baby/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"
	"sync"
)

//配偶宝宝接口处理
type BabyService interface {
	SyncBabyLearnLevel(pl player.Player, babyId int64, learnLevel int32)                                           //同步学习等级
	SyncBabyTalentList(pl player.Player, babyId int64, talentList []*babytypes.TalentInfo)                         //同步天赋技能
	BabyZhuanShi(pl player.Player, babyId int64)                                                                   //宝宝转世
	AddBaby(pl player.Player, babyId int64, quality, learnLevel, beishu int32, talentList []*babytypes.TalentInfo) //添加宝宝
	GetCoupleBabyInfo(spouseId int64) []*babytypes.CoupleBabyData                                                  //配偶宝宝信息
}

type babyService struct {
	//读写锁
	rwm sync.RWMutex
	//宝宝map
	babyMap map[int64]*CoupleBabyObject
}

//初始化
func (s *babyService) init() (err error) {
	s.babyMap = make(map[int64]*CoupleBabyObject)

	babyList, err := dao.GetBabyDao().GetCoupleBabyEntityList(global.GetGame().GetServerIndex())
	if err != nil {
		return
	}
	for _, entity := range babyList {
		bo := NewCoupleBabyObject()
		bo.FromEntity(entity)
		s.babyMap[bo.playerId] = bo
	}

	return
}

func (s *babyService) SyncBabyLearnLevel(pl player.Player, babyId int64, learnLevel int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	babyInfo := s.babyMap[pl.GetId()]
	if babyInfo == nil {
		return
	}

	_, babyData := s.getBaby(pl.GetId(), babyId)
	if babyData == nil {
		return
	}
	babyData.LearnLevel = learnLevel

	now := global.GetGame().GetTimeService().Now()
	babyInfo.updateTime = now
	babyInfo.SetModified()

	gameevent.Emit(babyeventtypes.EventTypeCoupleBabyChanged, pl, babyInfo.babyList)
}

func (s *babyService) SyncBabyTalentList(pl player.Player, babyId int64, talenList []*babytypes.TalentInfo) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	babyInfo := s.babyMap[pl.GetId()]
	if babyInfo == nil {
		return
	}

	_, babyData := s.getBaby(pl.GetId(), babyId)
	if babyData == nil {
		return
	}
	babyData.TalentList = talenList

	now := global.GetGame().GetTimeService().Now()
	babyInfo.updateTime = now
	babyInfo.SetModified()

	gameevent.Emit(babyeventtypes.EventTypeCoupleBabyChanged, pl, babyInfo.babyList)
}

func (s *babyService) BabyZhuanShi(pl player.Player, babyId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	playerId := pl.GetId()
	babyInfo, ok := s.babyMap[playerId]
	if !ok {
		return
	}

	delIndex, _ := s.getBaby(playerId, babyId)
	if delIndex == -1 {
		return
	}

	babyInfo.babyList = append(babyInfo.babyList[:delIndex], babyInfo.babyList[delIndex+1:]...)

	now := global.GetGame().GetTimeService().Now()
	babyInfo.updateTime = now
	babyInfo.SetModified()

	gameevent.Emit(babyeventtypes.EventTypeCoupleBabyChanged, pl, babyInfo.babyList)
}

func (s *babyService) AddBaby(pl player.Player, babyId int64, quality, learnLevel, danbei int32, talentList []*babytypes.TalentInfo) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	babyInfo, ok := s.babyMap[pl.GetId()]
	if !ok {
		babyInfo = s.newCoupleBabyObject(pl.GetId())
	}

	babyData := babytypes.NewCoupleBabyData(babyId, quality, learnLevel, danbei, talentList)
	babyInfo.babyList = append(babyInfo.babyList, babyData)
	gameevent.Emit(babyeventtypes.EventTypeCoupleBabyChanged, pl, babyInfo.babyList)
}

// 获取玩家宝宝信息
func (s *babyService) GetCoupleBabyInfo(spouseId int64) (babyList []*babytypes.CoupleBabyData) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	babyInfo, ok := s.babyMap[spouseId]
	if !ok {
		return
	}

	return babyInfo.babyList
}

func (s *babyService) getBaby(plId, babyId int64) (int, *babytypes.CoupleBabyData) {
	babyInfo := s.babyMap[plId]
	for index, baby := range babyInfo.babyList {
		if baby.BabyId != babyId {
			continue
		}

		return index, baby
	}

	return -1, nil
}

func (s *babyService) newCoupleBabyObject(playerId int64) *CoupleBabyObject {
	now := global.GetGame().GetTimeService().Now()
	bo := NewCoupleBabyObject()
	id, _ := idutil.GetId()
	bo.id = id
	bo.serverId = global.GetGame().GetServerIndex()
	bo.playerId = playerId
	bo.createTime = now
	s.babyMap[bo.playerId] = bo
	bo.SetModified()
	return bo
}

var (
	once sync.Once
	cs   *babyService
)

func Init() (err error) {
	once.Do(func() {
		cs = &babyService{}
		err = cs.init()
	})
	return err
}

func GetBabyService() BabyService {
	return cs
}
