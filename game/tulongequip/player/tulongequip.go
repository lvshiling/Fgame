package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/tulongequip/dao"
	tulongequipeventtypes "fgame/fgame/game/tulongequip/event/types"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fgame/fgame/pkg/idutil"
)

//玩家屠龙装备管理器
type PlayerTuLongEquipDataManager struct {
	p player.Player
	//屠龙装备背包
	tulongBagMap map[tulongequiptypes.TuLongSuitType]*BodyBag
	//套装技能数据
	suitSkillObjMap map[tulongequiptypes.TuLongSuitType]*PlayerTuLongSuitSkillObject
	//屠龙装数据
	tuLongEquipObject *PlayerTuLongEquipObject
}

func (m *PlayerTuLongEquipDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerTuLongEquipDataManager) Load() (err error) {
	//加载装备数据
	err = m.loadTuLongEquipSlot()
	if err != nil {
		return
	}

	// 加载装备技能
	err = m.loadTuLongEquipSkill()
	if err != nil {
		return
	}

	// 加载屠龙装备数据
	err = m.loadTuLongEquip()
	if err != nil {
		return
	}
	return nil
}

//加载后
func (m *PlayerTuLongEquipDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerTuLongEquipDataManager) Heartbeat() {}

//获取装备背包
func (m *PlayerTuLongEquipDataManager) GetTuLongEquipBag(suitType tulongequiptypes.TuLongSuitType) *BodyBag {
	return m.tulongBagMap[suitType]
}

//加载身上装备
func (m *PlayerTuLongEquipDataManager) loadTuLongEquipSlot() (err error) {
	//加载装备槽位
	tulongSlotList, err := dao.GetTuLongEquipDao().GetTuLongEquipSlotList(m.p.GetId())
	if err != nil {
		return
	}

	slotListMap := make(map[tulongequiptypes.TuLongSuitType][]*PlayerTuLongEquipSlotObject)
	for _, slot := range tulongSlotList {
		pio := NewPlayerTuLongEquipSlotObject(m.p)
		pio.FromEntity(slot)
		slotListMap[pio.suitType] = append(slotListMap[pio.suitType], pio)
	}

	m.tulongBagMap = make(map[tulongequiptypes.TuLongSuitType]*BodyBag)
	for initType := tulongequiptypes.MinSuitType; initType <= tulongequiptypes.MaxSuitType; initType++ {
		slotList := slotListMap[initType]
		m.tulongBagMap[initType] = createBodyBag(m.p, initType, slotList)
	}
	return
}

//加载装备技能
func (m *PlayerTuLongEquipDataManager) loadTuLongEquipSkill() (err error) {
	skillEntityList, err := dao.GetTuLongEquipDao().GetTuLongSuitSkillList(m.p.GetId())
	if err != nil {
		return
	}

	m.suitSkillObjMap = make(map[tulongequiptypes.TuLongSuitType]*PlayerTuLongSuitSkillObject)
	for _, entity := range skillEntityList {
		obj := NewPlayerTuLongSuitSkillObject(m.p)
		obj.FromEntity(entity)
		m.suitSkillObjMap[obj.suitType] = obj
	}
	return
}

//加载装备技能
func (m *PlayerTuLongEquipDataManager) loadTuLongEquip() (err error) {
	tuLongEntity, err := dao.GetTuLongEquipDao().GetTuLongEquipEntity(m.p.GetId())
	if err != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj := NewPlayerTuLongEquipObject(m.p)
	if tuLongEntity == nil {
		id, _ := idutil.GetId()
		obj.id = id
		obj.createTime = now
		obj.SetModified()
	} else {
		obj.FromEntity(tuLongEntity)
	}
	m.tuLongEquipObject = obj
	return
}

//使用装备
func (m *PlayerTuLongEquipDataManager) PutOn(suitType tulongequiptypes.TuLongSuitType, pos inventorytypes.BodyPositionType, itemId int32, bind itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) (flag bool) {
	tulongBag := m.GetTuLongEquipBag(suitType)
	if tulongBag == nil {
		return
	}

	flag = tulongBag.PutOn(pos, itemId, bind, propertyData)
	if flag {
		eventData := tulongequipeventtypes.CreatePlayerTuLongEquipChangedEventData(suitType, itemId)
		gameevent.Emit(tulongequipeventtypes.EventTypeTuLongEquipPutOn, m.p, eventData)
	}

	return
}

