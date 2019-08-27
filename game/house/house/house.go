package house

import (
	dummytemplate "fgame/fgame/game/dummy/template"
	"fgame/fgame/game/global"
	housetemplate "fgame/fgame/game/house/template"
	housetypes "fgame/fgame/game/house/types"
	"fgame/fgame/pkg/mathutils"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type HouseService interface {
	Heartbeat() //有定时任务，需要定义这个接口；
	//添加日志
	AddLog(playerName string, houseIndex, houseType, houseLevel int32, operateType housetypes.HouseOperateType)
	//获取日志
	GetLogByTime(time int64) []*HouseLogObject
	//清空日志
	GMClearLog()
}

type houseService struct {
	rwm sync.RWMutex
	//日志列表
	houseLogList []*HouseLogObject
	//上次系统插入日志时间
	lastAddDummyLogLime int64
}

//初始化
func (s *houseService) init() (err error) {
	return
}

//心跳
func (s *houseService) Heartbeat() {

	err := s.addDummyLog()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("house:系统生成假日志,错误")
		return
	}
}

//生成系统假日志
func (s *houseService) addDummyLog() (err error) {
	now := global.GetGame().GetTimeService().Now()
	diffTime := now - s.lastAddDummyLogLime
	randTime := s.getRandomLogTime()
	if diffTime < randTime {
		return
	}

	// 获取随机人物名称
	playerName := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	houseTemp := housetemplate.GetHouseTemplateService().GetRandomHouseTemp()
	if houseTemp == nil {
		return
	}

	var operateType housetypes.HouseOperateType
	if houseTemp.Level == housetypes.InitHouseLevel {
		// 等级为初始，只有激活和出售的情况
		operateType = housetypes.RandomOperateTypeExcludeUplevel()
	} else {
		// 等级不为初始，只有升级和出售的情况
		operateType = housetypes.RandomOperateTypeExcludeActivate()
	}

	s.AddLog(playerName, houseTemp.HouseIndex, int32(houseTemp.GetHouseType()), houseTemp.Level, operateType)
	s.lastAddDummyLogLime = now
	return
}

func (s *houseService) GetLogList() []*HouseLogObject {
	// 读锁
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.houseLogList
}

func (s *houseService) AddLog(playerName string, houseIndex, houseType, houseLevel int32, operateType housetypes.HouseOperateType) {
	// 写锁
	s.rwm.Lock()
	defer s.rwm.Unlock()

	obj := s.createLogObj(playerName, houseIndex, houseType, houseLevel, operateType)
	s.houseLogList = append(s.houseLogList, obj)
}

func (s *houseService) GetLogByTime(time int64) []*HouseLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	for index, log := range s.houseLogList {
		if time < log.updateTime {
			return s.houseLogList[index:]
		}
	}

	return nil
}

func (s *houseService) GMClearLog() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	var empty []*HouseLogObject
	s.houseLogList = empty
}

func (s *houseService) createLogObj(playerName string, houseIndex, houseType, houseLevel int32, operateType housetypes.HouseOperateType) *HouseLogObject {
	now := global.GetGame().GetTimeService().Now()
	maxLogLen := housetemplate.GetHouseTemplateService().GetHouseConstantTemplate().RiZhiCount
	var obj *HouseLogObject
	if len(s.houseLogList) >= int(maxLogLen) {
		obj = s.houseLogList[0]
		s.houseLogList = s.houseLogList[1:]
	} else {
		obj = NewHouseLogObject()
		obj.createTime = now
	}

	obj.playerName = playerName
	obj.houseIndex = houseIndex
	obj.houseType = houseType
	obj.houseLevel = houseLevel
	obj.operateType = int32(operateType)
	obj.updateTime = now
	return obj
}

//系统假日志生成间隔
func (s *houseService) getRandomLogTime() int64 {
	houseConstantTemp := housetemplate.GetHouseTemplateService().GetHouseConstantTemplate()
	min, max := houseConstantTemp.GetDummyLogTime()
	randTime := int64(mathutils.RandomRange(min, max)) //随机工具：范围随机，左闭区间右开区间
	return randTime
}

var (
	once sync.Once
	cs   *houseService
)

func Init() (err error) {
	once.Do(func() {
		cs = &houseService{}
		err = cs.init()
	})
	return err
}

func GetHouseService() HouseService {
	return cs
}
