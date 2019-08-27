package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/shenqi/dao"
	shenqieventtypes "fgame/fgame/game/shenqi/event/types"
	shenqitemplate "fgame/fgame/game/shenqi/template"
	shenqitypes "fgame/fgame/game/shenqi/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
)

//玩家神器管理器
type PlayerShenQiDataManager struct {
	p player.Player
	//玩家神器碎片
	shenQiDebrisMap map[shenqitypes.ShenQiType]map[shenqitypes.DebrisType]*PlayerShenQiDebrisObject
	//玩家淬炼
	shenQiSmeltMap map[shenqitypes.ShenQiType]map[shenqitypes.SmeltType]*PlayerShenQiSmeltObject
	//玩家器灵
	shenQiQiLingMap map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*PlayerShenQiQiLingObject
	//玩家神器
	shenQiObject *PlayerShenQiObject
	//玩家神器碎片最小等级
	shenQiDebrisMinLevelMap map[shenqitypes.ShenQiType]int32
	//玩家神器淬炼最小等级
	shenQiSmeltMinLevelMap map[shenqitypes.ShenQiType]int32
	//玩家神器器灵最小套装id
	shenQiQiLingTaoZhuangMap map[shenqitypes.ShenQiType]int
}

func (m *PlayerShenQiDataManager) Player() player.Player {
	return m.p
}

func (m *PlayerShenQiDataManager) GetAllShenQiLevel() int32 {
	totalLevel := int32(0)
	for _, level := range m.shenQiDebrisMinLevelMap {
		totalLevel += level
	}
	return totalLevel
}

//加载
func (m *PlayerShenQiDataManager) Load() (err error) {
	err = m.loadShenQi()
	if err != nil {
		return
	}

	err = m.loadShenQiDebris()
	if err != nil {
		return
	}

	err = m.loadShenQiSmelt()
	if err != nil {
		return
	}

	err = m.loadShenQiQiLing()
	if err != nil {
		return
	}

	return nil
}

func (m *PlayerShenQiDataManager) loadShenQi() (err error) {
	//加载玩家神器信息
	shenQiEntity, err := dao.GetShenQiDao().GetShenQiEntity(m.p.GetId())
	if err != nil {
		return
	}
	if shenQiEntity == nil {
		m.initPlayerShenQiObject()
	} else {
		m.shenQiObject = NewPlayerShenQiObject(m.p)
		m.shenQiObject.FromEntity(shenQiEntity)
	}
	m.shenQiDebrisMinLevelMap = make(map[shenqitypes.ShenQiType]int32)
	m.shenQiSmeltMinLevelMap = make(map[shenqitypes.ShenQiType]int32)
	m.shenQiQiLingTaoZhuangMap = make(map[shenqitypes.ShenQiType]int)
	return nil
}

func (m *PlayerShenQiDataManager) loadShenQiDebris() (err error) {
	m.shenQiDebrisMap = make(map[shenqitypes.ShenQiType]map[shenqitypes.DebrisType]*PlayerShenQiDebrisObject)
	debrisList, err := dao.GetShenQiDao().GetShenQiDebrisList(m.p.GetId())
	if err != nil {
		return
	}

	for _, info := range debrisList {
		obj := NewPlayerShenQiDebrisObject(m.p)
		obj.FromEntity(info)
		m.addShenQiDebris(obj)
	}

	for typ := shenqitypes.MinShenQiType; typ <= shenqitypes.MaxShenQiType; typ++ {
		_, ok := m.shenQiDebrisMap[typ]
		if !ok {
			continue
		}
		m.RefreshShenQiDebrisMinLevel(typ)
	}

	return
}

func (m *PlayerShenQiDataManager) loadShenQiSmelt() (err error) {
	m.shenQiSmeltMap = make(map[shenqitypes.ShenQiType]map[shenqitypes.SmeltType]*PlayerShenQiSmeltObject)
	smeltList, err := dao.GetShenQiDao().GetShenQiSmeltList(m.p.GetId())
	if err != nil {
		return
	}

	for _, info := range smeltList {
		obj := NewPlayerShenQiSmeltObject(m.p)
		obj.FromEntity(info)
		m.addShenQiSmelt(obj)
	}

	for typ := shenqitypes.MinShenQiType; typ <= shenqitypes.MaxShenQiType; typ++ {
		_, ok := m.shenQiSmeltMap[typ]
		if !ok {
			continue
		}
		m.RefreshShenQiSmeltMinLevel(typ)
	}

	return
}

