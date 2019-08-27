package player

import (
	additionsyscommon "fgame/fgame/game/additionsys/common"
	"fgame/fgame/game/additionsys/dao"
	additionsyseventtypes "fgame/fgame/game/additionsys/event/types"
	additionsystypes "fgame/fgame/game/additionsys/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//玩家 附加系统管理器
type PlayerAdditionSysDataManager struct {
	p player.Player
	// 装备背包
	additionSysEquipBagMap map[additionsystypes.AdditionSysType]*BodyBag
	// 系统灵珠
	additionSysLingZhuMap map[additionsystypes.AdditionSysType]map[additionsystypes.LingZhuType]*PlayerAdditionSysLingZhuObject
	// 系统等级
	additionSysLevelMap map[additionsystypes.AdditionSysType]*PlayerAdditionSysLevelObject
	// 系统觉醒
	additionSysAwakeMap map[additionsystypes.AdditionSysType]*PlayerAdditionSysAwakeObject
	// 系统通灵
	additionSysTongLingMap map[additionsystypes.AdditionSysType]*PlayerAdditionSysTongLingObject
	// 玩家圣痕
	playerShengHenObject *PlayerShengHenObject
}

func (m *PlayerAdditionSysDataManager) Player() player.Player {
	return m.p
}

//根据获取装备背包
func (m *PlayerAdditionSysDataManager) GetAdditionSysEquipBagByType(typ additionsystypes.AdditionSysType) *BodyBag {
	to, ok := m.additionSysEquipBagMap[typ]
	if !ok {
		return nil
	}
	return to
}

//系统灵珠信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysLingZhu(typ additionsystypes.AdditionSysType, lingzhutype additionsystypes.LingZhuType) *PlayerAdditionSysLingZhuObject {
	objMap, ok := m.additionSysLingZhuMap[typ]
	if !ok {
		return nil
	}
	obj, ook := objMap[lingzhutype]
	if !ook {
		return nil
	}
	return obj
}

//系统灵珠信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysLingZhuMap(typ additionsystypes.AdditionSysType) map[additionsystypes.LingZhuType]*PlayerAdditionSysLingZhuObject {
	objMap, ok := m.additionSysLingZhuMap[typ]
	if !ok {
		return nil
	}

	return objMap
}

//系统升级操作
func (m *PlayerAdditionSysDataManager) LingZhuUpLevel(typ additionsystypes.AdditionSysType, lingZhuType additionsystypes.LingZhuType, sucess bool, bless int64) {
	obj := m.GetAdditionSysLingZhu(typ, lingZhuType)
	if obj == nil {
		obj = createPlayerAdditionSysLingZhuObject(m.p, typ, lingZhuType)
		obj.SetModified()
		lingZhuMap, ok := m.additionSysLingZhuMap[typ]
		if !ok {
			lingZhuMap = make(map[additionsystypes.LingZhuType]*PlayerAdditionSysLingZhuObject)
			m.additionSysLingZhuMap[typ] = lingZhuMap
		}
		lingZhuMap[lingZhuType] = obj
	}
	if sucess {
		obj.level = obj.level + 1
		obj.times = 0
		obj.bless = 0
	} else {
		obj.times = obj.times + 1
		obj.bless += bless
	}
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
}

//系统觉醒信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysAwakeInfoByType(typ additionsystypes.AdditionSysType) *PlayerAdditionSysAwakeObject {
	return m.getAdditionSysAwakeInfo(typ)
}

//系统等级信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysLevelInfoByType(typ additionsystypes.AdditionSysType) *PlayerAdditionSysLevelObject {
	return m.getAdditionSysLevelInfo(typ)
}

//系统通灵信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysTongLingInfoByType(typ additionsystypes.AdditionSysType) *PlayerAdditionSysTongLingObject {
	to, ok := m.additionSysTongLingMap[typ]
	if !ok {
		return nil
	}
	return to
}

//系统等级信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysLevelInfoAll() map[additionsystypes.AdditionSysType]*PlayerAdditionSysLevelObject {
	return m.additionSysLevelMap
}

//系统通灵信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysTongLingInfoAll() map[additionsystypes.AdditionSysType]*PlayerAdditionSysTongLingObject {
	return m.additionSysTongLingMap
}

