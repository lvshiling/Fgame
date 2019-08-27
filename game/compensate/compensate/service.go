package compensate

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/compensate/dao"
	compensateeventtypes "fgame/fgame/game/compensate/event/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/idutil"
	"sync"
)

// 全服补偿
type CompensateService interface {
	//获取补偿列表
	GetCompensateList() []*CompensateObject
	//添加补偿
	AddCompensate(needLevel int32, needCreateTime int64, title, content string, attachement []*droptemplate.DropItemData)
}

type compensateService struct {
	rwm            sync.RWMutex
	compensateList []*CompensateObject
}

const (
	maxExpireDay = 15
)

func (s *compensateService) init() (err error) {
	err = s.loadCompensate()
	if err != nil {
		return
	}

	s.checkExpireCompensate()
	return
}

func (s *compensateService) loadCompensate() (err error) {
	// 加载补偿数据
	entityList, err := dao.GetCompensateDao().GetCompensateEntityList()
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := NewCompensateObject()
		obj.FromEntity(entity)
		s.compensateList = append(s.compensateList, obj)
	}

	return
}

func (s *compensateService) checkExpireCompensate() {
	now := global.GetGame().GetTimeService().Now()
	remainCompensateList := make([]*CompensateObject, 0, 1)
	for _, obj := range s.compensateList {
		expireTime := obj.createTime + int64(maxExpireDay)*int64(common.DAY)
		if now > expireTime {
			continue
		}
		remainCompensateList = append(remainCompensateList, obj)
	}
	s.compensateList = remainCompensateList
}

func (s *compensateService) GetCompensateList() []*CompensateObject {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.checkExpireCompensate()

	return s.compensateList
}

func (s *compensateService) AddCompensate(needLevel int32, needCreateTime int64, title, content string, attachement []*droptemplate.DropItemData) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	newObj := NewCompensateObject()
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	newObj.id = id
	newObj.title = title
	newObj.serverId = global.GetGame().GetServerIndex()
	newObj.content = content
	newObj.attachment = attachement
	newObj.roleLevel = needLevel
	newObj.roleCreateTime = needCreateTime
	newObj.createTime = now
	newObj.SetModified()
	s.compensateList = append(s.compensateList, newObj)

	gameevent.Emit(compensateeventtypes.EventTypeServerCompensateChanged, newObj, nil)
	return
}

var (
	once sync.Once
	cs   *compensateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &compensateService{}
		err = cs.init()
	})

	return
}

func GetCompensateSrevice() CompensateService {
	return cs
}