func (m *PlayerShenQiDataManager) loadShenQiQiLing() (err error) {
	m.shenQiQiLingMap = make(map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*PlayerShenQiQiLingObject)
	qiLingList, err := dao.GetShenQiDao().GetShenQiQiLingList(m.p.GetId())
	if err != nil {
		return
	}

	for _, info := range qiLingList {
		obj := NewPlayerShenQiQiLingObject(m.p)
		obj.FromEntity(info)
		m.addShenQiQiLing(obj)
	}

	for typ := shenqitypes.MinShenQiType; typ <= shenqitypes.MaxShenQiType; typ++ {
		_, ok := m.shenQiQiLingMap[typ]
		if !ok {
			continue
		}
		m.RefreshShenQiQiLingTaoZhuang(typ)
	}

	return
}

func (m *PlayerShenQiDataManager) initShenQiDebrisObjsMap(typ shenqitypes.ShenQiType) {
	for slotId := shenqitypes.MinDebrisType; slotId <= shenqitypes.MaxDebrisType; slotId++ {
		obj, ok := m.shenQiDebrisMap[typ][slotId]
		if ok {
			continue
		}
		now := global.GetGame().GetTimeService().Now()
		obj = createShenQiDebrisObject(m.p, typ, slotId, now)
		obj.SetModified()
		m.addShenQiDebris(obj)
	}
	m.RefreshShenQiDebrisMinLevel(typ)
	return
}

func (m *PlayerShenQiDataManager) initShenQiSmeltObjsMap(typ shenqitypes.ShenQiType) {
	for slotId := shenqitypes.MinSmeltType; slotId <= shenqitypes.MaxSmeltType; slotId++ {
		obj, ok := m.shenQiSmeltMap[typ][slotId]
		if ok {
			continue
		}
		now := global.GetGame().GetTimeService().Now()
		obj = createShenQiSmeltObject(m.p, typ, slotId, now)
		obj.SetModified()
		m.addShenQiSmelt(obj)
	}
	m.RefreshShenQiSmeltMinLevel(typ)
	return
}

func (m *PlayerShenQiDataManager) initShenQiQiLingObjsMap(typ shenqitypes.ShenQiType) {
	for subType, tempM := range shenqitypes.QiLingSubTypeStringMap {
		for slotId, _ := range tempM {
			obj, ok := m.shenQiQiLingMap[typ][subType][slotId]
			if ok {
				continue
			}
			now := global.GetGame().GetTimeService().Now()
			obj = createShenQiQiLingObject(m.p, typ, subType, slotId, now)
			obj.SetModified()
			m.addShenQiQiLing(obj)
		}
	}
	m.RefreshShenQiQiLingTaoZhuang(typ)
	return
}

func (m *PlayerShenQiDataManager) addShenQiDebris(obj *PlayerShenQiDebrisObject) {
	ojbM, ok := m.shenQiDebrisMap[obj.ShenQiType]
	if !ok {
		ojbM = make(map[shenqitypes.DebrisType]*PlayerShenQiDebrisObject)
		m.shenQiDebrisMap[obj.ShenQiType] = ojbM
	}
	ojbM[obj.SlotId] = obj
}

func (m *PlayerShenQiDataManager) addShenQiSmelt(obj *PlayerShenQiSmeltObject) {
	ojbM, ok := m.shenQiSmeltMap[obj.ShenQiType]
	if !ok {
		ojbM = make(map[shenqitypes.SmeltType]*PlayerShenQiSmeltObject)
		m.shenQiSmeltMap[obj.ShenQiType] = ojbM
	}
	ojbM[obj.SlotId] = obj
}

func (m *PlayerShenQiDataManager) addShenQiQiLing(obj *PlayerShenQiQiLingObject) {
	ojbMM, ok := m.shenQiQiLingMap[obj.ShenQiType]
	if !ok {
		ojbMM = make(map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*PlayerShenQiQiLingObject)
		m.shenQiQiLingMap[obj.ShenQiType] = ojbMM
	}
	ojbM, ok := ojbMM[obj.QiLingType]
	if !ok {
		ojbM = make(map[shenqitypes.QiLingSubType]*PlayerShenQiQiLingObject)
		ojbMM[obj.QiLingType] = ojbM
	}
	ojbM[obj.SlotId] = obj
}