//系统觉醒信息对象
func (m *PlayerAdditionSysDataManager) GetAdditionSysAwakeInfoAll() map[additionsystypes.AdditionSysType]*PlayerAdditionSysAwakeObject {
	return m.additionSysAwakeMap
}

//加载
func (m *PlayerAdditionSysDataManager) Load() (err error) {
	//加载装备数据
	err = m.loadAdditionSysBags()
	if err != nil {
		return
	}
	//加载系统等级数据
	err = m.loadAdditionSysLevel()
	if err != nil {
		return
	}

	//加载系统觉醒数据
	err = m.loadAdditionSysAwake()
	if err != nil {
		return
	}

	//加载系统其他数据
	err = m.loadAdditionSysTongLing()
	if err != nil {
		return
	}

	// 加载玩家圣痕数据
	err = m.loadPlayerShengHen()
	if err != nil {
		return
	}

	//加载系统灵珠数据
	err = m.loadAdditionSysLingZhu()
	if err != nil {
		return
	}

	return nil
}

//加载身上装备
func (m *PlayerAdditionSysDataManager) loadAdditionSysBags() (err error) {
	m.additionSysEquipBagMap = make(map[additionsystypes.AdditionSysType]*BodyBag)
	//加载槽位
	equipmentSlotList, err := dao.GetAdditionSysDao().GetAdditionSysSlotList(m.p.GetId())
	if err != nil {
		return
	}
	slotMap := make(map[additionsystypes.AdditionSysType]map[additionsystypes.SlotPositionType]*PlayerAdditionSysSlotObject)
	for _, slot := range equipmentSlotList {
		pio := NewPlayerAdditionSysSlotObject(m.p)
		pio.FromEntity(slot)
		tempM, ok := slotMap[pio.SysType]
		if !ok {
			tempM = make(map[additionsystypes.SlotPositionType]*PlayerAdditionSysSlotObject)
			slotMap[pio.SysType] = tempM
		}
		tempM[pio.SlotId] = pio
	}

	for sysTypeId := additionsystypes.MinType; sysTypeId <= additionsystypes.MaxType; sysTypeId++ {
		if !sysTypeId.Valid() {
			continue
		}
		m.additionSysEquipBagMap[sysTypeId] = createBodyBag(m.p, sysTypeId, slotMap[sysTypeId])
	}

	return
}

//加载系统等级
func (m *PlayerAdditionSysDataManager) loadAdditionSysLevel() (err error) {
	m.additionSysLevelMap = make(map[additionsystypes.AdditionSysType]*PlayerAdditionSysLevelObject)
	//加载数据库
	levelList, err := dao.GetAdditionSysDao().GetAdditionSysLevelList(m.p.GetId())
	if err != nil {
		return
	}

	for _, ety := range levelList {
		pio := NewPlayerAdditionSysLevelObject(m.p)
		pio.FromEntity(ety)
		m.additionSysLevelMap[pio.SysType] = pio
	}

	now := global.GetGame().GetTimeService().Now()
	for sysTypeId := additionsystypes.MinType; sysTypeId <= additionsystypes.MaxType; sysTypeId++ {
		if !sysTypeId.Valid() {
			continue
		}
		if m.additionSysLevelMap[sysTypeId] == nil {
			m.additionSysLevelMap[sysTypeId] = createAdditionSysLevelObject(m.p, sysTypeId, now)
		}
	}

	return
}

//加载系统觉醒
func (m *PlayerAdditionSysDataManager) loadAdditionSysAwake() (err error) {
	m.additionSysAwakeMap = make(map[additionsystypes.AdditionSysType]*PlayerAdditionSysAwakeObject)

	awakeList, err := dao.GetAdditionSysDao().GetAdditionSysAwakeList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range awakeList {
		obj := NewPlayerAdditionSysAwakeObject(m.p)
		obj.FromEntity(entity)
		m.additionSysAwakeMap[obj.SysType] = obj
	}
	now := global.GetGame().GetTimeService().Now()
	for sysTypeId := additionsystypes.MinType; sysTypeId <= additionsystypes.MaxType; sysTypeId++ {
		if !sysTypeId.Valid() {
			continue
		}
		if m.additionSysAwakeMap[sysTypeId] == nil {
			m.additionSysAwakeMap[sysTypeId] = createAdditionSysAwakeObject(m.p, sysTypeId, now)
		}
	}

	return
}

