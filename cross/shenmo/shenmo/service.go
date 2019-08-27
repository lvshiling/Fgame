package shenmo

import (
	"fgame/fgame/cross/shenmo/dao"
	"fgame/fgame/game/global"
	shenmotypes "fgame/fgame/game/shenmo/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sort"
	"sync"
)

type ShenMoService interface {
	Heartbeat()
	//获取上周排行榜
	GetLastRankList() (lastTime int64, lastDataList []*ShenMoRankObject)
	//获取本周排行榜
	GetThisRankList() (thisTime int64, thisDataList []*ShenMoRankObject)
	//获取排行榜积分
	GetJiFenNum(allianceId int64) int32
	//仙盟增加积分
	AddJiFenNum(serverId int32, allianceId int64, allianceName string, addNum int32)
}

type shenMoService struct {
	rwm sync.RWMutex
	//本周神魔战场排行榜数据
	thisShenMoRankList []*ShenMoRankObject
	//上周神魔战场排行榜数据
	lastShenMoRankList []*ShenMoRankObject

	//排行榜刷新时间对象
	shenMoRankTimeObj *ShenMoRankTimeObject
}

func (s *shenMoService) init() (err error) {
	s.thisShenMoRankList = make([]*ShenMoRankObject, 0, shenmotypes.ShenMoRankSize)
	s.lastShenMoRankList = make([]*ShenMoRankObject, 0, shenmotypes.ShenMoRankSize)
	platform := global.GetGame().GetPlatform()
	//排行榜列表
	shenMoRankList, err := dao.GetShenMoDao().GetShenMoRankList(platform)
	if err != nil {
		return
	}
	for _, shenMoRank := range shenMoRankList {
		obj := NewShenMoRankObject()
		obj.FromEntity(shenMoRank)
		s.thisShenMoRankList = append(s.thisShenMoRankList, obj)
		s.lastShenMoRankList = append(s.lastShenMoRankList, obj)
	}

	sort.Sort(sort.Reverse(ThisShenMoRankObjectList(s.thisShenMoRankList)))
	sort.Sort(sort.Reverse(LastShenMoRankObjectList(s.lastShenMoRankList)))

	//排行刷新时间
	shenMoRankTimeEntity, err := dao.GetShenMoDao().GetShenMoRankTimeEntity(platform)
	if err != nil {
		return
	}
	if shenMoRankTimeEntity == nil {
		s.initShenMoRankTimeObject(platform)
	} else {
		s.shenMoRankTimeObj = NewShenMoRankTimeObject()
		s.shenMoRankTimeObj.FromEntity(shenMoRankTimeEntity)
	}

	return nil
}

//第一次初始化
func (s *shenMoService) initShenMoRankTimeObject(platform int32) {
	po := NewShenMoRankTimeObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	po.Id = id
	//生成id
	po.Platform = platform
	po.ThisTime = 0
	po.LastTime = 0
	po.CreateTime = now
	s.shenMoRankTimeObj = po
	po.SetModified()
}

func (s *shenMoService) GetThisRankList() (thisTime int64, dataList []*ShenMoRankObject) {
	len := len(s.thisShenMoRankList)
	if len == 0 {
		return
	}
	thisTime = s.shenMoRankTimeObj.ThisTime
	addLen := int32(shenmotypes.ShenMoRankSize)
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return thisTime, s.thisShenMoRankList[0:addLen]
}

func (s *shenMoService) GetLastRankList() (lastTime int64, dataList []*ShenMoRankObject) {
	len := len(s.lastShenMoRankList)
	if len == 0 {
		return
	}
	lastTime = s.shenMoRankTimeObj.LastTime
	addLen := int32(shenmotypes.ShenMoRankSize)
	if addLen >= int32(len) {
		addLen = int32(len)
	}
	return lastTime, s.lastShenMoRankList[0:addLen]
}

func (s *shenMoService) GetJiFenNum(allianceId int64) (jiFenNum int32) {
	if len(s.thisShenMoRankList) == 0 {
		return
	}
	for _, shenMoRank := range s.thisShenMoRankList {
		if shenMoRank.AllianceId == allianceId {
			return shenMoRank.JiFenNum
		}
	}
	return
}

func (s *shenMoService) AddJiFenNum(serverId int32, allianceId int64, allianceName string, addNum int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj, flag := s.isExistAlliance(allianceId)
	if !flag {
		obj = initShenMoRankObject(serverId, allianceId, allianceName, addNum)
		s.thisShenMoRankList = append(s.thisShenMoRankList, obj)
		s.lastShenMoRankList = append(s.lastShenMoRankList, obj)
	} else {
		obj.JiFenNum += addNum
	}
	obj.SetModified()
	sort.Sort(sort.Reverse(ThisShenMoRankObjectList(s.thisShenMoRankList)))
}

func (s *shenMoService) isExistAlliance(allianceId int64) (shenMoRankObj *ShenMoRankObject, flag bool) {
	if len(s.thisShenMoRankList) == 0 {
		return
	}
	for _, shenMoRank := range s.thisShenMoRankList {
		if shenMoRank.AllianceId == allianceId {
			return shenMoRank, true
		}
	}
	return
}

func (s *shenMoService) Heartbeat() {
	s.checkRefreshWeekRank()
}

func (s *shenMoService) checkRefreshWeekRank() {
	now := global.GetGame().GetTimeService().Now()
	if s.shenMoRankTimeObj.ThisTime == 0 {
		timeStamp, _ := timeutils.MondayFivePointTime(now)
		s.shenMoRankTimeObj.ThisTime = timeStamp
		s.shenMoRankTimeObj.UpdateTime = now
		s.shenMoRankTimeObj.SetModified()
		return
	}

	diffTime := now - s.shenMoRankTimeObj.ThisTime
	if diffTime < 0 {
		diffTime *= -1
	}
	if diffTime < int64(7*timeutils.DAY) {
		return
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.refreshWeekRank(now)
}

func (s *shenMoService) refreshWeekRank(now int64) {
	timeStamp, _ := timeutils.MondayFivePointTime(now)
	s.shenMoRankTimeObj.LastTime = timeStamp - int64(7*timeutils.DAY)
	s.shenMoRankTimeObj.ThisTime = timeStamp
	s.shenMoRankTimeObj.UpdateTime = now
	s.shenMoRankTimeObj.SetModified()

	for _, rankObj := range s.thisShenMoRankList {
		rankObj.LastJiFenNum = rankObj.JiFenNum
		rankObj.JiFenNum = 0
		rankObj.LastTime = now
		rankObj.UpdateTime = now
		rankObj.SetModified()
	}

	if len(s.thisShenMoRankList) != 0 {
		sort.Sort(sort.Reverse(ThisShenMoRankObjectList(s.thisShenMoRankList)))
	}
	if len(s.lastShenMoRankList) != 0 {
		sort.Sort(sort.Reverse(LastShenMoRankObjectList(s.lastShenMoRankList)))
	}
}

var (
	once sync.Once
	cs   *shenMoService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shenMoService{}
		err = cs.init()
	})
	return err
}

func GetShenMoService() ShenMoService {
	return cs
}