func (m *PlayerShenQiDataManager) initPlayerShenQiObject() {
	obj := NewPlayerShenQiObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	obj.LingQiNum = 0
	obj.CreateTime = now
	obj.SetModified()

	m.shenQiObject = obj
}

//加载后
func (m *PlayerShenQiDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerShenQiDataManager) Heartbeat() {

}

//获取神器碎片信息
func (m *PlayerShenQiDataManager) GetShenQiDebrisMap() map[shenqitypes.ShenQiType]map[shenqitypes.DebrisType]*PlayerShenQiDebrisObject {
	return m.shenQiDebrisMap
}

//获取神器淬炼
func (m *PlayerShenQiDataManager) GetShenQiSmeltMap() map[shenqitypes.ShenQiType]map[shenqitypes.SmeltType]*PlayerShenQiSmeltObject {
	return m.shenQiSmeltMap
}

//获取神器器灵
func (m *PlayerShenQiDataManager) GetShenQiQiLingMap() map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*PlayerShenQiQiLingObject {
	return m.shenQiQiLingMap
}

//获取神器
func (m *PlayerShenQiDataManager) GetShenQiOjb() *PlayerShenQiObject {
	return m.shenQiObject
}

//根据神器获取神器碎片信息
func (m *PlayerShenQiDataManager) GetShenQiDebrisMapByShenQi(typ shenqitypes.ShenQiType) map[shenqitypes.DebrisType]*PlayerShenQiDebrisObject {
	return m.shenQiDebrisMap[typ]
}

//根据神器获取神器淬炼
func (m *PlayerShenQiDataManager) GetShenQiSmeltMapByShenQi(typ shenqitypes.ShenQiType) map[shenqitypes.SmeltType]*PlayerShenQiSmeltObject {
	return m.shenQiSmeltMap[typ]
}

//根据神器获取神器器灵
func (m *PlayerShenQiDataManager) GetShenQiQiLingMapByShenQi(typ shenqitypes.ShenQiType) map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*PlayerShenQiQiLingObject {
	return m.shenQiQiLingMap[typ]
}

//根据类型参数获取神器碎片信息
func (m *PlayerShenQiDataManager) GetShenQiDebrisMapByArg(typ shenqitypes.ShenQiType, subType shenqitypes.DebrisType) *PlayerShenQiDebrisObject {
	obj, ok := m.shenQiDebrisMap[typ][subType]
	if !ok {
		return nil
	}
	return obj
}

//根据类型参数获取神器淬炼
func (m *PlayerShenQiDataManager) GetShenQiSmeltMapByArg(typ shenqitypes.ShenQiType, subType shenqitypes.SmeltType) *PlayerShenQiSmeltObject {
	obj, ok := m.shenQiSmeltMap[typ][subType]
	if !ok {
		return nil
	}
	return obj
}

//根据类型参数获取神器器灵
func (m *PlayerShenQiDataManager) GetShenQiQiLingMapByArg(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType) *PlayerShenQiQiLingObject {
	obj, ok := m.shenQiQiLingMap[typ][subType][pos]
	if !ok {
		return nil
	}
	return obj
}

//根据类型参数获取神器器灵(不存在就初始化)
func (m *PlayerShenQiDataManager) GetShenQiQiLingOrInitByArg(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType) *PlayerShenQiQiLingObject {
	obj, ok := m.shenQiQiLingMap[typ][subType][pos]
	if !ok {
		m.initShenQiQiLingObjsMap(typ)
		obj = m.shenQiQiLingMap[typ][subType][pos]
	}
	return obj
}

//获取神器碎片最小等级
func (m *PlayerShenQiDataManager) GetShenQiDebrisMinLevelByShenQi(typ shenqitypes.ShenQiType) int32 {
	return m.shenQiDebrisMinLevelMap[typ]
}

//获取神器淬炼最小等级
func (m *PlayerShenQiDataManager) GetShenQiSmeltMinLevelByShenQi(typ shenqitypes.ShenQiType) int32 {
	return m.shenQiSmeltMinLevelMap[typ]
}

//获取神器淬炼最小等级
func (m *PlayerShenQiDataManager) GetShenQiQiLingTaoZhuangIdByShenQi(typ shenqitypes.ShenQiType) int {
	return m.shenQiQiLingTaoZhuangMap[typ]
}