//加载系统通灵
func (m *PlayerAdditionSysDataManager) loadAdditionSysTongLing() (err error) {
	m.additionSysTongLingMap = make(map[additionsystypes.AdditionSysType]*PlayerAdditionSysTongLingObject)
	//加载数据库
	tongLingList, err := dao.GetAdditionSysDao().GetAdditionSysTongLingList(m.p.GetId())

	for _, ety := range tongLingList {
		pio := NewPlayerAdditionSysTongLingObject(m.p)
		pio.FromEntity(ety)
		m.additionSysTongLingMap[pio.SysType] = pio
	}

	now := global.GetGame().GetTimeService().Now()
	for sysTypeId := additionsystypes.MinType; sysTypeId <= additionsystypes.MaxType; sysTypeId++ {
		if !sysTypeId.Valid() {
			continue
		}

		if m.additionSysTongLingMap[sysTypeId] == nil {
			m.additionSysTongLingMap[sysTypeId] = createAdditionSysTongLingObject(m.p, sysTypeId, now)
		}
	}

	return

}

//加载玩家圣痕数据
func (m *PlayerAdditionSysDataManager) loadPlayerShengHen() (err error) {
	//加载数据库
	shengHenEntity, err := dao.GetAdditionSysDao().GetPlayerShengHenEntity(m.p.GetId())
	if err != nil {
		return
	}

	if shengHenEntity == nil {
		now := global.GetGame().GetTimeService().Now()
		obj := createPlayerShengHenObject(m.p, now)
		obj.SetModified()
		m.playerShengHenObject = obj
	} else {
		obj := NewPlayerShengHenObject(m.p)
		obj.FromEntity(shengHenEntity)
		m.playerShengHenObject = obj
	}

	return

}

//加载系统灵珠
func (m *PlayerAdditionSysDataManager) loadAdditionSysLingZhu() (err error) {
	m.additionSysLingZhuMap = make(map[additionsystypes.AdditionSysType]map[additionsystypes.LingZhuType]*PlayerAdditionSysLingZhuObject)
	//加载数据库
	lingZhuList, err := dao.GetAdditionSysDao().GetAdditionSysLingZhuList(m.p.GetId())

	for _, ety := range lingZhuList {
		pio := NewPlayerAdditionSysLingZhuObject(m.p)
		pio.FromEntity(ety)

		tempM, ok := m.additionSysLingZhuMap[pio.sysType]
		if !ok {
			tempM = make(map[additionsystypes.LingZhuType]*PlayerAdditionSysLingZhuObject)
			m.additionSysLingZhuMap[pio.sysType] = tempM
		}
		tempM[pio.lingZhuType] = pio
	}

	// for sysType := additionsystypes.MinAdditionSysTypeLingTongEquipType; sysType <= additionsystypes.MaxAdditionSysTypeLingTongEquipType; sysType++ {
	// 	if !sysType.Valid() {
	// 		continue
	// 	}
	// 	for lingzhuid := additionsystypes.MinLingZhuType; lingzhuid <= additionsystypes.MaxLingZhuType; lingzhuid++ {
	// 		if m.additionSysLingZhuMap[sysType] == nil {
	// 			m.additionSysLingZhuMap[sysType] = make(map[additionsystypes.LingZhuType]*PlayerAdditionSysLingZhuObject)
	// 		}
	// 		if m.additionSysLingZhuMap[sysType][lingzhuid] == nil {
	// 			newobj := createPlayerAdditionSysLingZhuObject(m.p, sysType, lingzhuid)
	// 			m.additionSysLingZhuMap[sysType][lingzhuid] = newobj
	// 			newobj.SetModified()
	// 		}
	// 	}
	// }
	return
}

//加载后
func (m *PlayerAdditionSysDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerAdditionSysDataManager) Heartbeat() {}

//系统升级操作
func (m *PlayerAdditionSysDataManager) SystemUplevel(typ additionsystypes.AdditionSysType, pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	obj := m.getAdditionSysLevelInfo(typ)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		template := obj.GetNextShengJiTemplate()
		if template == nil {
			return
		}
		obj.Level += 1
		obj.UpNum = 0
		obj.UpPro = 0

		gameevent.Emit(additionsyseventtypes.EventTypeAdditionSysShengJi, m.p, typ)
	} else {
		obj.UpNum += addTimes
		obj.UpPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return
}