//脱下装备
func (m *PlayerTuLongEquipDataManager) TakeOff(suitType tulongequiptypes.TuLongSuitType, pos inventorytypes.BodyPositionType) (itemId int32) {
	tulongBag := m.GetTuLongEquipBag(suitType)
	if tulongBag == nil {
		return
	}

	flag, itemId := tulongBag.TakeOff(pos)
	if flag {
		eventData := tulongequipeventtypes.CreatePlayerTuLongEquipChangedEventData(suitType, itemId)
		gameevent.Emit(tulongequipeventtypes.EventTypeTuLongEquipTakeOff, m.p, eventData)
	}
	return
}

//获取装备
func (m *PlayerTuLongEquipDataManager) GetTuLongEquipByPos(suitType tulongequiptypes.TuLongSuitType, pos inventorytypes.BodyPositionType) *PlayerTuLongEquipSlotObject {
	tulongBag := m.GetTuLongEquipBag(suitType)
	if tulongBag == nil {
		return nil
	}

	return tulongBag.GetByPosition(pos)
}

//获取所有套装数量
func (m *PlayerTuLongEquipDataManager) GetAllTuLongEquipGroupNum() map[tulongequiptypes.TuLongSuitType]map[int32]int32 {
	allGroupMap := make(map[tulongequiptypes.TuLongSuitType]map[int32]int32)
	for suitType, tulongBag := range m.tulongBagMap {
		allGroupMap[suitType] = m.countSuitGroupNum(tulongBag.GetAll())
	}

	return allGroupMap
}

//获取屠龙类型套装数量
func (m *PlayerTuLongEquipDataManager) GetTuLongEquipGroupNumByType(suitType tulongequiptypes.TuLongSuitType) map[int32]int32 {
	tulongBag := m.tulongBagMap[suitType]
	return m.countSuitGroupNum(tulongBag.GetAll())
}

//获取阶数装备数量
func (m *PlayerTuLongEquipDataManager) GetTuLongEquipNumByJieShu(suitType tulongequiptypes.TuLongSuitType, jieshu int32) int32 {
	tulongBag := m.tulongBagMap[suitType]

	equipNum := int32(0)
	for _, slot := range tulongBag.GetAll() {
		if slot.IsEmpty() {
			continue
		}
		itemTemp := item.GetItemService().GetItem(int(slot.GetItemId()))
		number := itemTemp.GetTuLongEquipTemplate().Number
		if number < jieshu {
			continue
		}

		equipNum += 1
	}
	return equipNum
}

//获取装备总强化等级
func (m *PlayerTuLongEquipDataManager) GetTuLongEquipTotalLevel(suitType tulongequiptypes.TuLongSuitType) int32 {
	tulongBag := m.tulongBagMap[suitType]

	totalLevel := int32(0)
	for _, slot := range tulongBag.GetAll() {
		if slot.IsEmpty() {
			continue
		}

		totalLevel += slot.level
	}
	return totalLevel
}

func (m *PlayerTuLongEquipDataManager) countSuitGroupNum(slotList []*PlayerTuLongEquipSlotObject) map[int32]int32 {
	suitGroupMap := make(map[int32]int32)
	for _, slot := range slotList {
		if slot.IsEmpty() {
			continue
		}
		itemTemp := item.GetItemService().GetItem(int(slot.GetItemId()))
		groupId := itemTemp.GetTuLongEquipTemplate().SuitGroup
		if groupId == 0 {
			continue
		}

		_, ok := suitGroupMap[groupId]
		if ok {
			suitGroupMap[groupId] += int32(1)
		} else {
			suitGroupMap[groupId] = int32(1)
		}
	}
	return suitGroupMap
}

//装备改变
func (m *PlayerTuLongEquipDataManager) GetChangedEquipmentSlotAndResetMap() map[tulongequiptypes.TuLongSuitType][]*PlayerTuLongEquipSlotObject {
	changedMap := make(map[tulongequiptypes.TuLongSuitType][]*PlayerTuLongEquipSlotObject)
	for suitType, bag := range m.tulongBagMap {
		changedList := bag.GetChangedSlotAndReset()
		changedMap[suitType] = changedList
	}
	return changedMap
}