//获取下一个碎片等级配置
func (m *PlayerShenQiDataManager) GetNextDebrisUpTemplate(typ shenqitypes.ShenQiType, pos shenqitypes.DebrisType) *gametemplate.ShenQiLevelTemplate {
	slot, ok := m.shenQiDebrisMap[typ][pos]
	if !ok {
		m.initShenQiDebrisObjsMap(typ)
		slot = m.shenQiDebrisMap[typ][pos]
	}
	var nextTemplate *gametemplate.ShenQiLevelTemplate
	if slot.Level == 0 {
		nextTemplate = shenqitemplate.GetShenQiTemplateService().GetShenQiDebrisUpByArg(slot.ShenQiType, slot.SlotId, 1)
	} else {
		//判断槽位是否可以升级
		temp := shenqitemplate.GetShenQiTemplateService().GetShenQiDebrisUpByArg(slot.ShenQiType, slot.SlotId, slot.Level)
		nextTemplate = temp.GetNextTemplate()
	}
	return nextTemplate
}

//获取下一个碎片等级配置
func (m *PlayerShenQiDataManager) GetNextSmeltUpTemplate(typ shenqitypes.ShenQiType, pos shenqitypes.SmeltType) *gametemplate.ShenQiCuiLianLevelTemplate {
	slot, ok := m.shenQiSmeltMap[typ][pos]
	if !ok {
		m.initShenQiSmeltObjsMap(typ)
		slot = m.shenQiSmeltMap[typ][pos]
	}
	var nextTemplate *gametemplate.ShenQiCuiLianLevelTemplate
	if slot.Level == 0 {
		nextTemplate = shenqitemplate.GetShenQiTemplateService().GetShenQiSmeltUpByArg(slot.ShenQiType, slot.SlotId, 1)
	} else {
		//判断槽位是否可以升级
		temp := shenqitemplate.GetShenQiTemplateService().GetShenQiSmeltUpByArg(slot.ShenQiType, slot.SlotId, slot.Level)
		nextTemplate = temp.GetNextTemplate()
	}
	return nextTemplate
}

//获取下一个注灵配置
func (m *PlayerShenQiDataManager) GetNextZhuLingTemplate(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType) *gametemplate.ShenQiZhuLingTemplate {
	slot, ok := m.shenQiQiLingMap[typ][subType][pos]
	if !ok {
		m.initShenQiQiLingObjsMap(typ)
		slot = m.shenQiQiLingMap[typ][subType][pos]
	}
	var nextTemplate *gametemplate.ShenQiZhuLingTemplate
	if slot.Level == 0 {
		nextTemplate = shenqitemplate.GetShenQiTemplateService().GetShenQiZhuLingByArg(slot.ShenQiType, slot.QiLingType, slot.SlotId, 1)
	} else {
		//判断槽位是否可以注灵
		temp := shenqitemplate.GetShenQiTemplateService().GetShenQiZhuLingByArg(slot.ShenQiType, slot.QiLingType, slot.SlotId, slot.Level)
		nextTemplate = temp.GetNextTemplate()
	}
	return nextTemplate
}

//穿上
func (m *PlayerShenQiDataManager) QiLingPutOn(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType, itemId int32, bind itemtypes.ItemBindType) bool {
	slot, ok := m.shenQiQiLingMap[typ][subType][pos]
	if !ok {
		m.initShenQiQiLingObjsMap(typ)
		slot = m.shenQiQiLingMap[typ][subType][pos]
	}

	if slot.IsEmpty() {
		now := global.GetGame().GetTimeService().Now()
		slot.ItemId = itemId
		slot.BindType = bind
		slot.UpdateTime = now
		slot.SetModified()
		return true
	}

	return false
}

//脱下
func (m *PlayerShenQiDataManager) QiLingTakeOff(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType) (itemId int32) {
	slot, ok := m.shenQiQiLingMap[typ][subType][pos]
	if !ok {
		return
	}
	if slot.IsEmpty() {
		return
	}
	itemId = slot.ItemId
	defaultInitBind := itemtypes.ItemBindTypeUnBind
	now := global.GetGame().GetTimeService().Now()
	slot.ItemId = 0
	slot.BindType = defaultInitBind
	slot.UpdateTime = now
	slot.SetModified()
	return
}