//获取附装背包
func (m *PlayerAdditionSysDataManager) GetAdditionSysEquipBags() map[additionsystypes.AdditionSysType]*BodyBag {
	return m.additionSysEquipBagMap
}

//获取附装数量
func (m *PlayerAdditionSysDataManager) GetAdditionSysEquipNum(typ additionsystypes.AdditionSysType, qualityCondition int32) int32 {
	bag := m.GetAdditionSysEquipBagByType(typ)
	equipNum := int32(0)
	for _, obj := range bag.slotMap {
		if obj.IsEmpty() {
			continue
		}

		itemTemp := item.GetItemService().GetItem(int(obj.ItemId))
		if itemTemp.Quality < qualityCondition {
			continue
		}

		equipNum += 1
	}

	return equipNum
}

//获取部位信息
func (m *PlayerAdditionSysDataManager) GetAdditionSysByArg(typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) *PlayerAdditionSysSlotObject {
	equipBag, ok := m.additionSysEquipBagMap[typ]
	if !ok {
		return nil
	}
	item := equipBag.GetByPosition(pos)
	if item == nil {
		return nil
	}

	return item
}

//使用装备
func (m *PlayerAdditionSysDataManager) PutOn(typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType, itemId int32, bind itemtypes.ItemBindType) (flag bool) {
	equipBag, ok := m.additionSysEquipBagMap[typ]
	if !ok {
		return
	}
	flag = equipBag.PutOn(pos, itemId, bind)
	return
}

//脱下装备
func (m *PlayerAdditionSysDataManager) TakeOff(typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) (itemId int32) {
	equipBag, ok := m.additionSysEquipBagMap[typ]
	if !ok {
		return
	}
	//判断是否可以脱下
	flag := m.IfCanTakeOff(typ, pos)
	if !flag {
		return
	}

	itemId = equipBag.TakeOff(pos)
	return
}

//装备改变
func (m *PlayerAdditionSysDataManager) GetChangedEquipmentSlotAndReset(typ additionsystypes.AdditionSysType) (itemList []*PlayerAdditionSysSlotObject) {
	return m.additionSysEquipBagMap[typ].GetChangedSlotAndReset()
}

//是否可以卸下
func (m *PlayerAdditionSysDataManager) IfCanTakeOff(typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) bool {
	equipBag, ok := m.additionSysEquipBagMap[typ]
	if !ok {
		return false
	}
	bodySlot := equipBag.GetByPosition(pos)
	if bodySlot == nil {
		return false
	}
	if bodySlot.IsEmpty() {
		return false
	}
	return true
}

// 获取圣痕战斗力
func (m *PlayerAdditionSysDataManager) GetShengHenPower() int64 {
	return m.playerShengHenObject.power
}

// 设置圣痕战斗力
func (m *PlayerAdditionSysDataManager) SetShengHenPower(power int64) {
	if power <= 0 {
		return
	}
	if m.playerShengHenObject.power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerShengHenObject.power = power
	m.playerShengHenObject.updateTime = now
	m.playerShengHenObject.SetModified()
}

func CreatePlayerAdditionSysDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerAdditionSysDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerAdditionSysDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerAdditionSysDataManager))
}

//紧gm命令使用 强化等级修改
func (m *PlayerAdditionSysDataManager) GmSetSlotLev(typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType, lev int32) bool {
	equipBag, ok := m.additionSysEquipBagMap[typ]
	if !ok {
		return false
	}
	item := equipBag.GetByPosition(pos)
	if item == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	item.UpdateTime = now
	item.Level = lev
	item.SetModified()
	equipBag.changed(int(pos))
	return true
}