//获取所有屠龙装备
func (m *PlayerTuLongEquipDataManager) GetAllEquipSlotMap() map[tulongequiptypes.TuLongSuitType][]*PlayerTuLongEquipSlotObject {
	allSlotMap := make(map[tulongequiptypes.TuLongSuitType][]*PlayerTuLongEquipSlotObject)
	for suitType, tulongBag := range m.tulongBagMap {
		allSlotMap[suitType] = tulongBag.GetAll()
	}
	return allSlotMap
}

//更新身上装备等级
func (m *PlayerTuLongEquipDataManager) UpdateTuLongEquipLevel(suitType tulongequiptypes.TuLongSuitType, pos inventorytypes.BodyPositionType) bool {
	tulongBag := m.GetTuLongEquipBag(suitType)
	if tulongBag == nil {
		return false
	}

	slotIt := tulongBag.GetByPosition(pos)
	if slotIt == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	slotIt.level += 1
	slotIt.updateTime = now
	slotIt.SetModified()
	tulongBag.changed(int(pos))

	gameevent.Emit(tulongequipeventtypes.EventTypeTuLongEquipStrengSuccess, m.p, nil)
	return true
}

// func (m *PlayerTuLongEquipDataManager) ToTuLongEquipSlotList() (slotInfoList []*tulongequiptypes.TuLongEquipSlotInfo) {
// 	for _, slot := range m.tulongBag.GetAll() {
// 		slotInfo := &tulongequiptypes.TuLongEquipSlotInfo{}
// 		slotInfo.SlotId = int32(slot.GetSlotId())
// 		slotInfo.Level = slot.GetLevel()
// 		slotInfo.ItemId = slot.GetItemId()
// 		slotInfoList = append(slotInfoList, slotInfo)
// 		data, ok := slot.GetPropertyData().(*tulongequiptypes.TuLongEquipPropertyData)
// 		if !ok {
// 			//临时处理bug
// 			slotInfo.PropertyData = tulongequiptypes.NewTuLongEquipPropertyData()
// 			slotInfo.PropertyData.InitBase()
// 		} else {
// 			slotInfo.PropertyData = data
// 		}
// 	}
// 	return
// }

// 获取装备技能
func (m *PlayerTuLongEquipDataManager) GetSuitSkillList() map[tulongequiptypes.TuLongSuitType]*PlayerTuLongSuitSkillObject {
	return m.suitSkillObjMap
}

// 获取装备技能
func (m *PlayerTuLongEquipDataManager) GetSuitSkillLevel(suitType tulongequiptypes.TuLongSuitType) int32 {
	skillObj, ok := m.suitSkillObjMap[suitType]
	if !ok {
		return 0
	}
	return skillObj.level
}

// 升级装备技能
func (m *PlayerTuLongEquipDataManager) UpgradeSuitSkill(suitType tulongequiptypes.TuLongSuitType) {
	now := global.GetGame().GetTimeService().Now()
	skillObj, ok := m.suitSkillObjMap[suitType]
	if !ok {
		id, _ := idutil.GetId()
		skillObj = NewPlayerTuLongSuitSkillObject(m.p)
		skillObj.id = id
		skillObj.suitType = suitType
		skillObj.level = 0
		skillObj.createTime = now
		m.suitSkillObjMap[suitType] = skillObj
	}

	skillObj.level += 1
	skillObj.updateTime = now
	skillObj.SetModified()

	gameevent.Emit(tulongequipeventtypes.EventTypeTuLongEquipSkillUpgrade, m.p, skillObj)
}

func (m *PlayerTuLongEquipDataManager) GetPower() int64 {
	return m.tuLongEquipObject.power
}

func (m *PlayerTuLongEquipDataManager) SetPower(power int64) {
	if power < 0 {
		return
	}
	if m.tuLongEquipObject.power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.tuLongEquipObject.power = power
	m.tuLongEquipObject.updateTime = now
	m.tuLongEquipObject.SetModified()
}

func CreatePlayerTuLongEquipDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerTuLongEquipDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerTuLongEquipDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTuLongEquipDataManager))
}
