package hongbao

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/hongbao/dao"
	hongbaotypes "fgame/fgame/game/hongbao/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"
	"sync"
)

//红包接口处理
type HongBaoService interface {
	Heartbeat()
	CheckExpireHongBao()
	//获取红包
	GetHongBaoObj(id int64) *HongBaoObject
	//创建红包
	CreateHongBaoObj(awardArr []*AwardInfo, hongBaoType itemtypes.ItemHongBaoSubType, sendId int64) *HongBaoObject
	//抢红包
	SnatchHongBao(hongBaoObj *HongBaoObject, pl player.Player, keepTime int64) (result hongbaotypes.HongBaoResultType, nowCount int)
}

type hongBaoService struct {
	rwm sync.RWMutex
	//红包对象
	hongBaoMap map[int64]*HongBaoObject
	runner     heartbeat.HeartbeatTaskRunner
}

//获取红包
func (s *hongBaoService) Heartbeat() {
	s.runner.Heartbeat()
	return
}

//检查过期红包
func (s *hongBaoService) CheckExpireHongBao() {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	keepTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeHongBaoKeepTime))
	now := global.GetGame().GetTimeService().Now()
	for _, hongBaoObj := range s.hongBaoMap {
		if hongBaoObj.createTime+keepTime < now {
			delete(s.hongBaoMap, hongBaoObj.GetDBId())
		}
	}
	return
}

//获取红包
func (s *hongBaoService) GetHongBaoObj(id int64) *HongBaoObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	hongBaoObj, ok := s.hongBaoMap[id]
	if !ok {
		return nil
	}
	return hongBaoObj
}

//抢红包
func (s *hongBaoService) SnatchHongBao(hongBaoObj *HongBaoObject, pl player.Player, keepTime int64) (result hongbaotypes.HongBaoResultType, nowCount int) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	result = hongbaotypes.HongBaoResultTypeSucceed
	//是否抢完
	nowCount = len(hongBaoObj.snatchLog)
	countMax := len(hongBaoObj.awardList)
	if nowCount >= countMax {
		result = hongbaotypes.HongBaoResultTypeFinish
		return
	}
	//是否抢过
	for _, temp := range hongBaoObj.snatchLog {
		if temp.PlayerId == pl.GetId() {
			result = hongbaotypes.HongBaoResultTypeSnatched
			return
		}
	}
	//是否过期
	endTime := hongBaoObj.createTime + keepTime
	now := global.GetGame().GetTimeService().Now()
	if endTime < now {
		result = hongbaotypes.HongBaoResultTypeEndTime
		return
	}

	snatcher := &SnatcherInfo{}
	snatcher.PlayerId = pl.GetId()
	snatcher.Name = pl.GetName()
	snatcher.Role = pl.GetRole()
	snatcher.Sex = pl.GetSex()
	snatcher.Level = pl.GetLevel()
	hongBaoObj.snatchLog = append(hongBaoObj.snatchLog, snatcher)
	hongBaoObj.updateTime = now
	hongBaoObj.SetModified()
	return
}

//初始化服务器
func (s *hongBaoService) init() (err error) {
	now := global.GetGame().GetTimeService().Now()
	keepTime := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeHongBaoKeepTime))
	existTime := now - keepTime
	entityList, err := dao.GetHongBaoDao().GetAllHongBaoEntity(existTime)
	if err != nil {
		return err
	}
	s.hongBaoMap = make(map[int64]*HongBaoObject)
	if entityList != nil {
		for _, hongBaoEntity := range entityList {
			hongBaoObj := NewHongBaoObject()
			err = hongBaoObj.FromEntity(hongBaoEntity)
			if err != nil {
				return err
			}
			s.hongBaoMap[hongBaoObj.id] = hongBaoObj
		}
	}
	s.runner = heartbeat.NewHeartbeatTaskRunner()
	s.runner.AddTask(CreateHongBaoExpireTask(s))
	return
}

//创建新红包
func (s *hongBaoService) CreateHongBaoObj(awardArr []*AwardInfo, hongBaoType itemtypes.ItemHongBaoSubType, sendId int64) *HongBaoObject {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	peo := NewHongBaoObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	peo.id = id
	peo.serverId = global.GetGame().GetServerIndex()
	peo.createTime = now
	s.hongBaoMap[peo.id] = peo
	peo.hongBaoType = hongBaoType
	peo.sendId = sendId
	peo.awardList = awardArr

	peo.SetModified()
	return peo
}

var (
	once sync.Once
	cs   *hongBaoService
)

func Init() (err error) {
	once.Do(func() {
		cs = &hongBaoService{}
		err = cs.init()
	})
	return err
}

func GetHongBaoService() HongBaoService {
	return cs
}