//紧gm命令使用 神铸等级修改
func (m *PlayerAdditionSysDataManager) GmSetShenZhuLev(typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType, lev int32) bool {
	equipBag, ok := m.additionSysEquipBagMap[typ]
	if !ok {
		return false
	}
	item := equipBag.GetByPosition(pos)
	if item == nil {
		return false
	}

	if lev == item.ShenZhuLev {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	item.UpdateTime = now
	item.ShenZhuLev = lev
	item.ShenZhuNum = 0
	item.ShenZhuPro = 0
	item.SetModified()
	equipBag.changed(int(pos))
	return true
}

//紧gm命令使用 附加系统升级等级修改
func (m *PlayerAdditionSysDataManager) GmSetShengJi(typ additionsystypes.AdditionSysType, lev int32) bool {
	levInfo, ok := m.additionSysLevelMap[typ]
	if !ok {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	levInfo.Level = lev
	levInfo.UpNum = 0
	levInfo.UpPro = 0
	levInfo.UpdateTime = now
	levInfo.SetModified()
	return true
}

//紧gm命令使用 附加系统化灵丹等级修改
func (m *PlayerAdditionSysDataManager) GmSetHuaLingLevel(typ additionsystypes.AdditionSysType, lev int32) bool {
	levInfo, ok := m.additionSysLevelMap[typ]
	if !ok {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	levInfo.LingLevel = lev
	levInfo.LingNum = 0
	levInfo.LingPro = 0
	levInfo.UpdateTime = now
	levInfo.SetModified()
	return true
}

//紧gm命令使用 附加系统通灵等级修改
func (m *PlayerAdditionSysDataManager) GmSetTongLingLevel(typ additionsystypes.AdditionSysType, lev int32) bool {
	tongLingInfo, ok := m.additionSysTongLingMap[typ]
	if !ok {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	tongLingInfo.TongLingLev = lev
	tongLingInfo.TongLingNum = 0
	tongLingInfo.TongLingPro = 0
	tongLingInfo.UpdateTime = now
	tongLingInfo.SetModified()
	return true
}

//仅gm使用 附加系统觉醒
func (m *PlayerAdditionSysDataManager) GmSetAwake(typ additionsystypes.AdditionSysType) bool {
	awakeInfo, ok := m.additionSysAwakeMap[typ]
	if !ok {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	awakeInfo.IsAwake = 1
	awakeInfo.UpdateTime = now
	awakeInfo.SetModified()

	return true
}

func (m *PlayerAdditionSysDataManager) ToAllAdditionSysInfo() *additionsyscommon.AllAdditionSysInfo {
	allAdditionSysInfo := &additionsyscommon.AllAdditionSysInfo{}
	for typ, additionSysObj := range m.additionSysLevelMap {
		awakeObj := m.additionSysAwakeMap[typ]
		additionSysInfo := &additionsyscommon.AdditionSysInfo{
			SysType:   int32(typ),
			Level:     additionSysObj.Level,
			UpPro:     additionSysObj.UpPro,
			LingLevel: additionSysObj.LingLevel,
			LingPro:   additionSysObj.LingPro,
			IsAwake:   awakeObj.IsAwake,
		}
		tongLingObj := m.additionSysTongLingMap[typ]
		additionSysInfo.TongLingInfo = &additionsyscommon.AdditionSysTongLingInfo{
			TongLingLev: tongLingObj.TongLingLev,
			TongLingPro: tongLingObj.TongLingPro,
		}
		for subType, additionSysSlotObj := range m.additionSysEquipBagMap[typ].slotMap {
			additionSysSlotInfo := &additionsyscommon.AdditionSysSlotInfo{
				SlotId:     int32(subType),
				Level:      additionSysSlotObj.Level,
				ItemId:     additionSysSlotObj.ItemId,
				ShenZhuLev: additionSysSlotObj.ShenZhuLev,
				ShenZhuPro: additionSysSlotObj.ShenZhuPro,
			}
			additionSysInfo.SysTypeSlotList = append(additionSysInfo.SysTypeSlotList, additionSysSlotInfo)
		}
		allAdditionSysInfo.AdditionSysList = append(allAdditionSysInfo.AdditionSysList, additionSysInfo)
	}
	return allAdditionSysInfo
}

func (m *PlayerAdditionSysDataManager) getAdditionSysLevelInfo(typ additionsystypes.AdditionSysType) *PlayerAdditionSysLevelObject {
	to, ok := m.additionSysLevelMap[typ]
	if !ok {
		return nil
	}
	return to
}

func (m *PlayerAdditionSysDataManager) getAdditionSysAwakeInfo(typ additionsystypes.AdditionSysType) *PlayerAdditionSysAwakeObject {
	to, ok := m.additionSysAwakeMap[typ]
	if !ok {
		return nil
	}
	return to
}

func (o *PlayerAdditionSysLingZhuObject) UpLevel(sucess bool, bless int64) {
	if sucess {
		o.level = o.level + 1
		o.times = 0
		o.bless = 0
	} else {
		o.times = o.times + 1
		o.bless += bless
	}
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
}