//增加灵气值
func (m *PlayerShenQiDataManager) AddLingQiNum(val int64) (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	m.shenQiObject.LingQiNum += val
	m.shenQiObject.UpdateTime = now
	m.shenQiObject.SetModified()

	gameevent.Emit(shenqieventtypes.EventTypeShenQiLingQiNumChanged, m.p, m.shenQiObject)

	flag = true
	return
}

//减少灵气值
func (m *PlayerShenQiDataManager) SubLingQiNum(val int64) (flag bool) {
	if val > m.shenQiObject.LingQiNum {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.shenQiObject.LingQiNum -= val
	m.shenQiObject.UpdateTime = now
	m.shenQiObject.SetModified()

	gameevent.Emit(shenqieventtypes.EventTypeShenQiLingQiNumChanged, m.p, m.shenQiObject)

	flag = true
	return
}

//刷新神器等级
func (m *PlayerShenQiDataManager) RefreshShenQiDebrisMinLevel(typ shenqitypes.ShenQiType) (lev int32) {
	objMap, ok := m.shenQiDebrisMap[typ]
	if !ok {
		return
	}
	lev = m.shenQiDebrisMap[typ][shenqitypes.DebrisTypeOne].Level
	for _, obj := range objMap {
		if lev > obj.Level {
			lev = obj.Level
		}
	}
	oldLevel := m.shenQiDebrisMinLevelMap[typ]
	if oldLevel != lev {
		m.shenQiDebrisMinLevelMap[typ] = lev
	}

	return
}

//刷新神器淬炼等级
func (m *PlayerShenQiDataManager) RefreshShenQiSmeltMinLevel(typ shenqitypes.ShenQiType) (lev int32) {
	objMap, ok := m.shenQiSmeltMap[typ]
	if !ok {
		return
	}
	lev = m.shenQiSmeltMap[typ][shenqitypes.SmeltTypeLight].Level
	for _, obj := range objMap {
		if lev > obj.Level {
			lev = obj.Level
		}
	}
	oldLevel := m.shenQiSmeltMinLevelMap[typ]
	if oldLevel != lev {
		m.shenQiSmeltMinLevelMap[typ] = lev
	}
	return
}

//刷新神器器灵套装
func (m *PlayerShenQiDataManager) RefreshShenQiQiLingTaoZhuang(typ shenqitypes.ShenQiType) (id int) {
	objMM, ok := m.shenQiQiLingMap[typ]
	if !ok {
		m.shenQiQiLingTaoZhuangMap[typ] = 0
		return
	}
	//收集穿戴
	tempMap := make(map[int32]int32)
	for _, objM := range objMM {
		for _, obj := range objM {
			if obj.IsEmpty() {
				continue
			}
			itemTemplate := item.GetItemService().GetItem(int(obj.ItemId))
			if itemTemplate == nil {
				m.shenQiQiLingTaoZhuangMap[typ] = 0
				return
			}
			tempMap[itemTemplate.TypeFlag2] += 1
			nearNumber := int32(0)
			for number := range tempMap {
				if number < itemTemplate.TypeFlag2 {
					tempMap[number] += 1
				}
				if number > itemTemplate.TypeFlag2 && (nearNumber == 0 || nearNumber > number) {
					nearNumber = number
				}
			}
			if nearNumber != 0 {
				tempMap[itemTemplate.TypeFlag2] += tempMap[nearNumber]
			}
		}
	}
	//判断是套装几
	templateMap := shenqitemplate.GetShenQiTemplateService().GetShenQiTaoZhuangMap()
	for idx, temp := range templateMap {
		for number, count := range tempMap {
			if temp.NeedNumber <= number && temp.NeedCount <= count && id < idx {
				id = idx
			}
		}
	}

	if m.shenQiQiLingTaoZhuangMap[typ] != id {
		m.shenQiQiLingTaoZhuangMap[typ] = id
	}
	return
}

//神器碎片升级
func (m *PlayerShenQiDataManager) DebrisUpAdvanced(typ shenqitypes.ShenQiType, pos shenqitypes.DebrisType, pro int32, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	obj, ok := m.shenQiDebrisMap[typ][pos]
	if !ok {
		m.initShenQiDebrisObjsMap(typ)
		obj = m.shenQiDebrisMap[typ][pos]
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextTemplate := m.GetNextDebrisUpTemplate(typ, pos)
		if nextTemplate == nil {
			return
		}
		obj.Level = nextTemplate.Level
		obj.UpPro = 0
		obj.UpNum = 0
	} else {
		obj.UpNum += addTimes
		obj.UpPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return
}

//神器淬炼升级
func (m *PlayerShenQiDataManager) SmeltUpAdvanced(typ shenqitypes.ShenQiType, pos shenqitypes.SmeltType, pro int32, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	obj, ok := m.shenQiSmeltMap[typ][pos]
	if !ok {
		m.initShenQiSmeltObjsMap(typ)
		obj = m.shenQiSmeltMap[typ][pos]
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextTemplate := m.GetNextSmeltUpTemplate(typ, pos)
		if nextTemplate == nil {
			return
		}
		obj.Level = nextTemplate.Level
		obj.UpPro = 0
		obj.UpNum = 0
	} else {
		obj.UpNum += addTimes
		obj.UpPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return
}

//神器器灵注灵
func (m *PlayerShenQiDataManager) ZhuLingAdvanced(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType, pro int32, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	obj, ok := m.shenQiQiLingMap[typ][subType][pos]
	if !ok {
		m.initShenQiQiLingObjsMap(typ)
		obj = m.shenQiQiLingMap[typ][subType][pos]
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextTemplate := m.GetNextZhuLingTemplate(typ, subType, pos)
		if nextTemplate == nil {
			return
		}
		obj.Level = nextTemplate.Level
		obj.UpPro = 0
		obj.UpNum = 0
	} else {
		obj.UpNum += addTimes
		obj.UpPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return
}

// 设置神器战斗力
func (m *PlayerShenQiDataManager) SetShenQiPower(power int64) {
	if power < 0 {
		return
	}
	if m.shenQiObject.Power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.shenQiObject.Power = power
	m.shenQiObject.UpdateTime = now
	m.shenQiObject.SetModified()
}

//gm使用 修改神器碎片等级
func (m *PlayerShenQiDataManager) GmSetShenQiDebrisLevel(typ shenqitypes.ShenQiType, pos shenqitypes.DebrisType, lev int32) *PlayerShenQiDebrisObject {
	obj, ok := m.shenQiDebrisMap[typ][pos]
	if !ok {
		m.initShenQiDebrisObjsMap(typ)
		obj = m.shenQiDebrisMap[typ][pos]
	}
	now := global.GetGame().GetTimeService().Now()
	obj.Level = lev
	obj.UpPro = 0
	obj.UpNum = 0
	obj.UpdateTime = now
	obj.SetModified()
	return obj
}

//gm使用 修改神器淬炼等级
func (m *PlayerShenQiDataManager) GmSetShenQiSmeltLevel(typ shenqitypes.ShenQiType, pos shenqitypes.SmeltType, lev int32) *PlayerShenQiSmeltObject {
	obj, ok := m.shenQiSmeltMap[typ][pos]
	if !ok {
		m.initShenQiSmeltObjsMap(typ)
		obj = m.shenQiSmeltMap[typ][pos]
	}
	now := global.GetGame().GetTimeService().Now()
	obj.Level = lev
	obj.UpPro = 0
	obj.UpNum = 0
	obj.UpdateTime = now
	obj.SetModified()
	return obj
}

//gm使用 修改神器器灵注灵等级
func (m *PlayerShenQiDataManager) GmSetShenQiZhuLingLevel(typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType, lev int32) *PlayerShenQiQiLingObject {
	obj, ok := m.shenQiQiLingMap[typ][subType][pos]
	if !ok {
		m.initShenQiQiLingObjsMap(typ)
		obj = m.shenQiQiLingMap[typ][subType][pos]
	}
	now := global.GetGame().GetTimeService().Now()
	obj.Level = lev
	obj.UpPro = 0
	obj.UpNum = 0
	obj.UpdateTime = now
	obj.SetModified()
	return obj
}

//gm使用 修改灵气值
func (m *PlayerShenQiDataManager) GmSetLingQiNum(val int64) (flag bool) {
	if val < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.shenQiObject.LingQiNum = val
	m.shenQiObject.UpdateTime = now
	m.shenQiObject.SetModified()

	gameevent.Emit(shenqieventtypes.EventTypeShenQiLingQiNumChanged, m.p, m.shenQiObject)

	flag = true
	return
}

func CreatePlayerShenQiDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerShenQiDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerShenQiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerShenQiDataManager))
}
