package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/material/dao"
	materialeventtypes "fgame/fgame/game/material/event/types"
	materialtemplate "fgame/fgame/game/material/template"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

//玩家材料副本数据管理器
type PlayerMaterialDataManager struct {
	pl                player.Player
	materialObjectMap map[materialtypes.MaterialType]*PlayerMaterialObject
}

func (m *PlayerMaterialDataManager) Player() player.Player {
	return m.pl
}

//加载材料副本信息
func (m *PlayerMaterialDataManager) Load() (err error) {
	err = m.loadMaterialObject()
	if err != nil {
		return
	}

	m.initMaterialObject()
	return nil
}

//加载玩家材料副本数据
func (m *PlayerMaterialDataManager) loadMaterialObject() error {
	entityList, err := dao.GetMaterialDao().GetMaterialInfo(m.pl.GetId())
	if err != nil {
		return err
	}
	for _, entity := range entityList {
		newObj := CreateNewPlayerMaterialObject(m.pl)
		newObj.FromEntity(entity)
		m.materialObjectMap[newObj.materialType] = newObj
	}

	return nil
}

//生成玩家初始材料副本数据
func (m *PlayerMaterialDataManager) initMaterialObject() {
	for typ := materialtypes.MinMaterialType; typ <= materialtypes.MaxMaterialType; typ++ {
		if _, ok := m.materialObjectMap[typ]; !ok {
			now := global.GetGame().GetTimeService().Now()
			id, _ := idutil.GetId()
			newObj := CreateNewPlayerMaterialObject(m.pl)
			newObj.id = id
			newObj.materialType = typ
			newObj.useTimes = 0
			newObj.group = 0
			newObj.createTime = now
			m.materialObjectMap[typ] = newObj

			newObj.SetModified()
		}
	}
}

//加载后处理
func (m *PlayerMaterialDataManager) AfterLoad() (err error) {
	m.RefreshData()
	return
}

//心跳
func (m *PlayerMaterialDataManager) Heartbeat() {
}

//刷新材料副本数据
func (m *PlayerMaterialDataManager) RefreshData() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.materialObjectMap {
		obj.refreshUseTimes(now)
	}
	return
}

//获取所有材料副本信息
func (m *PlayerMaterialDataManager) GetPlayerMaterialInfoList() map[materialtypes.MaterialType]*PlayerMaterialObject {
	return m.materialObjectMap
}

//获取剩余次数
func (m *PlayerMaterialDataManager) IsEnoughAttendTimes(typ materialtypes.MaterialType, attendNum int32) bool {
	obj := m.GetPlayerMaterialInfo(typ)
	if obj == nil {
		return false
	}

	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(typ)
	maxTimes := materialTemplate.AllTimes
	return maxTimes >= obj.useTimes+attendNum
}

//获取免费挑战次数
func (m *PlayerMaterialDataManager) IsFreeTimes(typ materialtypes.MaterialType) bool {
	// obj := m.GetPlayerMaterialInfo(typ)
	// if obj == nil {
	// 	return false
	// }

	// materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(typ)
	// maxFree := materialTemplate.Free
	// if obj.useTimes >= maxFree {
	// 	return false
	// }
	freeTims := m.GetFreeAttendTimes(typ)
	return freeTims > 0
}

//获取免费次数
func (m *PlayerMaterialDataManager) GetFreeAttendTimes(typ materialtypes.MaterialType) int32 {
	obj := m.GetPlayerMaterialInfo(typ)
	if obj == nil {
		return 0
	}

	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(typ)
	maxFree := materialTemplate.Free
	if obj.useTimes >= maxFree {
		return 0
	}
	return maxFree - obj.useTimes
}

//完成扫荡或挑战
func (m *PlayerMaterialDataManager) UseTimes(typ materialtypes.MaterialType, num int32) {
	obj := m.GetPlayerMaterialInfo(typ)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.useTimes += num
	obj.updateTime = now
	obj.SetModified()

	m.EmitJoinEvent(typ, num)
	return
}

//发送挑战事件
func (m *PlayerMaterialDataManager) EmitJoinEvent(typ materialtypes.MaterialType, num int32) {
	eventData := materialeventtypes.CreateMaterialChallengeEventData(typ, num)
	gameevent.Emit(materialeventtypes.EventTypeMaterialChallenge, m.pl, eventData)
}

//发送完成事件
func (m *PlayerMaterialDataManager) EmitFinishEvent(typ materialtypes.MaterialType, num int32) {
	eventData := materialeventtypes.CreateMaterialFinishEventData(typ, num)
	gameevent.Emit(materialeventtypes.EventTypeMaterialFinish, m.pl, eventData)
}

// //发送扫荡事件
// func (m *PlayerMaterialDataManager) EmitSweepEvent(typ materialtypes.MaterialType, num int32) {
// 	eventData := materialeventtypes.CreateXianFuChallengeEventData(typ, num)
// 	gameevent.Emit(materialeventtypes.EventTypeXianFuSweep, m.pl, eventData)
// }

//刷新波数
func (m *PlayerMaterialDataManager) RefreshGroup(group int32, typ materialtypes.MaterialType) {
	obj := m.GetPlayerMaterialInfo(typ)
	if obj.group >= group {
		return
	}
	nowTime := global.GetGame().GetTimeService().Now()
	obj.group = group
	obj.updateTime = nowTime
	obj.SetModified()
	return
}

//获取材料副本
func (m *PlayerMaterialDataManager) GetPlayerMaterialInfo(typ materialtypes.MaterialType) (obj *PlayerMaterialObject) {
	return m.materialObjectMap[typ]
}

//获取材料副本所有剩余免费次数
func (m *PlayerMaterialDataManager) GetAllLeftTimes() (leftNum int32) {
	for typ := materialtypes.MinMaterialType; typ <= materialtypes.MaxMaterialType; typ++ {
		funcOpenType := typ.GetFuncOpenType()
		if !m.pl.IsFuncOpen(funcOpenType) {
			continue
		}
		obj := m.GetPlayerMaterialInfo(typ)
		if obj == nil {
			continue
		}

		materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(typ)
		maxFree := materialTemplate.Free
		if (maxFree - obj.useTimes) > 0 {
			leftNum += maxFree - obj.useTimes
		}
	}
	return
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerMaterialDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMaterialDataManager))
}

func CreatePlayerMaterialDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerMaterialDataManager{}
	m.pl = p
	m.materialObjectMap = make(map[materialtypes.MaterialType]*PlayerMaterialObject)

	return m
}
