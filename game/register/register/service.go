package register

import (
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/register/dao"
	"fgame/fgame/pkg/idutil"
	"sync"
)

type RegisterService interface {
	OpenRegister()
	CloseRegister()
	IsOpenRegister() bool
	Heartbeat()
}

type registerService struct {
	rwm sync.RWMutex
	//注册设置
	registerSettingObject *RegisterSettingObject
}

func (s *registerService) OpenRegister() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if s.registerSettingObject.open != 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	s.registerSettingObject.open = 1
	s.registerSettingObject.updateTime = now
	s.registerSettingObject.SetModified()
	//添加日志
	s.addLog(1)
	return
}

func (s *registerService) CloseRegister() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.registerSettingObject.open == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	s.registerSettingObject.open = 0
	s.registerSettingObject.updateTime = now
	s.registerSettingObject.SetModified()
	//添加日志
	s.addLog(0)
	return
}

func (s *registerService) IsOpenRegister() bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//开启gm的 就不关闭注册了
	if global.GetGame().GMOpen() {
		return true
	}
	return s.registerSettingObject.open != 0
}

const (
	closeTimes = 3 * common.DAY
)

func (s *registerService) Heartbeat() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.registerSettingObject.auto != 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if now < center.GetCenterService().GetStartTime()+int64(closeTimes) {
		return
	}
	//关闭
	s.registerSettingObject.open = 0
	s.registerSettingObject.auto = 1
	s.registerSettingObject.updateTime = now
	s.registerSettingObject.SetModified()
	//添加日志
	s.addLog(0)
}

func (s *registerService) addLog(open int32) {
	serverId := global.GetGame().GetServerIndex()
	now := global.GetGame().GetTimeService().Now()
	obj := createRegisterSettingLogObject()
	obj.id, _ = idutil.GetId()
	obj.serverId = serverId
	obj.open = open
	obj.createTime = now
	obj.SetModified()
	return
}

func (s *registerService) init() (err error) {
	serverId := global.GetGame().GetServerIndex()
	registerSettingEntity, err := dao.GetRegisterDao().GetRegisterSetting(serverId)
	if err != nil {
		return
	}

	if registerSettingEntity == nil {
		now := global.GetGame().GetTimeService().Now()
		obj := createRegisterSettingObject()
		obj.id, _ = idutil.GetId()
		obj.serverId = serverId
		obj.open = 1
		obj.auto = 0
		obj.createTime = now
		obj.SetModified()
		s.registerSettingObject = obj
		return
	}

	s.registerSettingObject = createRegisterSettingObject()
	err = s.registerSettingObject.FromEntity(registerSettingEntity)
	if err != nil {
		return
	}
	return nil
}

var (
	once sync.Once
	cs   *registerService
)

func Init() (err error) {
	once.Do(func() {
		cs = &registerService{}
		err = cs.init()
	})
	return err
}

func GetRegisterService() RegisterService {
	return cs
}
