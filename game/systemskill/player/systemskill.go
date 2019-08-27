package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	sysskillcommmon "fgame/fgame/game/systemskill/common"
	"fgame/fgame/game/systemskill/dao"
	systemskilleventtypes "fgame/fgame/game/systemskill/event/types"
	sysskilltemplate "fgame/fgame/game/systemskill/template"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	"fgame/fgame/pkg/idutil"
)

//玩家系统技能管理器
type PlayerSystemSkillDataManager struct {
	p player.Player
	//玩家系统技能map
	playerSystemSkillObjectMap map[sysskilltypes.SystemSkillType]*SystemSkill
}

func (pxfdm *PlayerSystemSkillDataManager) Player() player.Player {
	return pxfdm.p
}

//加载
func (pxfdm *PlayerSystemSkillDataManager) Load() (err error) {
	pxfdm.playerSystemSkillObjectMap = make(map[sysskilltypes.SystemSkillType]*SystemSkill)
	//加载玩家系统技能
	systemSkills, err := dao.GetSystemSkillDao().GetSystemSkillList(pxfdm.p.GetId())
	if err != nil {
		return
	}
	//系统技能信息
	for _, systemSkill := range systemSkills {
		pxfo := NewPlayerSystemSkillObject(pxfdm.p)
		pxfo.FromEntity(systemSkill)
		pxfdm.addSystemSkill(pxfo)
	}
	return nil
}

func (pxfdm *PlayerSystemSkillDataManager) addSystemSkill(obj *PlayerSystemSkillObject) {
	typ := obj.Type

	systemSkill, exist := pxfdm.playerSystemSkillObjectMap[typ]
	if !exist {
		systemSkill = createSystemSkill(pxfdm.p, obj)
		pxfdm.playerSystemSkillObjectMap[typ] = systemSkill
		return
	}
	systemSkill.addSystemSkill(obj)

}

//加载后
func (pxfdm *PlayerSystemSkillDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (pxfdm *PlayerSystemSkillDataManager) Heartbeat() {

}

func (pxfdm *PlayerSystemSkillDataManager) GetSystemSkillMap(typ sysskilltypes.SystemSkillType) *SystemSkill {
	if !typ.Valid() {
		return nil
	}
	systemSkill, exist := pxfdm.playerSystemSkillObjectMap[typ]
	if !exist {
		return nil
	}
	return systemSkill
}

//获取玩家系统技能map
func (pxfdm *PlayerSystemSkillDataManager) GetSystemSkillAllMap() map[sysskilltypes.SystemSkillType]*SystemSkill {
	return pxfdm.playerSystemSkillObjectMap
}

//获取系统技能通过系统技能类型
func (pxfdm *PlayerSystemSkillDataManager) GetSystemSkillByTyp(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) *PlayerSystemSkillObject {
	systemSkill, exist := pxfdm.playerSystemSkillObjectMap[typ]
	if !exist {
		return nil
	}
	return systemSkill.GetSystemSkillObjct(subType)
}

//获取等级通过系统技能类型
func (pxfdm *PlayerSystemSkillDataManager) GetSystemSkillLevelByTyp(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) int32 {
	obj := pxfdm.GetSystemSkillByTyp(typ, subType)
	if obj == nil {
		return 0
	}
	return obj.Level
}

//是否已激活
func (pxfdm *PlayerSystemSkillDataManager) IfSystemSkillExist(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) bool {
	obj := pxfdm.GetSystemSkillByTyp(typ, subType)
	if obj == nil {
		return false
	}
	return true
}

//是否达到满级
func (pxfdm *PlayerSystemSkillDataManager) ifFullLevel(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) bool {
	level := pxfdm.GetSystemSkillLevelByTyp(typ, subType)
	to := sysskilltemplate.GetSystemSkillTemplateService().GetSystemSkillTemplateByTypeAndLevel(typ, subType, level)
	if to.GetNextId() == 0 {
		return true
	}
	return false
}

//能否升级
func (pxfdm *PlayerSystemSkillDataManager) IfCanUpgrade(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) bool {
	flag := pxfdm.IfSystemSkillExist(typ, subType)
	if !flag {
		return false
	}
	flag = pxfdm.ifFullLevel(typ, subType)
	if flag {
		return false
	}
	return true
}

//系统技能激活
func (pxfdm *PlayerSystemSkillDataManager) SystemSkillActive(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) (*PlayerSystemSkillObject, bool) {
	flag := typ.Valid()
	if !flag {
		return nil, false
	}
	flag = subType.Valid()
	if !flag {
		return nil, false
	}
	flag = pxfdm.IfSystemSkillExist(typ, subType)
	if flag {
		return nil, false
	}

	skTemplate := sysskilltemplate.GetSystemSkillTemplateService().GetSystemSkillTemplateByTypeAndLevel(typ, subType, 1)
	if skTemplate == nil {
		return nil, false
	}
	id, _ := idutil.GetId()

	now := global.GetGame().GetTimeService().Now()
	pxfo := NewPlayerSystemSkillObject(pxfdm.p)
	pxfo.Id = id
	pxfo.PlayerId = pxfdm.p.GetId()
	pxfo.Type = typ
	pxfo.SubType = subType
	pxfo.Level = skTemplate.GetLevel()
	pxfo.CreateTime = now
	pxfo.SetModified()

	pxfdm.addSystemSkill(pxfo)
	eventData := systemskilleventtypes.CreateSystemSkillEventData(typ, subType)
	gameevent.Emit(systemskilleventtypes.EventTypeSystemSkillActive, pxfdm.p, eventData)
	return pxfo, true
}

//系统技能升级
func (pxfdm *PlayerSystemSkillDataManager) Upgrade(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType) (*PlayerSystemSkillObject, bool) {
	flag := pxfdm.IfCanUpgrade(typ, subType)
	if !flag {
		return nil, false
	}

	obj := pxfdm.GetSystemSkillByTyp(typ, subType)
	if obj == nil {
		return nil, false
	}
	now := global.GetGame().GetTimeService().Now()
	obj.UpdateTime = now
	obj.Level += 1
	obj.SetModified()
	eventData := systemskilleventtypes.CreateSystemSkillEventData(typ, subType)
	gameevent.Emit(systemskilleventtypes.EventTypeSystemSkillUpgrade, pxfdm.p, eventData)
	return obj, true
}

//系统技能信息
func (m *PlayerSystemSkillDataManager) ToAllSystemSkillInfo() (allSystemInfo *sysskillcommmon.AllSystemSkillInfo) {
	allSystemInfo = &sysskillcommmon.AllSystemSkillInfo{}
	for typ, systemSkillObj := range m.playerSystemSkillObjectMap {
		systemSkillInfo := &sysskillcommmon.SystemSkillInfo{
			Type: int32(typ),
		}
		for subType, systemSkillSubObj := range systemSkillObj.GetSysSkillMap() {
			systemSkillSubTypeInfo := &sysskillcommmon.SystemSkillSubTypeInfo{
				SubType: int32(subType),
				Level:   systemSkillSubObj.Level,
			}
			systemSkillInfo.SysSkillList = append(systemSkillInfo.SysSkillList, systemSkillSubTypeInfo)
		}
		allSystemInfo.SystemSkillList = append(allSystemInfo.SystemSkillList, systemSkillInfo)
	}
	return
}

func CreatePlayerSystemSkillDataManager(p player.Player) player.PlayerDataManager {
	pxfdm := &PlayerSystemSkillDataManager{}
	pxfdm.p = p
	return pxfdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSystemSkillDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSystemSkillDataManager))
}
