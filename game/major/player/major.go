package player

import (
	"fgame/fgame/game/global"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"fgame/fgame/game/major/dao"
	majortemplate "fgame/fgame/game/major/template"
)

//玩家双修管理器
type PlayerMajorDataManager struct {
	p                 player.Player
	playerMajorObjMap map[majortypes.MajorType]*PlayerMajorNumObject //双修map
	inviteTimeMap     map[majortypes.MajorType]int64                 //邀请时间
}

func (m *PlayerMajorDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerMajorDataManager) Load() (err error) {
	//加载玩家双休数信息
	m.playerMajorObjMap = make(map[majortypes.MajorType]*PlayerMajorNumObject)
	entityList, err := dao.GetMajorDao().GetMajorNumEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := NewPlayerMajorNumObject(m.p)
		obj.FromEntity(entity)
		m.playerMajorObjMap[obj.MajorType] = obj
	}

	// 初始化所有副本数据
	for initType := majortypes.MinType; initType <= majortypes.MaxType; initType++ {
		if m.getMajorNumObject(initType) != nil {
			continue
		}

		m.playerMajorObjMap[initType] = m.initPlayerMajorNumObject(initType)
	}

	return nil
}

//第一次初始化
func (m *PlayerMajorDataManager) initPlayerMajorNumObject(majoyType majortypes.MajorType) *PlayerMajorNumObject {
	obj := NewPlayerMajorNumObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	obj.PlayerId = m.p.GetId()
	obj.MajorType = majoyType
	obj.Num = int32(0)
	obj.CreateTime = now
	obj.SetModified()
	return obj
}

//加载后
func (m *PlayerMajorDataManager) AfterLoad() (err error) {
	err = m.refresh()
	return err
}

func (m *PlayerMajorDataManager) refresh() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.playerMajorObjMap {
		if obj.LastTime != 0 {
			flag, err := timeutils.IsSameFive(obj.LastTime, now)
			if err != nil {
				return err
			}
			if !flag {
				obj.Num = 0
				obj.LastTime = 0
				obj.UpdateTime = now
				obj.SetModified()
			}
		}
	}

	return
}

func (m *PlayerMajorDataManager) HasMajorNum(majorType majortypes.MajorType) bool {
	obj := m.getMajorNumObject(majorType)
	if obj == nil {
		return false
	}

	m.refresh()

	defaultMaxNum := majortemplate.GetMajorTemplateService().GetMajorDefaultMaxNum(majorType)
	if obj.Num >= defaultMaxNum {
		return false
	}
	return true
}

func (m *PlayerMajorDataManager) HasMajorNumByNum(majorType majortypes.MajorType, num int32) bool {
	obj := m.getMajorNumObject(majorType)
	if obj == nil {
		return false
	}

	m.refresh()

	defaultMaxNum := majortemplate.GetMajorTemplateService().GetMajorDefaultMaxNum(majorType)
	if obj.Num+num > defaultMaxNum {
		return false
	}
	return true
}

func (m *PlayerMajorDataManager) UseMajorNum(majorType majortypes.MajorType) (num int32) {
	obj := m.getMajorNumObject(majorType)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	defaultMaxNum := majortemplate.GetMajorTemplateService().GetMajorDefaultMaxNum(majorType)
	if obj.Num < defaultMaxNum {
		obj.Num++
	}
	obj.UpdateTime = now
	obj.LastTime = now
	obj.SetModified()
	num = obj.Num
	return
}

func (m *PlayerMajorDataManager) UseMajorNumByNum(majorType majortypes.MajorType, useNum int32) (num int32) {
	obj := m.getMajorNumObject(majorType)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	defaultMaxNum := majortemplate.GetMajorTemplateService().GetMajorDefaultMaxNum(majorType)
	if obj.Num+useNum < defaultMaxNum {
		obj.Num += useNum
	} else {
		obj.Num = defaultMaxNum
	}
	obj.UpdateTime = now
	obj.LastTime = now
	obj.SetModified()
	num = obj.Num
	return
}

func (m *PlayerMajorDataManager) getMajorNumObject(majoyType majortypes.MajorType) *PlayerMajorNumObject {
	obj, ok := m.playerMajorObjMap[majoyType]
	if !ok {
		return nil
	}

	return obj
}

func (m *PlayerMajorDataManager) GetMajorNum(majorType majortypes.MajorType) int32 {
	obj := m.getMajorNumObject(majorType)
	if obj == nil {
		return 0
	}

	m.refresh()

	return obj.Num
}

func (m *PlayerMajorDataManager) GetAllMajorNumObj() map[majortypes.MajorType]*PlayerMajorNumObject {
	m.refresh()
	return m.playerMajorObjMap
}

//心跳
func (m *PlayerMajorDataManager) Heartbeat() {
}

//是否邀请过于频繁
func (m *PlayerMajorDataManager) InviteFrequent(majorType majortypes.MajorType) bool {
	inviteTime := m.inviteTimeMap[majorType]
	if inviteTime == 0 {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	cdTime := majortemplate.GetMajorTemplateService().GetInvitePairCdTime(majorType)
	if now-inviteTime < cdTime {
		return true
	}
	return false
}

func (m *PlayerMajorDataManager) InviteTime(majorType majortypes.MajorType) int64 {
	now := global.GetGame().GetTimeService().Now()
	m.inviteTimeMap[majorType] = now
	return now
}

//Gm使用
func (m *PlayerMajorDataManager) GmClearMajorNum() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.playerMajorObjMap {
		obj.Num = 0
		obj.UpdateTime = now
		obj.LastTime = now
		obj.SetModified()
	}
}

func CreatePlayerMajorDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerMajorDataManager{}
	m.p = p
	m.inviteTimeMap = make(map[majortypes.MajorType]int64)
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerMajorDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMajorDataManager))
}
